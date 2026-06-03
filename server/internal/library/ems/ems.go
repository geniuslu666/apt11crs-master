// Package ems
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package ems

import (
	"APT/internal/model"
	"APT/utility/validate"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
	"gopkg.in/gomail.v2"
	"net/smtp"
	"strconv"
	"strings"
)

// Send 发送邮件入口
func Send(config *model.EmailConfig, to string, subject string, body string, attachmentPath string) error {
	//return sendToMail(config, to, subject, body, "html")
	return SendToSendCloudMail(config, []string{to}, subject, body, attachmentPath)
}

func sendToMail(config *model.EmailConfig, to, subject, body, mailType string) error {
	if config == nil {
		return gerror.New("邮件配置不能为空")
	}

	var (
		contentType string
		auth        = smtp.PlainAuth("", config.User, config.Password, config.Host)
		sendTo      = strings.Split(to, ";")
	)

	if len(sendTo) == 0 {
		return gerror.New("收件人不能为空")
	}

	for _, em := range sendTo {
		if !validate.IsEmail(em) {
			return gerror.Newf("邮件格式不正确，请检查：%v", em)
		}
	}

	if mailType == "html" {
		contentType = "Content-Type: text/" + mailType + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + config.SendName + "<" + config.User + ">" + "\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	return smtp.SendMail(config.Addr, auth, config.User, sendTo, msg)
}

func SendToSendCloudMail(config *model.EmailConfig, mailTo []string, subject string, body string, attachmentPath string) error {
	mailConn := map[string]string{
		"apiUser": config.User,     //apiuser
		"apiKey":  config.Password, //apikey
		"host":    config.Host,
		"port":    gvar.New(config.Port).String(),
		"from":    config.AdminMailbox, //The sender, use the correct email address instead
	}
	port, _ := strconv.Atoi(mailConn["port"]) //port type use int

	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(mailConn["from"], config.SendName)) //In this way, you can add aliases, that is, "fromname". If there are special characters such as Chinese, pay attention to character encoding
	m.SetHeader("To", mailTo...)
	m.SetHeader("Subject", subject) //Set message subject
	m.SetBody("text/html", body)    //Set message body
	if !g.IsEmpty(attachmentPath) {
		m.Attach(attachmentPath, gomail.Rename(guid.S()+".pdf")) //Attachment delivery
	}

	d := gomail.NewDialer(mailConn["host"], port, mailConn["apiUser"], mailConn["apiKey"])

	err := d.DialAndSend(m)
	return err
}
