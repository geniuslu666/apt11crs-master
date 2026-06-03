package logic_warning

import (
	"APT/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
)

type sWarningWorkWx struct{}

func NewWarningWorkWx() *sWarningWorkWx {
	return &sWarningWorkWx{}
}

func init() {
	service.RegisterWarningWorkWx(NewWarningWorkWx())
}

type SendParams struct {
	Msgtype  string   `json:"msgtype"`
	Markdown Markdown `json:"markdown"`
}
type Markdown struct {
	Content string `json:"content"`
}

func (s *sWarningWorkWx) SendWarningWorkWx(ctx context.Context, Content string) (err error) {
	var (
		ghttpclient   = g.Client()
		WorkWxHookUrl = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=5d4bd07d-deeb-42cf-8666-fce2e8b29e34"
		response      *gclient.Response
	)
	ghttpclient = ghttpclient.SetHeader("Content-Type", "application/json")
	if response, err = ghttpclient.Post(ctx, WorkWxHookUrl, &SendParams{
		Msgtype:  "markdown",
		Markdown: Markdown{Content: Content},
	}); err != nil {
		return
	}
	g.Log().Path("logs/HOOK/WOEKWX_HOOK").Info(ctx, response.Raw())
	return
}
