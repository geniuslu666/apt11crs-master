package admin

import (
	"APT/api/admin/pms"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"
)

func (c *ControllerPms) HelpCenterCategoryList(ctx context.Context, req *pms.HelpCenterCategoryListReq) (res *pms.HelpCenterCategoryListRes, err error) {
	list, totalCount, err := service.BasicsHelpcenterCategory().List(ctx, &req.PmsHelpcenterCategoryListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_basics.PmsHelpcenterCategoryListModel{}
	}

	res = new(pms.HelpCenterCategoryListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerPms) HelpCenterCategoryAll(ctx context.Context, req *pms.HelpCenterCategoryAllReq) (res *pms.HelpCenterCategoryAllRes, err error) {
	list, err := service.BasicsHelpcenterCategory().All(ctx, &req.PmsHelpcenterCategoryAllInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_basics.PmsHelpcenterCategoryAllModel{}
	}

	res = new(pms.HelpCenterCategoryAllRes)
	res.List = list
	return
}
func (c *ControllerPms) HelpCenterCategoryView(ctx context.Context, req *pms.HelpCenterCategoryViewReq) (res *pms.HelpCenterCategoryViewRes, err error) {
	data, err := service.BasicsHelpcenterCategory().View(ctx, &req.PmsHelpcenterCategoryViewInp)
	if err != nil {
		return
	}

	res = new(pms.HelpCenterCategoryViewRes)
	res.PmsHelpcenterCategoryViewModel = data
	return
}
func (c *ControllerPms) HelpCenterCategoryEdit(ctx context.Context, req *pms.HelpCenterCategoryEditReq) (res *pms.HelpCenterCategoryEditRes, err error) {
	err = service.BasicsHelpcenterCategory().Edit(ctx, &req.PmsHelpcenterCategoryEditInp)
	return
}
func (c *ControllerPms) HelpCenterCategoryDelete(ctx context.Context, req *pms.HelpCenterCategoryDeleteReq) (res *pms.HelpCenterCategoryDeleteRes, err error) {
	err = service.BasicsHelpcenterCategory().Delete(ctx, &req.PmsHelpcenterCategoryDeleteInp)
	return
}
