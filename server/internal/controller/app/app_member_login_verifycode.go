package app

import (
	"APT/api/app/member"
	"APT/internal/consts"
	"APT/internal/library/geetest"
	"APT/internal/model"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"APT/utility/charset"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerMember) EmailSendCode(ctx context.Context, req *member.EmailSendCodeReq) (res *member.EmailSendCodeRes, err error) {
	var (
		EmsParams *input_basics.SendEmsInp
	)
	EmsParams = new(input_basics.SendEmsInp)
	res = new(member.EmailSendCodeRes)
	EmsParams.Code = gvar.New(charset.RandomCreateBytes(6, []byte(`0123456789`)...)).String()
	EmsParams.Event = "login"
	EmsParams.Email = req.EMail
	if err = service.BasicsEmsLog().Send(ctx, EmsParams); err != nil {
		return
	}
	return
}

func (c *ControllerMember) SmsSendCode(ctx context.Context, req *member.SmsSendCodeReq) (res *member.SmsSendCodeRes, err error) {
	var (
		SmsParams input_basics.SendCodeInp
		gtConf    *model.GeeTestConfig
	)
	res = new(member.SmsSendCodeRes)
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

	return
}
