package param

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mm-ndj/api/param/types"
	"mm-ndj/config"
	"mm-ndj/pkg/errcode"
	"mm-ndj/pkg/xhttp"
	"mm-ndj/server/dao"
	"mm-ndj/server/view/startland/withdraw_record"
)

// 添加提款记录
func WithdrawRecordAdd(svc *dao.ServiceCtx) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.WithdrawRecordAddReq
		if err := ctx.BindJSON(&req); err != nil {
			config.Logger.Error("WithdrawRecordAdd", zap.Error(err))
			xhttp.Error(ctx, errcode.NewCustomErr(err.Error()))
			return
		}
		userId, _ := ctx.Keys["user_id"].(int64)
		err := withdraw_record.WithdrawRecordAdd(ctx.Request.Context(), svc, userId, req.Hash, req.Amount)
		if err != nil {
			xhttp.Error(ctx, errcode.NewCustomErr(err.Error()))
			return
		}
		xhttp.OkJson(ctx, nil)
	}
}

// 获取可提现金额
func GetValidWithdrawAmount(svc *dao.ServiceCtx) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req types.DepositRecordListReq
		if err := ctx.BindQuery(&req); err != nil {
			config.Logger.Error("DepositRecordList", zap.Error(err))
			xhttp.Error(ctx, errcode.NewCustomErr(err.Error()))
			return
		}
		userId, _ := ctx.Keys["user_id"].(int64)
		data, err := withdraw_record.GetValidWithdrawAmount(ctx.Request.Context(), svc, userId)
		if err != nil {
			xhttp.Error(ctx, errcode.NewCustomErr(err.Error()))
			return
		}
		xhttp.OkJson(ctx, data)
	}
}
