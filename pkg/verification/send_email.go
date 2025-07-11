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

// 封装jordan-wright/email包，发送验证码邮件
func SendCode(userEmail string, code string) error {
	// 处理邮件html模板
	// 解析模板
	tpl, err := template.New("verfication_email").Parse(emailTemplate)
	if err != nil {
		log.Println("解析模板失败: " + err.Error())
		return err
	}
	// 渲染模板,渲染后的结果会返回给body
	var body bytes.Buffer
	if err = tpl.Execute(&body, code); err != nil {
		log.Println("渲染模板失败: " + err.Error())
		return err
	}

	// 生成邮件实例
	e := email.NewEmail()
	// sender
	e.From = "Zhenqiii <noreply@foxmail.com> "
	// receiver:注册用户
	e.To = []string{userEmail}
	// subject：主题
	e.Subject = "IM邮箱验证码"
	// html内容:使用template
	e.HTML = body.Bytes()

	// 发送邮件:使用自己的qq邮箱授权码
	user := viper.GetString("smtp.user")
	authcode := viper.GetString("smtp.auth")
	err = e.Send("smtp.qq.com:25", smtp.PlainAuth("", user, authcode, "smtp.qq.com"))
	if err != nil {
		log.Println("邮件发送失败: " + err.Error())
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
