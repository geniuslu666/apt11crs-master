package app

import (
	"APT/internal/library/contexts"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"

	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/app/basics"
)

func (c *ControllerBasics) NotifyList(ctx context.Context, req *basics.NotifyListReq) (res *basics.NotifyListRes, err error) {
	var (
		MemberInfo = contexts.GetMemberUser(ctx)
	)
	if g.IsEmpty(MemberInfo.Id) {
		err = gerror.New(gi18n.T(ctx, "abnormal_user_identity"))
		return
	}
	req.MemberId = MemberInfo.Id
	list, totalCount, err := service.BasicsAppNotify().NotifyList(ctx, &req.PmsNotifyListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_basics.PmsNotifyListModel{}
	}

	res = new(basics.NotifyListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerBasics) NotifyView(ctx context.Context, req *basics.NotifyViewReq) (res *basics.NotifyViewRes, err error) {
	data, err := service.BasicsAppNotify().NotifyView(ctx, &req.PmsNotifyViewInp)
	if err != nil {
		return
	}

	res = new(basics.NotifyViewRes)
	res.PmsNotifyViewModel = data
	return
}
func (c *ControllerBasics) NotifyEdit(ctx context.Context, req *basics.NotifyEditReq) (res *basics.NotifyEditRes, err error) {
	err = service.BasicsAppNotify().NotifyEdit(ctx, &req.PmsNotifyEditInp)
	return
}
func (c *ControllerBasics) NotifyDelete(ctx context.Context, req *basics.NotifyDeleteReq) (res *basics.NotifyDeleteRes, err error) {
	err = service.BasicsAppNotify().NotifyDelete(ctx, &req.PmsNotifyDeleteInp)
	return
}
