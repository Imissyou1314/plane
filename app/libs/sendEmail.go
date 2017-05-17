package libs

import (
	"net/smtp"
	"strings"
)

// err := SendMail("1442919817@qq.com", "zfyxxxxxx252078", "smtp.qq.com:25", username+"@qq.com", "找回密码",
// "点击这里修改密码：http://www.xxx.com/updatepass?username="+username+"&time="+time.Now().Format("2006-01-0215:04:05"), "html")

// SendMail 发送邮件
func SendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}
