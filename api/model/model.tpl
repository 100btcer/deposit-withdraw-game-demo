{{$exportModelName := .ModelName | FirstCharUpper}}
package {{.PackageName}}

import (
	"context"
    "errors"
	"mm-ndj/pkg/errcode"
	"mm-ndj/pkg/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type {{$exportModelName}}Model = interface{
	Insert(data *{{$exportModelName}}) error
	InsertWithClauses(data *{{$exportModelName}}, clauses ...clause.Expression) error
	InsertBatch(dataset []*{{$exportModelName}}) error
	InsertBatchWithClauses(dataset []*{{$exportModelName}}, clauses ...clause.Expression) (int64, error)
	{{if .IsHaveId}}Update(data *{{$exportModelName}}, selects ...string) error{{end}}
	UpdateByCondition(data *{{$exportModelName}}, condition func(*gorm.DB) *gorm.DB) error
	UpdateWithMapByCondition(data map[string]interface{}, condition func(*gorm.DB) *gorm.DB) error
	FindOne(id int64) (*{{$exportModelName}}, error)
	FindOneByCondition(condition func(*gorm.DB) *gorm.DB) (*{{$exportModelName}}, error)
	FindByCondition(condition func(*gorm.DB) *gorm.DB) ([]*{{$exportModelName}}, error)
	FindMapByCondition(db *gorm.DB) ([]map[string]interface{}, error)
	Paginate(pq *model.PaginationQuery) ([]*{{$exportModelName}}, int64, error)
	PluckInt64s(column string, condition func(*gorm.DB) *gorm.DB) ([]int64, error)
	PluckStrings(column string, condition func(*gorm.DB) *gorm.DB) ([]string, error)
	CountByCondition(condition func(*gorm.DB) *gorm.DB) (int64, error)
	DeleteByCondition(condition func(*gorm.DB) *gorm.DB) error
	Init() error
}

type {{$exportModelName}} struct {
{{range .TableSchema}} {{.Field | ExportColumn}} {{. | TypeConvert}} {{. | Tags}}
{{end}}}


// TableName 返回用户表信息的数据库表名
func (o *{{$exportModelName}}) TableName() string {
	return "{{.TableName}}"
}

{{if .IsHaveId}}
// GetId 获取id
func (o *{{$exportModelName}}) GetId() int64 {
	return o.Id
}
{{end}}

// NewMintModel 新建用户表模型
func New{{$exportModelName}}Model(ctx context.Context, db *gorm.DB) {{$exportModelName}}Model {
	return &default{{$exportModelName}}Model{
		BaseModel: model.NewBaseModel(ctx, db),
	}
}

// defaultMintModel 默认用户表模型
type default{{$exportModelName}}Model struct {
	*model.BaseModel
}

// Insert 插入{{$exportModelName}}表信息
func (o *default{{$exportModelName}}Model) Insert(data *{{$exportModelName}}) error {
	return o.DB.Create(data).Error
}

// InsertWithClauses 使用子句插入{{$exportModelName}}表信息
func (o *default{{$exportModelName}}Model) InsertWithClauses(data *{{$exportModelName}}, clauses ...clause.Expression) error {
	return o.DB.Clauses(clauses...).Create(data).Error
}

// InsertBatch 批量插入{{$exportModelName}}表信息
func (o *default{{$exportModelName}}Model) InsertBatch(dataset []*{{$exportModelName}}) error {
	return o.DB.Create(&dataset).Error
}

// InsertBatchWithClauses 使用子句批量插入{{$exportModelName}}表信息
func (o *default{{$exportModelName}}Model) InsertBatchWithClauses(dataset []*{{$exportModelName}}, clauses ...clause.Expression) (int64, error) {
	db := o.DB.Clauses(clauses...).Create(&dataset)

	return db.RowsAffected, db.Error
}

{{if .IsHaveId}}
// Update 更新mint表信息
func (o *default{{$exportModelName}}Model) Update(data *{{$exportModelName}}, selects ...string) error {
	return o.BaseModel.Update(data, selects)
}
{{end}}

// UpdateByCondition 通过动态条件更新mint表信息
func (o *default{{$exportModelName}}Model) UpdateByCondition(data *{{$exportModelName}}, condition func(*gorm.DB) *gorm.DB) error {
	return o.BaseModel.UpdateByCondition(data, condition)
}

// UpdateWithMapByCondition 使用map通过动态条件更新{{$exportModelName}}表信息
func (o *default{{$exportModelName}}Model) UpdateWithMapByCondition(data map[string]interface{}, condition func(*gorm.DB) *gorm.DB) error {
	return o.BaseModel.UpdateWithMapByCondition(&{{$exportModelName}}{}, data, condition)
}

// FindOne 通过用户表id查找一个{{$exportModelName}}表信息
func (m *default{{$exportModelName}}Model) FindOne(id int64) (*{{$exportModelName}}, error) {
	if id {{ .SpecialStr | unescaped}} 1 {
		return nil, errcode.ErrInvalidParams
	}

	var o {{$exportModelName}}
	err := m.DB.Where("`id` = ?", id).First(&o).Error
	if err != nil {
		if model.IsRecordNotFound(err) {
			return nil, errors.New("{{$exportModelName}} not exist")
		}
		return nil, err
	}

	return &o, nil
}

// FindOneByCondition 通过动态条件查找一个{{$exportModelName}}表信息
func (m *default{{$exportModelName}}Model) FindOneByCondition(condition func(*gorm.DB) *gorm.DB) (*{{$exportModelName}}, error) {
	var o {{$exportModelName}}
	err := m.DB.Scopes(condition).First(&o).Error
	if err != nil {
		if model.IsRecordNotFound(err) {
			return nil, errors.New("{{$exportModelName}} not exist")
		}
		return nil, err
	}

	return &o, nil
}

// FindByCondition 通过动态条件查找{{$exportModelName}}表信息
func (o *default{{$exportModelName}}Model) FindByCondition(condition func(*gorm.DB) *gorm.DB) ([]*{{$exportModelName}}, error) {
	var op []*{{$exportModelName}}
	err := o.DB.Scopes(condition).Find(&op).Error
	if err != nil {
		return nil, err
	}

	return op, nil
}

func (o *default{{$exportModelName}}Model) FindMapByCondition(db *gorm.DB) ([]map[string]interface{}, error) {
	var op []map[string]interface{}
	err := db.Find(&op).Error
	if err != nil {
		return op, err
	}
	return op, nil
}

// Paginate 分页查找{{$exportModelName}}表信息
func (o *default{{$exportModelName}}Model) Paginate(pq *model.PaginationQuery) ([]*{{$exportModelName}}, int64, error) {
	var total int64
	err := o.DB.Scopes(pq.Queries()).Model(&{{$exportModelName}}{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	var op []*{{$exportModelName}}
	err = o.DB.Scopes(pq.Paginate()).Find(&op).Error
	if err != nil {
		return nil, 0, err
	}

	return op, total, nil
}

// PluckInt64s 通过动态条件查找单个列并将结果扫描到int64切片中
func (o *default{{$exportModelName}}Model) PluckInt64s(column string, condition func(*gorm.DB) *gorm.DB) ([]int64, error) {
	return o.BaseModel.PluckInt64s(&{{$exportModelName}}{}, column, condition)
}

// PluckStrings 通过动态条件查找单个列并将结果扫描到string切片中
func (o *default{{$exportModelName}}Model) PluckStrings(column string, condition func(*gorm.DB) *gorm.DB) ([]string, error) {
	return o.BaseModel.PluckStrings(&{{$exportModelName}}{}, column, condition)
}

// CountByCondition 通过动态条件计数用户表信息
func (o *default{{$exportModelName}}Model) CountByCondition(condition func(*gorm.DB) *gorm.DB) (int64, error) {
	return o.BaseModel.CountByCondition(&{{$exportModelName}}{}, condition)
}

// DeleteByCondition 通过动态条件删除用户表信息
func (o *default{{$exportModelName}}Model) DeleteByCondition(condition func(*gorm.DB) *gorm.DB) error {
	return o.BaseModel.DeleteByCondition(&{{$exportModelName}}{}, condition)
}

// Init 初始化{{$exportModelName}}表信息
func (o *default{{$exportModelName}}Model) Init() error {
	var num int64
	err := o.DB.Model(&{{$exportModelName}}{}).Count(&num).Error
	if err != nil {
		return err
	}

	if num == 0 {
		// 执行初始化方法
		return nil
	}
	return nil
}