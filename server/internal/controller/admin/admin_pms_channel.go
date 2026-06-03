package admin

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"

	"APT/api/admin/pms"
)

func (c *ControllerPms) ChannelList(ctx context.Context, req *pms.ChannelListReq) (res *pms.ChannelListRes, err error) {
	list, totalCount, err := service.BasicsChannel().List(ctx, &req.PmsChannelListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_basics.PmsChannelListModel{}
	}

	res = new(pms.ChannelListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerPms) ChannelView(ctx context.Context, req *pms.ChannelViewReq) (res *pms.ChannelViewRes, err error) {
	data, err := service.BasicsChannel().View(ctx, &req.PmsChannelViewInp)
	if err != nil {
		return
	}

	res = new(pms.ChannelViewRes)
	res.PmsChannelViewModel = data
	return
}
func (c *ControllerPms) ChannelEdit(ctx context.Context, req *pms.ChannelEditReq) (res *pms.ChannelEditRes, err error) {
	err = service.BasicsChannel().Edit(ctx, &req.PmsChannelEditInp)
	return
}
func (c *ControllerPms) ChannelDelete(ctx context.Context, req *pms.ChannelDeleteReq) (res *pms.ChannelDeleteRes, err error) {
	err = service.BasicsChannel().Delete(ctx, &req.PmsChannelDeleteInp)
	return
}
func (c *ControllerPms) ChannelStatus(ctx context.Context, req *pms.ChannelStatusReq) (res *pms.ChannelStatusRes, err error) {
	err = service.BasicsChannel().Status(ctx, &req.PmsChannelStatusInp)
	return
}
func (c *ControllerPms) ChannelBind(ctx context.Context, req *pms.ChannelBindReq) (res *pms.ChannelBindRes, err error) {
	err = service.BasicsChannel().Bind(ctx, &req.PmsChannelBindInp)
	return
}
func (c *ControllerPms) ChannelUnbind(ctx context.Context, req *pms.ChannelUnbindReq) (res *pms.ChannelUnbindRes, err error) {
	err = service.BasicsChannel().Unbind(ctx, &req.PmsChannelUnbindInp)
	return
}
