package admin

import (
	"APT/internal/library/contexts"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/admin/basics"
)

func (c *ControllerBasics) MemberUpdateEmail(ctx context.Context, req *basics.MemberUpdateEmailReq) (res *basics.MemberUpdateEmailRes, err error) {
	err = service.BasicsAdminMember().UpdateEmail(ctx, &req.MemberUpdateEmailInp)
	return
}
func (c *ControllerBasics) MemberUpdateMobile(ctx context.Context, req *basics.MemberUpdateMobileReq) (res *basics.MemberUpdateMobileRes, err error) {
	err = service.BasicsAdminMember().UpdateMobile(ctx, &req.MemberUpdateMobileInp)
	return
}
func (c *ControllerBasics) MemberUpdateProfile(ctx context.Context, req *basics.MemberUpdateProfileReq) (res *basics.MemberUpdateProfileRes, err error) {
	err = service.BasicsAdminMember().UpdateProfile(ctx, &req.MemberUpdateProfileInp)
	return
}
func (c *ControllerBasics) MemberUpdatePwd(ctx context.Context, req *basics.MemberUpdatePwdReq) (res *basics.MemberUpdatePwdRes, err error) {
	user := contexts.GetUser(ctx)
	if user == nil {
		err = gerror.New("获取用户信息失败！")
		return
	}
	memberId := user.Id
	if memberId <= 0 {
		err = gerror.New("获取用户信息失败！")
		return
	}

	err = service.BasicsAdminMember().UpdatePwd(ctx, &input_basics.MemberUpdatePwdInp{
		Id:          memberId,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})
	return
}
func (c *ControllerBasics) MemberResetPwd(ctx context.Context, req *basics.MemberResetPwdReq) (res *basics.MemberResetPwdRes, err error) {
	err = service.BasicsAdminMember().ResetPwd(ctx, &req.MemberResetPwdInp)
	return
}
func (c *ControllerBasics) MemberList(ctx context.Context, req *basics.MemberListReq) (res *basics.MemberListRes, err error) {
	list, totalCount, err := service.BasicsAdminMember().List(ctx, &req.MemberListInp)
	if err != nil {
		return
	}

	res = new(basics.MemberListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerBasics) MemberView(ctx context.Context, req *basics.MemberViewReq) (res *basics.MemberViewRes, err error) {
	data, err := service.BasicsAdminMember().View(ctx, &req.MemberViewInp)
	if err != nil {
		return
	}

	res = new(basics.MemberViewRes)
	res.MemberViewModel = data
	return
}
func (c *ControllerBasics) MemberEdit(ctx context.Context, req *basics.MemberEditReq) (res *basics.MemberEditRes, err error) {
	err = service.BasicsAdminMember().Edit(ctx, &req.MemberEditInp)
	return
}
func (c *ControllerBasics) MemberDelete(ctx context.Context, req *basics.MemberDeleteReq) (res *basics.MemberDeleteRes, err error) {
	err = service.BasicsAdminMember().Delete(ctx, &req.MemberDeleteInp)
	return
}
func (c *ControllerBasics) MemberStatus(ctx context.Context, req *basics.MemberStatusReq) (res *basics.MemberStatusRes, err error) {
	err = service.BasicsAdminMember().Status(ctx, &req.MemberStatusInp)
	return
}
func (c *ControllerBasics) MemberSelect(ctx context.Context, req *basics.MemberSelectReq) (res *basics.MemberSelectRes, err error) {
	data, err := service.BasicsAdminMember().Select(ctx, &req.MemberSelectInp)
	if err != nil {
		return
	}

	res = (*basics.MemberSelectRes)(&data)
	return
}
func (c *ControllerBasics) MemberInfo(ctx context.Context, req *basics.MemberInfoReq) (res *basics.MemberInfoRes, err error) {
	data, err := service.BasicsAdminMember().LoginMemberInfo(ctx)
	if err != nil {
		return
	}

	res = new(basics.MemberInfoRes)
	res.LoginMemberInfoModel = data
	return
}
