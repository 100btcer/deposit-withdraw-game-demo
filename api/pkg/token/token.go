package token

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"mm-ndj/config"
	"mm-ndj/pkg/jwt"
	"mm-ndj/server/dao"
	"time"
)

func CreateToken(svc *dao.ServiceCtx, userId int64, address, code string) (string, error) {
	//生成Token结构初始化
	jt, err := jwt.NewJWT(&jwt.Config{
		Issuer:         config.Conf.JwtC.Issuer,
		SecretKey:      config.Conf.JwtC.SecretKey,
		ExpirationTime: time.Duration(config.Conf.JwtC.ExpirationTime),
	})
	if err != nil {
		config.Logger.Error("Login NewJWT", zap.Error(err))
		return "", err
	}

	//生成token
	token, err := jt.CreateToken(code, address, 7*24*60*60*time.Second)
	if err != nil {
		config.Logger.Error("Login CreateToken", zap.Error(err))
		return "", err
	}
	fmt.Printf("token : %+v\n", token)
	//缓存token信息
	if err := CacheData(svc, userId, token, address); err != nil {
		config.Logger.Error("Login Rds.SetNx", zap.Error(err))
		return "", err
	}
	return token, nil
}

type CacheDataStruct struct {
	UserId  int64  `json:"user_id"`
	Address string `json:"address"`
	Token   string `json:"token"`
}

func CacheData(svc *dao.ServiceCtx, userId int64, token, address string) error {
	bytes, err := json.Marshal(CacheDataStruct{
		UserId:  userId,
		Token:   token,
		Address: address,
	})
	fmt.Printf("bytes : %+v\n", string(bytes))
	if err != nil {
		return err
	}
	if err := svc.Rds.SetNx(token, string(bytes), 7*24*60*60*time.Second); err != nil {
		config.Logger.Error("Login Rds.SetNx", zap.Error(err))
		return err
	}
	return nil
}

// 解析token
func ParseToken(svc *dao.ServiceCtx, token string) (data *CacheDataStruct, err error) {
	cacheData, _ := svc.Rds.GetValue(token)
	err = json.Unmarshal([]byte(cacheData), &data)
	return
}

// 验证token是否获取
func VfyToken(svc *dao.ServiceCtx, token string) (*CacheDataStruct, error) {
	if token == "" {
		return nil, errors.New("token expired")
	}
	val, err := svc.Rds.GetValue(token)
	if err != nil {
		config.Logger.Info("验证token信息", zap.Any("验证失败：", err))
		return nil, errors.New("token expired")
	}
	if val == "" {
		return nil, errors.New("token expired")
	}
	var tokenData *CacheDataStruct
	err = json.Unmarshal([]byte(val), &tokenData)
	if err != nil {
		return nil, err
	}
	return tokenData, nil
}
