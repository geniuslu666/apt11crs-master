package admin

import (
	"APT/api/admin/travel"
	"APT/internal/model/input/input_travel"
	"APT/internal/service"
	"context"
)

func (c *ControllerTravel) VerifyRecordList(ctx context.Context, req *travel.VerifyRecordListReq) (res *travel.VerifyRecordListRes, err error) {
	list, totalCount, err := service.TravelVerifyRecord().List(ctx, &req.TravelVerifyRecordListInp)
	if err != nil {
		return
	}
	if list == nil {
		list = []*input_travel.TravelVerifyRecordListModel{}
	}
	res = new(travel.VerifyRecordListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
