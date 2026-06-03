package admin

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"

	"APT/api/admin/basics"
)

func (c *ControllerBasics) SysAttachmentKindList(ctx context.Context, req *basics.SysAttachmentKindListReq) (res *basics.SysAttachmentKindListRes, err error) {
	list, totalCount, err := service.BasicsAttachmentKind().List(ctx, &req.SysAttachmentKindListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_basics.SysAttachmentKindListModel{}
	}

	res = new(basics.SysAttachmentKindListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerBasics) SysAttachmentKindView(ctx context.Context, req *basics.SysAttachmentKindViewReq) (res *basics.SysAttachmentKindViewRes, err error) {
	data, err := service.BasicsAttachmentKind().View(ctx, &req.SysAttachmentKindViewInp)
	if err != nil {
		return
	}

	res = new(basics.SysAttachmentKindViewRes)
	res.SysAttachmentKindViewModel = data
	return
}
func (c *ControllerBasics) SysAttachmentKindEdit(ctx context.Context, req *basics.SysAttachmentKindEditReq) (res *basics.SysAttachmentKindEditRes, err error) {
	err = service.BasicsAttachmentKind().Edit(ctx, &req.SysAttachmentKindEditInp)
	return
}
func (c *ControllerBasics) SysAttachmentKindDelete(ctx context.Context, req *basics.SysAttachmentKindDeleteReq) (res *basics.SysAttachmentKindDeleteRes, err error) {
	err = service.BasicsAttachmentKind().Delete(ctx, &req.SysAttachmentKindDeleteInp)
	return
}
