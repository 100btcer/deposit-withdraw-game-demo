package param

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mm-ndj/api/param/types"
	"mm-ndj/config"
	"mm-ndj/pkg/errcode"
	"mm-ndj/pkg/xhttp"
	"mm-ndj/server/dao"
	"mm-ndj/server/view/startland/lottery"
)

// 获取开奖结果
func LotteryResultGet(svc *dao.ServiceCtx) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req types.LotteryResultReq
		if err := ctx.BindQuery(&req); err != nil {
			config.Logger.Error("LotteryResultGet", zap.Error(err))
			xhttp.Error(ctx, errcode.NewCustomErr(err.Error()))
			return
		}
		userId, _ := ctx.Keys["user_id"].(int64)
		data, err := lottery.GetLotteryResult(ctx.Request.Context(), svc, userId, req.Hash)
		if err != nil {
			xhttp.Error(ctx, errcode.NewCustomErr(err.Error()))
			return
		}
		xhttp.OkJson(ctx, data)
	}
}

// 获取中奖结果
func LotteryResultList(svc *dao.ServiceCtx) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req types.LotteryResultListReq
		if err := ctx.BindQuery(&req); err != nil {
			config.Logger.Error("LotteryResultList", zap.Error(err))
			xhttp.Error(ctx, errcode.NewCustomErr(err.Error()))
			return
		}
		userId, _ := ctx.Keys["user_id"].(int64)
		var data interface{}
		var err error
		if req.Type != nil {
			//查询token或NFT
			data, err = lottery.GetLotteryPrizeList(ctx.Request.Context(), svc, userId, req.Type, req.Page, req.PageSize)
		} else {
			//查询记录
			data, err = lottery.GetLotteryPrizeRecordList(ctx.Request.Context(), svc, userId, req.Type, req.Page, req.PageSize)
		}
		if err != nil {
			xhttp.Error(ctx, errcode.NewCustomErr(err.Error()))
			return
		}
		xhttp.OkJson(ctx, data)
	}
}

// 获取奖品证明
func LotteryPrizeProof(svc *dao.ServiceCtx) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req types.PrizeProofReq
		if err := ctx.BindQuery(&req); err != nil {
			config.Logger.Error("LotteryResultList", zap.Error(err))
			xhttp.Error(ctx, errcode.NewCustomErr(err.Error()))
			return
		}
		userId, _ := ctx.Keys["user_id"].(int64)
		data, err := lottery.GetPrizeProof(ctx.Request.Context(), svc, userId, req.Id)
		if err != nil {
			xhttp.Error(ctx, errcode.NewCustomErr(err.Error()))
			return
		}
		xhttp.OkJson(ctx, data)
	}
}

// 系统发奖
func GiveOutAward(svc *dao.ServiceCtx) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req types.PrizeProofReq
		if err := ctx.BindQuery(&req); err != nil {
			config.Logger.Error("LotteryResultList", zap.Error(err))
			xhttp.Error(ctx, errcode.NewCustomErr(err.Error()))
			return
		}
		userId, _ := ctx.Keys["user_id"].(int64)
		data, err := lottery.GiveOutAward(ctx.Request.Context(), svc, userId, req.Id)
		if err != nil {
			xhttp.Error(ctx, errcode.NewCustomErr(err.Error()))
			return
		}
		xhttp.OkJson(ctx, data)
	}
}
