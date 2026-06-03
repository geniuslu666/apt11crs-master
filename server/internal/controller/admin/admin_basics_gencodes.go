package admin

import (
	"APT/api/admin/basics"
	"APT/internal/service"
	"context"
)

func (c *ControllerBasics) GenCodesList(ctx context.Context, req *basics.GenCodesListReq) (res *basics.GenCodesListRes, err error) {
	list, totalCount, err := service.BasicsGenCodes().List(ctx, &req.GenCodesListInp)
	if err != nil {
		return
	}

	res = new(basics.GenCodesListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerBasics) GenCodesView(ctx context.Context, req *basics.GenCodesViewReq) (res *basics.GenCodesViewRes, err error) {
	data, err := service.BasicsGenCodes().View(ctx, &req.GenCodesViewInp)
	if err != nil {
		return
	}

	res = new(basics.GenCodesViewRes)
	res.GenCodesViewModel = data
	return
}
func (c *ControllerBasics) GenCodesEdit(ctx context.Context, req *basics.GenCodesEditReq) (res *basics.GenCodesEditRes, err error) {
	data, err := service.BasicsGenCodes().Edit(ctx, &req.GenCodesEditInp)
	if err != nil {
		return
	}

	res = new(basics.GenCodesEditRes)
	res.GenCodesEditModel = data
	return
}
func (c *ControllerBasics) GenCodesDelete(ctx context.Context, req *basics.GenCodesDeleteReq) (res *basics.GenCodesDeleteRes, err error) {
	err = service.BasicsGenCodes().Delete(ctx, &req.GenCodesDeleteInp)
	return
}
func (c *ControllerBasics) GenCodesMaxSort(ctx context.Context, req *basics.GenCodesMaxSortReq) (res *basics.GenCodesMaxSortRes, err error) {
	res = new(basics.GenCodesMaxSortRes)
	res.GenCodesMaxSortModel, err = service.BasicsGenCodes().MaxSort(ctx, &req.GenCodesMaxSortInp)
	return
}
func (c *ControllerBasics) GenCodesStatus(ctx context.Context, req *basics.GenCodesStatusReq) (res *basics.GenCodesStatusRes, err error) {
	err = service.BasicsGenCodes().Status(ctx, &req.GenCodesStatusInp)
	return
}
func (c *ControllerBasics) GenCodesSelects(ctx context.Context, req *basics.GenCodesSelectsReq) (res *basics.GenCodesSelectsRes, err error) {
	data, err := service.BasicsGenCodes().Selects(ctx, &req.GenCodesSelectsInp)
	if err != nil {
		return
	}

	res = new(basics.GenCodesSelectsRes)
	res.GenCodesSelectsModel = data
	return
}
func (c *ControllerBasics) GenCodesTableSelect(ctx context.Context, req *basics.GenCodesTableSelectReq) (res *basics.GenCodesTableSelectRes, err error) {
	data, err := service.BasicsGenCodes().TableSelect(ctx, &req.GenCodesTableSelectInp)
	if err != nil {
		return
	}

	res = (*basics.GenCodesTableSelectRes)(&data)
	return
}
func (c *ControllerBasics) GenCodesColumnSelect(ctx context.Context, req *basics.GenCodesColumnSelectReq) (res *basics.GenCodesColumnSelectRes, err error) {
	data, err := service.BasicsGenCodes().ColumnSelect(ctx, &req.GenCodesColumnSelectInp)
	if err != nil {
		return
	}

	res = (*basics.GenCodesColumnSelectRes)(&data)
	return
}
func (c *ControllerBasics) GenCodesColumnList(ctx context.Context, req *basics.GenCodesColumnListReq) (res *basics.GenCodesColumnListRes, err error) {
	data, err := service.BasicsGenCodes().ColumnList(ctx, &req.GenCodesColumnListInp)
	if err != nil {
		return
	}

	res = (*basics.GenCodesColumnListRes)(&data)
	return
}
func (c *ControllerBasics) GenCodesPreview(ctx context.Context, req *basics.GenCodesPreviewReq) (res *basics.GenCodesPreviewRes, err error) {
	data, err := service.BasicsGenCodes().Preview(ctx, &req.GenCodesPreviewInp)
	if err != nil {
		return
	}

	res = new(basics.GenCodesPreviewRes)
	res.GenCodesPreviewModel = data
	return
}
func (c *ControllerBasics) GenCodesBuild(ctx context.Context, req *basics.GenCodesBuildReq) (res *basics.GenCodesBuildRes, err error) {
	err = service.BasicsGenCodes().Build(ctx, &req.GenCodesBuildInp)
	//return nil, gerror.New("提交生成代码已停用")
	return
}
