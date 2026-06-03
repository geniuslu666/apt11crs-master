package admin

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"

	"APT/api/admin/pms"
)

func (c *ControllerPms) HelpCenterList(ctx context.Context, req *pms.HelpCenterListReq) (res *pms.HelpCenterListRes, err error) {
	list, totalCount, err := service.BasicsHelpcenter().List(ctx, &req.PmsHelpcenterListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_basics.PmsHelpcenterListModel{}
	}

	res = new(pms.HelpCenterListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerPms) HelpCenterView(ctx context.Context, req *pms.HelpCenterViewReq) (res *pms.HelpCenterViewRes, err error) {
	data, err := service.BasicsHelpcenter().View(ctx, &req.PmsHelpcenterViewInp)
	if err != nil {
		return
	}

	res = new(pms.HelpCenterViewRes)
	res.PmsHelpcenterViewModel = data
	return
}
func (c *ControllerPms) HelpCenterEdit(ctx context.Context, req *pms.HelpCenterEditReq) (res *pms.HelpCenterEditRes, err error) {
	err = service.BasicsHelpcenter().Edit(ctx, &req.PmsHelpcenterEditInp)
	return
}
func (c *ControllerPms) HelpCenterDelete(ctx context.Context, req *pms.HelpCenterDeleteReq) (res *pms.HelpCenterDeleteRes, err error) {
	err = service.BasicsHelpcenter().Delete(ctx, &req.PmsHelpcenterDeleteInp)
	return
}
