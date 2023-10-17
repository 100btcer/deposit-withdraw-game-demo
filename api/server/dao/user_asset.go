package dao

import (
	"context"
	"gorm.io/gorm"
	"mm-ndj/constant"
	"mm-ndj/model"
	"time"
)

func GetUserAssetByType(ctx context.Context, svc *ServiceCtx, userId int64, assetType int) (*model.UserAsset, error) {
	m := model.NewUserAssetModel(ctx, svc.Db)
	return m.FindOneByCondition(func(d *gorm.DB) *gorm.DB {
		return d.Where("user_id=? and type = ?", userId, assetType)
	})
}

// 生成用户资产记录
func InsertUserAsset(ctx context.Context, tx *gorm.DB, userId int64) error {
	amodel := model.NewUserAssetModel(ctx, tx)
	arg := &model.UserAsset{
		UserId:     userId,
		Type:       constant.AssetTypeTicket,
		Balance:    0,
		Freeze:     0,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	return amodel.Insert(arg)
}
