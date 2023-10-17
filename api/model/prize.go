package model

import (
	"context"
	"errors"
	"mm-ndj/pkg/errcode"
	"mm-ndj/pkg/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PrizeModel = interface {
	Insert(data *Prize) error
	InsertWithClauses(data *Prize, clauses ...clause.Expression) error
	InsertBatch(dataset []*Prize) error
	InsertBatchWithClauses(dataset []*Prize, clauses ...clause.Expression) (int64, error)
	Update(data *Prize, selects ...string) error
	UpdateByCondition(data *Prize, condition func(*gorm.DB) *gorm.DB) error
	UpdateWithMapByCondition(data map[string]interface{}, condition func(*gorm.DB) *gorm.DB) error
	FindOne(id int64) (*Prize, error)
	FindOneByCondition(condition func(*gorm.DB) *gorm.DB) (*Prize, error)
	FindByCondition(condition func(*gorm.DB) *gorm.DB) ([]*Prize, error)
	FindMapByCondition(db *gorm.DB) ([]map[string]interface{}, error)
	Paginate(pq *model.PaginationQuery) ([]*Prize, int64, error)
	PluckInt64s(column string, condition func(*gorm.DB) *gorm.DB) ([]int64, error)
	PluckStrings(column string, condition func(*gorm.DB) *gorm.DB) ([]string, error)
	CountByCondition(condition func(*gorm.DB) *gorm.DB) (int64, error)
	DeleteByCondition(condition func(*gorm.DB) *gorm.DB) error
	Init() error
}

type Prize struct {
	Id          int64   `json:"id" gorm:"primaryKey;column:id;"`
	UserId      int64   `json:"user_id" gorm:"column:user_id;int"`
	PrizeId     int     `json:"prize_id" gorm:"column:prize_id;int"`
	TokenId     int     `json:"token_id" gorm:"column:token_id;int"`
	Amount      float64 `json:"amount" gorm:"column:amount;decimal(22,4)"`
	Status      int8    `json:"status" gorm:"column:status;tinyint"`
	UpdateTime  int64   `json:"update_time" gorm:"column:update_time;int"`
	CreateTime  int64   `json:"create_time" gorm:"column:create_time;int"`
	PrizeType   int     `json:"prize_type" gorm:"column:prize_type;int"`
	ReceiveHash string  `json:"receive_hash" gorm:"column:receive_hash;varchar(128)"`
	Sort        int     `json:"sort" gorm:"column:sort;int"`
}

// TableName 返回用户表信息的数据库表名
func (o *Prize) TableName() string {
	return "prize"
}

// GetId 获取id
func (o *Prize) GetId() int64 {
	return o.Id
}

// NewMintModel 新建用户表模型
func NewPrizeModel(ctx context.Context, db *gorm.DB) PrizeModel {
	return &defaultPrizeModel{
		BaseModel: model.NewBaseModel(ctx, db),
	}
}

// defaultMintModel 默认用户表模型
type defaultPrizeModel struct {
	*model.BaseModel
}

// Insert 插入Prize表信息
func (o *defaultPrizeModel) Insert(data *Prize) error {
	return o.DB.Create(data).Error
}

// InsertWithClauses 使用子句插入Prize表信息
func (o *defaultPrizeModel) InsertWithClauses(data *Prize, clauses ...clause.Expression) error {
	return o.DB.Clauses(clauses...).Create(data).Error
}

// InsertBatch 批量插入Prize表信息
func (o *defaultPrizeModel) InsertBatch(dataset []*Prize) error {
	return o.DB.Create(&dataset).Error
}

// InsertBatchWithClauses 使用子句批量插入Prize表信息
func (o *defaultPrizeModel) InsertBatchWithClauses(dataset []*Prize, clauses ...clause.Expression) (int64, error) {
	db := o.DB.Clauses(clauses...).Create(&dataset)

	return db.RowsAffected, db.Error
}

// Update 更新mint表信息
func (o *defaultPrizeModel) Update(data *Prize, selects ...string) error {
	return o.BaseModel.Update(data, selects)
}

// UpdateByCondition 通过动态条件更新mint表信息
func (o *defaultPrizeModel) UpdateByCondition(data *Prize, condition func(*gorm.DB) *gorm.DB) error {
	return o.BaseModel.UpdateByCondition(data, condition)
}

// UpdateWithMapByCondition 使用map通过动态条件更新Prize表信息
func (o *defaultPrizeModel) UpdateWithMapByCondition(data map[string]interface{}, condition func(*gorm.DB) *gorm.DB) error {
	return o.BaseModel.UpdateWithMapByCondition(&Prize{}, data, condition)
}

// FindOne 通过用户表id查找一个Prize表信息
func (m *defaultPrizeModel) FindOne(id int64) (*Prize, error) {
	if id < 1 {
		return nil, errcode.ErrInvalidParams
	}

	var o Prize
	err := m.DB.Where("`id` = ?", id).First(&o).Error
	if err != nil {
		if model.IsRecordNotFound(err) {
			return nil, errors.New("Prize not exist")
		}
		return nil, err
	}

	return &o, nil
}

// FindOneByCondition 通过动态条件查找一个Prize表信息
func (m *defaultPrizeModel) FindOneByCondition(condition func(*gorm.DB) *gorm.DB) (*Prize, error) {
	var o Prize
	err := m.DB.Scopes(condition).First(&o).Error
	if err != nil {
		if model.IsRecordNotFound(err) {
			return nil, err
		}
		return nil, err
	}

	return &o, nil
}

// FindByCondition 通过动态条件查找Prize表信息
func (o *defaultPrizeModel) FindByCondition(condition func(*gorm.DB) *gorm.DB) ([]*Prize, error) {
	var op []*Prize
	err := o.DB.Scopes(condition).Find(&op).Error
	if err != nil {
		return nil, err
	}

	return op, nil
}

func (o *defaultPrizeModel) FindMapByCondition(db *gorm.DB) ([]map[string]interface{}, error) {
	var op []map[string]interface{}
	err := db.Find(&op).Error
	if err != nil {
		return op, err
	}
	return op, nil
}

// Paginate 分页查找Prize表信息
func (o *defaultPrizeModel) Paginate(pq *model.PaginationQuery) ([]*Prize, int64, error) {
	var total int64
	err := o.DB.Scopes(pq.Queries()).Model(&Prize{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	var op []*Prize
	err = o.DB.Scopes(pq.Paginate()).Find(&op).Error
	if err != nil {
		return nil, 0, err
	}

	return op, total, nil
}

// PluckInt64s 通过动态条件查找单个列并将结果扫描到int64切片中
func (o *defaultPrizeModel) PluckInt64s(column string, condition func(*gorm.DB) *gorm.DB) ([]int64, error) {
	return o.BaseModel.PluckInt64s(&Prize{}, column, condition)
}

// PluckStrings 通过动态条件查找单个列并将结果扫描到string切片中
func (o *defaultPrizeModel) PluckStrings(column string, condition func(*gorm.DB) *gorm.DB) ([]string, error) {
	return o.BaseModel.PluckStrings(&Prize{}, column, condition)
}

// CountByCondition 通过动态条件计数用户表信息
func (o *defaultPrizeModel) CountByCondition(condition func(*gorm.DB) *gorm.DB) (int64, error) {
	return o.BaseModel.CountByCondition(&Prize{}, condition)
}

// DeleteByCondition 通过动态条件删除用户表信息
func (o *defaultPrizeModel) DeleteByCondition(condition func(*gorm.DB) *gorm.DB) error {
	return o.BaseModel.DeleteByCondition(&Prize{}, condition)
}

// Init 初始化Prize表信息
func (o *defaultPrizeModel) Init() error {
	var num int64
	err := o.DB.Model(&Prize{}).Count(&num).Error
	if err != nil {
		return err
	}

	if num == 0 {
		// 执行初始化方法
		return nil
	}
	return nil
}
