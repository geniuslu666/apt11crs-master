// Package middleware
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package middleware

import (
	"APT/internal/global"
	"APT/internal/library/contexts"
	"APT/internal/model"
	"APT/internal/service"
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"mime"
	"net/http"
	"reflect"
)

//type DefaultHandlerResponse struct {
//	Code      int         `json:"code"      dc:"Error code"`
//	Message   string      `json:"message"   dc:"Error message"`
//	Data      interface{} `json:"data"      dc:"Result data for certain request according API definition"`
//	Timestamp int64       `json:"timestamp" dc:"服务器时间戳"`
//	TraceID   string      `json:"traceID"   dc:"链路ID"`
//}

const (
	contentTypeEventStream  = "text/event-stream"
	contentTypeOctetStream  = "application/octet-stream"
	contentTypeMixedReplace = "multipart/x-mixed-replace"
)

var (
	streamContentType = []string{contentTypeEventStream, contentTypeOctetStream, contentTypeMixedReplace}
)

// HandleResponse 统一响应参数
func HandleResponse(r *ghttp.Request) {
	var (
		Resp *model.Response
	)
	r.Middleware.Next()
	Resp = new(model.Response)
	// 安全性配置
	r.Response.Header().Set("X-Frame-Options", "SAMEORIGIN")
	r.Response.Header().Set("X-XSS-Protection", "1; mode=block")
	r.Response.Header().Set("X-Content-Type-Options", "nosniff")
	r.Response.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'")
	if r.Response.BufferLength() > 0 && !gstr.Contains(r.Response.BufferString(), "exception recovered") {
		return
	}
	if gstr.Contains(r.Response.BufferString(), "exception recovered") {
		r.Response.ClearBuffer()
		_ = service.WarningWorkWx().SendWarningWorkWx(r.Context(), `
		### 系统异常
		@高杨
		*  链路ID：`+gctx.CtxId(r.Context())+`
		*  请求地址：`+r.Request.URL.String()+`
		*  请求参数：`+r.GetBodyString()+`
		*  响应内容：`+r.Response.BufferString())
		r.SetError(gerror.New("exception"))
	}
	mediaType, _, _ := mime.ParseMediaType(r.GetHeader("Content-Type"))
	for _, ct := range streamContentType {
		if mediaType == ct {
			return
		}
	}
	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	if err != nil {
		// exception recovered: runtime error: invalid memory address or nil pointer dereference
		if gstr.Contains(err.Error(), "exception recovered") {
			_ = service.WarningWorkWx().SendWarningWorkWx(r.Context(), `
		### 系统异常
		@高杨
		*  链路ID：`+gctx.CtxId(r.Context())+`
		*  请求地址：`+r.Request.URL.String()+`
		*  请求参数：`+r.GetBodyString()+`
        *  响应内容：`+err.Error())
			err = gerror.New("exception")
		}
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		Resp.Error = err
		msg = err.Error()
	} else {
		if r.Response.Status != http.StatusOK {
			msg = http.StatusText(r.Response.Status)
			switch r.Response.Status {
			case http.StatusNotFound:
				code = gcode.CodeNotFound
			case http.StatusForbidden:
				code = gcode.CodeNotAuthorized
			default:
				code = gcode.CodeUnknown
			}
			// It creates error as it can be retrieved by other middlewares.
			err = gerror.NewCode(code, msg)
			Resp.Error = err
		} else {
			code = gcode.CodeOK
			msg = "SUCCESS"
		}
	}
	I18nMsg := global.I18ngjson.Get(fmt.Sprintf("%s.%s", msg, contexts.GetLanguage(r.Context())))

	if !I18nMsg.IsEmpty() {
		msg = I18nMsg.String()
	}
	Resp.Code = code.Code()
	Resp.Message = msg
	Resp.Data = checkResponseData(res)
	Resp.Timestamp = gtime.Timestamp()
	Resp.TraceID = gctx.CtxId(r.Context())
	r.Response.WriteJson(Resp)
}

func checkResponseData(resp interface{}) (res interface{}) {
	res = resp
	if g.IsEmpty(res) {
		return
	}
	// 获取Person实例的反射值
	value := reflect.ValueOf(resp)
	// 获取User类型
	typ := reflect.TypeOf(resp)
	// 如果类型是指针类型的话需处理
	for value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if value.Kind() != reflect.Struct {
		return
	}
	// 遍历结构体的所有字段
	for i := 0; i < typ.NumField(); i++ {
		// 获取字段类型信息
		fieldType := typ.Field(i)
		// 获取字段的标签字符串
		tag := fieldType.Tag
		// 使用Get方法获取特定前缀的标签值，例如获取json标签
		ResponseTag := tag.Get("response")
		if ResponseTag == "arr" {
			res = value.Field(i).Interface()
			break
		}
	}
	return
}
