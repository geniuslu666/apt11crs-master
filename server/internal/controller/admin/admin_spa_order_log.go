package admin

import (
	"APT/api/admin/spa"
	"APT/internal/model/input/input_spa"
	"APT/internal/service"
	"context"
)

func (c *ControllerSpa) OrderLogList(ctx context.Context, req *spa.OrderLogListReq) (res *spa.OrderLogListRes, err error) {
	list, totalCount, err := service.SpaOrderLog().List(ctx, &req.SpaOrderLogListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_spa.SpaOrderLogListModel{}
	}

	res = new(spa.OrderLogListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
