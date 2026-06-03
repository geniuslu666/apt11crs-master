package airhousePublicApi

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	Host         string
	XAHPartner   = "033e644c-b2f1-4f94-ae5a-54b2c09c0488"
	XAccessToken = "0a4c363b46bfd6a96cb09b195ac8c55a"
	LANG         = "ja"
)

func init() {
	var (
		mode      = g.Cfg().MustGet(gctx.New(), "airhousePublicApi.mode").String()
		ConfigMap g.MapStrStr
	)
	if mode == "develop" {
		ConfigMap = g.Cfg().MustGet(gctx.New(), "airhousePublicApi.develop").MapStrStr()
	} else if mode == "produce" {
		ConfigMap = g.Cfg().MustGet(gctx.New(), "airhousePublicApi.produce").MapStrStr()
	} else {
		panic(gerror.New("airhostPublicApi mode error"))
	}
	Host = ConfigMap["host"]
	XAHPartner = ConfigMap["X-AH-Partner"]
	XAccessToken = ConfigMap["X-Access-Token"]

}

// CurlAirHostPublicApiGet get方式请求
func CurlAirHostPublicApiGet(ctx context.Context, path string, params interface{}) (res string, err error, httpStatus int) {
	gclient := g.Client()
	gclient.SetHeader("Accept", "*/*")
	gclient.SetHeader("X-AH-Partner", XAHPartner)
	gclient.SetHeader("X-Access-Token", XAccessToken)
	gclient.SetHeader("Accept-Language", LANG)
	r, err := gclient.Get(ctx, Host+path, params)
	if err != nil {
		return
	}
	defer r.Close()
	httpStatus = r.StatusCode
	g.Log().Path("logs/SDK/AIRHOST_PUBLIC_SDK").Infof(ctx, r.Raw())
	if httpStatus == 429 {
		err = gerror.New("429 Too Many Requests")
		return
	}
	res = r.ReadAllString()
	if !g.IsEmpty(gjson.New(res).Get("error").String()) {
		// 做企业微信群通知
		_ = SendWarningWorkWx(ctx, `
		### 系统异常 AIRHOST API
		*  链路ID：`+gctx.CtxId(ctx)+`
		*  请求地址：`+Host+path+`
		*  请求参数：`+gjson.New(params).String()+`
		*  响应内容：`+res)
	}
	return
}

// CurlAirHostPublicApiPost post方式请求
func CurlAirHostPublicApiPost(ctx context.Context, path string, params interface{}) (res string, err error, httpStatus int) {
	gclient := g.Client()
	gclient.SetHeader("Accept", "*/*")
	gclient.SetHeader("X-AH-Partner", XAHPartner)
	gclient.SetHeader("X-Access-Token", XAccessToken)
	gclient.SetHeader("Accept-Language", LANG)
	gclient.SetHeader("Content-Type", "application/json")
	r, err := gclient.Post(ctx, Host+path, params)
	if err != nil {
		return
	}
	defer r.Close()
	httpStatus = r.StatusCode
	g.Log().Path("logs/SDK/AIRHOST_PUBLIC_SDK").Infof(ctx, r.Raw())
	res = r.ReadAllString()
	if !g.IsEmpty(gjson.New(res).Get("error").String()) {
		// 做企业微信群通知
		_ = SendWarningWorkWx(ctx, `
		### 系统异常 AIRHOST API
		*  链路ID：`+gctx.CtxId(ctx)+`
		*  请求地址：`+Host+path+`
		*  请求参数：`+gjson.New(params).String()+`
		*  响应内容：`+res)
	}
	return
}

// CurlAirHostPublicApiPut put方式请求
func CurlAirHostPublicApiPut(ctx context.Context, path string, params interface{}) (res string, err error, httpStatus int) {
	gclient := g.Client()
	gclient.SetHeader("Accept", "*/*")
	gclient.SetHeader("X-AH-Partner", XAHPartner)
	gclient.SetHeader("X-Access-Token", XAccessToken)
	gclient.SetHeader("Accept-Language", LANG)
	r, err := gclient.ContentJson().Put(ctx, Host+path, params)
	if err != nil {
		return
	}
	defer r.Close()
	httpStatus = r.StatusCode
	g.Log().Path("logs/SDK/AIRHOST_PUBLIC_SDK").Infof(ctx, r.Raw())
	res = r.ReadAllString()
	if !g.IsEmpty(gjson.New(res).Get("error").String()) {
		// 做企业微信群通知
		_ = SendWarningWorkWx(ctx, `
		### 系统异常 AIRHOST API
		*  链路ID：`+gctx.CtxId(ctx)+`
		*  请求地址：`+Host+path+`
		*  请求参数：`+gjson.New(params).String()+`
		*  响应内容：`+res)
	}
	return
}

type SendParams struct {
	Msgtype  string   `json:"msgtype"`
	Markdown Markdown `json:"markdown"`
}
type Markdown struct {
	Content string `json:"content"`
}

func SendWarningWorkWx(ctx context.Context, Content string) (err error) {
	var (
		ghttpclient   = g.Client()
		WorkWxHookUrl = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=9a15e3cf-2938-4832-b43d-d2fb01238530"
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
