package verification

import (
	"bytes"
	"html/template"
	"log"
	"math/rand/v2"
	"net/smtp"
	"strconv"

	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
)

// 注册邮件模板
const emailTemplate = `<!DOCTYPE html>
<html>
<body>
    <p>你好,你正在注册IM即时通讯：</p>
    <p>你的注册验证码是：</p>
    <div style="font-size:24px; font-weight:bold;">{{.Code}}</div>
</body>
</html>
`

type templateData struct {
	Code string `json:"code"`
}

// 发送验证码邮件，使用jordan-wright/email包
// 但是这样用自己的邮箱发送，一定时间内次数多了会被限制
func SendCode(userEmail string, code string) error {
	// 处理邮件html模板
	// 解析模板
	tpl, err := template.New("verfication_email").Parse(emailTemplate)
	if err != nil {
		log.Println("解析模板失败: " + err.Error())
		return err
	}
	// 渲染模板,渲染后的结果会返回给body
	data := templateData{
		Code: code,
	}
	var body bytes.Buffer
	if err = tpl.Execute(&body, data); err != nil {
		log.Println("渲染模板失败: " + err.Error())
		return err
	}

	// // 使用smtp库
	// user := viper.GetString("smtp.user")
	// authcode := viper.GetString("smtp.authcode")
	// // 设置plainAuth
	// auth := smtp.PlainAuth("", user, authcode, "smtp.qq.com")
	// // 收件人
	// receiver := []string{userEmail}
	// // 邮件内容
	// from := "From: Zhenqiiii <" + user + ">\r\n"
	// to := "To: " + userEmail + "\r\n"
	// subject := "Subject: IM邮箱验证码 \r\n"
	// contentType := "Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n"
	// content := body.Bytes()
	// msg := []byte(from + to + subject + contentType + string(content))

	// // 发送
	// err = smtp.SendMail("smtp.qq.com:465", auth, user, receiver, msg)
	// if err != nil {
	// 	log.Printf("邮件发送失败： %v\n", err)
	// 	return err
	// }

	// 读取配置
	user := viper.GetString("smtp.user")
	authcode := viper.GetString("smtp.authcode")
	// 生成邮件实例
	e := email.NewEmail()
	// sender
	e.From = "Zhenqiiii <" + user + "> "
	// receiver:注册用户
	e.To = []string{userEmail}
	// subject：主题
	e.Subject = "IM邮箱验证码"
	// html内容:使用template
	e.HTML = body.Bytes()

	// 发送邮件:使用自己的邮箱授权码
	err = e.Send("smtp.163.com:25", smtp.PlainAuth("", user, authcode, "smtp.163.com"))
	// e.SendWithTLS("smtp.qq.com:465",
	// 	smtp.PlainAuth("", user, authcode, "smtp.qq.com"),
	// 	&tls.Config{
	// 		InsecureSkipVerify: true, ServerName: "smtp.qq.com",
	// 	})
	if err != nil {
		log.Printf("邮件发送失败: %v\n", err)
		return err
	}
	return nil

}

// 生成随机验证码
func GenCode() string {
	// 固定seed
	code := strconv.FormatInt(rand.Int64N(100000)+100000, 10)
	return code
}
