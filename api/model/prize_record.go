package model

import (
	"context"
	"errors"
	"mm-ndj/pkg/errcode"
	"mm-ndj/pkg/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PrizeRecordModel = interface {
	Insert(data *PrizeRecord) error
	InsertWithClauses(data *PrizeRecord, clauses ...clause.Expression) error
	InsertBatch(dataset []*PrizeRecord) error
	InsertBatchWithClauses(dataset []*PrizeRecord, clauses ...clause.Expression) (int64, error)
	Update(data *PrizeRecord, selects ...string) error
	UpdateByCondition(data *PrizeRecord, condition func(*gorm.DB) *gorm.DB) error
	UpdateWithMapByCondition(data map[string]interface{}, condition func(*gorm.DB) *gorm.DB) error
	FindOne(id int64) (*PrizeRecord, error)
	FindOneByCondition(condition func(*gorm.DB) *gorm.DB) (*PrizeRecord, error)
	FindByCondition(condition func(*gorm.DB) *gorm.DB) ([]*PrizeRecord, error)
	FindMapByCondition(db *gorm.DB) ([]map[string]interface{}, error)
	Paginate(pq *model.PaginationQuery) ([]*PrizeRecord, int64, error)
	PluckInt64s(column string, condition func(*gorm.DB) *gorm.DB) ([]int64, error)
	PluckStrings(column string, condition func(*gorm.DB) *gorm.DB) ([]string, error)
	CountByCondition(condition func(*gorm.DB) *gorm.DB) (int64, error)
	DeleteByCondition(condition func(*gorm.DB) *gorm.DB) error
	Init() error
}

type PrizeRecord struct {
	Id          int64   `json:"id" gorm:"primaryKey;column:id;"`
	UserId      int64   `json:"user_id" gorm:"column:user_id;int"`
	PrizeId     int     `json:"prize_id" gorm:"column:prize_id;int"`
	Amount      float64 `json:"amount" gorm:"column:amount;int"`
	Status      int8    `json:"status" gorm:"column:status;tinyint"`
	UpdateTime  int64   `json:"update_time" gorm:"column:update_time;int"`
	CreateTime  int64   `json:"create_time" gorm:"column:create_time;int"`
	BurnHash    string  `json:"burn_hash" gorm:"column:burn_hash;varchar(128)"`
	PrizeType   int     `json:"prize_type" gorm:"column:prize_type;int"`
	Proof       string  `json:"proof" gorm:"column:proof;text"`
	ReceiveHash string  `json:"receive_hash" gorm:"column:receive_hash;varchar(128)"`
}

// TableName 返回用户表信息的数据库表名
func (o *PrizeRecord) TableName() string {
	return "prize_record"
}

// GetId 获取id
func (o *PrizeRecord) GetId() int64 {
	return o.Id
}

// NewMintModel 新建用户表模型
func NewPrizeRecordModel(ctx context.Context, db *gorm.DB) PrizeRecordModel {
	return &defaultPrizeRecordModel{
		BaseModel: model.NewBaseModel(ctx, db),
	}
}

// defaultMintModel 默认用户表模型
type defaultPrizeRecordModel struct {
	*model.BaseModel
}

// Insert 插入PrizeRecord表信息
func (o *defaultPrizeRecordModel) Insert(data *PrizeRecord) error {
	return o.DB.Create(data).Error
}

// InsertWithClauses 使用子句插入PrizeRecord表信息
func (o *defaultPrizeRecordModel) InsertWithClauses(data *PrizeRecord, clauses ...clause.Expression) error {
	return o.DB.Clauses(clauses...).Create(data).Error
}

// InsertBatch 批量插入PrizeRecord表信息
func (o *defaultPrizeRecordModel) InsertBatch(dataset []*PrizeRecord) error {
	return o.DB.Create(&dataset).Error
}

// InsertBatchWithClauses 使用子句批量插入PrizeRecord表信息
func (o *defaultPrizeRecordModel) InsertBatchWithClauses(dataset []*PrizeRecord, clauses ...clause.Expression) (int64, error) {
	db := o.DB.Clauses(clauses...).Create(&dataset)

	return db.RowsAffected, db.Error
}

// Update 更新mint表信息
func (o *defaultPrizeRecordModel) Update(data *PrizeRecord, selects ...string) error {
	return o.BaseModel.Update(data, selects)
}

// UpdateByCondition 通过动态条件更新mint表信息
func (o *defaultPrizeRecordModel) UpdateByCondition(data *PrizeRecord, condition func(*gorm.DB) *gorm.DB) error {
	return o.BaseModel.UpdateByCondition(data, condition)
}

// UpdateWithMapByCondition 使用map通过动态条件更新PrizeRecord表信息
func (o *defaultPrizeRecordModel) UpdateWithMapByCondition(data map[string]interface{}, condition func(*gorm.DB) *gorm.DB) error {
	return o.BaseModel.UpdateWithMapByCondition(&PrizeRecord{}, data, condition)
}

// FindOne 通过用户表id查找一个PrizeRecord表信息
func (m *defaultPrizeRecordModel) FindOne(id int64) (*PrizeRecord, error) {
	if id < 1 {
		return nil, errcode.ErrInvalidParams
	}

	var o PrizeRecord
	err := m.DB.Where("`id` = ?", id).First(&o).Error
	if err != nil {
		if model.IsRecordNotFound(err) {
			return nil, errors.New("PrizeRecord not exist")
		}
		return nil, err
	}

	return &o, nil
}

// FindOneByCondition 通过动态条件查找一个PrizeRecord表信息
func (m *defaultPrizeRecordModel) FindOneByCondition(condition func(*gorm.DB) *gorm.DB) (*PrizeRecord, error) {
	var o PrizeRecord
	err := m.DB.Scopes(condition).First(&o).Error
	if err != nil {
		if model.IsRecordNotFound(err) {
			return nil, errors.New("PrizeRecord not exist")
		}
		return nil, err
	}

	return &o, nil
}

// FindByCondition 通过动态条件查找PrizeRecord表信息
func (o *defaultPrizeRecordModel) FindByCondition(condition func(*gorm.DB) *gorm.DB) ([]*PrizeRecord, error) {
	var op []*PrizeRecord
	err := o.DB.Scopes(condition).Find(&op).Error
	if err != nil {
		return nil, err
	}

	return op, nil
}

func (o *defaultPrizeRecordModel) FindMapByCondition(db *gorm.DB) ([]map[string]interface{}, error) {
	var op []map[string]interface{}
	err := db.Find(&op).Error
	if err != nil {
		return op, err
	}
	return op, nil
}

// Paginate 分页查找PrizeRecord表信息
func (o *defaultPrizeRecordModel) Paginate(pq *model.PaginationQuery) ([]*PrizeRecord, int64, error) {
	var total int64
	err := o.DB.Scopes(pq.Queries()).Model(&PrizeRecord{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	var op []*PrizeRecord
	err = o.DB.Scopes(pq.Paginate()).Find(&op).Error
	if err != nil {
		return nil, 0, err
	}

	return op, total, nil
}

// PluckInt64s 通过动态条件查找单个列并将结果扫描到int64切片中
func (o *defaultPrizeRecordModel) PluckInt64s(column string, condition func(*gorm.DB) *gorm.DB) ([]int64, error) {
	return o.BaseModel.PluckInt64s(&PrizeRecord{}, column, condition)
}

// PluckStrings 通过动态条件查找单个列并将结果扫描到string切片中
func (o *defaultPrizeRecordModel) PluckStrings(column string, condition func(*gorm.DB) *gorm.DB) ([]string, error) {
	return o.BaseModel.PluckStrings(&PrizeRecord{}, column, condition)
}

// CountByCondition 通过动态条件计数用户表信息
func (o *defaultPrizeRecordModel) CountByCondition(condition func(*gorm.DB) *gorm.DB) (int64, error) {
	return o.BaseModel.CountByCondition(&PrizeRecord{}, condition)
}

// DeleteByCondition 通过动态条件删除用户表信息
func (o *defaultPrizeRecordModel) DeleteByCondition(condition func(*gorm.DB) *gorm.DB) error {
	return o.BaseModel.DeleteByCondition(&PrizeRecord{}, condition)
}

// Init 初始化PrizeRecord表信息
func (o *defaultPrizeRecordModel) Init() error {
	var num int64
	err := o.DB.Model(&PrizeRecord{}).Count(&num).Error
	if err != nil {
		return err
	}

	if num == 0 {
		// 执行初始化方法
		return nil
	}
	return nil
}
