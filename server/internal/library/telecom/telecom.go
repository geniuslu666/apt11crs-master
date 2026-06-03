package telecom

import (
	"context"
	"encoding/xml"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

//userid	企业id	企业ID
//account	发送用户帐号	用户帐号，由系统管理员
//password	发送帐号密码	用户账号对应的密码
//mobile	全部被叫号码	发信发送的目的号码.多个号码之间用半角逗号隔开
//content	发送内容	短信的内容，内容需要UTF-8编码
//sendTime	定时发送时间	为空表示立即发送，定时发送格式2010-10-24 09:08:10
//action	发送任务命令	设置为固定的:send
//extno	扩展子号	请先询问配置的通道是否支持扩展子号，如果不支持，请填空。子号只能为数字，且最多10位数。

var (
	URL      = "http://60.205.206.26:8888/sms.aspx"
	USERID   = "1200"
	ACCOUNT  = "yeebok"
	PASSWORD = "yeebok@123"
	ACTION   = "send"
)

func SendSms(ctx context.Context, mobile string, content string) (err error) {
	var (
		res       string
		returnsms Returnsms
	)
	if res, err = CookieCurlGet(ctx, URL, g.MapStrAny{
		"userid":   USERID,
		"action":   ACTION,
		"account":  ACCOUNT,
		"password": PASSWORD,
		"mobile":   mobile,
		"content":  content,
	}); err != nil {
		return
	}

	if err = xml.Unmarshal([]byte(res), &returnsms); err != nil {
		return
	}
	if returnsms.Returnstatus != "Success" {
		g.Log().Path("logs/SDK/TELECOM").Error(ctx, returnsms.Message)
		err = gerror.New("短信发送失败")
	}
	return
}

type Returnsms struct {
	XMLName       xml.Name `xml:"returnsms"`
	Text          string   `xml:",chardata"`
	Returnstatus  string   `xml:"returnstatus"`
	Message       string   `xml:"message"`
	Remainpoint   string   `xml:"remainpoint"`
	TaskID        string   `xml:"taskID"`
	SuccessCounts string   `xml:"successCounts"`
}

func CookieCurlGet(ctx context.Context, path string, params interface{}) (res string, err error) {
	gclient := g.Client()
	gclient.SetHeader("Accept", "*/*")
	r, err := gclient.Get(ctx, path, params)
	if err != nil {
		return
	}
	defer r.Close()
	g.Log().Path("logs/SDK/TELECOM").Info(ctx, err)
	g.Log().Path("logs/SDK/TELECOM").Infof(ctx, r.Raw())
	res = r.ReadAllString()
	return
}
