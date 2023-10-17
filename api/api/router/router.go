package router

import (
	"mm-ndj/api/param"
	"mm-ndj/server/dao"

	"github.com/gin-gonic/gin"
)

func NewRouter(svcCtx *dao.ServiceCtx) *gin.Engine {
	gin.ForceConsoleColor()
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.Use(corsMiddleware())
	loadV1(r, svcCtx)

	return r
}

func loadV1(r *gin.Engine, svcCtx *dao.ServiceCtx) {
	{
		apiv1 := r.Group("/api")
		u := apiv1.Group("/user")
		u.POST("/wallet/login", param.WalletLogin(svcCtx))
		u.Use(checkToken(svcCtx))
		{
			u.GET("/info", param.GetUserInfo(svcCtx))
		}

		deposit := apiv1.Group("/deposit").Use(checkToken(svcCtx))
		{
			deposit.POST("/record/add", param.DepositRecordAdd(svcCtx))  //添加质押记录
			deposit.GET("/record/list", param.DepositRecordList(svcCtx)) //获取质押记录列表
		}

		withdraw := apiv1.Group("/withdraw").Use(checkToken(svcCtx))
		{
			withdraw.POST("/record/add", param.WithdrawRecordAdd(svcCtx))           //添加提款记录
			withdraw.GET("/valid_amount/get", param.GetValidWithdrawAmount(svcCtx)) //获取可提现金额
		}

		config := apiv1.Group("/config").Use(checkToken(svcCtx))
		{
			config.GET("/get", param.GetConfig(svcCtx)) //获取配置
		}

		lottery := apiv1.Group("/lottery").Use(checkToken(svcCtx))
		{
			lottery.GET("/result/get", param.LotteryResultGet(svcCtx))       //获取配置
			lottery.GET("/result/list", param.LotteryResultList(svcCtx))     //获取列表
			lottery.GET("/prize/proof/get", param.LotteryPrizeProof(svcCtx)) //获取证明
			lottery.GET("/prize/award/giveout", param.GiveOutAward(svcCtx))  //系统发奖
		}
	}
}
