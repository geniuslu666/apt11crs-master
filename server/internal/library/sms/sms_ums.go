package sms

import (
	"APT/internal/model/input/input_basics"
	"APT/utility/uuid"
	"bytes"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/url"
)

var (
	SpCode    = "278371"
	AppKey    = "sz_ybj"
	AppSecret = "01a4c27693cceee6f46364bd17116537"
	Domain    = "https://api.ums86.com:9600"
)

type UmsDrive struct{}

func (d *UmsDrive) SendCode(ctx context.Context, in *input_basics.SendCodeInp) (err error) {
	var (
		ghttpclient = g.Client()
		Response    *gclient.Response
		SendContent []byte
	)

	if SendContent, err = Utf8ToGbk([]byte(fmt.Sprintf("您的动态验证码为：%s，请勿向任何人提供此验证码，如非本人操作，请忽略本短信！", in.Code))); err != nil {
		return
	}

	// 构建请求体
	requestBody := map[string]interface{}{
		"SpCode":         SpCode,
		"MessageContent": string(SendContent),
		"templateId":     in.Template,
		"UserNumber":     in.Mobile[3:],
		"LoginName":      AppKey,
		"Password":       AppSecret,
		"SerialNumber":   uuid.CreateOrderCode("S"),
	}

	ghttpclient = ghttpclient.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	if Response, err = ghttpclient.Post(ctx, Domain+"/sms/Api/Send.do", requestBody); err != nil {
		return
	}
	defer Response.Close()
	g.Log().Debug(ctx, Response.Raw())

	// 使用 url.ParseQuery 解析字符串
	params, err := url.ParseQuery(Response.ReadAllString())
	if err != nil {
		fmt.Println("发送短信失败")
		return
	}
	if params.Get("result") != "0" {
		return fmt.Errorf("发送验证码失败: %s", params.Get("description"))
	}
	return
}

func (d *UmsDrive) SendMsg(ctx context.Context, in *input_basics.SendMsgInp) (err error) {
	var (
		ghttpclient = g.Client()
		Response    *gclient.Response
		SendContent []byte
	)

	g.Log().Info(ctx, in.Content)

	if SendContent, err = Utf8ToGbk([]byte(in.Content)); err != nil {
		return
	}

	// 构建请求体
	requestBody := map[string]interface{}{
		"SpCode":         SpCode,
		"MessageContent": string(SendContent),
		"templateId":     in.Template,
		"UserNumber":     in.Mobile[3:],
		"LoginName":      AppKey,
		"Password":       AppSecret,
		"SerialNumber":   uuid.CreateOrderCode("S"),
	}

	ghttpclient = ghttpclient.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	if Response, err = ghttpclient.Post(ctx, Domain+"/sms/Api/Send.do", requestBody); err != nil {
		return
	}
	defer Response.Close()
	g.Log().Debug(ctx, Response.Raw())

	// 使用 url.ParseQuery 解析字符串
	params, err := url.ParseQuery(Response.ReadAllString())
	if err != nil {
		fmt.Println("发送短信失败")
		return
	}
	if params.Get("result") != "0" {
		return fmt.Errorf("发送验证码失败: %s", params.Get("description"))
	}
	return
}

// GbkToUtf8 GBK 转 UTF-8
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// Utf8ToGbk UTF-8 转 GBK
func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
