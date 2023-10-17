package dao

import (
	"mm-ndj/config"
	"mm-ndj/pkg/db/mysql"
	"mm-ndj/pkg/db/redis"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ServiceCtx struct {
	C   *config.Config
	Db  *gorm.DB
	Rds redis.RedisFactory
}

var serviceCtx *ServiceCtx
var once sync.Once

func NewServiceContext() (*ServiceCtx, error) {
	//初始化日志
	config.LoadConfig()
	//初始化mysql
	initMsql := mysql.NewMySQL()
	db, err := initMsql.NewMysqlDB()
	if err != nil {
		config.Logger.Error("INIT MYSQL CONFIG FAILED", zap.Error(err))
		return nil, err
	}
	//初始化redis
	rds := redis.InitRds(config.Conf.RedisC.RedisHost, config.Conf.RedisC.Password, 1)

	serviceCtxs := NewServerCtx(WithDB(db))
	serviceCtxs.C = config.Conf
	serviceCtxs.Rds = rds
	return serviceCtxs, nil
}

func GetServiceCtx() *ServiceCtx {
	once.Do(func() {
		serviceCtx, _ = NewServiceContext()
	})
	return serviceCtx
}
