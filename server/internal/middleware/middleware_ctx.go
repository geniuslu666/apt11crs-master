package middleware

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/location"
	"APT/internal/model"
	"APT/internal/model/entity"
	"context"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"go.opentelemetry.io/otel"
)

var (
	NotRecordRequest = g.Map{
		"/admin/upload/file":       struct{}{}, // 上传文件
		"/admin/upload/uploadPart": struct{}{}, // 上传分片
	}
)

func Ctx(r *ghttp.Request) {
	var (
		CtxLogger = g.Log()
		SysLog    *entity.SysLog
		ctx       = r.Context()
	)
	tracer := otel.Tracer("otel-go-tracer")
	ctx, span := tracer.Start(ctx, "[Middleware][Ctx]全局日志入口")
	defer span.End()
	SysLog = new(entity.SysLog)
	SysLog.ReqId = gctx.CtxId(r.Context())
	CtxLogger.Info(r.Context(), "--[REQUEST -> ]----------------------------------------------------------------")
	CtxLogger.Info(r.Context(), r.Request.URL.Path)
	CtxLogger.Info(r.Context(), r.Request.Method)
	SysLog.Method = r.Request.Method
	for k, v := range r.Header {
		CtxLogger.Info(r.Context(), " -H ", k, v[0])
	}
	if len(r.GetBodyString()) > 20000 {
		CtxLogger.Info(r.Context(), "request.body too long")
	} else {
		CtxLogger.Info(r.Context(), r.GetBodyString())
	}
	data := make(g.Map)
	if _, ok := NotRecordRequest[r.URL.Path]; ok {
		data["request.body"] = gjson.New(nil)
	} else {
		data["request.body"] = gjson.New(r.GetBodyString())
	}

	contexts.Init(r, &model.Context{
		Data:   data,
		Module: getModule(r.URL.Path),
	})
	SetLanguage(r)
	if len(r.Cookie.GetSessionId()) == 0 {
		r.Cookie.SetSessionId(gctx.CtxId(r.Context()))
	}

	r.SetCtx(r.GetNeverDoneCtx())
	r.Middleware.Next()
	CtxLogger.Info(r.Context(), "--[RESPONSE <-]----------------------------------------------------------------")
	CtxLogger.Info(r.Context(), r.Response.Status)
	for k, v := range r.Response.Header() {
		CtxLogger.Info(r.Context(), " -H ", k, v[0])
	}
	CtxLogger.Info(r.Context(), r.Response.Buffer())
}

func AnalysisLog(ctx context.Context) {
	var (
		greq       = ghttp.RequestFromCtx(ctx)
		CtxId      = gctx.CtxId(ctx)
		mctx       = contexts.Get(ctx)
		request    = ghttp.RequestFromCtx(ctx)
		postData   = gjson.New(consts.NilJsonToString)
		getData    = gjson.New(request.URL.Query())
		headerData = gjson.New(consts.NilJsonToString)
		errorData  = gjson.New(consts.NilJsonToString)
		data       entity.SysLog
		memberId   int64
		errorCode  int
		errorMsg   string
		traceID    string
		timestamp  int64
		appId      string
		takeUpTime int64
		module     string
		clientIp   = location.GetClientIp(request)
	)

	// 安全检查：如果上下文为空则直接返回
	if mctx == nil {
		g.Log().Warning(ctx, "AnalysisLog: mctx is nil, skip logging")
		return
	}

	var (
		response        = mctx.Response
		gatewayResponse = mctx.GatewayCode
		user            = mctx.User
		MemberUser      = mctx.MemberUser
	)
	module = mctx.Module

	// 响应数据
	if response != nil {
		errorCode = response.Code
		errorMsg = response.Message
		traceID = CtxId
		timestamp = response.Timestamp
		if len(gconv.String(response.Error)) > 0 {
			errorData = gjson.New(response.Error)
		}
	}
	if gatewayResponse != nil {
		errorCode = gatewayResponse.ResCode
		errorMsg = gatewayResponse.ResMessage
		traceID = CtxId
		timestamp = gtime.Now().Timestamp()
		if !g.IsEmpty(greq.GetError()) {
			errorData = gjson.New(greq.GetError())
		}
	}

	if timestamp == 0 {
		timestamp = gtime.Timestamp()
	}

	// 请求头
	if reqHeadersBytes, _ := gjson.New(request.Header).MarshalJSON(); len(reqHeadersBytes) > 0 {
		headerData = gjson.New(reqHeadersBytes)
	}

	// post参数
	if mctx.Data != nil {
		if body, ok := mctx.Data["request.body"].(*gjson.Json); ok {
			postData = body
		}
	}

	// post表单
	postForm := gjson.New(gconv.String(request.PostForm)).Map()
	if len(postForm) > 0 {
		for k, v := range postForm {
			postData.MustSet(k, v)
		}
	}

	if postData.IsNil() || len(postData.Map()) == 0 {
		postData = gjson.New(consts.NilJsonToString)
	}

	// 当前登录用户
	if user != nil {
		memberId = user.Id
		appId = user.App
	}

	// 当前App登录用户
	if MemberUser != nil {
		memberId = gvar.New(MemberUser.MemberId).Int64()
		appId = MemberUser.App
	}

	//ipData, err := location.GetLocation(ctx, clientIp)
	//if err != nil {
	//    g.Log().Debugf(ctx, "location.GetLocation clientIp:%v, err:%+v", clientIp, err)
	//}

	ipData := new(location.IpLocationData)

	// 请求耗时
	if mctx.Data != nil {
		if tt, ok := mctx.Data["request.takeUpTime"].(int64); ok {
			takeUpTime = tt
		}
	}

	headerData.MustAppend("Enter-Time", request.EnterTime.String())

	data = entity.SysLog{
		AppId:       appId,
		AppMemberId: 0,
		MemberId:    memberId,
		Method:      request.Method,
		Module:      module,
		Url:         request.URL.Path,
		GetData:     getData,
		PostData:    postData,
		HeaderData:  SimplifyHeaderParams(headerData),
		Ip:          clientIp,
		ProvinceId:  ipData.ProvinceCode,
		CityId:      ipData.CityCode,
		ErrorCode:   errorCode,
		ErrorMsg:    errorMsg,
		ErrorData:   errorData,
		ReqId:       traceID,
		Timestamp:   timestamp,
		UserAgent:   request.Header.Get("User-Agent"),
		Status:      consts.StatusEnabled,
		TakeUpTime:  takeUpTime,
		UpdatedAt:   gtime.Now(),
		CreatedAt:   request.EnterTime,
	}
	_, logInsertErr := dao.SysLog.Ctx(ctx).Data(data).Insert()
	if logInsertErr != nil {
		g.Log().Error(ctx, logInsertErr)
	}
	return
}

func SimplifyHeaderParams(data *gjson.Json) *gjson.Json {
	if data == nil || data.IsNil() {
		return data
	}
	var filters = []string{"Accept", "Cookie"}
	for _, filter := range filters {
		v := data.Get(filter)
		if len(v.String()) > 128 {
			data.MustRemove(filter)
		}
	}
	return data
}
