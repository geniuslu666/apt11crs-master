package admin

import (
	"APT/api/admin/travel"
	"APT/internal/model/input/input_travel"
	"APT/internal/service"
	"context"
)

func (c *ControllerTravel) VerifyStaffList(ctx context.Context, req *travel.VerifyStaffListReq) (res *travel.VerifyStaffListRes, err error) {
	list, totalCount, err := service.TravelVerifyStaff().List(ctx, &req.TravelVerifyStaffListInp)
	if err != nil {
		return
	}
	if list == nil {
		list = []*input_travel.TravelVerifyStaffListModel{}
	}
	res = new(travel.VerifyStaffListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

func (c *ControllerTravel) VerifyStaffView(ctx context.Context, req *travel.VerifyStaffViewReq) (res *travel.VerifyStaffViewRes, err error) {
	data, err := service.TravelVerifyStaff().View(ctx, &req.TravelVerifyStaffViewInp)
	if err != nil {
		return
	}
	res = new(travel.VerifyStaffViewRes)
	res.TravelVerifyStaffViewModel = data
	return
}

func (c *ControllerTravel) VerifyStaffEdit(ctx context.Context, req *travel.VerifyStaffEditReq) (res *travel.VerifyStaffEditRes, err error) {
	err = service.TravelVerifyStaff().Edit(ctx, &req.TravelVerifyStaffEditInp)
	return
}

func (c *ControllerTravel) VerifyStaffDelete(ctx context.Context, req *travel.VerifyStaffDeleteReq) (res *travel.VerifyStaffDeleteRes, err error) {
	err = service.TravelVerifyStaff().Delete(ctx, &req.TravelVerifyStaffDeleteInp)
	return
}

func (c *ControllerTravel) VerifyStaffStatus(ctx context.Context, req *travel.VerifyStaffStatusReq) (res *travel.VerifyStaffStatusRes, err error) {
	err = service.TravelVerifyStaff().Status(ctx, &req.TravelVerifyStaffStatusInp)
	return
}

func (c *ControllerTravel) VerifyStaffScopeOptions(ctx context.Context, req *travel.VerifyStaffScopeOptionsReq) (res *travel.VerifyStaffScopeOptionsRes, err error) {
	list, err := service.TravelVerifyStaff().ScopeOptions(ctx)
	if err != nil {
		return
	}
	if list == nil {
		list = []*input_travel.TravelVerifyStaffScopeOptionModel{}
	}
	res = &travel.VerifyStaffScopeOptionsRes{List: list}
	return
}
