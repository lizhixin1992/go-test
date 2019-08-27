package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

//发送有问题

const (
	EmailTo   = "18500950209@163.com" //发送给谁
	EmailFrom = "510717270@qq.com"    //谁发的
	EmailPass = "cjxhvryczyuebihj"    //密码
	EmailHost = "smtp.qq.com"         //一般是25端口
	EmailPort = "25"                  //一般是25端口
)

type loginAuth struct {
	username, password, host string
}

func LoginAuth(username, password, host string) smtp.Auth {
	return &loginAuth{username, password, host}
}

//需要使用Login作为参数
func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", nil, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	command := string(fromServer)
	command = strings.TrimSpace(command)
	command = strings.TrimSuffix(command, ":")
	command = strings.ToLower(command)

	if more {
		if command == "username" {
			return []byte(fmt.Sprintf("%s", a.username)), nil
		} else if command == "password" {
			return []byte(fmt.Sprintf("%s", a.password)), nil
		} else if command == "host" {
			return []byte(fmt.Sprintf("%s", a.host)), nil
		} else {
			// We've already sent everything.
			return nil, fmt.Errorf("unexpected server challenge: %s", command)
		}
	}
	return nil, nil
}

func SendEmail(subject, body string) error {
	send_to := strings.Split(EmailTo, ";")
	content_type := "Content-Type: text/plain; charset=UTF-8"
	msg := []byte("To: All \r\nFrom: " + EmailFrom + " >\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)

	//auth := LoginAuth(EmailFrom, EmailPass,EmailHost)
	auth := smtp.PlainAuth("", "510717270@qq.com", "cjxhvryczyuebihj", "smtp.qq.com")
	err := smtp.SendMail(EmailHost+":"+EmailPort, auth, EmailFrom, send_to, msg)
	return err
}

func main() {
	SendEmail("mail测试", "go实现mail测试")
}
