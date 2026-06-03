package admin

import (
	"APT/api/admin/pms"
	"APT/internal/model/input/input_hotel"
	"APT/internal/service"
	"context"
)

func (c *ControllerPms) PricePlan(ctx context.Context, req *pms.PricePlanReq) (res *pms.PricePlanRes, err error) {
	list, totalCount, err := service.HotelService().HotelPricePlanList(ctx, req.SearchHotelPricePlanInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_hotel.SearchHotelPricePlanRes{}
	}

	res = new(pms.PricePlanRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerPms) PricePlanEdit(ctx context.Context, req *pms.PricePlanEditReq) (res *pms.PricePlanEditRes, err error) {
	res = new(pms.PricePlanEditRes)
	if err = service.HotelService().HotelPricePlanEdit(ctx, req.EditHotelPricePlanInp); err != nil {
		return
	}
	return
}
func (c *ControllerPms) PricePlanDel(ctx context.Context, req *pms.PricePlanDelReq) (res *pms.PricePlanDelRes, err error) {
	res = new(pms.PricePlanDelRes)
	if err = service.HotelService().HotelPricePlanDelete(ctx, req.DeleteHotelPricePlanInp); err != nil {
		return
	}
	return
}
func (c *ControllerPms) PricePlanOne(ctx context.Context, req *pms.PricePlanOneReq) (res *pms.PricePlanOneRes, err error) {
	res = new(pms.PricePlanOneRes)
	if res.FindOneHotelPricePlanRes, err = service.HotelService().HotelPricePlanGet(ctx, req.FindOneHotelPricePlanInp); err != nil {
		return
	}
	return
}
func (c *ControllerPms) PricePlanStatus(ctx context.Context, req *pms.PricePlanStatusReq) (res *pms.PricePlanStatusRes, err error) {
	err = service.HotelService().HotelPricePlanStatus(ctx, &req.PricePlanStatusInp)
	return
}
func (c *ControllerPms) PricePlanRecycle(ctx context.Context, req *pms.PricePlanRecycleReq) (res *pms.PricePlanRecycleRes, err error) {
	err = service.HotelService().HotelPricePlanRecycle(ctx, &req.DeleteHotelPricePlanInp)
	return
}
