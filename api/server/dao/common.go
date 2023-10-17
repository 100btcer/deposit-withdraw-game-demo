package dao

import (
	"mm-ndj/pkg/errcode"
	"mm-ndj/pkg/utils"
	"reflect"
	"time"

	"github.com/jinzhu/now"
	"gorm.io/gorm"
)

const (
	ReportDateToday = iota + 1
	ReportDateYesterday
	ReportDateLast3Day
	ReportDateLast7Day
	ReportDateLastMonth
	ReportDateLast6Month
	ReportDateAll     = 10
	ReportDateAllDesc = "all"
	DaySecond         = 86400
)

func GetSearchTime(dateType int) (startStr, endStr string) {
	start := time.Now().Unix()
	end := time.Now().Unix()
	switch dateType {
	case ReportDateToday:
		start = now.BeginningOfDay().Unix()
	case ReportDateYesterday:
		start = now.BeginningOfDay().Unix() - 1*DaySecond
		end = now.BeginningOfDay().Unix()
	case ReportDateLast3Day:
		start = now.BeginningOfDay().Unix() - 3*DaySecond
	case ReportDateLast7Day:
		start = now.BeginningOfDay().Unix() - 7*DaySecond
	case ReportDateLastMonth:
		start = now.BeginningOfDay().Unix() - 30*DaySecond
	case ReportDateLast6Month:
		start = now.BeginningOfDay().Unix() - 180*DaySecond
	}
	startStr = utils.GetUnixToDateLong(start)
	endStr = utils.GetUnixToDateLong(end)
	return
}

func WhereWithUserId(userId int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", userId)
	}
}

func WhereWithKey(key string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("`key` = ? ", key)
	}
}

func WhereWithId(id int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ? ", id)
	}
}

func CheckExist(err error, infos interface{}) (bool, error) {

	//参数来自FindOneByCondition的返回,返回是否记录存在
	valueStr := reflect.ValueOf(infos)
	method := valueStr.MethodByName("TableName")
	resultValue := method.Call([]reflect.Value{})
	result := resultValue[0].Interface()
	tableName := result.(string)
	if err != nil && err.Error() == errcode.GetDbNotExistMsg(tableName) {
		return false, nil
	}
	if err != nil && err.Error() != errcode.GetDbNotExistMsg(tableName) {
		return false, err
	}
	return true, nil
}
