package admin

import (
	"APT/internal/model/input/input_spa"
	"APT/internal/service"
	"context"

	"APT/api/admin/spa"
)

func (c *ControllerSpa) LabelList(ctx context.Context, req *spa.LabelListReq) (res *spa.LabelListRes, err error) {
	list, totalCount, err := service.SpaLabel().List(ctx, &req.SpaLabelListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_spa.SpaLabelListModel{}
	}

	res = new(spa.LabelListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerSpa) LabelView(ctx context.Context, req *spa.LabelViewReq) (res *spa.LabelViewRes, err error) {
	data, err := service.SpaLabel().View(ctx, &req.SpaLabelViewInp)
	if err != nil {
		return
	}

	res = new(spa.LabelViewRes)
	res.SpaLabelViewModel = data
	return
}
func (c *ControllerSpa) LabelEdit(ctx context.Context, req *spa.LabelEditReq) (res *spa.LabelEditRes, err error) {
	err = service.SpaLabel().Edit(ctx, &req.SpaLabelEditInp)
	return
}
func (c *ControllerSpa) LabelDelete(ctx context.Context, req *spa.LabelDeleteReq) (res *spa.LabelDeleteRes, err error) {
	err = service.SpaLabel().Delete(ctx, &req.SpaLabelDeleteInp)
	return
}
func (c *ControllerSpa) LabelMaxSort(ctx context.Context, req *spa.LabelMaxSortReq) (res *spa.LabelMaxSortRes, err error) {
	data, err := service.SpaLabel().MaxSort(ctx, &req.SpaLabelMaxSortInp)
	if err != nil {
		return
	}

	res = new(spa.LabelMaxSortRes)
	res.SpaLabelMaxSortModel = data
	return
}
func (c *ControllerSpa) LabelStatus(ctx context.Context, req *spa.LabelStatusReq) (res *spa.LabelStatusRes, err error) {
	err = service.SpaLabel().Status(ctx, &req.SpaLabelStatusInp)
	return
}
