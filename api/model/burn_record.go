package model

import (
	"context"
	"errors"
	"mm-ndj/pkg/errcode"
	"mm-ndj/pkg/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BurnRecordModel = interface {
	Insert(data *BurnRecord) error
	InsertWithClauses(data *BurnRecord, clauses ...clause.Expression) error
	InsertBatch(dataset []*BurnRecord) error
	InsertBatchWithClauses(dataset []*BurnRecord, clauses ...clause.Expression) (int64, error)
	Update(data *BurnRecord, selects ...string) error
	UpdateByCondition(data *BurnRecord, condition func(*gorm.DB) *gorm.DB) error
	UpdateWithMapByCondition(data map[string]interface{}, condition func(*gorm.DB) *gorm.DB) error
	FindOne(id int64) (*BurnRecord, error)
	FindOneByCondition(condition func(*gorm.DB) *gorm.DB) (*BurnRecord, error)
	FindByCondition(condition func(*gorm.DB) *gorm.DB) ([]*BurnRecord, error)
	FindMapByCondition(db *gorm.DB) ([]map[string]interface{}, error)
	Paginate(pq *model.PaginationQuery) ([]*BurnRecord, int64, error)
	PluckInt64s(column string, condition func(*gorm.DB) *gorm.DB) ([]int64, error)
	PluckStrings(column string, condition func(*gorm.DB) *gorm.DB) ([]string, error)
	CountByCondition(condition func(*gorm.DB) *gorm.DB) (int64, error)
	DeleteByCondition(condition func(*gorm.DB) *gorm.DB) error
	Init() error
}

type BurnRecord struct {
	Id          int64   `json:"id" gorm:"primaryKey;column:id;"`
	UserId      int64   `json:"user_id" gorm:"column:user_id;int"`
	Address     string  `json:"address" gorm:"column:address;varchar(64)"`
	Hash        string  `json:"hash" gorm:"column:hash;varchar(128)"`
	BlockHeight int64   `json:"block_height" gorm:"column:block_height;bigint"`
	TxStatus    int     `json:"tx_status" gorm:"column:tx_status;tinyint"`
	Amount      float64 `json:"amount" gorm:"column:amount;int"`
	Status      int8    `json:"status" gorm:"column:status;tinyint"`
	CreateTime  int64   `json:"create_time" gorm:"column:create_time;int"`
	UpdateTime  int64   `json:"update_time" gorm:"column:update_time;int"`
}

// TableName 返回用户表信息的数据库表名
func (o *BurnRecord) TableName() string {
	return "burn_record"
}

// GetId 获取id
func (o *BurnRecord) GetId() int64 {
	return o.Id
}

// NewMintModel 新建用户表模型
func NewBurnRecordModel(ctx context.Context, db *gorm.DB) BurnRecordModel {
	return &defaultBurnRecordModel{
		BaseModel: model.NewBaseModel(ctx, db),
	}
}

// defaultMintModel 默认用户表模型
type defaultBurnRecordModel struct {
	*model.BaseModel
}

// Insert 插入BurnRecord表信息
func (o *defaultBurnRecordModel) Insert(data *BurnRecord) error {
	return o.DB.Create(data).Error
}

// InsertWithClauses 使用子句插入BurnRecord表信息
func (o *defaultBurnRecordModel) InsertWithClauses(data *BurnRecord, clauses ...clause.Expression) error {
	return o.DB.Clauses(clauses...).Create(data).Error
}

// InsertBatch 批量插入BurnRecord表信息
func (o *defaultBurnRecordModel) InsertBatch(dataset []*BurnRecord) error {
	return o.DB.Create(&dataset).Error
}

// InsertBatchWithClauses 使用子句批量插入BurnRecord表信息
func (o *defaultBurnRecordModel) InsertBatchWithClauses(dataset []*BurnRecord, clauses ...clause.Expression) (int64, error) {
	db := o.DB.Clauses(clauses...).Create(&dataset)

	return db.RowsAffected, db.Error
}

// Update 更新mint表信息
func (o *defaultBurnRecordModel) Update(data *BurnRecord, selects ...string) error {
	return o.BaseModel.Update(data, selects)
}

// UpdateByCondition 通过动态条件更新mint表信息
func (o *defaultBurnRecordModel) UpdateByCondition(data *BurnRecord, condition func(*gorm.DB) *gorm.DB) error {
	return o.BaseModel.UpdateByCondition(data, condition)
}

// UpdateWithMapByCondition 使用map通过动态条件更新BurnRecord表信息
func (o *defaultBurnRecordModel) UpdateWithMapByCondition(data map[string]interface{}, condition func(*gorm.DB) *gorm.DB) error {
	return o.BaseModel.UpdateWithMapByCondition(&BurnRecord{}, data, condition)
}

// FindOne 通过用户表id查找一个BurnRecord表信息
func (m *defaultBurnRecordModel) FindOne(id int64) (*BurnRecord, error) {
	if id < 1 {
		return nil, errcode.ErrInvalidParams
	}

	var o BurnRecord
	err := m.DB.Where("`id` = ?", id).First(&o).Error
	if err != nil {
		if model.IsRecordNotFound(err) {
			return nil, errors.New("BurnRecord not exist")
		}
		return nil, err
	}

	return &o, nil
}

// FindOneByCondition 通过动态条件查找一个BurnRecord表信息
func (m *defaultBurnRecordModel) FindOneByCondition(condition func(*gorm.DB) *gorm.DB) (*BurnRecord, error) {
	var o BurnRecord
	err := m.DB.Scopes(condition).First(&o).Error
	if err != nil {
		if model.IsRecordNotFound(err) {
			return nil, errors.New("BurnRecord not exist")
		}
		return nil, err
	}

	return &o, nil
}

// FindByCondition 通过动态条件查找BurnRecord表信息
func (o *defaultBurnRecordModel) FindByCondition(condition func(*gorm.DB) *gorm.DB) ([]*BurnRecord, error) {
	var op []*BurnRecord
	err := o.DB.Scopes(condition).Find(&op).Error
	if err != nil {
		return nil, err
	}

	return op, nil
}

func (o *defaultBurnRecordModel) FindMapByCondition(db *gorm.DB) ([]map[string]interface{}, error) {
	var op []map[string]interface{}
	err := db.Find(&op).Error
	if err != nil {
		return op, err
	}
	return op, nil
}

// Paginate 分页查找BurnRecord表信息
func (o *defaultBurnRecordModel) Paginate(pq *model.PaginationQuery) ([]*BurnRecord, int64, error) {
	var total int64
	err := o.DB.Scopes(pq.Queries()).Model(&BurnRecord{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	var op []*BurnRecord
	err = o.DB.Scopes(pq.Paginate()).Find(&op).Error
	if err != nil {
		return nil, 0, err
	}

	return op, total, nil
}

// PluckInt64s 通过动态条件查找单个列并将结果扫描到int64切片中
func (o *defaultBurnRecordModel) PluckInt64s(column string, condition func(*gorm.DB) *gorm.DB) ([]int64, error) {
	return o.BaseModel.PluckInt64s(&BurnRecord{}, column, condition)
}

// PluckStrings 通过动态条件查找单个列并将结果扫描到string切片中
func (o *defaultBurnRecordModel) PluckStrings(column string, condition func(*gorm.DB) *gorm.DB) ([]string, error) {
	return o.BaseModel.PluckStrings(&BurnRecord{}, column, condition)
}

// CountByCondition 通过动态条件计数用户表信息
func (o *defaultBurnRecordModel) CountByCondition(condition func(*gorm.DB) *gorm.DB) (int64, error) {
	return o.BaseModel.CountByCondition(&BurnRecord{}, condition)
}

// DeleteByCondition 通过动态条件删除用户表信息
func (o *defaultBurnRecordModel) DeleteByCondition(condition func(*gorm.DB) *gorm.DB) error {
	return o.BaseModel.DeleteByCondition(&BurnRecord{}, condition)
}

// Init 初始化BurnRecord表信息
func (o *defaultBurnRecordModel) Init() error {
	var num int64
	err := o.DB.Model(&BurnRecord{}).Count(&num).Error
	if err != nil {
		return err
	}

	if num == 0 {
		// 执行初始化方法
		return nil
	}
	return nil
}
