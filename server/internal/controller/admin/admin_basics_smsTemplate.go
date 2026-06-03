package admin

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"

	"APT/api/admin/basics"
)

func (c *ControllerBasics) SmsTemplateList(ctx context.Context, req *basics.SmsTemplateListReq) (res *basics.SmsTemplateListRes, err error) {
	list, totalCount, err := service.BasicsSmsTemplate().List(ctx, &req.SmsTemplateListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_basics.SmsTemplateListModel{}
	}

	res = new(basics.SmsTemplateListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerBasics) SmsTemplateView(ctx context.Context, req *basics.SmsTemplateViewReq) (res *basics.SmsTemplateViewRes, err error) {
	data, err := service.BasicsSmsTemplate().View(ctx, &req.SmsTemplateViewInp)
	if err != nil {
		return
	}

	res = new(basics.SmsTemplateViewRes)
	res.SmsTemplateViewModel = data
	return
}
func (c *ControllerBasics) SmsTemplateEdit(ctx context.Context, req *basics.SmsTemplateEditReq) (res *basics.SmsTemplateEditRes, err error) {
	err = service.BasicsSmsTemplate().Edit(ctx, &req.SmsTemplateEditInp)
	return
}
func (c *ControllerBasics) SmsTemplateDelete(ctx context.Context, req *basics.SmsTemplateDeleteReq) (res *basics.SmsTemplateDeleteRes, err error) {
	err = service.BasicsSmsTemplate().Delete(ctx, &req.SmsTemplateDeleteInp)
	return
}
func (c *ControllerBasics) SmsTemplateLog(ctx context.Context, req *basics.SmsTemplateLogReq) (res *basics.SmsTemplateLogRes, err error) {
	list, totalCount, err := service.BasicsSmsTemplate().Log(ctx, &req.SmsTemplateLogInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_basics.SmsTemplateLogModel{}
	}

	res = new(basics.SmsTemplateLogRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
