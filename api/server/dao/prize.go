package dao

import (
	"context"
	"gorm.io/gorm"
	"mm-ndj/constant"
	"mm-ndj/model"
	model2 "mm-ndj/pkg/model"
)

// 获取奖品
func GetPrizeById(ctx context.Context, svc *ServiceCtx, id int) (*model.Prize, error) {
	m := model.NewPrizeModel(ctx, svc.Db)
	return m.FindOneByCondition(func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	})
}

// 根据tokenid获取奖品
func GetPrizeByTokenId(ctx context.Context, svc *ServiceCtx, userId int64, tokenId int, status int) (*model.Prize, error) {
	m := model.NewPrizeModel(ctx, svc.Db)
	return m.FindOneByCondition(func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ? and token_id = ? and status = ?", userId, tokenId, status)
	})
}

// 查询所有未领取的奖励
func GetPrizeByPrizeIdAndStatus(ctx context.Context, svc *ServiceCtx, prizeId int, status int) ([]*model.Prize, error) {
	m := model.NewPrizeModel(ctx, svc.Db)
	return m.FindByCondition(func(db *gorm.DB) *gorm.DB {
		return db.Where("prize_id = ? and status = ?", prizeId, status)
	})
}

// 获取奖品记录
func GetLotteryPrizeList(ctx context.Context, svc *ServiceCtx, userId int64, ty *int, page, pageSize int64) ([]*model.Prize, int64, error) {
	m := model.NewPrizeModel(ctx, svc.Db)
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
	pg.Add(func(db *gorm.DB) *gorm.DB {
		return db.Order("sort asc")
	})
	return m.Paginate(pg)
}
