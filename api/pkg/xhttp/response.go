package xhttp

import (
	"mm-ndj/config"
	"mm-ndj/pkg/errcode"
	"mm-ndj/pkg/kit/convert"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Response struct {
	TraceId string      `json:"trace_id" example:"a1b2c3d4e5f6g7h8" extensions:"x-order=000"` // 链路追踪id
	Code    uint32      `json:"code" example:"200" extensions:"x-order=001"`                  // 状态码
	Msg     string      `json:"msg" example:"OK" extensions:"x-order=002"`                    // 消息
	Data    interface{} `json:"data" extensions:"x-order=003"`                                // 数据
}

type ResponseErr struct {
	TraceId string      `json:"trace_id" example:"a1b2c3d4e5f6g7h8" extensions:"x-order=000"` // 链路追踪id
	Data    interface{} `json:"data" extensions:"x-order=003"`                                // 数据
}

type ResponseErrUnexpected struct {
	ResponseErr
	Code uint32 `json:"code" example:"7777" extensions:"x-order=001"`                                 // 状态码
	Msg  string `json:"msg" example:"Network error, please try again later" extensions:"x-order=002"` // 消息
}

type ResponseErrParamInvalid struct {
	ResponseErr
	Code uint32 `json:"code" example:"10002" extensions:"x-order=001"`               // 状态码
	Msg  string `json:"msg" example:"Parameter is illegal" extensions:"x-order=002"` // 消息
}

type ResponseErrTokenNotValidYet struct {
	ResponseErr
	Code uint32 `json:"code" example:"9999" extensions:"x-order=001"`         // 状态码
	Msg  string `json:"msg" example:"Token illegal" extensions:"x-order=002"` // 消息
}

func ResponseData(code uint32, msg string, data interface{}) (*Response, error) {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}, nil
}

// // GetTraceId 获取链路追踪id
// func GetTraceId(ctx context.Context) string {
// 	spanCtx := trace.SpanContextFromContext(ctx)
// 	if spanCtx.HasTraceID() {
// 		return spanCtx.TraceID().String()
// 	}

// 	return ""
// }

// WriteHeader 写入自定义响应header
func WriteHeader(w http.ResponseWriter, err ...error) {
	var ee error
	if len(err) > 0 {
		ee = err[0]
	}

	e := errcode.ParseErr(ee)
	w.Header().Set(HeaderGWErrorCode, convert.ToString(e.Code()))
	w.Header().Set(HeaderGWErrorMessage, url.QueryEscape(e.Error()))
}

// OkJson 成功json响应返回
func OkJson(c *gin.Context, v interface{}) {
	WriteHeader(c.Writer)

	c.JSON(http.StatusOK, &Response{
		TraceId: "",
		Code:    errcode.CodeOK,
		Msg:     errcode.MsgOK,
		Data:    v,
	})
}

// Error 错误响应返回
func Error(c *gin.Context, err error) {
	// ctx := c.Request.Context()
	e := errcode.ParseErr(err)
	if e == errcode.ErrUnexpected || e == errcode.ErrCustom {
		config.Logger.Error("request handler err code:", zap.Any("", e.Code()), zap.Error(err))
	}

	WriteHeader(c.Writer, e)
	c.JSON(e.HTTPCode(), &Response{
		TraceId: "",
		Code:    e.Code(),
		Msg:     e.Error(),
		Data:    nil,
	})
}
