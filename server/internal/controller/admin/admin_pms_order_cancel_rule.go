package admin

import (
	"APT/internal/model/input/input_hotel"
	"APT/internal/service"
	"context"

	"APT/api/admin/pms"
)

func (c *ControllerPms) CancelOrderRuleList(ctx context.Context, req *pms.CancelOrderRuleListReq) (res *pms.CancelOrderRuleListRes, err error) {
	list, _, err := service.HotelService().CancelRateList(ctx, nil)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_hotel.PmsCancelRateListModel{}
	}

	res = new(pms.CancelOrderRuleListRes)
	res.List = list
	return
}
func (c *ControllerPms) CancelOrderRuleView(ctx context.Context, req *pms.CancelOrderRuleViewReq) (res *pms.CancelOrderRuleViewRes, err error) {
	data, err := service.HotelService().CancelRateView(ctx, &req.PmsCancelRateViewInp)
	if err != nil {
		return
	}

	res = new(pms.CancelOrderRuleViewRes)
	res.PmsCancelRateViewModel = data
	return
}
func (c *ControllerPms) CancelOrderRuleEdit(ctx context.Context, req *pms.CancelOrderRuleEditReq) (res *pms.CancelOrderRuleEditRes, err error) {
	err = service.HotelService().CancelRateEdit(ctx, &req.PmsCancelRateEditInp)
	return
}
func (c *ControllerPms) CancelOrderRuleDelete(ctx context.Context, req *pms.CancelOrderRuleDeleteReq) (res *pms.CancelOrderRuleDeleteRes, err error) {
	err = service.HotelService().CancelRateDelete(ctx, &req.PmsCancelRateDeleteInp)
	return
}
func (c *ControllerPms) CancelOrderRuleMaxSort(ctx context.Context, req *pms.CancelOrderRuleMaxSortReq) (res *pms.CancelOrderRuleMaxSortRes, err error) {
	data, err := service.HotelService().CancelRateMaxSort(ctx, &req.PmsCancelRateMaxSortInp)
	if err != nil {
		return
	}

	res = new(pms.CancelOrderRuleMaxSortRes)
	res.PmsCancelRateMaxSortModel = data
	return
}
