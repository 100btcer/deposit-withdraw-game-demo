package dao

import (
	"context"
	"gorm.io/gorm"
	"mm-ndj/model"
)

// 获取奖品
//func GetPrizeById(ctx context.Context, svc *ServiceCtx, id int) (*model.PrizeRecord, error) {
//	m := model.NewPrizeRecordModel(ctx, svc.Db)
//	return m.FindOneByCondition(func(db *gorm.DB) *gorm.DB {
//		return db.Where("id = ?", id)
//	})
//}

// 根据hash获取开奖结果
func GetPrizeByUserIdAndBurnHash(ctx context.Context, svc *ServiceCtx, userId int64, hash string) ([]*model.PrizeRecord, error) {
	m := model.NewPrizeRecordModel(ctx, svc.Db)
	return m.FindByCondition(func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ? and burn_hash = ?", userId, hash)
	})
}

// 查询所有token类型中奖列表
func GetAllTokenPrizeList(ctx context.Context, svc *ServiceCtx) ([]*model.PrizeRecord, error) {
	m := model.NewPrizeRecordModel(ctx, svc.Db)
	return m.FindByCondition(func(db *gorm.DB) *gorm.DB {
		return db.Where("prize_type = ?", 1)
	})
}

// 记录用户奖品
//func AddUserPrize(ctx context.Context, svc *ServiceCtx, userId int64, prizeId int, amount int, userInfo *model.User) error {
//	return svc.Db.Transaction(func(tx *gorm.DB) error {
//		//记录奖品
//		prizeM := model.NewPrizeRecordModel(ctx, svc.Db)
//		err := prizeM.Insert(&model.PrizeRecord{
//			UserId:     userId,
//			PrizeId:    prizeId,
//			Amount:     amount,
//			Status:     0,
//			UpdateTime: time.Now().Unix(),
//			CreateTime: time.Now().Unix(),
//		})
//		if err != nil {
//			return err
//		}
//		//更新用户领奖次数
//		userM := model.NewUserModel(ctx, svc.Db)
//		userInfo.LotteryNum -= 1
//		userInfo.UpdateTime = time.Now().Unix()
//		version := userInfo.Version
//		userInfo.Version += 1
//		err = userM.UpdateByCondition(userInfo, func(db *gorm.DB) *gorm.DB {
//			return db.Where("id = ? and version = ?", userId, version)
//		})
//		if err != nil {
//			return err
//		}
//		return nil
//	})
//}

// 获取用户中奖次数
func GetUserPrizeCount(ctx context.Context, svc *ServiceCtx, userId int64, prizeId int) (int64, error) {
	type Amount struct {
		Amount float64
	}
	sql := `select sum(amount) amount from prize_record where user_id = ? and prize_id = ?`
	var amount Amount
	err := svc.Db.Raw(sql, userId, prizeId).Scan(&amount).Error
	return int64(amount.Amount), err
}

// 获取奖品总领取次数
func GetPrizeCount(ctx context.Context, svc *ServiceCtx, prizeId int) (int64, error) {
	type Amount struct {
		Amount float64
	}
	sql := `select sum(amount) amount from prize_record where prize_id = ?`
	var amount Amount
	err := svc.Db.Raw(sql, prizeId).Scan(&amount).Error
	return int64(amount.Amount), err
}
