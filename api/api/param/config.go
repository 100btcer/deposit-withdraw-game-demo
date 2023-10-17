package param

import (
	"github.com/gin-gonic/gin"
	"mm-ndj/pkg/errcode"
	"mm-ndj/pkg/xhttp"
	"mm-ndj/server/dao"
	"mm-ndj/server/view/startland/config"
)

// 获取配置
func GetConfig(svc *dao.ServiceCtx) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, _ := ctx.Keys["user_id"].(int64)
		data, err := config.GetConfig(ctx.Request.Context(), svc, userId)
		if err != nil {
			xhttp.Error(ctx, errcode.NewCustomErr(err.Error()))
			return
		}
		xhttp.OkJson(ctx, data)
	}
}
