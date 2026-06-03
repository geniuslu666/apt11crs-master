package app

import (
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/token"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_travel"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/app/travel"
)

func (c *ControllerTravel) StaffLogin(ctx context.Context, req *travel.StaffLoginReq) (res *travel.StaffLoginRes, err error) {
	var (
		TravelVerifyStaff *entity.TravelVerifyStaff
		StaffTokenInfo    *model.TravelStaffIdentity
	)

	res = new(travel.StaffLoginRes)

	// 一日游核销人员
	if err = dao.TravelVerifyStaff.Ctx(ctx).Where(dao.TravelVerifyStaff.Columns().Username, req.Username).Scan(&TravelVerifyStaff); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}

	// 核销账号不存在
	if g.IsEmpty(TravelVerifyStaff) {
		err = gerror.New(gi18n.T(ctx, "account_does_not_exist"))
		return
	}

	// 密码不正确
	if TravelVerifyStaff.PasswordHash != gmd5.MustEncryptString(req.Password+TravelVerifyStaff.Salt) {
		err = gerror.New(gi18n.T(ctx, "password_incorrect"))
		return
	}

	res.Name = TravelVerifyStaff.Name
	res.Mobile = TravelVerifyStaff.Mobile
	res.Username = TravelVerifyStaff.Username

	if err = gvar.New(TravelVerifyStaff).Struct(&StaffTokenInfo); err != nil {
		return
	}
	StaffTokenInfo.LoginAt = gtime.Now()
	if res.Token, res.Expires, err = token.TravelStaffLogin(ctx, StaffTokenInfo); err != nil {
		return
	}

	return
}
func (c *ControllerTravel) StaffLogout(ctx context.Context, req *travel.StaffLogoutReq) (res *travel.StaffLogoutRes, err error) {
	err = token.TravelStaffLogout(ghttp.RequestFromCtx(ctx))
	return
}
func (c *ControllerTravel) VerifyLog(ctx context.Context, req *travel.VerifyLogReq) (res *travel.VerifyLogRes, err error) {
	res = new(travel.VerifyLogRes)

	MemberInfo := contexts.GetTravelStaffUser(ctx)

	if res.List, res.Count, err = service.TravelVerifyRecord().VerifyList(ctx, &input_travel.VerifyListInp{
		PageReq: input_form.PageReq{
			Page:       req.PageNum,
			PerPage:    req.PageSize,
			Pagination: true,
		},
		VerifyStaffId: int64(MemberInfo.Id),
	}); err != nil {
		return
	}
	return
}
func (c *ControllerTravel) VerifyLogView(ctx context.Context, req *travel.VerifyLogViewReq) (res *travel.VerifyLogViewRes, err error) {
	res = new(travel.VerifyLogViewRes)
	if res.VerifyLogViewModel, err = service.TravelVerifyRecord().VerifyView(ctx, &req.VerifyLogViewInp); err != nil {
		return
	}
	return
}
func (c *ControllerTravel) CodeView(ctx context.Context, req *travel.CodeViewReq) (res *travel.CodeViewRes, err error) {
	res = new(travel.CodeViewRes)
	if res.CodeViewModel, err = service.TravelVerifyRecord().CodeView(ctx, &req.CodeViewInp); err != nil {
		return
	}
	return
}
func (c *ControllerTravel) CodeVerify(ctx context.Context, req *travel.CodeVerifyReq) (res *travel.CodeVerifyRes, err error) {
	res = new(travel.CodeVerifyRes)
	if res.CodeVerifyModel, err = service.TravelVerifyRecord().CodeVerify(ctx, &req.CodeVerifyInp); err != nil {
		return
	}
	return
}
func (c *ControllerTravel) StaffConfig(ctx context.Context, req *travel.StaffConfigReq) (res *travel.StaffConfigRes, err error) {
	res = new(travel.StaffConfigRes)

	// 从sys_config表获取一日游其他配置（参考app_travel_order_info.go）
	config, err := service.BasicsConfig().GetConfigByGroup(ctx, &input_basics.GetConfigInp{
		Group: "travelothersetting",
	})
	// 设置客服电话内容
	res.ContactMobile = "" // 默认空值
	if err == nil && config != nil && config.List != nil {
		if contactMobileValue := config.List["contactMobile"]; contactMobileValue != nil {
			res.ContactMobile = gvar.New(contactMobileValue).String()
		}
	}

	return
}
