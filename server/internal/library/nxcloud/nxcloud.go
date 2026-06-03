package nxcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
)

var (
	HostUrl      = "https://api2.nxcloud.com/api/sms/mtsend"
	appKey       = "vz0pFeCL"
	secretKey    = "08cYqmas"
	accessKey    = "0CfRCfHS"
	accessSecret = "DKH2RWa0BTn0"
)

type SmsRes struct {
	Result    string `json:"result"`
	Messageid string `json:"messageid"`
	Code      string `json:"code"`
}

func SendSms(ctx context.Context, mobile string, content string) (err error) {
	var (
		SmsReq      *gclient.Response
		smsRes      SmsRes
		body        string
		ghttpclient = g.Client()
	)

	ghttpclient.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	if SmsReq, err = ghttpclient.Post(ctx, HostUrl, g.MapStrAny{
		"appkey":    appKey,
		"secretkey": secretKey,
		"phone":     mobile,
		"content":   fmt.Sprintf("【%s】%s", "Apartment11", content),
	}); err != nil {
		return
	}
	defer SmsReq.Close()
	body = SmsReq.ReadAllString()
	g.Log().Path("logs/SDK/NX_CLOUD").Info(ctx, SmsReq.Raw())
	g.Log().Path("logs/SDK/NX_CLOUD").Info(ctx, body)
	if err = json.Unmarshal([]byte(body), &smsRes); err != nil {
		return
	}
	if smsRes.Code != "0" {
		err = gerror.New(smsRes.Result)
	}
	return
}
