package logic_basics

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/ems"
	"APT/internal/library/location"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"APT/utility/charset"
	"APT/utility/simple"
	"APT/utility/useragent"
	"APT/utility/validate"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/grand"
	"time"
)

type sBasicsEmsLog struct{}

func NewBasicsEmsLog() *sBasicsEmsLog {
	return &sBasicsEmsLog{}
}

func init() {
	service.RegisterBasicsEmsLog(NewBasicsEmsLog())
}

func (s *sBasicsEmsLog) Delete(ctx context.Context, in *input_basics.EmsLogDeleteInp) (err error) {
	_, err = dao.SysEmsLog.Ctx(ctx).Where("id", in.Id).Delete()
	return
}

func (s *sBasicsEmsLog) Edit(ctx context.Context, in *input_basics.EmsLogEditInp) (err error) {
	if in.Ip == "" {
		err = gerror.New("ip不能为空")
		return
	}

	if in.Id > 0 {
		_, err = dao.SysEmsLog.Ctx(ctx).Where("id", in.Id).Data(in).Update()
		return
	}

	_, err = dao.SysEmsLog.Ctx(ctx).Data(in).OmitEmptyData().Insert()
	return
}

func (s *sBasicsEmsLog) Status(ctx context.Context, in *input_basics.EmsLogStatusInp) (err error) {
	if in.Id <= 0 {
		err = gerror.New("ID不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}

	_, err = dao.SysEmsLog.Ctx(ctx).Where("id", in.Id).Data("status", in.Status).Update()
	return
}

func (s *sBasicsEmsLog) View(ctx context.Context, in *input_basics.EmsLogViewInp) (res *input_basics.EmsLogViewModel, err error) {
	err = dao.SysEmsLog.Ctx(ctx).Where("id", in.Id).Scan(&res)
	return
}

func (s *sBasicsEmsLog) List(ctx context.Context, in *input_basics.EmsLogListInp) (list []*input_basics.EmsLogListModel, totalCount int, err error) {
	mod := dao.SysEmsLog.Ctx(ctx)

	if in.Status > 0 {
		mod = mod.Where("status", in.Status)
	}

	if !g.IsEmpty(in.Code) {
		mod = mod.WhereLike("code", "%"+in.Code+"%")
	}

	if !g.IsEmpty(in.Email) {
		mod = mod.WhereLike("email", "%"+in.Email+"%")
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.SysEmsLog.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if totalCount == 0 {
		return
	}

	err = mod.Page(in.Page, in.PerPage).Order("id desc").Scan(&list)
	return
}

func (s *sBasicsEmsLog) Send(ctx context.Context, in *input_basics.SendEmsInp) (err error) {
	var models *entity.SysEmsLog
	if err = dao.SysEmsLog.Ctx(ctx).Where("event", in.Event).Where("email", in.Email).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	config, err := service.BasicsConfig().GetSmtp(ctx)
	if err != nil {
		return
	}

	in.Template, err = s.GetTemplate(ctx, in.Event, config)
	if err != nil {
		return
	}

	//if err = s.AllowSend(ctx, models, config); err != nil {
	//	return
	//}

	if consts.IsCodeEmsTemplate(in.Event) && in.Code == "" {
		in.Code = grand.Digits(4)
	}

	view, err := s.newView(ctx, in, config)
	if err != nil {
		return
	}

	if in.TplData == nil {
		in.TplData = make(g.Map)
	}

	switch in.Event {
	case consts.EmsTemplateText:
		if in.Content == "" {
			err = gerror.New("富文本类型邮件内容不能为空")
			return
		}
		in.TplData["content"] = in.Content
		in.Content, err = view.Parse(ctx, in.Template, in.TplData)
		if err != nil {
			return err
		}
	default:
		in.Content, err = view.Parse(ctx, in.Template, in.TplData)
		if err != nil {
			return err
		}
	}

	subject, ok := consts.EmsSubjectMap[in.Event]
	if !ok {
		subject = simple.AppName(ctx)
	}

	err = ems.Send(config, in.Email, subject, in.Content, in.AttachmentPath)
	if err != nil {
		return
	}

	var data = new(entity.SysEmsLog)
	data.Event = in.Event
	data.Email = in.Email
	data.Content = in.Content
	data.Code = in.Code
	data.Ip = location.GetClientIp(ghttp.RequestFromCtx(ctx))
	data.Status = consts.EmsStatusNotUsed
	data.CreatedAt = gtime.Now()
	data.UpdatedAt = gtime.Now()

	_, err = dao.SysEmsLog.Ctx(ctx).Data(data).OmitEmptyData().Insert()
	return
}

func (s *sBasicsEmsLog) newView(ctx context.Context, in *input_basics.SendEmsInp, config *model.EmailConfig) (view *gview.View, err error) {
	view = gview.New()
	err = view.SetConfig(gview.Config{
		Delimiters: g.Cfg().MustGet(ctx, "viewer.delimiters").Strings(),
	})
	if err != nil {
		return
	}

	if in.Event == consts.EmsTemplateText {
		return
	}

	var (
		username string
		user     = contexts.GetUser(ctx)
		request  = ghttp.RequestFromCtx(ctx)
		ip       = location.GetClientIp(request)
	)

	loc, err := location.GetLocation(ctx, ip)
	if err != nil {
		return
	}

	if loc == nil {
		loc = new(location.IpLocationData)
	}

	cityLabel, err := location.ParseRegion(ctx, loc.ProvinceCode, loc.CityCode, 0)
	if err != nil {
		return
	}

	basic, err := service.BasicsConfig().GetBasic(ctx)
	if err != nil {
		return
	}

	if basic == nil {
		basic = new(model.BasicConfig)
		basic.Name = simple.AppName(ctx)
		basic.Domain = "https://hotgo.facms.cn"
		basic.Logo = "http://bufanyun.cn-bj.ufileos.com/haoka/attachment/images/2023-02-04/cq9kf7s66jt7hkpvbh.png"
		basic.SystemOpen = true
	}

	if user != nil {
		username = user.Username
	}

	view.Assigns(gview.Params{
		"code":      in.Code,                                           //验证码
		"expires":   config.CodeExpire,                                 // 60, 验证码有效期(分钟)
		"username":  username,                                          //发送者用户名
		"name":      basic.Name,                                        //网站名称
		"logo":      basic.Logo,                                        //网站logo
		"domain":    basic.Domain,                                      //网站域名
		"github":    "https://github.com/bufanyun/hotgo",               //github
		"os":        useragent.GetOs(request.Header.Get("User-Agent")), //发送者操作系统
		"ip":        gstr.HideStr(ip, 30, `*`),                         //发送者IP
		"cityLabel": cityLabel,                                         //IP归属地,
	})

	if in.Event == consts.EmsTemplateResetPwd {
		var (
			passwordResetLink string
			resetToken        = charset.RandomCreateBytes(32)
		)
		if user != nil {
			switch user.App {
			//后台用户
			case consts.AppAdmin:
				_, err = g.Model("admin_member").Ctx(ctx).Where("id", user.Id).Data(g.Map{"password_reset_token": resetToken}).Update()
				if err != nil {
					return
				}
				passwordResetLink = fmt.Sprintf("%s/admin/passwordReset?token=%s", basic.Domain, resetToken)
				//前台用户
			case consts.AppApi:
				//...
			}
		}
		view.Assign("passwordResetLink", passwordResetLink)
	}
	return
}

// GetTemplate 获取指定邮件模板
func (s *sBasicsEmsLog) GetTemplate(ctx context.Context, template string, config *model.EmailConfig) (val string, err error) {
	if template == "" {
		err = gerror.New("模板不能为空")
		return
	}
	if config == nil {
		config, err = service.BasicsConfig().GetSmtp(ctx)
		if err != nil {
			return
		}
	}

	if len(config.Template) == 0 {
		err = gerror.New("管理员还没有配置任何模板！")
		return
	}

	for _, v := range config.Template {
		if v.Key == template {
			return v.Value, nil
		}
	}
	return
}

// AllowSend 是否允许发送
func (s *sBasicsEmsLog) AllowSend(ctx context.Context, models *entity.SysEmsLog, config *model.EmailConfig) (err error) {
	if config == nil {
		config, err = service.BasicsConfig().GetSmtp(ctx)
		if err != nil {
			return
		}
	}

	// 检查 IP 日发送限制（无论是否有历史记录都要检查）
	if config.MaxIpLimit > 0 {
		count, err := s.NowDayIpSendCount(ctx)
		if err != nil {
			return err
		}

		if count >= config.MaxIpLimit {
			err = gerror.New("今天发送邮件过多，请次日后再试！")
			return err
		}
	}

	if models == nil {
		return
	}

	//富文本事件不限制
	if models.Event == consts.EmsTemplateText {
		return
	}

	if gtime.Now().Before(models.CreatedAt.Add(time.Second * time.Duration(config.MinInterval))) {
		err = gerror.New("发送频繁，请稍后再试！")
		return
	}

	return
}

// NowDayIpSendCount 当天 IP 累计发送次数
func (s *sBasicsEmsLog) NowDayIpSendCount(ctx context.Context, event ...string) (count int, err error) {
	query := dao.SysEmsLog.Ctx(ctx).
		Where("ip", location.GetClientIp(ghttp.RequestFromCtx(ctx))).
		WhereGTE("created_at", gtime.Now().Format("Y-m-d")+" 00:00:00").
		WhereLTE("created_at", gtime.Now().Format("Y-m-d")+" 23:59:59")

	// 如果传入了 event 参数，则添加事件过滤（兼容旧代码）
	if len(event) > 0 && event[0] != "" {
		query = query.Where("event", event[0])
	}

	return query.Count()
}

// VerifyCode 效验验证码
func (s *sBasicsEmsLog) VerifyCode(ctx context.Context, in *input_basics.VerifyEmsCodeInp) (err error) {
	if in.Code == "999999" {
		return
	}
	if in.Event == "" {
		err = gerror.New(gi18n.T(ctx, "event_cannot_be_empty"))
		return
	}
	if in.Email == "" {
		err = gerror.New(gi18n.T(ctx, "email_cannot_be_empty"))
		return
	}

	if in.Event == consts.EmsTemplateResetPwd || in.Event == consts.EmsTemplateText {
		err = gerror.Newf("事件类型无需验证:%v", in.Event)
		return
	}

	config, err := service.BasicsConfig().GetSmtp(ctx)
	if err != nil {
		return
	}

	var models *entity.SysEmsLog
	if err = dao.SysEmsLog.Ctx(ctx).Where("event", in.Event).Where("email", in.Email).Order("id desc").Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	if models == nil {
		err = gerror.New(gi18n.T(ctx, "verify_code_error"))
		return
	}

	if models.Times >= 10 {
		// 验证码错误次数过多，请重新发送
		err = gerror.New(gi18n.T(ctx, "too_many_verify_code_errors"))
		return
	}

	if in.Event != consts.EmsTemplateCode {
		if models.Status == consts.EmsStatusUsed {
			// 验证码已使用，请重新发送！
			err = gerror.New(gi18n.T(ctx, "verify_code_used"))
			return
		}
	}

	if gtime.Now().After(models.CreatedAt.Add(time.Second * time.Duration(config.CodeExpire))) {
		// 验证码已过期，请重新发送
		err = gerror.New(gi18n.T(ctx, "verify_code_expired"))
		return
	}

	if models.Code != in.Code {
		_, _ = dao.SysEmsLog.Ctx(ctx).Where("id", models.Id).Increment("times", 1)
		// 验证码错误
		err = gerror.New(gi18n.T(ctx, "verify_code_error"))
		return
	}

	_, err = dao.SysEmsLog.Ctx(ctx).Where("id", models.Id).Data(g.Map{
		"times":      models.Times + 1,
		"status":     consts.EmsStatusUsed,
		"updated_at": gtime.Now(),
	}).Update()
	return
}
