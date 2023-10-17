package user

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/storyicon/sigverify"
	"mm-ndj/api/param/types"
	"mm-ndj/config"
	"mm-ndj/constant"
	"mm-ndj/server/dao"

	"mm-ndj/pkg/token"

	"go.uber.org/zap"
)

type SendEmailData struct {
	Ctx   context.Context
	Svc   *dao.ServiceCtx
	Email string
	Ty    int
}

func WalletLogin(ctx context.Context, svc *dao.ServiceCtx, req types.WalletLoginReq) (*types.RegisterResp, error) {
	valid, err := sigverify.VerifyEllipticCurveHexSignatureEx(
		common.HexToAddress(req.Address),
		[]byte(req.Message),
		req.Sign,
	)
	if !valid || err != nil {
		config.Logger.Error("WalletLogin valid err", zap.Error(err))
		return nil, err
	}

	//var err error
	//验证用户是否存在
	usr, _ := dao.GetUserByAddress(ctx, svc, req.Address)
	userId := int64(0)
	if usr == nil {
		if err != nil {
			config.Logger.Error("Login GenID", zap.Error(err))
			return nil, err
		}
		//插入数据
		if userId, err = dao.InsertUserData(ctx, svc, req.Address); err != nil {
			return nil, err
		}

		//创建token
		tokenStr, err := token.CreateToken(svc, userId, req.Address, req.Address)
		if err != nil {
			return nil, err
		}

		return &types.RegisterResp{
			Token:  tokenStr,
			UserId: userId,
		}, nil
	}
	userId = usr.Id
	tokenStr, err := token.CreateToken(svc, usr.Id, req.Address, fmt.Sprint(req.Address))
	if err != nil {
		return nil, err
	}

	return &types.RegisterResp{
		Token:  tokenStr,
		UserId: userId,
	}, nil
}

// 用户信息
func GetUserInfo(ctx context.Context, svc *dao.ServiceCtx, userId int64) (*types.UserInfoResp, error) {
	userInfo, err := dao.GetUserById(ctx, svc, userId)
	if err != nil {
		return nil, err
	}
	userAsset, err := dao.GetUserAssetByType(ctx, svc, userId, constant.AssetTypeTicket)
	if err != nil {
		return nil, err
	}
	depositAmount, err := dao.GetDepositSum(ctx, svc, userId)
	if err != nil {
		return nil, err
	}
	withdrawAmount, err := dao.GetWithdrawSum(ctx, svc, userId)
	if err != nil {
		return nil, err
	}
	return &types.UserInfoResp{
		UserId:        userId,
		Address:       userInfo.Address,
		Ticket:        userAsset.Balance,
		DepositAmount: depositAmount - withdrawAmount,
	}, nil
}
