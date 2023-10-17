package dao

import (
	"context"
	"mm-ndj/model"
	"time"

	"gorm.io/gorm"
)

func GetUserByAddress(ctx context.Context, svc *ServiceCtx, address string) (*model.User, error) {
	uModel := model.NewUserModel(ctx, svc.Db)
	return uModel.FindOneByCondition(func(d *gorm.DB) *gorm.DB {
		return d.Where("address=?", address)
	})
}

func GetUserById(ctx context.Context, svc *ServiceCtx, userId int64) (*model.User, error) {
	uModel := model.NewUserModel(ctx, svc.Db)
	return uModel.FindOneByCondition(func(d *gorm.DB) *gorm.DB {
		return d.Where("id=?", userId)
	})
}

// 插入用户信息
func InsertUserData(ctx context.Context, svc *ServiceCtx, address string) (userId int64, err error) {
	err = svc.Db.Transaction(func(tx *gorm.DB) error {
		//用户数据
		amodel := model.NewUserModel(ctx, tx)
		data := &model.User{
			Address:    address,
			CreateTime: time.Now().Unix(),
		}
		err = amodel.Insert(data)
		if err != nil {
			return err
		}
		userId = data.Id

		//创建用户资产数据
		err = InsertUserAsset(ctx, tx, userId)
		if err != nil {
			return err
		}
		return nil
	})
	return
}
