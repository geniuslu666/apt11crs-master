package admin

import (
	"APT/api/admin/spa"
	"APT/internal/model/input/input_spa"
	"APT/internal/service"
	"context"
)

func (c *ControllerSpa) TechnicianList(ctx context.Context, req *spa.TechnicianListReq) (res *spa.TechnicianListRes, err error) {
	list, totalCount, err := service.SpaTechnician().List(ctx, &req.SpaTechnicianListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_spa.SpaTechnicianListModel{}
	}

	res = new(spa.TechnicianListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerSpa) TechnicianAllList(ctx context.Context, req *spa.TechnicianAllListReq) (res *spa.TechnicianAllListRes, err error) {
	list, err := service.SpaTechnician().All(ctx, &req.SpaTechnicianListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_spa.SpaTechnicianAllListModel{}
	}

	res = new(spa.TechnicianAllListRes)
	res.List = list
	return
}
func (c *ControllerSpa) TechnicianView(ctx context.Context, req *spa.TechnicianViewReq) (res *spa.TechnicianViewRes, err error) {
	data, err := service.SpaTechnician().View(ctx, &req.SpaTechnicianViewInp)
	if err != nil {
		return
	}

	res = new(spa.TechnicianViewRes)
	res.SpaTechnicianViewModel = data
	return
}
func (c *ControllerSpa) TechnicianEdit(ctx context.Context, req *spa.TechnicianEditReq) (res *spa.TechnicianEditRes, err error) {
	err = service.SpaTechnician().Edit(ctx, &req.SpaTechnicianEditInp)
	return
}
func (c *ControllerSpa) TechnicianDelete(ctx context.Context, req *spa.TechnicianDeleteReq) (res *spa.TechnicianDeleteRes, err error) {
	err = service.SpaTechnician().Delete(ctx, &req.SpaTechnicianDeleteInp)
	return
}
func (c *ControllerSpa) TechnicianStatus(ctx context.Context, req *spa.TechnicianStatusReq) (res *spa.TechnicianStatusRes, err error) {
	err = service.SpaTechnician().Status(ctx, &req.SpaTechnicianStatusInp)
	return
}
func (c *ControllerSpa) TechnicianWorkStatus(ctx context.Context, req *spa.TechnicianWorkStatusReq) (res *spa.TechnicianWorkStatusRes, err error) {
	err = service.SpaTechnician().WorkStatus(ctx, &req.SpaTechnicianWorkStatusInp)
	return
}
func (c *ControllerSpa) TechnicianBind(ctx context.Context, req *spa.TechnicianBindReq) (res *spa.TechnicianBindRes, err error) {
	err = service.SpaTechnician().Bind(ctx, &req.SpaTechnicianBindInp)
	return
}
func (c *ControllerSpa) TechnicianUnbind(ctx context.Context, req *spa.TechnicianUnbindReq) (res *spa.TechnicianUnbindRes, err error) {
	err = service.SpaTechnician().Unbind(ctx, &req.SpaTechnicianUnbindInp)
	return
}
