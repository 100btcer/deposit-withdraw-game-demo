package withdraw_record

import (
	"context"
	"mm-ndj/server/dao"
	"mm-ndj/server/task"
)

// 添加提款记录
func WithdrawRecordAdd(ctx context.Context, svc *dao.ServiceCtx, userId int64, hash string, amount float64) error {
	return dao.WithdrawRecordAdd(ctx, svc, userId, hash, amount)
}

type ValidWithdrawAmountResp struct {
	Address string `json:"address"`
	Amount  string `json:"amount"`
}

// 获取可提现金额
func GetValidWithdrawAmount(ctx context.Context, svc *dao.ServiceCtx, userId int64) (*ValidWithdrawAmountResp, error) {
	userInfo, err := dao.GetUserById(ctx, svc, userId)
	if err != nil {
		return nil, err
	}
	amount, err := task.GetWithdrawAmount(ctx, svc, userInfo.Address)
	if err != nil {
		return nil, err
	}
	return &ValidWithdrawAmountResp{
		Address: userInfo.Address,
		Amount:  amount,
	}, err
}
