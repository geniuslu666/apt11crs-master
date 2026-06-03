package admin

import (
	"APT/internal/consts"
	"APT/internal/service"
	"context"

	"APT/api/admin/basics"
)

func (c *ControllerBasics) NoticeList(ctx context.Context, req *basics.NoticeListReq) (res *basics.NoticeListRes, err error) {
	list, totalCount, err := service.BasicsAdminNotice().List(ctx, &req.NoticeListInp)
	if err != nil {
		return nil, err
	}

	res = new(basics.NoticeListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerBasics) NoticeView(ctx context.Context, req *basics.NoticeViewReq) (res *basics.NoticeViewRes, err error) {
	data, err := service.BasicsAdminNotice().View(ctx, &req.NoticeViewInp)
	if err != nil {
		return
	}

	res = new(basics.NoticeViewRes)
	res.NoticeViewModel = data
	return
}
func (c *ControllerBasics) NoticeEdit(ctx context.Context, req *basics.NoticeEditReq) (res *basics.NoticeEditRes, err error) {
	err = service.BasicsAdminNotice().Edit(ctx, &req.NoticeEditInp)
	return
}
func (c *ControllerBasics) NoticeDelete(ctx context.Context, req *basics.NoticeDeleteReq) (res *basics.NoticeDeleteRes, err error) {
	err = service.BasicsAdminNotice().Delete(ctx, &req.NoticeDeleteInp)
	return
}
func (c *ControllerBasics) NoticeMaxSort(ctx context.Context, req *basics.NoticeMaxSortReq) (res *basics.NoticeMaxSortRes, err error) {
	res = new(basics.NoticeMaxSortRes)
	res.NoticeMaxSortModel, err = service.BasicsAdminNotice().MaxSort(ctx, &req.NoticeMaxSortInp)
	return
}
func (c *ControllerBasics) NoticeStatus(ctx context.Context, req *basics.NoticeStatusReq) (res *basics.NoticeStatusRes, err error) {
	err = service.BasicsAdminNotice().Status(ctx, &req.NoticeStatusInp)
	return
}
func (c *ControllerBasics) NoticeEditNotify(ctx context.Context, req *basics.NoticeEditNotifyReq) (res *basics.NoticeEditNotifyRes, err error) {
	req.Type = consts.NoticeTypeNotify
	err = service.BasicsAdminNotice().Edit(ctx, &req.NoticeEditInp)
	return
}
func (c *ControllerBasics) NoticeEditNotice(ctx context.Context, req *basics.NoticeEditNoticeReq) (res *basics.NoticeEditNoticeRes, err error) {
	req.Type = consts.NoticeTypeNotice
	err = service.BasicsAdminNotice().Edit(ctx, &req.NoticeEditInp)
	return
}
func (c *ControllerBasics) NoticeEditLetter(ctx context.Context, req *basics.NoticeEditLetterReq) (res *basics.NoticeEditLetterRes, err error) {
	req.Type = consts.NoticeTypeLetter
	err = service.BasicsAdminNotice().Edit(ctx, &req.NoticeEditInp)
	return
}
func (c *ControllerBasics) NoticePullMessages(ctx context.Context, req *basics.NoticePullMessagesReq) (res *basics.NoticePullMessagesRes, err error) {
	data, err := service.BasicsAdminNotice().PullMessages(ctx, &req.PullMessagesInp)
	if err != nil {
		return
	}

	res = new(basics.NoticePullMessagesRes)
	res.PullMessagesModel = data
	return
}
func (c *ControllerBasics) NoticeReadAll(ctx context.Context, req *basics.NoticeReadAllReq) (res *basics.NoticeReadAllRes, err error) {
	err = service.BasicsAdminNotice().ReadAll(ctx, &req.NoticeReadAllInp)
	return
}
func (c *ControllerBasics) NoticeUpRead(ctx context.Context, req *basics.NoticeUpReadReq) (res *basics.NoticeUpReadRes, err error) {
	err = service.BasicsAdminNotice().UpRead(ctx, &req.NoticeUpReadInp)
	return
}
func (c *ControllerBasics) NoticeMessageList(ctx context.Context, req *basics.NoticeMessageListReq) (res *basics.NoticeMessageListRes, err error) {
	list, totalCount, err := service.BasicsAdminNotice().MessageList(ctx, &req.NoticeMessageListInp)
	if err != nil {
		return
	}

	res = new(basics.NoticeMessageListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
