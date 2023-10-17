package dehero

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"mm-ndj/config"
	"time"

	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
)

func DoSendEmail(email string, t int) (int, error) {
	var start, end int64
	start = time.Now().UnixMilli()
	var template_verify_code string
	code := GenVerifyCode(email)
	if t == 1 { //英文
		template_verify_code = fmt.Sprintf("<span style=\"font-weight: bold\"> Welcome to DeHeroGame! Your verify code is %v. </span> ", code)
	}
	if t == 2 { //泰语
		template_verify_code = fmt.Sprintf("<span style=\"font-weight: bold\"> ยินดีต้อนรับสู่ dehero รหัสยืนยันของคุณคือ %v.</span>", code)
	}

	if t == 3 { //葡萄牙语
		template_verify_code = fmt.Sprintf("<span style=\"font-weight: bold\"> Bem-vindo ao dehero O seu código de verificação é %v.</span>", code)
	}

	if t == 4 { //印尼语
		template_verify_code = fmt.Sprintf("<span style=\"font-weight: bold\"> Selamat datang di dehero  kode verifikasi Anda adalah %v. </span>", code)
	}

	//Gmail 邮箱：
	host := "smtp.gmail.com"
	port := 465
	userName := "noreply@dehero.co"
	app_specific_password := "lmzpwdtzeygsrypr"
	m := gomail.NewMessage()
	m.SetHeader("From", "DeHeroGame"+"<"+userName+">") // 增加发件人别名
	m.SetHeader("To", email)                           // 收件人，可以多个收件人，但必须使用相同的 SMTP 连接
	m.SetHeader("Subject", "DeHeroGame")               // 邮件主题
	m.SetBody("text/html", fmt.Sprintf(template_verify_code))

	d := gomail.NewDialer(
		host,
		port,
		userName,
		app_specific_password,
	)

	// 关闭SSL协议认证
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return 0, err
	}
	end = time.Now().UnixMilli()
	backend := end - start
	config.Logger.Info("send message ", zap.Any("发送邮件消耗时间:", backend))
	return code, nil
}

func GenVerifyCode(email string) int {

	//rand.Seed(time.Now().UnixNano())
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := rnd.Intn(10000)
	if code < 1000 {
		code = code + 6000
	}

	return code
}
