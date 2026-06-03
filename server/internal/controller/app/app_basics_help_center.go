package app

import (
	"APT/internal/library/contexts"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"APT/internal/service"
	"context"

	"APT/api/app/basics"
)

func (c *ControllerBasics) HelpCenterCategoryList(ctx context.Context, req *basics.HelpCenterCategoryListReq) (res *basics.HelpCenterCategoryListRes, err error) {
	res = new(basics.HelpCenterCategoryListRes)
	if res.List, _, err = service.BasicsHelpcenterCategory().List(ctx, &input_basics.PmsHelpcenterCategoryListInp{
		Language: contexts.GetLanguage(ctx),
		PageReq: input_form.PageReq{
			Pagination: false,
		}}); err != nil {
		return
	}
	return
}

func (c *ControllerBasics) HelpCenterList(ctx context.Context, req *basics.HelpCenterListReq) (res *basics.HelpCenterListRes, err error) {
	res = new(basics.HelpCenterListRes)
	if res.List, _, err = service.BasicsHelpcenter().List(ctx, &input_basics.PmsHelpcenterListInp{
		CategoryId: uint(req.CategoryId),
		Title:      req.Title,
		Language:   contexts.GetLanguage(ctx),
		PageReq: input_form.PageReq{
			Pagination: false,
		}}); err != nil {
		return
	}
	return
}

func (c *ControllerBasics) HelpCenterView(ctx context.Context, req *basics.HelpCenterViewReq) (res *basics.HelpCenterViewRes, err error) {
	res = new(basics.HelpCenterViewRes)
	if res.PmsHelpcenterViewModel, err = service.BasicsHelpcenter().View(ctx, &input_basics.PmsHelpcenterViewInp{Id: req.Id}); err != nil {
		return
	}
	return
}
