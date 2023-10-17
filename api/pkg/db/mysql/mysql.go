package mysql

import (
	"context"
	"mm-ndj/config"

	"gorm.io/gorm"
)

type MySQL struct {
	c  *Config
	DB *gorm.DB
}

func NewMySQL() *MySQL {
	config := &Config{
		User:               config.Conf.MySQLC.User,
		Host:               config.Conf.MySQLC.Host,
		Password:           config.Conf.MySQLC.Password,
		Port:               config.Conf.MySQLC.Port,
		Database:           config.Conf.MySQLC.Database,
		MaxIdleConns:       config.Conf.MySQLC.MaxIdleConns,
		MaxOpenConns:       config.Conf.MySQLC.MaxOpenConns,
		MaxConnMaxLifetime: config.Conf.MySQLC.MaxConnMaxLifetime,
		LogLevel:           config.Conf.MySQLC.LogLevel,
	}
	return &MySQL{
		c: config,
	}
}

func (msl *MySQL) NewMysqlDB() (*gorm.DB, error) {
	db := MustNewDB(msl.c)
	ctx := context.Background()
	if err := InitModel(ctx, db); err != nil {
		return nil, err
	}
	return db, nil
}

func (msl *MySQL) Close() {

}

func InitModel(ctx context.Context, db *gorm.DB) error {
	err := db.Set("gorm:table_options",
		"ENGINE=InnoDB AUTO_INCREMENT=1 CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci").AutoMigrate()
	if err != nil {
		return err
	}
	return nil
}
