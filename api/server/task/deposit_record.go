package task

import (
	"context"
	"mm-ndj/server/dao"
	"time"
)

// 更新存款记录
func UpdateDepositRecordStatus(ctx context.Context, svc *dao.ServiceCtx, address string, hash string, amount float64) error {
	_, err := dao.GetUserByAddress(ctx, svc, address)
	if err != nil {
		//用户地址不存在
		return err
	}
	//获取质押记录
	depositRecord, err := dao.GetDepositRecordByHash(ctx, svc, hash)
	if err != nil {
		//质押记录不存在
		return err
	}
	depositRecord.DepositAmount = amount
	depositRecord.UpdateTime = time.Now().Unix()
	//更新质押记录
	err = dao.UpdateDepositRecordByHash(ctx, svc, depositRecord, hash)
	if err != nil {
		//更新状态失败
		return err
	}
	return nil
}
