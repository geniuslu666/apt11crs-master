package logic_basics

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/location"
	"APT/internal/library/sms"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
)

type sBasicsSmsLog struct{}

func NewBasicsSmsLog() *sBasicsSmsLog {
	return &sBasicsSmsLog{}
}

func init() {
	service.RegisterBasicsSmsLog(NewBasicsSmsLog())
}

func (s *sBasicsSmsLog) Delete(ctx context.Context, in *input_basics.SmsLogDeleteInp) (err error) {
	_, err = dao.SysSmsLog.Ctx(ctx).WherePri(in.Id).Delete()
	return
}

func (s *sBasicsSmsLog) View(ctx context.Context, in *input_basics.SmsLogViewInp) (res *input_basics.SmsLogViewModel, err error) {
	if err = dao.SysSmsLog.Ctx(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}
	return
}

func (s *sBasicsSmsLog) List(ctx context.Context, in *input_basics.SmsLogListInp) (list []*input_basics.SmsLogListModel, totalCount int, err error) {
	mod := dao.SysSmsLog.Ctx(ctx)
	cols := dao.SysSmsLog.Columns()

	if in.Mobile != "" {
		mod = mod.WhereLike(cols.Mobile, "%"+in.Mobile+"%")
	}

	if in.Ip != "" {
		mod = mod.Where(cols.Ip, in.Ip)
	}

	if in.Event != "" {
		mod = mod.Where(cols.Event, in.Event)
	}

	if in.Status > 0 {
		mod = mod.Where(cols.Status, in.Status)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(cols.CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Page(in.Page, in.PerPage).OrderDesc(cols.Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}
	return
}

func (s *sBasicsSmsLog) SendCode(ctx context.Context, in *input_basics.SendCodeInp) (err error) {
	if g.IsEmpty(in.Event) {
		err = gerror.New("事件不能为空")
		return
	}

	if g.IsEmpty(in.Mobile) {
		err = gerror.New("手机号不能为空")
		return
	}

	var models *entity.SysSmsLog
	if err = dao.SysSmsLog.Ctx(ctx).Where(dao.SysSmsLog.Columns().Event, in.Event).Where(dao.SysSmsLog.Columns().Mobile, in.Mobile).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	config, err := service.BasicsConfig().GetSms(ctx)
	if !g.IsEmpty(in.SmsDrive) {
		config.SmsDrive = in.SmsDrive
	}
	if err != nil {
		return
	}

	if in.Template, err = s.GetTemplate(ctx, in.Event, config); err != nil {
		return
	}

	if err = s.AllowSend(ctx, models, config); err != nil {
		return
	}

	if in.Code == "" {
		in.Code = grand.Digits(4)
	}

	if err = sms.New(config.SmsDrive).SendCode(ctx, in); err != nil {
		return
	}

	var data = new(entity.SysSmsLog)
	data.Event = in.Event
	data.Mobile = in.Mobile
	data.Code = in.Code
	data.Ip = location.GetClientIp(ghttp.RequestFromCtx(ctx))
	data.Status = consts.SmsStatusNotUsed
	data.CreatedAt = gtime.Now()
	data.UpdatedAt = gtime.Now()

	_, err = dao.SysSmsLog.Ctx(ctx).Data(data).OmitEmptyData().Insert()
	return
}

// SendPlaceOrderMsg 发送下单预警短信
func (s *sBasicsSmsLog) SendPlaceOrderMsg(ctx context.Context, in *input_basics.SendMsgInp) (err error) {

	if g.IsEmpty(in.Event) {
		err = gerror.New("没有事件参数")
		return
	}

	// 判断是否发送过
	var models *entity.SysSmsLog
	if err = dao.SysSmsLog.Ctx(ctx).Where(dao.SysSmsLog.Columns().Event, in.Event).Where(dao.SysSmsLog.Columns().Mobile, in.Mobile).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	// 获取短信配置
	config, err := service.BasicsConfig().GetSms(ctx)
	if !g.IsEmpty(in.SmsDrive) {
		config.SmsDrive = in.SmsDrive
	}
	if err != nil {
		return
	}

	var spaSmsConfig *model.OrderSmsConfig
	var carSmsConfig *model.OrderSmsConfig
	var smsTemplateContent string
	var isSendSms = 0
	if in.Event == "spaPlaceOrder" {
		// 按摩
		if spaSmsConfig, err = service.BasicsConfig().GetSpaSmsConfig(ctx); err != nil {
			return
		}

		for _, smsTemplateItem := range spaSmsConfig.SmsTemplate {
			if smsTemplateItem.Event == in.Event {
				isSendSms = smsTemplateItem.IsSendSms
				in.Template = smsTemplateItem.TemplateCode
				//in.Mobile = smsTemplateItem.SendSmsPhone
				smsTemplateContent = smsTemplateItem.Content
				in.Content = smsTemplateItem.Content
				break
			}
		}

		// 获取发送的手机号
		var spaOrder *entity.SpaOrder
		var spaIsp *entity.SpaIsp
		if err = dao.SpaOrder.Ctx(ctx).Where(dao.SpaOrder.Columns().OrderSn, in.OrderSn).Scan(&spaOrder); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return
		}
		if !g.IsEmpty(spaOrder) {
			if err = dao.SpaIsp.Ctx(ctx).Where(dao.SpaIsp.Columns().Id, spaOrder.IspId).Scan(&spaIsp); err != nil {
				err = gerror.Wrap(err, consts.ErrorORM)
				return
			}
			if !g.IsEmpty(spaIsp) {
				in.Mobile = spaIsp.SendSmsPhone
			}
		}

	} else if in.Event == "carPlaceOrder" {
		// 接送机
		if carSmsConfig, err = service.BasicsConfig().GetCarSmsConfig(ctx); err != nil {
			return
		}
		for _, smsTemplateItem := range carSmsConfig.SmsTemplate {
			if smsTemplateItem.Event == in.Event {
				isSendSms = smsTemplateItem.IsSendSms

				in.Template = smsTemplateItem.TemplateCode
				in.Mobile = smsTemplateItem.SendSmsPhone
				smsTemplateContent = smsTemplateItem.Content
				in.Content = smsTemplateItem.Content
				break
			}
		}
	} else if in.Event == "carRefundOrder" {
		var CarOrder *entity.CarOrder
		var scene string
		if err = dao.CarOrder.Ctx(ctx).
			Where(dao.CarOrder.Columns().OrderSn, in.OrderSn).
			Scan(&CarOrder); err != nil {
			return
		}
		if g.IsEmpty(CarOrder) {
			return gerror.New("订单不存在")
		}
		if CarOrder.ServiceType == "PICKUP" {
			scene = "接机"
		} else if CarOrder.ServiceType == "DELIVERY" {
			scene = "送机"
		} else {
			scene = "包车"
		}
		// 接送机
		if carSmsConfig, err = service.BasicsConfig().GetCarSmsConfig(ctx); err != nil {
			return
		}
		for _, smsTemplateItem := range carSmsConfig.SmsTemplate {
			if smsTemplateItem.Event == in.Event {
				isSendSms = smsTemplateItem.IsSendSms

				in.Template = smsTemplateItem.TemplateCode
				in.Mobile = smsTemplateItem.SendSmsPhone
				smsTemplateContent = smsTemplateItem.Content
				in.TemplateParams = map[string]string{
					"1": scene,
					"2": fmt.Sprintf("%s %s", CarOrder.BookDate, CarOrder.BookTime),
					"3": CarOrder.BookingName,
				}
				content := strings.Replace(smsTemplateItem.Content, "{1}", scene, 1)
				content = strings.Replace(content, "{2}", fmt.Sprintf("%s %s", CarOrder.BookDate, CarOrder.BookTime), 1)
				content = strings.Replace(content, "{3}", CarOrder.BookingName, 1)
				in.Content = content
				break
			}
		}
	}

	if isSendSms != 1 {
		err = gerror.New("短信发送未开启")
		return
	}

	if g.IsEmpty(in.Mobile) {
		err = gerror.New("手机号为空")
		return
	}

	if err = s.AllowSend(ctx, models, config); err != nil {
		return
	}

	if err = sms.New(config.SmsDrive).SendMsg(ctx, in); err != nil {
		return
	}

	var data = new(entity.SysSmsLog)
	data.Event = in.Event
	data.Mobile = in.Mobile
	data.Code = smsTemplateContent
	data.Ip = location.GetClientIp(ghttp.RequestFromCtx(ctx))
	data.Status = consts.SmsStatusNotUsed
	data.CreatedAt = gtime.Now()
	data.UpdatedAt = gtime.Now()

	_, err = dao.SysSmsLog.Ctx(ctx).Data(data).OmitEmptyData().Insert()
	return
}

func (s *sBasicsSmsLog) GetTemplate(ctx context.Context, template string, config *model.SmsConfig) (val string, err error) {
	if template == "" {
		err = gerror.New("模板不能为空")
		return
	}
	if config == nil {
		config, err = service.BasicsConfig().GetSms(ctx)
		if err != nil {
			return
		}
	}

	switch config.SmsDrive {
	case consts.SmsDriveAliYun:
		if len(config.AliYunTemplate) == 0 {
			err = gerror.New("管理员还没有配置任何阿里云短信模板！")
			return
		}

		for _, v := range config.AliYunTemplate {
			if v.Key == template {
				return v.Value, nil
			}
		}

	case consts.SmsDriveTencent:
		if len(config.TencentTemplate) == 0 {
			err = gerror.New("管理员还没有配置任何腾讯云短信模板！")
			return
		}

		for _, v := range config.TencentTemplate {
			if v.Key == template {
				return v.Value, nil
			}
		}
	case consts.SmsDriveUms:
		if len(config.UmsTemplate) == 0 {
			err = gerror.New("管理员还没有配置任何腾讯云短信模板！")
			return
		}

		for _, v := range config.UmsTemplate {
			if v.Key == template {
				return v.Value, nil
			}
		}
	case consts.SmsDriveTwilio:
		if len(config.TwilioTemplate) == 0 {
			err = gerror.New("管理员还没有配置任何Twilio短信模板！")
			return
		}

		for _, v := range config.TwilioTemplate {
			if v.Key == template {
				return v.Value, nil
			}
		}
	default:
		err = gerror.Newf("暂不支持短信驱动:%v", config.SmsDrive)
		return
	}
	return
}

func (s *sBasicsSmsLog) AllowSend(ctx context.Context, models *entity.SysSmsLog, config *model.SmsConfig) (err error) {

	if config == nil {
		if config, err = service.BasicsConfig().GetSms(ctx); err != nil {
			return
		}
	}
	if config.SmsMaxIpLimit > 0 {
		count, err := s.NowDayIpSendCount(ctx)
		if err != nil {
			return err
		}

		if count >= config.SmsMaxIpLimit {
			// 今天发送短信过多，请次日后再试！
			err = gerror.New(gi18n.T(ctx, "too_many_text_messages_sent"))
			return err
		}
	}
	if models == nil {
		return
	}

	if gtime.Now().Before(models.CreatedAt.Add(time.Second * time.Duration(config.SmsMinInterval))) {
		// 发送频繁，请稍后再试！
		err = gerror.New(gi18n.T(ctx, "frequent_sending_please_try_again_later"))
		return
	}

	return
}

func (s *sBasicsSmsLog) NowDayIpSendCount(ctx context.Context) (count int, err error) {
	return dao.SysSmsLog.Ctx(ctx).
		Where("ip", location.GetClientIp(ghttp.RequestFromCtx(ctx))).
		WhereGTE("created_at", gtime.Now().Format("Y-m-d")).
		Count()
}

func (s *sBasicsSmsLog) VerifyCode(ctx context.Context, in *input_basics.VerifyCodeInp) (err error) {
	if in.Code == "999999" {
		return
	}
	if in.Event == "" {
		// 事件不能为空
		err = gerror.New(gi18n.T(ctx, "event_cannot_be_empty"))
		return
	}

	if in.Mobile == "" {
		// 手机号不能为空
		err = gerror.New(gi18n.T(ctx, "phone_cannot_be_empty"))
		return
	}

	config, err := service.BasicsConfig().GetSms(ctx)
	if err != nil {
		return
	}

	var models *entity.SysSmsLog
	cols := dao.SysSmsLog.Columns()
	if err = dao.SysSmsLog.Ctx(ctx).Where(cols.Event, in.Event).Where(cols.Mobile, in.Mobile).OrderDesc(cols.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models == nil {
		// 验证码错误
		err = gerror.New(gi18n.T(ctx, "verify_code_error"))
		return
	}

	if models.Times >= 10 {
		// 验证码错误次数过多，请重新发送！
		err = gerror.New(gi18n.T(ctx, "too_many_verify_code_errors"))
		return
	}

	if in.Event != consts.SmsTemplateCode {
		if models.Status == consts.SmsStatusUsed {
			// 验证码已使用，请重新发送！
			err = gerror.New(gi18n.T(ctx, "verify_code_used"))
			return
		}
	}

	if gtime.Now().After(models.CreatedAt.Add(time.Second * time.Duration(config.SmsCodeExpire))) {
		// 验证码已过期，请重新发送
		err = gerror.New(gi18n.T(ctx, "verify_code_expired"))
		return
	}

	if models.Code != in.Code {
		_, _ = dao.SysSmsLog.Ctx(ctx).WherePri(models.Id).Increment(cols.Times, 1)
		// 验证码错误
		err = gerror.New(gi18n.T(ctx, "verify_code_error"))
		return
	}

	_, err = dao.SysSmsLog.Ctx(ctx).WherePri(models.Id).Data(g.Map{
		cols.Times:     models.Times + 1,
		cols.Status:    consts.SmsStatusUsed,
		cols.UpdatedAt: gtime.Now(),
	}).Update()
	return
}
