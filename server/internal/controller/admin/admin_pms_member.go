package admin

import (
	"APT/api/admin/pms"
	"APT/internal/consts"
	"APT/internal/library/contexts"
	"APT/internal/library/geetest"
	"APT/internal/library/nxcloud"
	"APT/internal/library/telecom"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_app_member"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"APT/utility/charset"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerPms) MemberList(ctx context.Context, req *pms.MemberListReq) (res *pms.MemberListRes, err error) {
	list, totalCount, err := service.AppMember().MemberList(ctx, &req.PmsMemberListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_app_member.PmsMemberListModel{}
	}

	res = new(pms.MemberListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerPms) MemberSelect(ctx context.Context, req *pms.MemberSelectReq) (res *pms.MemberSelectRes, err error) {
	list, err := service.AppMember().MemberSelectList(ctx, &req.PmsMemberSelectInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_app_member.PmsMemberListModel{}
	}

	res = new(pms.MemberSelectRes)
	res.List = list
	return
}
func (c *ControllerPms) MemberExport(ctx context.Context, req *pms.MemberExportReq) (res *pms.MemberExportRes, err error) {
	err = service.AppMember().MemberExport(ctx, &req.PmsMemberListInp)
	return
}
func (c *ControllerPms) MemberView(ctx context.Context, req *pms.MemberViewReq) (res *pms.MemberViewRes, err error) {
	data, err := service.AppMember().MemberView(ctx, &req.PmsMemberViewInp)
	if err != nil {
		return
	}

	res = new(pms.MemberViewRes)
	res.PmsMemberViewModel = data
	return
}
func (c *ControllerPms) MemberEdit(ctx context.Context, req *pms.MemberEditReq) (res *pms.MemberEditRes, err error) {
	err = service.AppMember().MemberEdit(ctx, &req.PmsMemberEditInp)
	return
}
func (c *ControllerPms) MemberBaseEdit(ctx context.Context, req *pms.MemberBaseEditReq) (res *pms.MemberBaseEditRes, err error) {
	err = service.AppMember().MemberBaseEdit(ctx, &req.PmsMemberBaseEditInp)
	return
}
func (c *ControllerPms) MemberBalanceEdit(ctx context.Context, req *pms.MemberBalanceEditReq) (res *pms.MemberBalanceEditRes, err error) {
	err = service.AppMember().MemberBalanceEdit(ctx, &req.PmsMemberBalanceEditInp)
	return
}
func (c *ControllerPms) MemberExpEdit(ctx context.Context, req *pms.MemberExpEditReq) (res *pms.MemberExpEditRes, err error) {
	err = service.AppMember().MemberExpEdit(ctx, &req.PmsMemberExpEditInp)
	return
}
func (c *ControllerPms) MemberDelete(ctx context.Context, req *pms.MemberDeleteReq) (res *pms.MemberDeleteRes, err error) {
	err = service.AppMember().MemberDelete(ctx, &req.PmsMemberDeleteInp)
	return
}
func (c *ControllerPms) MemberSendEmail(ctx context.Context, req *pms.MemberSendEmailReq) (res *pms.MemberSendEmailRes, err error) {
	var (
		EmsParams input_basics.SendEmsInp
	)

	res = new(pms.MemberSendEmailRes)
	EmsParams.Event = "text"
	EmsParams.Email = req.Email
	EmsParams.Content = req.Content
	if err = service.BasicsEmsLog().Send(ctx, &EmsParams); err != nil {
		return
	}
	return
}
func (c *ControllerPms) MemberSendNotify(ctx context.Context, req *pms.MemberSendNotifyReq) (res *pms.MemberSendNotifyRes, err error) {
	err = service.BasicsAppNotify().NotifyEdit(ctx, &input_basics.PmsNotifyEditInp{
		PmsNotify: entity.PmsNotify{
			MemberId:      req.MemberId,
			NotifyTitle:   "SYSTEM",
			NotifyType:    req.Type, // 当前默认系统消息
			NotifyContent: req.Content,
		},
	})
	return
}
func (c *ControllerPms) MemberSendSms(ctx context.Context, req *pms.MemberSendSmsReq) (res *pms.MemberSendSmsRes, err error) {
	var (
		SmsParams input_basics.SendCodeInp
	)
	if req.AreaNo == "+86" {
		SmsParams.Code = gvar.New(charset.RandomCreateBytes(6, []byte(`0123456789`)...)).String()
		SmsParams.Event = "login"
		SmsParams.Mobile = req.Phone
		if err = telecom.SendSms(ctx, req.Phone, req.Content); err != nil {
			return
		}
	} else {
		Phone := fmt.Sprintf("%s%s", req.AreaNo[1:], req.Phone)
		err = nxcloud.SendSms(ctx, Phone, req.Content)
	}

	return
}
func (c *ControllerPms) MemberStat(ctx context.Context, req *pms.MemberStatReq) (res *pms.MemberStatRes, err error) {
	data, err := service.AppMember().MemberStat(ctx, &req.PmsMemberStatInp)

	res = new(pms.MemberStatRes)
	res.PmsMemberStatModel = data
	return
}
func (c *ControllerPms) MemberStatus(ctx context.Context, req *pms.MemberStatusReq) (res *pms.MemberStatusRes, err error) {
	err = service.AppMember().Status(ctx, &req.PmsMemberStatusInp)
	return
}
func (c *ControllerPms) MemberCancel(ctx context.Context, req *pms.MemberCancelReq) (res *pms.MemberCancelRes, err error) {
	err = service.AppMember().Cancel(ctx, &req.PmsMemberCancelInp)
	return
}
func (c *ControllerPms) MemberSendMsg(ctx context.Context, req *pms.MemberSendMsgReq) (res *pms.MemberSendMsgRes, err error) {
	// 发送到消息队列
	err = service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
		SystemMessageTitle: map[string]string{
			"zh":    req.Title,
			"en":    req.Title,
			"ja":    req.Title,
			"ko":    req.Title,
			"zh_CN": req.Title,
		},
		SystemMessageContent: map[string]string{
			"zh":    req.Content,
			"en":    req.Content,
			"ja":    req.Content,
			"ko":    req.Content,
			"zh_CN": req.Content,
		},
		Scene:        "system",
		Type:         "im",
		MemberId:     req.MemberId,
		EnablePush:   true,
		EnableSms:    false,
		PushTitle:    req.Title,
		PushContent:  req.Content,
		OperatorId:   int(contexts.GetUserId(ctx)),
		OperatorRole: "ADMIN",
	})
	if err != nil {
		return
	}

	return
}
func (c *ControllerPms) MemberSendRegisterCode(ctx context.Context, req *pms.MemberSendRegisterCodeReq) (res *pms.MemberSendRegisterCodeRes, err error) {
	var (
		SmsParams input_basics.SendCodeInp
		gtConf    *model.GeeTestConfig
	)
	res = new(pms.MemberSendRegisterCodeRes)

	if req.Type == 1 {
		if gtConf, err = service.BasicsConfig().GetGeeTest(ctx); err != nil {
			return
		}
		if gtConf != nil && gtConf.Enabled {
			if _, err = geetest.Validate(ctx, geetest.Config{
				Enabled:     gtConf.Enabled,
				CaptchaID:   gtConf.CaptchaID,
				CaptchaKey:  gtConf.CaptchaKey,
				ValidateURL: gtConf.ValidateURL,
			}, geetest.ValidateParams{
				CaptchaOutput: req.CaptchaOutput,
				LotNumber:     req.LotNumber,
				PassToken:     req.PassToken,
				GenTime:       req.GenTime,
			}); err != nil {
				return nil, gerror.Wrap(err, "verify failed")
			}
		}

		// 发送短信
		SmsParams.Code = gvar.New(charset.RandomCreateBytes(6, []byte(`0123456789`)...)).String()
		switch req.AreaNo {
		case "+86":
			SmsParams.Event = "login"
			SmsParams.SmsDrive = consts.SmsDriveUms
			break
		case "+81":
			SmsParams.Event = "login-ja"
			break
		case "+82":
			SmsParams.Event = "login-ko"
			break
		default:
			SmsParams.Event = "login-en"
			break
		}
		SmsParams.Mobile = fmt.Sprintf("%s%s", req.AreaNo, req.Phone)
		if err = service.BasicsSmsLog().SendCode(ctx, &SmsParams); err != nil {
			return
		}
	} else {
		// 发送邮件
		var (
			EmsParams *input_basics.SendEmsInp
		)
		EmsParams = new(input_basics.SendEmsInp)
		EmsParams.Code = gvar.New(charset.RandomCreateBytes(6, []byte(`0123456789`)...)).String()
		EmsParams.Event = "login"
		EmsParams.Email = req.EMail
		if err = service.BasicsEmsLog().Send(ctx, EmsParams); err != nil {
			return
		}
	}
	return
}
