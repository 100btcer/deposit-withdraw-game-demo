package dao

import (
	"context"
	"mm-ndj/constant"
	"mm-ndj/model"
	"time"
)

// 保存销毁记录
func BurnRecordAdd(ctx context.Context, svc *ServiceCtx, userId int64, hash string, amount float64) error {
	m := model.NewBurnRecordModel(ctx, svc.Db)
	return m.Insert(&model.BurnRecord{
		UserId:     userId,
		Amount:     amount,
		Hash:       hash,
		Status:     constant.TradeStatus1,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	})
}
