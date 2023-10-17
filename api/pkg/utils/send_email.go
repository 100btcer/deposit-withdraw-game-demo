package utils

import (
	"fmt"
	"math/rand"
	"mm-ndj/config"
	"net/smtp"
	"os"
	"time"

	"github.com/jordan-wright/email"
	"go.uber.org/zap"
)

var count int

type (
	SendEmail interface {
		SendEamil() (string, error)
	}

	Email struct {
		To   string `json:"to"`
		Code string `json:"code"`
		Ty   int    `json:"ty"`
	}

	EmailService struct {
		ServiceName string //邮箱服务器名称+端口
		Email       string //发送邮箱
		Password    string //邮箱密码或授权码
		EmService   string //邮箱服务器
	}
)

func InitEmail(to, code string, ty int) SendEmail {
	return &Email{
		To:   to,
		Code: code,
		Ty:   ty,
	}
}

func (em *Email) SendEamil() (string, error) {
	email_list := emailService()
	env := os.Getenv(config.Env)
	index := randInt(0, len(email_list))
	if env == config.EnvTest {
		index = 1
	}
	eml := email_list[index]
	e := email.NewEmail()
	to := []string{}
	to = append(to, em.To)
	e.Subject = config.Conf.EmailC.Subject
	e.From = fmt.Sprintf("Staland <%s>", eml.Email)
	e.To = to
	text := GetEmailText(fmt.Sprint(em.Ty), em.Code)
	e.HTML = []byte(text)
	if err := e.Send(eml.ServiceName, smtp.PlainAuth("", eml.Email, eml.Password, eml.EmService)); err != nil {
		config.Logger.Error("[ERROR]", zap.Error(err))
		if count <= 3 {
			count++
			em.SendEamil()
		} else {
			return "", err
		}
	}
	config.Logger.Info("[INFO]", zap.String("send email code", "success~"))
	return eml.Email, nil
}

func emailService() []EmailService {
	return []EmailService{
		{"smtp.qq.com:25", "3230806479@qq.com", "aurnrwlvqsdychbg", "smtp.qq.com"},
		{"smtp.gmail.com:587", "noreply@dehero.co", "lmzpwdtzeygsrypr", "smtp.gmail.com"},
	}
}

// 随机数
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func GetVerifyCode() int {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := rnd.Intn(100000)
	if code < 100000 {
		code = code + 600000
	}

	return code
}
