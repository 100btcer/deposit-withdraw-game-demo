package router

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mm-ndj/config"
	"mm-ndj/pkg/errcode"
	"mm-ndj/pkg/token"
	"mm-ndj/pkg/xhttp"
	"mm-ndj/server/dao"
	"net/http"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Origin,Content-Type,AccessToken,X-CSRF-Token, Authorization, token,Cookie,Set-Cookie")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func checkToken(svc *dao.ServiceCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token
		tokenStr := c.Request.Header.Get("token")
		if tokenStr == "" {
			config.Logger.Error("token", zap.String("get token", "token is null"))
			xhttp.Error(c, errcode.ErrTokenNotValidYet)
			c.Abort()
			return
		}
		if err := VfyToken(svc, tokenStr); err != nil {
			config.Logger.Error("token", zap.String("get token", "token错误"))
			xhttp.Error(c, errcode.ErrTokenNotValidYet)
			c.Abort()
			return
		}
		cacheData, err := token.ParseToken(svc, tokenStr)
		if err != nil {
			config.Logger.Error("token", zap.String("get token", "读取缓存数据失败"))
			xhttp.Error(c, errcode.ErrTokenNotValidYet)
			c.Abort()
			return
		}
		c.Set("user_id", cacheData.UserId)
		c.Set("address", cacheData.Address)
		c.Next()
	}
}

// 验证token是否获取
func VfyToken(svc *dao.ServiceCtx, token string) error {
	if token == "" {
		return errors.New("token expired")
	}
	val, err := svc.Rds.GetValue(token)
	if err != nil {
		config.Logger.Info("验证token信息", zap.Any("验证失败：", err))
		return errors.New("token expired")
	}
	if val == "" {
		return errors.New("token expired")
	}
	return nil
}
