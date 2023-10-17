package model

import (
	"context"
	"errors"
	"mm-ndj/pkg/errcode"
	"mm-ndj/pkg/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserAssetModel = interface {
	Insert(data *UserAsset) error
	InsertWithClauses(data *UserAsset, clauses ...clause.Expression) error
	InsertBatch(dataset []*UserAsset) error
	InsertBatchWithClauses(dataset []*UserAsset, clauses ...clause.Expression) (int64, error)
	Update(data *UserAsset, selects ...string) error
	UpdateByCondition(data *UserAsset, condition func(*gorm.DB) *gorm.DB) error
	UpdateWithMapByCondition(data map[string]interface{}, condition func(*gorm.DB) *gorm.DB) error
	FindOne(id int64) (*UserAsset, error)
	FindOneByCondition(condition func(*gorm.DB) *gorm.DB) (*UserAsset, error)
	FindByCondition(condition func(*gorm.DB) *gorm.DB) ([]*UserAsset, error)
	FindMapByCondition(db *gorm.DB) ([]map[string]interface{}, error)
	Paginate(pq *model.PaginationQuery) ([]*UserAsset, int64, error)
	PluckInt64s(column string, condition func(*gorm.DB) *gorm.DB) ([]int64, error)
	PluckStrings(column string, condition func(*gorm.DB) *gorm.DB) ([]string, error)
	CountByCondition(condition func(*gorm.DB) *gorm.DB) (int64, error)
	DeleteByCondition(condition func(*gorm.DB) *gorm.DB) error
	Init() error
}

type UserAsset struct {
	Id         int64   `json:"id" gorm:"primaryKey;column:id;"`
	UserId     int64   `json:"user_id" gorm:"column:user_id;int"`
	Balance    float64 `json:"balance" gorm:"column:balance;decimal(22,4)"`
	Type       int     `json:"type" gorm:"column:type;int"`
	CreateTime int64   `json:"create_time" gorm:"column:create_time;int"`
	UpdateTime int64   `json:"update_time" gorm:"column:update_time;int"`
	Freeze     float64 `json:"freeze" gorm:"column:freeze;decimal(22,4)"`
}

// TableName 返回用户表信息的数据库表名
func (o *UserAsset) TableName() string {
	return "user_asset"
}

// GetId 获取id
func (o *UserAsset) GetId() int64 {
	return o.Id
}

// NewMintModel 新建用户表模型
func NewUserAssetModel(ctx context.Context, db *gorm.DB) UserAssetModel {
	return &defaultUserAssetModel{
		BaseModel: model.NewBaseModel(ctx, db),
	}
}

// defaultMintModel 默认用户表模型
type defaultUserAssetModel struct {
	*model.BaseModel
}

// Insert 插入UserAsset表信息
func (o *defaultUserAssetModel) Insert(data *UserAsset) error {
	return o.DB.Create(data).Error
}

// InsertWithClauses 使用子句插入UserAsset表信息
func (o *defaultUserAssetModel) InsertWithClauses(data *UserAsset, clauses ...clause.Expression) error {
	return o.DB.Clauses(clauses...).Create(data).Error
}

// InsertBatch 批量插入UserAsset表信息
func (o *defaultUserAssetModel) InsertBatch(dataset []*UserAsset) error {
	return o.DB.Create(&dataset).Error
}

// InsertBatchWithClauses 使用子句批量插入UserAsset表信息
func (o *defaultUserAssetModel) InsertBatchWithClauses(dataset []*UserAsset, clauses ...clause.Expression) (int64, error) {
	db := o.DB.Clauses(clauses...).Create(&dataset)

	return db.RowsAffected, db.Error
}

// Update 更新mint表信息
func (o *defaultUserAssetModel) Update(data *UserAsset, selects ...string) error {
	return o.BaseModel.Update(data, selects)
}

// UpdateByCondition 通过动态条件更新mint表信息
func (o *defaultUserAssetModel) UpdateByCondition(data *UserAsset, condition func(*gorm.DB) *gorm.DB) error {
	return o.BaseModel.UpdateByCondition(data, condition)
}

// UpdateWithMapByCondition 使用map通过动态条件更新UserAsset表信息
func (o *defaultUserAssetModel) UpdateWithMapByCondition(data map[string]interface{}, condition func(*gorm.DB) *gorm.DB) error {
	return o.BaseModel.UpdateWithMapByCondition(&UserAsset{}, data, condition)
}

// FindOne 通过用户表id查找一个UserAsset表信息
func (m *defaultUserAssetModel) FindOne(id int64) (*UserAsset, error) {
	if id < 1 {
		return nil, errcode.ErrInvalidParams
	}

	var o UserAsset
	err := m.DB.Where("`id` = ?", id).First(&o).Error
	if err != nil {
		if model.IsRecordNotFound(err) {
			return nil, errors.New("UserAsset not exist")
		}
		return nil, err
	}

	return &o, nil
}

// FindOneByCondition 通过动态条件查找一个UserAsset表信息
func (m *defaultUserAssetModel) FindOneByCondition(condition func(*gorm.DB) *gorm.DB) (*UserAsset, error) {
	var o UserAsset
	err := m.DB.Scopes(condition).First(&o).Error
	if err != nil {
		if model.IsRecordNotFound(err) {
			return nil, errors.New("UserAsset not exist")
		}
		return nil, err
	}

	return &o, nil
}

// FindByCondition 通过动态条件查找UserAsset表信息
func (o *defaultUserAssetModel) FindByCondition(condition func(*gorm.DB) *gorm.DB) ([]*UserAsset, error) {
	var op []*UserAsset
	err := o.DB.Scopes(condition).Find(&op).Error
	if err != nil {
		return nil, err
	}

	return op, nil
}

func (o *defaultUserAssetModel) FindMapByCondition(db *gorm.DB) ([]map[string]interface{}, error) {
	var op []map[string]interface{}
	err := db.Find(&op).Error
	if err != nil {
		return op, err
	}
	return op, nil
}

// Paginate 分页查找UserAsset表信息
func (o *defaultUserAssetModel) Paginate(pq *model.PaginationQuery) ([]*UserAsset, int64, error) {
	var total int64
	err := o.DB.Scopes(pq.Queries()).Model(&UserAsset{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	var op []*UserAsset
	err = o.DB.Scopes(pq.Paginate()).Find(&op).Error
	if err != nil {
		return nil, 0, err
	}

	return op, total, nil
}

// PluckInt64s 通过动态条件查找单个列并将结果扫描到int64切片中
func (o *defaultUserAssetModel) PluckInt64s(column string, condition func(*gorm.DB) *gorm.DB) ([]int64, error) {
	return o.BaseModel.PluckInt64s(&UserAsset{}, column, condition)
}

// PluckStrings 通过动态条件查找单个列并将结果扫描到string切片中
func (o *defaultUserAssetModel) PluckStrings(column string, condition func(*gorm.DB) *gorm.DB) ([]string, error) {
	return o.BaseModel.PluckStrings(&UserAsset{}, column, condition)
}

// CountByCondition 通过动态条件计数用户表信息
func (o *defaultUserAssetModel) CountByCondition(condition func(*gorm.DB) *gorm.DB) (int64, error) {
	return o.BaseModel.CountByCondition(&UserAsset{}, condition)
}

// DeleteByCondition 通过动态条件删除用户表信息
func (o *defaultUserAssetModel) DeleteByCondition(condition func(*gorm.DB) *gorm.DB) error {
	return o.BaseModel.DeleteByCondition(&UserAsset{}, condition)
}

// Init 初始化UserAsset表信息
func (o *defaultUserAssetModel) Init() error {
	var num int64
	err := o.DB.Model(&UserAsset{}).Count(&num).Error
	if err != nil {
		return err
	}

	if num == 0 {
		// 执行初始化方法
		return nil
	}
	return nil
}
