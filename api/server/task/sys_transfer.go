package task

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
	"math/big"
	"mm-ndj/config"
	"mm-ndj/pkg/contract"
	"mm-ndj/pkg/utils"
	"mm-ndj/server/dao"
)

func SysTransfer(ctx context.Context, svc *dao.ServiceCtx, tokenContract string, amount float64, id int64) (string, error) {
	client, err := ethclient.Dial(config.Conf.RangersC.RPC)
	if err != nil {
		config.Logger.Error("SysTransfer ethclient.Dial", zap.Error(err))
		return "", err
	}
	contractAddress := config.Conf.RangersC.MM
	privateKey, err := crypto.HexToECDSA(config.Conf.RangersC.PrivateKey)
	if err != nil {
		config.Logger.Error("SysTransfer crypto.HexToECDSA", zap.Error(err))
		return "", err
	}
	mm, err := contract.NewMM(common.HexToAddress(contractAddress), client)
	if err != nil {
		config.Logger.Error("SysTransfer contract.NewMM", zap.Error(err))
		return "", err
	}

	//查询余额
	tokenAmount, err := mm.BalanceOf(nil, common.HexToAddress(tokenContract))
	if err != nil {
		config.Logger.Error("SysTransfer 查询余额", zap.Error(err))
		return "", err
	}
	if tokenAmount.Div(tokenAmount, big.NewInt(1e18)).Int64() < int64(amount) {
		config.Logger.Error("SysTransfer 合约余额不足", zap.Error(err), zap.String("tokenAmount:", tokenAmount.String()), zap.Float64("amount:", amount))
		return "", errors.New("合约余额不足")
	}

	opt, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(config.Conf.RangersC.ChainId))
	res, err := mm.SysTransfer(opt, common.HexToAddress(tokenContract), utils.ToWei(amount), big.NewInt(id))
	if err != nil {
		config.Logger.Info("SysTransfer 发送交易失败", zap.Error(err))
		return "", err
	}
	return res.Hash().String(), nil
}

// 获取可提现
func GetWithdrawAmount(ctx context.Context, svc *dao.ServiceCtx, address string) (string, error) {
	client, err := ethclient.Dial(config.Conf.RangersC.RPC)
	if err != nil {
		config.Logger.Error("SysTransfer ethclient.Dial", zap.Error(err))
		return "", err
	}
	contractAddress := config.Conf.RangersC.MM
	mm, err := contract.NewMM(common.HexToAddress(contractAddress), client)
	if err != nil {
		config.Logger.Error("SysTransfer contract.NewMM", zap.Error(err))
		return "", err
	}

	opt := bind.CallOpts{From: common.HexToAddress(address)}
	res, err := mm.GetValidWithdrawLp(&opt)
	if err != nil {
		config.Logger.Info("SysTransfer 发送交易失败", zap.Error(err))
		return "", err
	}
	return res.Div(res, big.NewInt(1e18)).String(), nil
}
