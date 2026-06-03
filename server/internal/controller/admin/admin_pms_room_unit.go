package admin

import (
	"APT/internal/model/input/input_hotel"
	"APT/internal/service"
	"context"

	"APT/api/admin/pms"
)

func (c *ControllerPms) RoomUnitList(ctx context.Context, req *pms.RoomUnitListReq) (res *pms.RoomUnitListRes, err error) {
	list, totalCount, err := service.HotelService().RoomUnitList(ctx, &req.PmsRoomUnitListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_hotel.PmsRoomUnitListModel{}
	}

	res = new(pms.RoomUnitListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerPms) RoomUnitExport(ctx context.Context, req *pms.RoomUnitExportReq) (res *pms.RoomUnitExportRes, err error) {
	err = service.HotelService().RoomUnitExport(ctx, &req.PmsRoomUnitListInp)
	return
}
func (c *ControllerPms) RoomUnitView(ctx context.Context, req *pms.RoomUnitViewReq) (res *pms.RoomUnitViewRes, err error) {
	data, err := service.HotelService().RoomUnitView(ctx, &req.PmsRoomUnitViewInp)
	if err != nil {
		return
	}

	res = new(pms.RoomUnitViewRes)
	res.PmsRoomUnitViewModel = data
	return
}
func (c *ControllerPms) RoomUnitEdit(ctx context.Context, req *pms.RoomUnitEditReq) (res *pms.RoomUnitEditRes, err error) {
	err = service.HotelService().RoomUnitEdit(ctx, &req.PmsRoomUnitEditInp)
	return
}
func (c *ControllerPms) RoomUnitDelete(ctx context.Context, req *pms.RoomUnitDeleteReq) (res *pms.RoomUnitDeleteRes, err error) {
	err = service.HotelService().RoomUnitDelete(ctx, &req.PmsRoomUnitDeleteInp)
	return
}
