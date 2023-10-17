package app

import (
	"mm-ndj/config"

	"mm-ndj/server/dao"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Platform struct {
	config    *config.Config
	router    *gin.Engine
	serverCtx *dao.ServiceCtx
}

func NewPlatform(config *config.Config, router *gin.Engine, server *dao.ServiceCtx) (*Platform, error) {
	return &Platform{
		config:    config,
		router:    router,
		serverCtx: server,
	}, nil
}

func (p *Platform) AppStart() error {
	config.Logger.Info("[INFO]", zap.String("service is", " starting ..."), zap.Any("address:", p.config.ServerC.Addr))
	if err := p.router.Run(p.config.ServerC.Addr); err != nil {
		config.Logger.Error("start service faild", zap.Error(err))
		return err
	}
	return nil
}

func (p *Platform) AppClose() error {
	p.serverCtx.Rds.Close()
	db, _ := p.serverCtx.Db.DB()
	db.Close()
	return nil
}
