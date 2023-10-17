package dao

import (
	"context"
	"gorm.io/gorm"
	"mm-ndj/config"
	"mm-ndj/constant"
	"mm-ndj/model"
	model2 "mm-ndj/pkg/model"
	"time"
)

// 添加质押记录
func DepositRecordAdd(ctx context.Context, svc *ServiceCtx, userId int64, hash string, amount float64, poolId int) error {
	m := model.NewDepositRecordModel(ctx, svc.Db)
	lockDay := config.PoolConfig[poolId].Days
	return m.Insert(&model.DepositRecord{
		UserId:        userId,
		CreateTime:    time.Now().Unix(),
		LockDay:       lockDay,
		DepositAmount: amount,
		Hash:          hash,
		ConfirmNum:    0,
		Status:        constant.TradeStatus1,
	})
}

// 获取质押记录
func DepositRecordList(ctx context.Context, svc *ServiceCtx, userId int64, page, pageSize int64) ([]*model.DepositRecord, int64, error) {
	m := model.NewDepositRecordModel(ctx, svc.Db)
	pg := model2.NewPaginationQuery(page, pageSize, "id desc")
	pg.Add(func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", userId)
	})
	return m.Paginate(pg)
}

// 根据hash获取质押记录
func GetDepositRecordByHash(ctx context.Context, svc *ServiceCtx, hash string) (*model.DepositRecord, error) {
	m := model.NewDepositRecordModel(ctx, svc.Db)
	return m.FindOneByCondition(func(db *gorm.DB) *gorm.DB {
		return db.Where("hash = ?", hash)
	})
}

// 更新质押记录
func UpdateDepositRecordByHash(ctx context.Context, svc *ServiceCtx, record *model.DepositRecord, hash string) error {
	m := model.NewDepositRecordModel(ctx, svc.Db)
	return m.UpdateByCondition(record, func(db *gorm.DB) *gorm.DB {
		return db.Where("hash = ?", hash)
	})
}

// 获取质押总额
func GetDepositSum(ctx context.Context, svc *ServiceCtx, userId int64) (float64, error) {
	var res []float64
	var totalAmount float64
	sql := "select deposit_amount from deposit_record where user_id = ?"
	err := svc.Db.Raw(sql, userId).Pluck("depost_amount", &res).Error
	if err != nil {
		return 0, err
	}
	for _, v := range res {
		totalAmount += v
	}
	return totalAmount, err
}

// 获取奖品记录
func GetLotteryPrizeRecordList(ctx context.Context, svc *ServiceCtx, userId int64, ty *int, page, pageSize int64) ([]*model.PrizeRecord, int64, error) {
	m := model.NewPrizeRecordModel(ctx, svc.Db)
	pg := model2.NewPaginationQuery(page, pageSize, "id desc,prize_id asc")
	pg.Add(func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", userId)
	})
	if ty != nil {
		pg.Add(func(db *gorm.DB) *gorm.DB {
			return db.Where("status = ?", constant.Status0)
		})
	}
	if ty != nil {
		pg.Add(func(db *gorm.DB) *gorm.DB {
			return db.Where("prize_type = ?", &ty)
		})
	}
	return m.Paginate(pg)
}
