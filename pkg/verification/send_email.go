package verification

import (
	"bytes"
	"crypto/tls"
	"html/template"
	"log"
	"math/rand/v2"
	"net/smtp"
	"strconv"

	"github.com/spf13/viper"
)

// 注册邮件模板
const RegisterTemplate = `<!DOCTYPE html>
<html>
<body>
    <p>你好,你正在注册IM即时通讯：</p>
    <p>你的注册验证码是：</p>
    <div style="font-size:24px; font-weight:bold;">{{.Code}}</div>
	<p> 如果这不是你本人的操作，请忽略该邮件 </p>
</body>
</html>
`

// 密码修改模板
const ResetTemplate = `<!DOCTYPE html>
<html>
<body>
    <p>你好,你正在修改叮当账号的密码：</p>
    <p>你的验证码是：</p>
    <div style="font-size:24px; font-weight:bold;">{{.Code}}</div>
	<p> 如果这不是你本人的操作，请忽略该邮件 </p>
</body>
</html>
`

// 找回密码模板
const RetrieveTemplate = `<!DOCTYPE html>
<html>
<body>
    <p>你好,你正在找回叮当账号的密码：</p>
    <p>你的验证码是：</p>
    <div style="font-size:24px; font-weight:bold;">{{.Code}}</div>
	<p> 如果这不是你本人的操作，请忽略该邮件 </p>
</body>
</html>
`

type templateData struct {
	Code string `json:"code"`
}

// 邮件模式
const (
	RegisterMode = "register"
	ResetMode    = "reset"
	RetrieveMode = "retrieve"
)

// 使用qq邮箱时需要设置tls加密，否则就会出现发送验证码成功但依旧报错的情况
// 而使用网易163邮箱可以直接使用jordan-wright/email包，同时也无需加密

// 这里使用smtp库并显式启用tls加密
// 发送验证码邮件
func SendCode(userEmail string, code string, mode string) error {
	// 处理邮件html模板
	// 解析模板:注册、重置、找回
	var tpl *template.Template
	var err error
	if mode == RegisterMode {
		tpl, err = template.New("verfication_email").Parse(RegisterTemplate)
	} else if mode == ResetMode { // Reset Pwd
		tpl, err = template.New("verfication_email").Parse(ResetTemplate)
	} else { //Retrieve Pwd
		tpl, err = template.New("verfication_email").Parse(RetrieveTemplate)
	}
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

	// 使用smtp库:qq邮箱
	sender := viper.GetString("smtp.sender")
	authcode := viper.GetString("smtp.authcode")
	// 设置plainAuth
	auth := smtp.PlainAuth("", sender, authcode, "smtp.qq.com")
	// 邮件内容
	from := "From: Zhenqiiii <" + sender + ">\r\n"
	to := "To: " + userEmail + "\r\n"
	subject := "Subject: IM邮箱验证码 \r\n"
	contentType := "Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n"
	content := body.Bytes()
	msg := []byte(from + to + subject + contentType + string(content))

	// 发送
	// 创建 tls 配置
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "smtp.qq.com",
	}
	conn, err := tls.Dial("tcp", "smtp.qq.com:465", tlsconfig)
	if err != nil {
		log.Printf("TLS连接失败: %v\n", err)
		return err
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, "smtp.qq.com")
	if err != nil {
		log.Printf("smtp客户端创建失败: %v\n", err)
		return err
	}
	defer client.Quit()

	// 使用 auth 进行认证
	if err = client.Auth(auth); err != nil {
		log.Printf("邮箱smtp认证失败: %v\n", err)
		return err
	}

	// 设置发件人和收件人
	if err = client.Mail(sender); err != nil {
		log.Printf("发件人设置失败: %v\n", err)
		return err
	}
	if err = client.Rcpt(userEmail); err != nil {
		log.Printf("收件人设置失败: %v\n", err)
		return err
	}

	// 写入邮件内容
	wc, err := client.Data()
	if err != nil {
		log.Printf("数据写入失败: %v\n", err)
		return err
	}
	defer wc.Close()

	_, err = wc.Write(msg)
	if err != nil {
		log.Printf("邮件发送失败: %v\n", err)
		return err
	}

	return nil

	// 以下是使用jordan-wright/email包的邮件发送代码,但qq邮箱使用不了,163邮箱可以直接使用
	// // 读取配置
	// user := viper.GetString("smtp.sender")
	// authcode := viper.GetString("smtp.authcode")
	// // 生成邮件实例
	// e := email.NewEmail()
	// // sender
	// e.From = "Zhenqiiii <" + sender + "> "
	// // receiver:注册用户
	// e.To = []string{userEmail}
	// // subject：主题
	// e.Subject = "IM邮箱验证码"
	// // html内容:使用template
	// e.HTML = body.Bytes()

	// // 发送邮件:使用自己的邮箱授权码
	// err = e.Send("smtp.163.com:25", smtp.PlainAuth("", sender, authcode, "smtp.163.com"))
	// if err != nil {
	// 	log.Printf("邮件发送失败: %v\n", err)
	// 	return err
	// }
	// return nil

}

// 生成随机验证码
func GenCode() string {
	// 固定seed
	code := strconv.FormatInt(rand.Int64N(100000)+100000, 10)
	return code
}
