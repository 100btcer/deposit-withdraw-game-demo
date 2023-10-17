package dao

import (
	"context"
	"gorm.io/gorm"
	"mm-ndj/constant"
	"mm-ndj/model"
	"time"
)

// 添加提款记录
func WithdrawRecordAdd(ctx context.Context, svc *ServiceCtx, userId int64, hash string, amount float64) error {
	m := model.NewWithdrawRecordModel(ctx, svc.Db)
	return m.Insert(&model.WithdrawRecord{
		UserId:     userId,
		Amount:     amount,
		Hash:       hash,
		Status:     constant.TradeStatus1,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	})
}

// 根据hash获取提现记录
func GetWithdrawRecordByHash(ctx context.Context, svc *ServiceCtx, hash string) (*model.WithdrawRecord, error) {
	m := model.NewWithdrawRecordModel(ctx, svc.Db)
	return m.FindOneByCondition(func(db *gorm.DB) *gorm.DB {
		return db.Where("hash = ?", hash)
	})
}

// 更新提现记录
func UpdateWithdrawRecordByHash(ctx context.Context, svc *ServiceCtx, record *model.WithdrawRecord, hash string) error {
	m := model.NewWithdrawRecordModel(ctx, svc.Db)
	return m.UpdateByCondition(record, func(db *gorm.DB) *gorm.DB {
		return db.Where("hash = ?", hash)
	})
}

// 获取提款总额
func GetWithdrawSum(ctx context.Context, svc *ServiceCtx, userId int64) (float64, error) {
	var res []float64
	var totalAmount float64
	sql := "select amount from withdraw_record where user_id = ?"
	err := svc.Db.Raw(sql, userId).Pluck("amount", &res).Error
	if err != nil {
		return 0, err
	}
	for _, v := range res {
		totalAmount += v
	}
	return totalAmount, err
}
