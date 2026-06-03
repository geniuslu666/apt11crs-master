package logic_basics

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/global"
	"APT/internal/library/contexts"
	"APT/internal/library/dict"
	"APT/internal/library/hgorm/handler"
	"APT/internal/library/hgorm/hook"
	"APT/internal/library/location"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"APT/utility/excel"
	"APT/utility/simple"
	"APT/utility/validate"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type sBasicsLog struct{}

func NewBasicsLog() *sBasicsLog {
	return &sBasicsLog{}
}

func init() {
	service.RegisterBasicsLog(NewBasicsLog())
}

func (s *sBasicsLog) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.SysLog.Ctx(ctx), option...)
}

func (s *sBasicsLog) Export(ctx context.Context, in *input_basics.LogListInp) (err error) {

	type exportImage struct {
		Id         int64       `json:"id"           description:""`
		AppId      string      `json:"app_id"       description:"应用id"`
		Method     string      `json:"method"       description:"提交类型"`
		Module     string      `json:"module"       description:"模块"`
		Url        string      `json:"url"          description:"提交url"`
		Ip         string      `json:"ip"           description:"ip地址"`
		ErrorCode  int         `json:"error_code"   description:"报错code"`
		ErrorMsg   string      `json:"error_msg"    description:"报错信息"`
		ReqId      string      `json:"req_id"       description:"对外id"`
		TakeUpTime int64       `json:"take_up_time" description:"请求耗时"`
		CreatedAt  *gtime.Time `json:"created_at"   description:"创建时间"`
		MemberName string      `json:"memberName"`
		Region     string      `json:"region"`
	}

	var (
		titleList  = []string{"ID", "应用", "提交类型", "模块", "提交url", "ip地址", "报错code", "报错信息", "对外id", "请求耗时", "创建时间", "用户", "访问地"}
		fileName   = "访问日志导出-" + gctx.CtxId(ctx)
		sheetName  = simple.AppName(ctx)
		exportList []exportImage
		row        exportImage
	)

	list, _, err := s.List(ctx, in)
	if err != nil {
		return err
	}

	for i := 0; i < len(list); i++ {
		row.Id = list[i].Id
		row.AppId = list[i].AppId
		row.Module = list[i].Module
		row.Method = list[i].Method
		row.Url = list[i].Url
		row.Ip = list[i].Ip
		row.ReqId = list[i].ReqId
		row.ErrorCode = list[i].ErrorCode
		row.ErrorMsg = list[i].ErrorMsg
		row.TakeUpTime = list[i].TakeUpTime
		row.CreatedAt = list[i].CreatedAt
		row.MemberName = list[i].MemberName
		row.Region = list[i].Region
		exportList = append(exportList, row)
	}

	err = excel.ExportByStructs(ctx, titleList, exportList, fileName, sheetName)
	return
}

func (s *sBasicsLog) RealWrite(ctx context.Context, log entity.SysLog) (err error) {
	_, err = dao.SysLog.Ctx(ctx).FieldsEx(dao.SysLog.Columns().Id).Data(log).Unscoped().OmitEmptyData().Insert()
	return
}

func (s *sBasicsLog) AutoLog(ctx context.Context) error {
	return g.Try(ctx, func(ctx context.Context) {
		var err error
		defer func() {
			if err != nil {
				g.Log().Error(ctx, "autoLog err:%+v", err)
			}
		}()

		config, err := service.BasicsConfig().GetLoadLog(ctx)
		if err != nil {
			return
		}

		if config == nil || !config.Switch {
			return
		}

		data := s.AnalysisLog(ctx)
		if ok := validate.InSliceExistStr(config.Module, data.Module); !ok {
			return
		}

		if ok := validate.InSliceExistStr(config.SkipCode, gconv.String(data.ErrorCode)); ok {
			return
		}

		err = s.RealWrite(ctx, data)
	})
}

func (s *sBasicsLog) AnalysisLog(ctx context.Context) entity.SysLog {
	var (
		mctx       = contexts.Get(ctx)
		request    = ghttp.RequestFromCtx(ctx)
		clientIp   = location.GetClientIp(request)
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
	)

	// 安全检查：如果上下文为空则返回空日志
	if mctx == nil {
		g.Log().Warning(ctx, "AnalysisLog: mctx is nil, skip logging")
		return data
	}

	var (
		response = mctx.Response
		user     = mctx.User
	)
	module = mctx.Module

	if response != nil {
		errorCode = response.Code
		errorMsg = response.Message
		traceID = response.TraceID
		timestamp = response.Timestamp
		if len(gconv.String(response.Error)) > 0 {
			errorData = gjson.New(response.Error)
		}
	}

	if timestamp == 0 {
		timestamp = gtime.Timestamp()
	}

	if reqHeadersBytes, _ := gjson.New(request.Header).MarshalJSON(); len(reqHeadersBytes) > 0 {
		headerData = gjson.New(reqHeadersBytes)
	}

	if mctx.Data != nil {
		if body, ok := mctx.Data["request.body"].(*gjson.Json); ok {
			postData = body
		}
	}

	postForm := gjson.New(gconv.String(request.PostForm)).Map()
	if len(postForm) > 0 {
		for k, v := range postForm {
			postData.MustSet(k, v)
		}
	}

	if postData.IsNil() || len(postData.Map()) == 0 {
		postData = gjson.New(consts.NilJsonToString)
	}

	if user != nil {
		memberId = user.Id
		appId = user.App
	}

	ipData, err := location.GetLocation(ctx, clientIp)
	if err != nil {
		g.Log().Debugf(ctx, "location.GetLocation clientIp:%v, err:%+v", clientIp, err)
	}

	if ipData == nil {
		ipData = new(location.IpLocationData)
	}

	if mctx.Data != nil {
		if tt, ok := mctx.Data["request.takeUpTime"].(int64); ok {
			takeUpTime = tt
		}
	}

	headerData.MustSet("qqq", request.EnterTime.String())

	data = entity.SysLog{
		AppId:       appId,
		AppMemberId: 0,
		MemberId:    memberId,
		Method:      request.Method,
		Module:      module,
		Url:         request.URL.Path,
		GetData:     getData,
		PostData:    postData,
		HeaderData:  headerData,
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
	return data
}

func (s *sBasicsLog) View(ctx context.Context, in *input_basics.LogViewInp) (res *input_basics.LogViewModel, err error) {
	mod := s.Model(ctx)

	count, err := service.BasicsLoginLog().Model(ctx).
		LeftJoinOnFields(dao.SysLog.Table(), dao.SysLoginLog.Columns().ReqId, "=", dao.SysLog.Columns().ReqId).
		WherePrefix(dao.SysLog.Table(), dao.SysLog.Columns().Id, in.Id).Count()
	if err != nil {
		return nil, err
	}

	if count > 0 {
		mod = dao.SysLog.Ctx(ctx)
	}

	if err = mod.Hook(hook.CityLabel).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if res == nil {
		return
	}

	routes := global.LoadHTTPRoutes(ghttp.RequestFromCtx(ctx))
	key := global.GenRouteKey(res.Method, res.Url)
	route, ok := routes[key]
	if ok {
		res.Tags = route.Tags
		res.Summary = route.Summary
		res.Description = route.Description
	}

	if simple.IsDemo(ctx) {
		res.HeaderData = gjson.New(`{
		   "none": [
		       "` + consts.DemoTips + `"
		   ]
		}`)
	}
	return
}

func (s *sBasicsLog) Delete(ctx context.Context, in *input_basics.LogDeleteInp) (err error) {
	_, err = s.Model(ctx).WherePri(in.Id).Delete()
	return
}

func (s *sBasicsLog) List(ctx context.Context, in *input_basics.LogListInp) (list []*input_basics.LogListModel, totalCount int, err error) {
	mod := s.Model(ctx).FieldsEx("get_data", "header_data", "post_data")

	if in.Url != "" {
		mod = mod.WhereLike("url", "%"+in.Url+"%")
	}

	if in.Module != "" {
		mod = mod.Where("module", in.Module)
	}

	if in.ReqId != "" {
		mod = mod.Where("req_id", in.ReqId)
	}

	if in.Method != "" {
		mod = mod.Where("method", in.Method)
	}

	if in.MemberId > 0 {
		mod = mod.Where("member_id", in.MemberId)
	}

	if in.Ip != "" {
		mod = mod.Where("ip", in.Ip)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween("created_at", gtime.New(in.CreatedAt[0]), gtime.New(in.CreatedAt[1]))
	}

	if in.ErrorCode != "" {
		mod = mod.Where("error_code", in.ErrorCode)
	}

	if dict.HasOptionKey(consts.HTTPHandlerTimeOptions, in.TakeUpTime) {
		mod = mod.Where(fmt.Sprintf("`take_up_time` %v", in.TakeUpTime))
	}

	totalCount, err = mod.Count()
	if err != nil || totalCount == 0 {
		return
	}

	if err = mod.Page(in.Page, in.PerPage).Hook(hook.CityLabel).Order("id desc").Scan(&list); err != nil {
		return
	}

	routes := global.LoadHTTPRoutes(ghttp.RequestFromCtx(ctx))
	for _, v := range list {
		if v.AppId == consts.AppAdmin {
			memberName, err := dao.AdminMember.Ctx(ctx).Fields("realname").WherePri(v.MemberId).Value()
			if err != nil {
				err = gerror.Wrap(err, consts.ErrorORM)
				return list, totalCount, err
			}
			v.MemberName = memberName.String()
		}

		if v.MemberName == "" {
			v.MemberName = "游客"
		}

		if gstr.Contains(v.Url, "?") {
			v.Url = gstr.StrTillEx(v.Url, "?")
		}

		key := global.GenRouteKey(v.Method, v.Url)
		route, ok := routes[key]
		if ok {
			v.Tags = route.Tags
			v.Summary = route.Summary
			v.Description = route.Description
		}

		if simple.IsDemo(ctx) {
			v.HeaderData = gjson.New(`{
			   "none": [
			       "` + consts.DemoTips + `"
			   ]
			}`)
		}
	}
	return
}
