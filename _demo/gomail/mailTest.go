package main

import (
	"github.com/jordan-wright/email"
	"net/smtp"
	"net/textproto"
)

func main() {
	e := &email.Email{
		To:      []string{"li.zhixin@chinaott.net"},
		From:    "Jordan Wright <510717270@qq.com>",
		Subject: "Awesome Subject",
		Text:    []byte("Text Body is, of course, supported!"),
		HTML:    []byte("<h1>Fancy HTML is supported, too!</h1>"),
		Headers: textproto.MIMEHeader{},
	}
	e.Send("smtp.qq.com:25", smtp.PlainAuth("", "510717270@qq.com", "cjxhvryczyuebihj", "smtp.qq.com"))
}
