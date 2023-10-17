package model

import (
	"context"
	"errors"
	"mm-ndj/pkg/errcode"
	"mm-ndj/pkg/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DepositRecordModel = interface {
	Insert(data *DepositRecord) error
	InsertWithClauses(data *DepositRecord, clauses ...clause.Expression) error
	InsertBatch(dataset []*DepositRecord) error
	InsertBatchWithClauses(dataset []*DepositRecord, clauses ...clause.Expression) (int64, error)
	Update(data *DepositRecord, selects ...string) error
	UpdateByCondition(data *DepositRecord, condition func(*gorm.DB) *gorm.DB) error
	UpdateWithMapByCondition(data map[string]interface{}, condition func(*gorm.DB) *gorm.DB) error
	FindOne(id int64) (*DepositRecord, error)
	FindOneByCondition(condition func(*gorm.DB) *gorm.DB) (*DepositRecord, error)
	FindByCondition(condition func(*gorm.DB) *gorm.DB) ([]*DepositRecord, error)
	FindMapByCondition(db *gorm.DB) ([]map[string]interface{}, error)
	Paginate(pq *model.PaginationQuery) ([]*DepositRecord, int64, error)
	PluckInt64s(column string, condition func(*gorm.DB) *gorm.DB) ([]int64, error)
	PluckStrings(column string, condition func(*gorm.DB) *gorm.DB) ([]string, error)
	CountByCondition(condition func(*gorm.DB) *gorm.DB) (int64, error)
	DeleteByCondition(condition func(*gorm.DB) *gorm.DB) error
	Init() error
}

type DepositRecord struct {
	Id            int64   `json:"id" gorm:"primaryKey;column:id;"`
	CreateTime    int64   `json:"create_time" gorm:"column:create_time;int"`
	UpdateTime    int64   `json:"update_time" gorm:"column:update_time;int"`
	LockDay       int     `json:"lock_day" gorm:"column:lock_day;int"`
	DepositAmount float64 `json:"deposit_amount" gorm:"column:deposit_amount;decimal(22,4)"`
	TicketAmount  float64 `json:"ticket_amount" gorm:"column:ticket_amount;decimal(22,4)"`
	Hash          string  `json:"hash" gorm:"column:hash;varchar(128)"`
	BlockHeight   int64   `json:"block_height" gorm:"column:block_height;bigint"`
	TxStatus      int     `json:"tx_status" gorm:"column:tx_status;tinyint"`
	ConfirmNum    int     `json:"confirm_num" gorm:"column:confirm_num;int"`
	UserId        int64   `json:"user_id" gorm:"column:user_id;int"`
	Address       string  `json:"address" gorm:"column:address;varchar(64)"`
	Status        int8    `json:"status" gorm:"column:status;tinyint"`
}

// TableName 返回用户表信息的数据库表名
func (o *DepositRecord) TableName() string {
	return "deposit_record"
}

// GetId 获取id
func (o *DepositRecord) GetId() int64 {
	return o.Id
}

// NewMintModel 新建用户表模型
func NewDepositRecordModel(ctx context.Context, db *gorm.DB) DepositRecordModel {
	return &defaultDepositRecordModel{
		BaseModel: model.NewBaseModel(ctx, db),
	}
}

// defaultMintModel 默认用户表模型
type defaultDepositRecordModel struct {
	*model.BaseModel
}

// Insert 插入DepositRecord表信息
func (o *defaultDepositRecordModel) Insert(data *DepositRecord) error {
	return o.DB.Create(data).Error
}

// InsertWithClauses 使用子句插入DepositRecord表信息
func (o *defaultDepositRecordModel) InsertWithClauses(data *DepositRecord, clauses ...clause.Expression) error {
	return o.DB.Clauses(clauses...).Create(data).Error
}

// InsertBatch 批量插入DepositRecord表信息
func (o *defaultDepositRecordModel) InsertBatch(dataset []*DepositRecord) error {
	return o.DB.Create(&dataset).Error
}

// InsertBatchWithClauses 使用子句批量插入DepositRecord表信息
func (o *defaultDepositRecordModel) InsertBatchWithClauses(dataset []*DepositRecord, clauses ...clause.Expression) (int64, error) {
	db := o.DB.Clauses(clauses...).Create(&dataset)

	return db.RowsAffected, db.Error
}

// Update 更新mint表信息
func (o *defaultDepositRecordModel) Update(data *DepositRecord, selects ...string) error {
	return o.BaseModel.Update(data, selects)
}

// UpdateByCondition 通过动态条件更新mint表信息
func (o *defaultDepositRecordModel) UpdateByCondition(data *DepositRecord, condition func(*gorm.DB) *gorm.DB) error {
	return o.BaseModel.UpdateByCondition(data, condition)
}

// UpdateWithMapByCondition 使用map通过动态条件更新DepositRecord表信息
func (o *defaultDepositRecordModel) UpdateWithMapByCondition(data map[string]interface{}, condition func(*gorm.DB) *gorm.DB) error {
	return o.BaseModel.UpdateWithMapByCondition(&DepositRecord{}, data, condition)
}

// FindOne 通过用户表id查找一个DepositRecord表信息
func (m *defaultDepositRecordModel) FindOne(id int64) (*DepositRecord, error) {
	if id < 1 {
		return nil, errcode.ErrInvalidParams
	}

	var o DepositRecord
	err := m.DB.Where("`id` = ?", id).First(&o).Error
	if err != nil {
		if model.IsRecordNotFound(err) {
			return nil, errors.New("DepositRecord not exist")
		}
		return nil, err
	}

	return &o, nil
}

// FindOneByCondition 通过动态条件查找一个DepositRecord表信息
func (m *defaultDepositRecordModel) FindOneByCondition(condition func(*gorm.DB) *gorm.DB) (*DepositRecord, error) {
	var o DepositRecord
	err := m.DB.Scopes(condition).First(&o).Error
	if err != nil {
		if model.IsRecordNotFound(err) {
			return nil, errors.New("DepositRecord not exist")
		}
		return nil, err
	}

	return &o, nil
}

// FindByCondition 通过动态条件查找DepositRecord表信息
func (o *defaultDepositRecordModel) FindByCondition(condition func(*gorm.DB) *gorm.DB) ([]*DepositRecord, error) {
	var op []*DepositRecord
	err := o.DB.Scopes(condition).Find(&op).Error
	if err != nil {
		return nil, err
	}

	return op, nil
}

func (o *defaultDepositRecordModel) FindMapByCondition(db *gorm.DB) ([]map[string]interface{}, error) {
	var op []map[string]interface{}
	err := db.Find(&op).Error
	if err != nil {
		return op, err
	}
	return op, nil
}

// Paginate 分页查找DepositRecord表信息
func (o *defaultDepositRecordModel) Paginate(pq *model.PaginationQuery) ([]*DepositRecord, int64, error) {
	var total int64
	err := o.DB.Scopes(pq.Queries()).Model(&DepositRecord{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	var op []*DepositRecord
	err = o.DB.Scopes(pq.Paginate()).Find(&op).Error
	if err != nil {
		return nil, 0, err
	}

	return op, total, nil
}

// PluckInt64s 通过动态条件查找单个列并将结果扫描到int64切片中
func (o *defaultDepositRecordModel) PluckInt64s(column string, condition func(*gorm.DB) *gorm.DB) ([]int64, error) {
	return o.BaseModel.PluckInt64s(&DepositRecord{}, column, condition)
}

// PluckStrings 通过动态条件查找单个列并将结果扫描到string切片中
func (o *defaultDepositRecordModel) PluckStrings(column string, condition func(*gorm.DB) *gorm.DB) ([]string, error) {
	return o.BaseModel.PluckStrings(&DepositRecord{}, column, condition)
}

// CountByCondition 通过动态条件计数用户表信息
func (o *defaultDepositRecordModel) CountByCondition(condition func(*gorm.DB) *gorm.DB) (int64, error) {
	return o.BaseModel.CountByCondition(&DepositRecord{}, condition)
}

// DeleteByCondition 通过动态条件删除用户表信息
func (o *defaultDepositRecordModel) DeleteByCondition(condition func(*gorm.DB) *gorm.DB) error {
	return o.BaseModel.DeleteByCondition(&DepositRecord{}, condition)
}

// Init 初始化DepositRecord表信息
func (o *defaultDepositRecordModel) Init() error {
	var num int64
	err := o.DB.Model(&DepositRecord{}).Count(&num).Error
	if err != nil {
		return err
	}

	if num == 0 {
		// 执行初始化方法
		return nil
	}
	return nil
}
