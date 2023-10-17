package task

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/robfig/cron"
	"gorm.io/gorm"
	"math/big"
	"mm-ndj/config"
	model2 "mm-ndj/model"
	"mm-ndj/pkg/contract"
	"mm-ndj/pkg/merkle"
	"mm-ndj/pkg/utils"
	"mm-ndj/server/dao"
	"time"
)

// 定时更新默克尔树
type MerkleTask struct {
	cron   *cron.Cron
	svcCtx *dao.ServiceCtx
}

func NewMerkleTask(svcCtx *dao.ServiceCtx) *MerkleTask {
	merkleTask := &MerkleTask{
		svcCtx: svcCtx,
	}
	//merkleTask.initCron()
	//merkleTask.Run()
	return merkleTask
}

//func (m *MerkleTask) initCron() {
//	m.cron = cron.New()
//	m.cron.Start()
//}
//
//func (m *MerkleTask) addTask() {
//	m.cron.AddFunc("* * * * * *", func() {
//
//	})
//}

func (m *MerkleTask) Run() ([]merkle.Claim, error) {
	allPrize, err := dao.GetAllTokenPrizeList(context.Background(), m.svcCtx)
	if err != nil {
		return nil, err
	}
	//计算默克尔树
	balances := []merkle.Balance{}
	for _, v := range allPrize {
		userInfo, err := dao.GetUserById(context.Background(), m.svcCtx, v.UserId)
		if err != nil {
			return nil, err
		}
		prize, err := config.GetPrizeConfigById(v.PrizeId)
		if err != nil {
			return nil, err
		}
		balances = append(balances, merkle.Balance{
			Account:       common.HexToAddress(userInfo.Address),
			TokenContract: common.HexToAddress(prize.ContractAddress),
			Amount:        utils.ToWei(v.Amount),
			Id:            int(v.Id),
		})
	}

	info, err := merkle.ParseBalanceMap(balances)
	if err != nil {
		return nil, err
	}
	//更新数据
	go func() {
		err = m.svcCtx.Db.Transaction(func(tx *gorm.DB) error {
			model := model2.NewPrizeRecordModel(context.Background(), m.svcCtx.Db)
			for _, v := range info.Claims {
				err := model.UpdateWithMapByCondition(map[string]interface{}{
					"proof":       encodeToJson(v.Proof),
					"update_time": time.Now().Unix(),
				}, func(db *gorm.DB) *gorm.DB {
					return db.Where("id = ?", v.Id)
				})
				if err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			//提交事务失败
			return
		}
	}()
	//调用合约，更新默克尔树
	err = m.UpdateMerkle(info.MerkleRoot)
	return info.Claims, err
}

func encodeToJson(data []common.Hash) string {
	bytes, _ := json.Marshal(data)
	return string(bytes)
}

// 更新默克尔树
func (m *MerkleTask) UpdateMerkle(merkleRoot common.Hash) error {
	client, err := ethclient.Dial(config.Conf.RangersC.RPC)
	if err != nil {
		fmt.Println("Failed to Dial ", err)
		return err
	}
	contractAddress := config.Conf.RangersC.MM
	privateKey, err := crypto.HexToECDSA(config.Conf.RangersC.PrivateKey)
	if err != nil {
		fmt.Println("Failed to Dial ", err)
		return err
	}
	mm, err := contract.NewMM(common.HexToAddress(contractAddress), client)
	if err != nil {
		fmt.Println("Failed to Dial ", err)
		return err
	}

	opt, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(config.Conf.RangersC.ChainId))
	_, err = mm.SetMerkleRoot(opt, merkleRoot)
	if err != nil {
		fmt.Println("SetMerkleRoot failed ", err)
		return err
	}
	return nil
}
