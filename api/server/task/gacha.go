package task

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"math/big"
	"mm-ndj/config"
	"mm-ndj/contract/gacha"
	"mm-ndj/contract/ticket"
	"mm-ndj/model"
	mathUtil "mm-ndj/pkg/kit/math"
	"mm-ndj/server/dao"
	"strings"
	"time"
)

var (
	gachaNotFirstFlags          = false
	gachaStartBlockHeight int64 = 0
	gachaQueryNotEnded          = false
	gachaQueryBlock       int64 = 1000
	serverCtx                   = dao.GetServiceCtx()
	tokenDecimals               = 18

	poolLockedList = map[int64]int{
		1: 7,
		2: 14,
		3: 30,
		4: 90,
		5: 180,
		6: 356,
		7: 1,
	}
)

//    event Deposit(address indexed user, uint256 indexed pid, uint256 amount, uint256 lockTimestamp);
//    event Withdraw(address indexed user, uint256 amount);
//	  event Transfer(address indexed from, address indexed to, uint256 value);

func MonitorGacha() {
	prefix := "[MonitorGacha]"

	var (
		query     ethereum.FilterQuery
		fromBlock int64
		toBlock   int64
	)

	client, err := NewETHClient()
	if err != nil {
		config.Logger.Error("connect rpc error", zap.Error(err))
		return
	}
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		config.Logger.Error("get header error ", zap.Error(err))
		return
	}

	// 判断上一次query有没有结束
	if !gachaQueryNotEnded {
		gachaQueryNotEnded = true
		if gachaNotFirstFlags {

			mmContract := serverCtx.C.RangersC.MM
			ticketContract := serverCtx.C.RangersC.Ticket

			depositEvent := crypto.Keccak256Hash([]byte("Deposit(address,uint256,uint256,uint256,uint256)"))
			withdrawEvent := crypto.Keccak256Hash([]byte("Withdraw(address,uint256)"))
			transferEvent := crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))
			fmt.Println("depositEvent", depositEvent.Hex())
			fmt.Println("withdrawEvent", withdrawEvent.Hex())
			fmt.Println("transferEvent", transferEvent.Hex())
			startBlock := gachaStartBlockHeight
			currentBlock := header.Number.Int64()
			if startBlock >= currentBlock {
				fmt.Println(prefix, "currentBlock <= startBlock, continue", currentBlock, "<=", startBlock)
				gachaQueryNotEnded = false
				return
			} else {
				if currentBlock-startBlock > gachaQueryBlock {
					fromBlock = startBlock + 1
					toBlock = startBlock + gachaQueryBlock

					gachaStartBlockHeight = startBlock + gachaQueryBlock
				} else {
					fromBlock = startBlock + 1
					toBlock = currentBlock

					gachaStartBlockHeight = currentBlock
				}
			}

			fmt.Println(prefix, "from/to", fromBlock, "/", toBlock)
			// 过滤日志
			query = ethereum.FilterQuery{
				FromBlock: big.NewInt(fromBlock),
				ToBlock:   big.NewInt(toBlock),
				Addresses: []common.Address{
					common.HexToAddress(mmContract),
					common.HexToAddress(ticketContract), // TODO: 是否可以过滤到数据
				},
			}

			logs, err := client.FilterLogs(context.Background(), query)
			if err != nil {
				fmt.Println(prefix, "filter logs error ", err)
				gachaQueryNotEnded = false
				return
			}

			for _, log := range logs {
				tx := serverCtx.Db.Begin()

				switch log.Topics[0].Hex() {
				case depositEvent.Hex():
					config.Logger.Info("deposit event")
					if err := deposit(tx, &log, client); err != nil {
						fmt.Println(prefix, "deposit error", err)
						tx.Rollback()
						continue
					}
				case withdrawEvent.Hex():
					config.Logger.Info("withdraw event")
					if err := withdraw(tx, &log, client); err != nil {
						fmt.Println(prefix, "withdraw error", err)
						tx.Rollback()
						continue
					}

				case transferEvent.Hex():
					config.Logger.Info("transfer event")
					if err := transfer(tx, &log, client); err != nil {
						fmt.Println(prefix, "transfer error ", err)
						tx.Rollback()
						continue
					}
				}

				tx.Commit()
			}

		} else {
			// 第一次从数据库获取区块高度，没有则从链上获取

			b := &model.BurnRecord{}
			w := &model.WithdrawRecord{}
			d := &model.DepositRecord{}

			err1 := serverCtx.Db.Table(b.TableName()).Order("block_height DESC").First(b).Error
			err2 := serverCtx.Db.Table(w.TableName()).Order("block_height DESC").First(w).Error
			err3 := serverCtx.Db.Table(d.TableName()).Order("block_height DESC").First(d).Error

			fmt.Println(err1, err2, err3)
			//err1 := serverCtx.Db.Model(&model.BurnRecord{}).Order("block_height DESC").First(b).Error
			//err2 := serverCtx.Db.Model(&model.WithdrawRecord{}).Order("block_height DESC").First(b).Error
			//err3 := serverCtx.Db.Model(&model.DepositRecord{}).Order("block_height DESC").First(b).Error

			//b, err1 := model.NewBurnRecordModel(context.Background(), serverCtx.Db).FindOneByCondition(func(db *gorm.DB) *gorm.DB {
			//	return db.Order("block_height DESC")
			//})
			//
			//w, err2 := model.NewWithdrawRecordModel(context.Background(), serverCtx.Db).FindOneByCondition(func(db *gorm.DB) *gorm.DB {
			//	return db.Order("block_height DESC")
			//})
			//
			//d, err3 := model.NewDepositRecordModel(context.Background(), serverCtx.Db).FindOneByCondition(func(db *gorm.DB) *gorm.DB {
			//	return db.Order("block_height DESC")
			//})

			if err1 != nil && err2 != nil && err3 != nil {
				gachaStartBlockHeight = header.Number.Int64()
			} else {
				gachaStartBlockHeight = mathUtil.MaxInt64(b.BlockHeight, w.BlockHeight, d.BlockHeight)
			}

			gachaNotFirstFlags = true
		}

		gachaQueryNotEnded = false
	} else {
		// 如果上一次没有执行结束，则直接跳过这次query
		fmt.Println(prefix, " previous filter not ended")
		return
	}

}

// event Deposit(address indexed user, uint256 indexed pid, uint256 amount);
func deposit(tx *gorm.DB, log *types.Log, client *ethclient.Client) error {
	var (
		txStatus int
	)

	depositIndexedLog := struct {
		User common.Address
		PId  *big.Int
	}{}

	depositLog := struct {
		DepositAmount *big.Int
		TicketAmount  *big.Int
		LockTimestamp *big.Int
	}{}

	depositIndexedLog.User = common.BytesToAddress(log.Topics[1][:])
	depositIndexedLog.PId = log.Topics[2].Big()

	gachaABI, err := abi.JSON(strings.NewReader(gacha.MMABI))
	if err != nil {
		config.Logger.Info("read gacha abi error", zap.Error(err))
		return err
	}

	err = gachaABI.UnpackIntoInterface(&depositLog, "Deposit", log.Data)
	if err != nil {
		config.Logger.Info("unpack into interface error", zap.Error(err))
		return err
	}

	d := &model.DepositRecord{}

	if err := tx.Table(d.TableName()).
		Where("hash=?", log.TxHash.Hex()).
		First(&d).Error; err != nil && err != gorm.ErrRecordNotFound {

		config.Logger.Info("query deposit record error", zap.Error(err))
		return err
	}

	// TODO: 判断条件是否正确
	// hash存在，直接跳过，不存在，插入数据
	if d != nil && d.Id != 0 {
		return nil
	}

	txStatus = 1
	receipt, err := client.TransactionReceipt(context.Background(), log.TxHash)
	if err != nil {
		config.Logger.Info("query tx receipt error", zap.Error(err))
		return err
	}
	if receipt.Status == 0 {
		txStatus = 2
	}
	userInfo, err := dao.GetUserByAddress(context.Background(), serverCtx, depositIndexedLog.User.Hex())
	if err != nil {
		config.Logger.Info("query user error", zap.Error(err))
		return err
	}
	depositAmount := decimal.NewFromBigInt(depositLog.DepositAmount, 0).Div(decimal.New(1, int32(tokenDecimals))).InexactFloat64()
	ticketAmount := decimal.NewFromBigInt(depositLog.TicketAmount, 0).Div(decimal.New(1, int32(tokenDecimals))).InexactFloat64()

	i := &model.DepositRecord{
		CreateTime:    time.Now().Unix(),
		LockDay:       poolLockedList[depositIndexedLog.PId.Int64()],
		DepositAmount: depositAmount,
		TicketAmount:  ticketAmount,
		Hash:          log.TxHash.Hex(),
		BlockHeight:   int64(log.BlockNumber),
		UserId:        userInfo.Id,
		Address:       depositIndexedLog.User.Hex(),
		TxStatus:      txStatus,
		Status:        2,
	}

	if err := tx.Table(d.TableName()).Create(i).Error; err != nil {
		config.Logger.Info("create error", zap.Error(err))
		return err
	}

	return nil
}

// event Withdraw(address indexed user, uint256 amount);
func withdraw(tx *gorm.DB, log *types.Log, client *ethclient.Client) error {
	var (
		txStatus int
	)

	withdrawIndexedLog := struct {
		User common.Address
	}{}

	withdrawLog := struct {
		Amount *big.Int
	}{}

	withdrawIndexedLog.User = common.BytesToAddress(log.Topics[1][:])

	gachaABI, err := abi.JSON(strings.NewReader(gacha.MMABI))
	if err != nil {
		config.Logger.Info("read gacha abi error", zap.Error(err))
		return err
	}

	err = gachaABI.UnpackIntoInterface(&withdrawLog, "Withdraw", log.Data)
	if err != nil {
		config.Logger.Info("unpack into interface error", zap.Error(err))
		return err
	}

	w := &model.WithdrawRecord{}

	if err := tx.Table(w.TableName()).
		Where("hash=?", log.TxHash.Hex()).
		First(&w).Error; err != nil && err != gorm.ErrRecordNotFound {

		config.Logger.Info("query withdraw record error", zap.Error(err))
		return err
	}

	// TODO: 判断条件是否正确
	// hash存在，直接跳过，不存在，插入数据
	if w != nil && w.Id != 0 {
		return nil
	}

	txStatus = 1
	receipt, err := client.TransactionReceipt(context.Background(), log.TxHash)
	if err != nil {
		config.Logger.Info("query tx receipt error", zap.Error(err))
		return err
	}
	if receipt.Status == 0 {
		txStatus = 2
	}
	userInfo, err := dao.GetUserByAddress(context.Background(), serverCtx, withdrawIndexedLog.User.Hex())
	if err != nil {
		config.Logger.Info("query user error", zap.Error(err))
		return err
	}
	amount := decimal.NewFromBigInt(withdrawLog.Amount, 0).Div(decimal.New(1, int32(tokenDecimals))).InexactFloat64()
	i := &model.WithdrawRecord{
		UserId:      userInfo.Id,
		Address:     withdrawIndexedLog.User.Hex(),
		Amount:      amount,
		Hash:        log.TxHash.Hex(),
		BlockHeight: int64(log.BlockNumber),
		TxStatus:    txStatus,
		Status:      2,
		CreateTime:  time.Now().Unix(),
	}

	if err := tx.Table(i.TableName()).Create(i).Error; err != nil {
		config.Logger.Info("create error", zap.Error(err))
		return err
	}

	return nil
}

// event Transfer(address indexed from, address indexed to, uint256 value);

func transfer(tx *gorm.DB, log *types.Log, client *ethclient.Client) error {
	var (
		txStatus int
	)

	transferIndexedLog := struct {
		From common.Address
		To   common.Address
	}{}

	transferLog := struct {
		Value *big.Int
	}{}

	transferIndexedLog.From = common.BytesToAddress(log.Topics[1][:])
	transferIndexedLog.To = common.BytesToAddress(log.Topics[2][:])

	ticketABI, err := abi.JSON(strings.NewReader(ticket.ContractABI))
	if err != nil {
		config.Logger.Info("read ticket abi error", zap.Error(err))
		return err
	}

	err = ticketABI.UnpackIntoInterface(&transferLog, "Transfer", log.Data)
	if err != nil {
		config.Logger.Info("unpack into interface error", zap.Error(err))
		return err
	}

	if (transferIndexedLog.To != common.Address{}) {
		return errors.New("not burn tx")
	}

	b := &model.BurnRecord{}

	if err := tx.Table(b.TableName()).
		Where("hash=?", log.TxHash.Hex()).
		First(&b).Error; err != nil && err != gorm.ErrRecordNotFound {

		config.Logger.Info("query burn record error", zap.Error(err))
		return err
	}

	// TODO: 判断条件是否正确
	// hash存在，直接跳过，不存在，插入数据
	if b != nil && b.Id != 0 {
		return nil
	}

	txStatus = 1
	receipt, err := client.TransactionReceipt(context.Background(), log.TxHash)
	if err != nil {
		config.Logger.Info("query tx receipt error", zap.Error(err))
		return err
	}
	if receipt.Status == 0 {
		txStatus = 2
	}

	userInfo, err := dao.GetUserByAddress(context.Background(), serverCtx, transferIndexedLog.From.Hex())
	if err != nil {
		config.Logger.Info("query user error", zap.Error(err))
		return err
	}
	amount := decimal.NewFromBigInt(transferLog.Value, 0).Div(decimal.New(1, int32(tokenDecimals))).InexactFloat64()

	i := &model.BurnRecord{
		UserId:      userInfo.Id,
		Address:     transferIndexedLog.From.Hex(),
		Hash:        log.TxHash.Hex(),
		BlockHeight: int64(log.BlockNumber),
		TxStatus:    txStatus,
		Amount:      amount,
		Status:      1,
		CreateTime:  time.Now().Unix(),
	}

	if err := tx.Table(i.TableName()).Create(i).Error; err != nil {
		config.Logger.Info("create error", zap.Error(err))
		return err
	}

	newServerCtx := &dao.ServiceCtx{
		C:   serverCtx.C,
		Db:  tx,
		Rds: serverCtx.Rds,
	}
	err = BurnTicketSuccess(context.Background(), newServerCtx, transferIndexedLog.From.Hex(), log.TxHash.Hex(), amount)
	if err != nil {
		return err
	}

	return nil
}

func NewETHClient() (*ethclient.Client, error) {
	url := serverCtx.C.RangersC.RPC
	conn, err := ethclient.Dial(url)
	if err != nil {
		config.Logger.Info("connect to RPC error: ", zap.Error(err))
		return nil, err
	}

	return conn, nil
}
