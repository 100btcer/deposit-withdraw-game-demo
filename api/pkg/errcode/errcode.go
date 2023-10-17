package errcode

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"mm-ndj/config"
	"net/http"
	"strings"

	"go.uber.org/zap"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"google.golang.org/grpc/status"
)

const (
	// CodeOK 请求成功业务状态码
	CodeOK = 200
	// MsgOK 请求成功消息
	MsgOK = "Successful"

	// CodeCustom 自定义错误业务状态码
	CodeCustom = 7000
	// MsgCustom 自定义错误消息
	MsgCustom = "Custom error"
)

// Err 业务错误结构
type Err struct {
	code     uint32
	httpCode int
	msg      string
}

// Code 业务状态码
func (e *Err) Code() uint32 {
	return e.code
}

// HTTPCode HTTP状态码
func (e *Err) HTTPCode() int {
	return e.httpCode
}

// Error 消息
func (e *Err) Error() string {
	return e.msg
}

// 业务错误
var (
	NoErr = NewErr(CodeOK, MsgOK)

	ErrCustom     = NewErr(CodeCustom, MsgCustom)
	ErrUnexpected = NewErr(7777, "Network error, please try again later", http.StatusInternalServerError)

	ErrTokenNotValidYet    = NewErr(9999, "Token illegal", http.StatusOK)
	ErrInvalidUrl          = NewErr(10000, "URL is illegal")
	ErrInvalidHeader       = NewErr(10001, "Invalid request header")
	ErrInvalidParams       = NewErr(10002, "Parameter is illegal")
	ErrTokenVerify         = NewErr(10003, "Token check error", http.StatusUnauthorized)
	ErrTokenExpire         = NewErr(10004, "Expired token", http.StatusOK)
	ErrUserLogin           = NewErr(10005, "User log out", http.StatusUnauthorized)
	ErrUserPrivilegeChange = NewErr(10006, "Permission changed", http.StatusUnauthorized)
	ErrLockNotAcquire      = NewErr(10007, "Lock not released")
	ErrLockAcquire         = NewErr(10008, "Lock acquisition error")
	ErrLockNotRelease      = NewErr(10009, "Lock is not released")
	ErrLockRelease         = NewErr(10010, "Lock released err")
	ErrTWitterAddress      = NewErr(10011, "Twitter address illeage")
	ErrDiscordAddress      = NewErr(10012, "Discord address illeage")
	ErrAddress             = NewErr(10013, "Address illeage")

	ErrUserNotExist      = NewErr(10014, "User not exist")
	ErrCampaignNotStart  = NewErr(10015, "Campaign Not Start")
	ErrCampaignEnded     = NewErr(10016, "Campaign Ended")
	ErrTwitterNotBind    = NewErr(10017, "Get \"\": unsupported protocol scheme \"\"")
	ErrTokenRefreshFaild = NewErr(10018, "refresh token error")
)

var codeToErr = map[uint32]*Err{
	200: NoErr,

	7000: ErrCustom,
	7777: ErrUnexpected,

	9999:  ErrTokenNotValidYet,
	10000: ErrInvalidUrl,
	10001: ErrInvalidHeader,
	10002: ErrInvalidParams,
	10003: ErrTokenVerify,
	10004: ErrTokenExpire,
	10005: ErrUserLogin,
	10006: ErrUserPrivilegeChange,
	10007: ErrLockNotAcquire,
	10008: ErrLockAcquire,
	10009: ErrLockNotRelease,
	10010: ErrLockRelease,
	10011: ErrTWitterAddress,
	10012: ErrDiscordAddress,
	10013: ErrAddress,
	10014: ErrUserNotExist,
	10015: ErrCampaignNotStart,
	10016: ErrCampaignEnded,
	10017: ErrTwitterNotBind,
	10018: ErrTokenRefreshFaild,
}

// NewErr 创建新的业务错误
func NewErr(code uint32, msg string, httpCode ...int) *Err {
	hc := http.StatusOK
	if len(httpCode) != 0 {
		hc = httpCode[0]
	}

	return &Err{code: code, httpCode: hc, msg: msg}
}

func GetCodeToErr() map[uint32]*Err {
	return codeToErr
}

func SetCodeToErr(code uint32, err *Err) error {
	if _, ok := codeToErr[code]; ok {
		return errors.New("has exist")
	}

	codeToErr[code] = err
	return nil
}

// NewCustomErr 创建新的自定义错误
func NewCustomErr(msg string, httpCode ...int) *Err {
	return NewErr(CodeCustom, msg, httpCode...)
}

// IsErr 判断是否为业务错误
func IsErr(err error) bool {
	if err == nil {
		return true
	}

	_, ok := err.(*Err)
	return ok
}

// ParseErr 解析业务错误
func ParseErr(err error) *Err {
	if err == nil {
		return NoErr
	}

	if e, ok := err.(*Err); ok {
		return e
	}

	s, _ := status.FromError(err)
	c := uint32(s.Code())
	if c == CodeCustom {
		return NewCustomErr(s.Message())
	}

	return ParseCode(c)
}

// ParseCode 解析业务状态码对应的业务错误
func ParseCode(code uint32) *Err {
	if e, ok := codeToErr[code]; ok {
		return e
	}

	return ErrUnexpected
}

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

// MD5加密
// plaintext string	明文
// string	密文，32个字符16进制string
func MD5(plaintext string) string {
	m := md5.New()
	_, err := io.WriteString(m, plaintext)
	if err != nil {
		config.Logger.Fatal("[MD5]", zap.Error(err))
	}
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}

// 两次MD5加密
// plaintext string	明文
// string	密文，32个字符16进制string
func MD5Double(plaintext string) string {
	return MD5(MD5(plaintext))
}

func GetDbNotExistMsg(table string) (newS string) {

	c := cases.Title(language.Und)
	tableS := strings.Split(table, "_")
	for i, v := range tableS {
		if i == 0 {
			continue
		}
		newS += c.String(v)
	}
	newS += " not exist"
	return
}
