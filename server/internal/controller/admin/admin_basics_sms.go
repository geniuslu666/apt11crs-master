package admin

import (
	"APT/api/admin/basics"
	"APT/internal/consts"
	"APT/internal/library/contexts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerBasics) SendTestSms(ctx context.Context, req *basics.SendTestSmsReq) (res *basics.SendTestSmsRes, err error) {
	err = service.BasicsSmsLog().SendCode(ctx, &req.SendCodeInp)
	return
}
func (c *ControllerBasics) SendBindSms(ctx context.Context, _ *basics.SendBindSmsReq) (res *basics.SendBindSmsRes, err error) {
	var (
		memberId = contexts.GetUserId(ctx)
		models   *entity.AdminMember
	)

	if memberId <= 0 {
		err = gerror.New("用户身份异常，请重新登录！")
		return
	}

	if err = g.Model("admin_member").Fields("mobile").Where("id", memberId).Scan(&models); err != nil {
		return
	}

	if models == nil {
		err = gerror.New("用户信息不存在")
		return
	}

	if models.Mobile == "" {
		err = gerror.New("未绑定手机号无需发送")
		return
	}

	err = service.BasicsSmsLog().SendCode(ctx, &input_basics.SendCodeInp{
		Event:  consts.SmsTemplateBind,
		Mobile: models.Mobile,
	})
	return
}
func (c *ControllerBasics) SendSms(ctx context.Context, req *basics.SendSmsReq) (res *basics.SendSmsRes, err error) {
	err = service.BasicsSmsLog().SendCode(ctx, &req.SendCodeInp)
	return
}
