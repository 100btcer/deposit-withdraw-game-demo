package task

import (
	"context"
	"mm-ndj/server/dao"
	"time"
)

// 更新赎回记录
func UpdateWithdrawRecordStatus(ctx context.Context, svc *dao.ServiceCtx, address string, hash string, amount float64) error {
	_, err := dao.GetUserByAddress(ctx, svc, address)
	if err != nil {
		//用户地址不存在
		return err
	}
	//获取提现记录
	withdrawRecord, err := dao.GetWithdrawRecordByHash(ctx, svc, hash)
	if err != nil {
		//提现记录不存在
		return err
	}
	withdrawRecord.Amount = amount
	withdrawRecord.UpdateTime = time.Now().Unix()
	//更新提现记录
	err = dao.UpdateWithdrawRecordByHash(ctx, svc, withdrawRecord, hash)
	if err != nil {
		//更新状态失败
		return err
	}
	return nil
}
