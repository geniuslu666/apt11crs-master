package admin

import (
	"APT/api/admin/basics"
	"APT/internal/service"
	"context"
)

func (c *ControllerBasics) SmsLogList(ctx context.Context, req *basics.SmsLogListReq) (res *basics.SmsLogListRes, err error) {
	list, totalCount, err := service.BasicsSmsLog().List(ctx, &req.SmsLogListInp)
	if err != nil {
		return
	}

	res = new(basics.SmsLogListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerBasics) SmsLogView(ctx context.Context, req *basics.SmsLogViewReq) (res *basics.SmsLogViewRes, err error) {
	data, err := service.BasicsSmsLog().View(ctx, &req.SmsLogViewInp)
	if err != nil {
		return
	}

	res = new(basics.SmsLogViewRes)
	res.SmsLogViewModel = data
	return
}
func (c *ControllerBasics) SmsLogDelete(ctx context.Context, req *basics.SmsLogDeleteReq) (res *basics.SmsLogDeleteRes, err error) {
	err = service.BasicsSmsLog().Delete(ctx, &req.SmsLogDeleteInp)
	return
}
