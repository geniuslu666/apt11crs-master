package logic_basics

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	"APT/internal/library/hgorm/hook"
	"APT/internal/library/location"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"APT/utility/convert"
	"APT/utility/excel"
	"APT/utility/useragent"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sBasicsLoginLog struct{}

func NewBasicsLoginLog() *sBasicsLoginLog {
	return &sBasicsLoginLog{}
}

func init() {
	service.RegisterBasicsLoginLog(NewBasicsLoginLog())
}

func (s *sBasicsLoginLog) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.SysLoginLog.Ctx(ctx), option...)
}

func (s *sBasicsLoginLog) List(ctx context.Context, in *input_basics.LoginLogListInp) (list []*input_basics.LoginLogListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	if in.Status > 0 {
		mod = mod.Where(dao.SysLoginLog.Columns().Status, in.Status)
	}

	if len(in.LoginAt) == 2 {
		mod = mod.WhereBetween(dao.SysLoginLog.Columns().LoginAt, in.LoginAt[0], in.LoginAt[1])
	}

	if in.LoginIp != "" {
		mod = mod.Where(dao.SysLoginLog.Columns().LoginIp, in.LoginIp)
	}

	if in.Username != "" {
		mod = mod.Where(dao.SysLoginLog.Columns().Username, in.Username)
	}

	totalCount, err = mod.Clone().Count(1)
	if err != nil || totalCount == 0 {
		return
	}

	if err = mod.Fields(input_basics.LoginLogListModel{}).Hook(hook.CityLabel).Page(in.Page, in.PerPage).OrderDesc(dao.SysLoginLog.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	for _, v := range list {
		v.Os = useragent.GetOs(v.UserAgent)
		v.Browser = useragent.GetBrowser(v.UserAgent)
		v.SysLogId, err = dao.SysLog.Ctx(ctx).Fields("id").Where("req_id", v.ReqId).Value()
		if err != nil {
			return nil, 0, err
		}
	}
	return
}

func (s *sBasicsLoginLog) Export(ctx context.Context, in *input_basics.LoginLogListInp) (err error) {
	list, _, err := s.List(ctx, in)
	if err != nil {
		return
	}

	tags, err := convert.GetEntityDescTags(input_basics.LoginLogExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出登录日志-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("登录日志")
		exports   []input_basics.LoginLogExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

func (s *sBasicsLoginLog) Delete(ctx context.Context, in *input_basics.LoginLogDeleteInp) (err error) {
	_, err = s.Model(ctx).WherePri(in.Id).Delete()
	return
}

func (s *sBasicsLoginLog) Push(ctx context.Context, in *input_basics.LoginLogPushInp) {
	if in.Response == nil {
		in.Response = new(input_basics.LoginModel)
	}

	r := ghttp.RequestFromCtx(ctx)
	if r == nil {
		g.Log().Warningf(ctx, "ctx not http request")
		return
	}

	clientIp := location.GetClientIp(r)
	ipData, err := location.GetLocation(ctx, clientIp)
	if err != nil {
		g.Log().Debugf(ctx, "location.GetLocation clientIp:%v, err:%+v", clientIp, err)
	}

	if ipData == nil {
		ipData = new(location.IpLocationData)
	}

	var models entity.SysLoginLog
	models.ReqId = gctx.CtxId(ctx)
	models.MemberId = in.Response.Id
	models.Username = in.Response.Username
	models.LoginAt = gtime.Now()
	models.LoginIp = clientIp
	models.UserAgent = r.UserAgent()
	models.ProvinceId = ipData.ProvinceCode
	models.CityId = ipData.CityCode
	models.Status = consts.StatusEnabled

	if in.Err != nil {
		models.Status = consts.StatusDisable
		models.ErrMsg = in.Err.Error()
	}

	models.Response = gjson.New(consts.NilJsonToString)
	if in.Response != nil {
		models.Response = gjson.New(in.Response)
	}

	//if err = queue.Push(consts.QueueLoginLogTopic, models); err != nil {
	//	g.Log().Warningf(ctx, "push LoginLog err:%+v, models:%v", err, gjson.New(models).String())
	//}
}

func (s *sBasicsLoginLog) RealWrite(ctx context.Context, models entity.SysLoginLog) (err error) {
	_, err = dao.SysLoginLog.Ctx(ctx).Data(models).OmitEmptyData().Insert()
	return
}
