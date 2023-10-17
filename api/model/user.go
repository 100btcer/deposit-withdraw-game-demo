package model

import (
	"context"
	"errors"
	"mm-ndj/pkg/errcode"
	"mm-ndj/pkg/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserModel = interface {
	Insert(data *User) error
	InsertWithClauses(data *User, clauses ...clause.Expression) error
	InsertBatch(dataset []*User) error
	InsertBatchWithClauses(dataset []*User, clauses ...clause.Expression) (int64, error)
	Update(data *User, selects ...string) error
	UpdateByCondition(data *User, condition func(*gorm.DB) *gorm.DB) error
	UpdateWithMapByCondition(data map[string]interface{}, condition func(*gorm.DB) *gorm.DB) error
	FindOne(id int64) (*User, error)
	FindOneByCondition(condition func(*gorm.DB) *gorm.DB) (*User, error)
	FindByCondition(condition func(*gorm.DB) *gorm.DB) ([]*User, error)
	FindMapByCondition(db *gorm.DB) ([]map[string]interface{}, error)
	Paginate(pq *model.PaginationQuery) ([]*User, int64, error)
	PluckInt64s(column string, condition func(*gorm.DB) *gorm.DB) ([]int64, error)
	PluckStrings(column string, condition func(*gorm.DB) *gorm.DB) ([]string, error)
	CountByCondition(condition func(*gorm.DB) *gorm.DB) (int64, error)
	DeleteByCondition(condition func(*gorm.DB) *gorm.DB) error
	Init() error
}

type User struct {
	Id         int64  `json:"id" gorm:"primaryKey;column:id;"`
	Address    string `json:"address" gorm:"column:address;varchar(191)"`
	Name       string `json:"name" gorm:"column:name;varchar(20)"`
	AvatarUri  string `json:"avatar_uri" gorm:"column:avatar_uri;varchar(512)"`
	CreateTime int64  `json:"create_time" gorm:"column:create_time;bigint"`
	UpdateTime int64  `json:"update_time" gorm:"column:update_time;bigint"`
	LotteryNum int    `json:"lottery_num" gorm:"column:lottery_num;int"`
	Version    int64  `json:"version" gorm:"column:version;bigint"`
}

// TableName 返回用户表信息的数据库表名
func (o *User) TableName() string {
	return "user"
}

// GetId 获取id
func (o *User) GetId() int64 {
	return o.Id
}

// NewMintModel 新建用户表模型
func NewUserModel(ctx context.Context, db *gorm.DB) UserModel {
	return &defaultUserModel{
		BaseModel: model.NewBaseModel(ctx, db),
	}
}

// defaultMintModel 默认用户表模型
type defaultUserModel struct {
	*model.BaseModel
}

// Insert 插入User表信息
func (o *defaultUserModel) Insert(data *User) error {
	return o.DB.Create(data).Error
}

// InsertWithClauses 使用子句插入User表信息
func (o *defaultUserModel) InsertWithClauses(data *User, clauses ...clause.Expression) error {
	return o.DB.Clauses(clauses...).Create(data).Error
}

// InsertBatch 批量插入User表信息
func (o *defaultUserModel) InsertBatch(dataset []*User) error {
	return o.DB.Create(&dataset).Error
}

// InsertBatchWithClauses 使用子句批量插入User表信息
func (o *defaultUserModel) InsertBatchWithClauses(dataset []*User, clauses ...clause.Expression) (int64, error) {
	db := o.DB.Clauses(clauses...).Create(&dataset)

	return db.RowsAffected, db.Error
}

// Update 更新mint表信息
func (o *defaultUserModel) Update(data *User, selects ...string) error {
	return o.BaseModel.Update(data, selects)
}

// UpdateByCondition 通过动态条件更新mint表信息
func (o *defaultUserModel) UpdateByCondition(data *User, condition func(*gorm.DB) *gorm.DB) error {
	return o.BaseModel.UpdateByCondition(data, condition)
}

// UpdateWithMapByCondition 使用map通过动态条件更新User表信息
func (o *defaultUserModel) UpdateWithMapByCondition(data map[string]interface{}, condition func(*gorm.DB) *gorm.DB) error {
	return o.BaseModel.UpdateWithMapByCondition(&User{}, data, condition)
}

// FindOne 通过用户表id查找一个User表信息
func (m *defaultUserModel) FindOne(id int64) (*User, error) {
	if id < 1 {
		return nil, errcode.ErrInvalidParams
	}

	var o User
	err := m.DB.Where("`id` = ?", id).First(&o).Error
	if err != nil {
		if model.IsRecordNotFound(err) {
			return nil, errors.New("User not exist")
		}
		return nil, err
	}

	return &o, nil
}

// FindOneByCondition 通过动态条件查找一个User表信息
func (m *defaultUserModel) FindOneByCondition(condition func(*gorm.DB) *gorm.DB) (*User, error) {
	var o User
	err := m.DB.Scopes(condition).First(&o).Error
	if err != nil {
		if model.IsRecordNotFound(err) {
			return nil, errors.New("User not exist")
		}
		return nil, err
	}

	return &o, nil
}

// FindByCondition 通过动态条件查找User表信息
func (o *defaultUserModel) FindByCondition(condition func(*gorm.DB) *gorm.DB) ([]*User, error) {
	var op []*User
	err := o.DB.Scopes(condition).Find(&op).Error
	if err != nil {
		return nil, err
	}

	return op, nil
}

func (o *defaultUserModel) FindMapByCondition(db *gorm.DB) ([]map[string]interface{}, error) {
	var op []map[string]interface{}
	err := db.Find(&op).Error
	if err != nil {
		return op, err
	}
	return op, nil
}

// Paginate 分页查找User表信息
func (o *defaultUserModel) Paginate(pq *model.PaginationQuery) ([]*User, int64, error) {
	var total int64
	err := o.DB.Scopes(pq.Queries()).Model(&User{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	var op []*User
	err = o.DB.Scopes(pq.Paginate()).Find(&op).Error
	if err != nil {
		return nil, 0, err
	}

	return op, total, nil
}

// PluckInt64s 通过动态条件查找单个列并将结果扫描到int64切片中
func (o *defaultUserModel) PluckInt64s(column string, condition func(*gorm.DB) *gorm.DB) ([]int64, error) {
	return o.BaseModel.PluckInt64s(&User{}, column, condition)
}

// PluckStrings 通过动态条件查找单个列并将结果扫描到string切片中
func (o *defaultUserModel) PluckStrings(column string, condition func(*gorm.DB) *gorm.DB) ([]string, error) {
	return o.BaseModel.PluckStrings(&User{}, column, condition)
}

// CountByCondition 通过动态条件计数用户表信息
func (o *defaultUserModel) CountByCondition(condition func(*gorm.DB) *gorm.DB) (int64, error) {
	return o.BaseModel.CountByCondition(&User{}, condition)
}

// DeleteByCondition 通过动态条件删除用户表信息
func (o *defaultUserModel) DeleteByCondition(condition func(*gorm.DB) *gorm.DB) error {
	return o.BaseModel.DeleteByCondition(&User{}, condition)
}

// Init 初始化User表信息
func (o *defaultUserModel) Init() error {
	var num int64
	err := o.DB.Model(&User{}).Count(&num).Error
	if err != nil {
		return err
	}

	if num == 0 {
		// 执行初始化方法
		return nil
	}
	return nil
}
