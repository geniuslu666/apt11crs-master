package admin

import (
	"APT/internal/model/input/input_hotel"
	"APT/internal/service"
	"context"

	"APT/api/admin/pms"
)

func (c *ControllerPms) OrderLogList(ctx context.Context, req *pms.OrderLogListReq) (res *pms.OrderLogListRes, err error) {
	list, totalCount, err := service.HotelService().AppStayLogList(ctx, &req.AppStayLogListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_hotel.AppStayLogListModel{}
	}

	res = new(pms.OrderLogListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
