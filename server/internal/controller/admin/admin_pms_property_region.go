package admin

import (
	"APT/api/admin/pms"
	"APT/internal/model/input/input_hotel"
	"APT/internal/service"
	"context"
)

func (c *ControllerPms) RegionList(ctx context.Context, req *pms.RegionListReq) (res *pms.RegionListRes, err error) {
	list, totalCount, err := service.HotelService().RegionList(ctx, &req.PropertyRegionListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_hotel.PropertyRegionListModel{}
	}

	res = new(pms.RegionListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
