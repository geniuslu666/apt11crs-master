package admin

import (
	"APT/internal/model/input/input_hotel"
	"APT/internal/service"
	"context"

	"APT/api/admin/pms"
)

func (c *ControllerPms) OrderChangeList(ctx context.Context, req *pms.OrderChangeListReq) (res *pms.OrderChangeListRes, err error) {
	list, totalCount, err := service.HotelService().OrderChangeList(ctx, &req.OrderChangeListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_hotel.OrderChangeListModel{}
	}

	res = new(pms.OrderChangeListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerPms) OrderChangeExport(ctx context.Context, req *pms.OrderChangeExportReq) (res *pms.OrderChangeExportRes, err error) {
	err = service.HotelService().OrderChangeExport(ctx, &req.OrderChangeListInp)
	return
}
func (c *ControllerPms) OrderChangeView(ctx context.Context, req *pms.OrderChangeViewReq) (res *pms.OrderChangeViewRes, err error) {
	data, err := service.HotelService().OrderChangeView(ctx, &req.OrderChangeViewInp)
	if err != nil {
		return
	}

	res = new(pms.OrderChangeViewRes)
	res.OrderChangeViewModel = data
	return
}
func (c *ControllerPms) OrderChangeEdit(ctx context.Context, req *pms.OrderChangeEditReq) (res *pms.OrderChangeEditRes, err error) {
	err = service.HotelService().OrderChangeEdit(ctx, &req.OrderChangeEditInp)
	return
}
func (c *ControllerPms) OrderChangeDelete(ctx context.Context, req *pms.OrderChangeDeleteReq) (res *pms.OrderChangeDeleteRes, err error) {
	err = service.HotelService().OrderChangeDelete(ctx, &req.OrderChangeDeleteInp)
	return
}
