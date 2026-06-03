package admin

import (
	"APT/internal/model/input/input_hotel"
	"APT/internal/service"
	"context"

	"APT/api/admin/pms"
)

func (c *ControllerPms) RoomTypeList(ctx context.Context, req *pms.RoomTypeListReq) (res *pms.RoomTypeListRes, err error) {
	list, totalCount, err := service.HotelService().RoomTypeList(ctx, &req.PmsRoomTypeListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_hotel.PmsRoomTypeListModel{}
	}

	res = new(pms.RoomTypeListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerPms) RoomTypeView(ctx context.Context, req *pms.RoomTypeViewReq) (res *pms.RoomTypeViewRes, err error) {
	data, err := service.HotelService().RoomTypeView(ctx, &req.PmsRoomTypeViewInp)
	if err != nil {
		return
	}

	res = new(pms.RoomTypeViewRes)
	res.PmsRoomTypeViewModel = data
	return
}
func (c *ControllerPms) RoomTypeEdit(ctx context.Context, req *pms.RoomTypeEditReq) (res *pms.RoomTypeEditRes, err error) {
	err = service.HotelService().RoomTypeEdit(ctx, &req.PmsRoomTypeEditInp)
	return
}
func (c *ControllerPms) RoomTypeDelete(ctx context.Context, req *pms.RoomTypeDeleteReq) (res *pms.RoomTypeDeleteRes, err error) {
	err = service.HotelService().RoomTypeDelete(ctx, &req.PmsRoomTypeDeleteInp)
	return
}
func (c *ControllerPms) RoomTypeIsShow(ctx context.Context, req *pms.RoomTypeIsShowReq) (res *pms.RoomTypeIsShowRes, err error) {
	err = service.HotelService().RoomTypeStatus(ctx, &req.PmsRoomTypeIsShowInp)
	return
}
