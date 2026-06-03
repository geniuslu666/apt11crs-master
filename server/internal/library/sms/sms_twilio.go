package sms

import (
	"APT/internal/model/input/input_basics"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
	"os"
	"strings"
)

var (
	FromPhone = os.Getenv("TWILIO_FROM_PHONE")
	Username  = os.Getenv("TWILIO_ACCOUNT_SID")
	Password  = os.Getenv("TWILIO_AUTH_TOKEN")
)

type TwilioDrive struct{}

func (d *TwilioDrive) SendCode(ctx context.Context, in *input_basics.SendCodeInp) (err error) {
	var (
		SendContent []byte
	)
	// 您的动态验证码为：{1}，请勿向任何人提供此验证码，如非本人操作，请忽略本短信！
	content := strings.Replace(in.Template, "{1}", in.Code, -1)
	SendContent = []byte(content)

	if err = d.Send(ctx, in.Mobile, string(SendContent)); err != nil {
		return
	}
	return
}

func (d *TwilioDrive) SendMsg(ctx context.Context, in *input_basics.SendMsgInp) (err error) {
	if err = d.Send(ctx, in.Mobile, in.Content); err != nil {
		return
	}
	return
}

func (d *TwilioDrive) Send(ctx context.Context, toPhone string, Content string) (err error) {
	// Find your Account SID and Auth Token at twilio.com/console
	// and set the environment variables. See http://twil.io/secure
	// Make sure TWILIO_ACCOUNT_SID and TWILIO_AUTH_TOKEN exists in your environment
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: Username,
		Password: Password,
	})

	params := &api.CreateMessageParams{}
	params.SetBody(Content)
	params.SetFrom(FromPhone)
	params.SetTo(toPhone)
	g.Log().Info(ctx, Content)
	_, err = client.Api.CreateMessage(params)

	if err != nil {
		return
	}
	return
}
