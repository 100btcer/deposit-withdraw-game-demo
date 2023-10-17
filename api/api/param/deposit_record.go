package param

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mm-ndj/api/param/types"
	"mm-ndj/config"
	"mm-ndj/pkg/errcode"
	"mm-ndj/pkg/xhttp"
	"mm-ndj/server/dao"
	"mm-ndj/server/view/startland/deposit_record"
)

// 添加质押记录
func DepositRecordAdd(svc *dao.ServiceCtx) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.DepositRecordAddReq
		if err := ctx.BindJSON(&req); err != nil {
			config.Logger.Error("DepositRecordAdd", zap.Error(err))
			xhttp.Error(ctx, errcode.NewCustomErr(err.Error()))
			return
		}
		userId, _ := ctx.Keys["user_id"].(int64)
		err := deposit_record.DepositRecordAdd(ctx.Request.Context(), svc, userId, req.Hash, req.Amount, req.PoolId)
		if err != nil {
			xhttp.Error(ctx, errcode.NewCustomErr(err.Error()))
			return
		}
		xhttp.OkJson(ctx, nil)
	}
}

// 获取质押记录
func DepositRecordList(svc *dao.ServiceCtx) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req types.DepositRecordListReq
		if err := ctx.BindQuery(&req); err != nil {
			config.Logger.Error("DepositRecordList", zap.Error(err))
			xhttp.Error(ctx, errcode.NewCustomErr(err.Error()))
			return
		}
		userId, _ := ctx.Keys["user_id"].(int64)
		data, err := deposit_record.DepositRecordList(ctx.Request.Context(), svc, userId, req.Page, req.PageSize)
		if err != nil {
			xhttp.Error(ctx, errcode.NewCustomErr(err.Error()))
			return
		}
		xhttp.OkJson(ctx, data)
	}
}
