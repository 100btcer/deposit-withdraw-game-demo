package param

import (
	"mm-ndj/api/param/types"
	"mm-ndj/config"
	"mm-ndj/pkg/errcode"
	"mm-ndj/pkg/xhttp"
	"mm-ndj/server/dao"
	"mm-ndj/server/view/startland/user"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func WalletLogin(svc *dao.ServiceCtx) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req types.WalletLoginReq
		if err := ctx.BindJSON(&req); err != nil {
			config.Logger.Error("WalletLogin", zap.Error(err))
			xhttp.Error(ctx, errcode.NewCustomErr(err.Error()))
			return
		}
		//获取token
		data, err := user.WalletLogin(ctx.Request.Context(), svc, req)
		if err != nil {
			xhttp.Error(ctx, errcode.NewCustomErr(err.Error()))
			return
		}
		xhttp.OkJson(ctx, data)
	}
}

// 获取用户资料
func GetUserInfo(svc *dao.ServiceCtx) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userId, _ := ctx.Keys["user_id"].(int64)
		//获取token
		data, err := user.GetUserInfo(ctx.Request.Context(), svc, userId)
		if err != nil {
			xhttp.Error(ctx, errcode.NewCustomErr(err.Error()))
			return
		}
		xhttp.OkJson(ctx, data)
	}
}
