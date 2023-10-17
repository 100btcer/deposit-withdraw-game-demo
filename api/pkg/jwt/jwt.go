package jwt

import (
	"encoding/base64"
	"encoding/json"
	"mm-ndj/pkg/errcode"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/pkg/errors"
)

const (
	// PrivatePayloadName 私有载荷名称
	PrivatePayloadName = "x_user_info"
)

// Config JWT相关配置
type Config struct {
	Issuer         string        // 签发者
	SecretKey      string        // 密钥
	ExpirationTime time.Duration // 过期时间
}

// JWT JWT结构详情
type JWT struct {
	c *Config
}

type Claims struct {
	WalletAddr string
	jwt.RegisteredClaims
}

// NewJWT 新建JWT
func NewJWT(c *Config) (*JWT, error) {
	return &JWT{c: c}, nil
}

// MustNewJWT 新建JWT
func MustNewJWT(c *Config) *JWT {
	j, err := NewJWT(c)
	if err != nil {
		panic(err)
	}

	return j
}

// CreateToken 创建JWT字符串
func (j *JWT) CreateToken(secretKey, address string, expirationTime time.Duration) (string, error) {
	claims := Claims{
		WalletAddr: address,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 生效时间
		}}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ts, err := t.SignedString([]byte(secretKey))
	if err != nil {
		return "", errors.WithMessage(err, "sign token err")
	}

	return ts, nil
}

// ParseToken 解析JWT字符串
func (j *JWT) ParseToken(tokenStr string, token interface{}) error {
	tokenStr = strings.Replace(tokenStr, "Bearer ", "", 1)

	t, err := jwt.Parse(tokenStr, jwtKeyFunc(j.c.SecretKey))
	if err != nil {
		if e, ok := err.(*jwt.ValidationError); ok {
			if e.Errors&jwt.ValidationErrorMalformed != 0 {
				return errcode.ErrTokenVerify
			} else if e.Errors&jwt.ValidationErrorExpired != 0 {
				return errcode.ErrTokenExpire
			} else if e.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return errcode.ErrTokenNotValidYet
			} else {
				return errcode.ErrTokenVerify
			}
		}
		return errcode.ErrTokenVerify
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return errcode.ErrTokenVerify
	}

	s, ok := claims[PrivatePayloadName].(string)
	if !ok {
		return errcode.ErrTokenVerify
	}

	payload, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return errcode.ErrTokenVerify
	}

	err = json.Unmarshal(payload, token)
	if err != nil {
		return errcode.ErrTokenVerify
	}

	return nil
}

func Secret(key string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	}
}

func (j *JWT) ParseToken1(token string, address, salt string) (*Claims, error) {
	parsed, err := jwt.ParseWithClaims(token, &Claims{
		WalletAddr: address,
	}, Secret(salt))

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("ErrTokenNotToken")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("ErrTokenExpire")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("ErrTokenNotValidYet")
			} else {
				return nil, errors.New("ErrTokenOthers1")
			}
		}
	}
	if claims, ok := parsed.Claims.(*Claims); ok && parsed.Valid {
		return claims, nil
	}
	return nil, errors.New("ErrTokenOthers")
}

// jwtKeyFunc JWT签名密钥函数
func jwtKeyFunc(key string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	}
}
