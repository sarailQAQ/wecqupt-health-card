package wecqupt_health_card

import (
	"github.com/go-gomail/gomail"
	"strconv"
	"strings"
)

func SendMail(subject, body string, mailInfo EmailConfig) error {
	if strings.ToLower(mailInfo.Enable) == "false" || mailInfo.Enable == "0" {
		return nil
	}

	m := gomail.NewMessage()
	m.SetHeader("To", mailInfo.Address)                   // 收件人
	m.SetAddressHeader("From", mailInfo.Address, "打卡助手菌") // 发件人
	m.SetHeader("Subject", subject)                       // 主题
	m.SetBody("text/html", body)                          // 正文

	port, err := strconv.Atoi(mailInfo.Port)
	if err != nil {
		return err
	}

	d := gomail.NewDialer(mailInfo.Host, port, mailInfo.Address, mailInfo.Password)
	err = d.DialAndSend(m)

	return err
}
