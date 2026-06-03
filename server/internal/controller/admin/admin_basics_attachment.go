package admin

import (
	"APT/api/admin/basics"
	"APT/internal/dao"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"
)

func (c *ControllerBasics) AttachmentList(ctx context.Context, req *basics.AttachmentListReq) (res *basics.AttachmentListRes, err error) {
	list, totalCount, err := service.BasicsAttachment().List(ctx, &req.AttachmentListInp)
	if err != nil {
		return
	}

	res = new(basics.AttachmentListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerBasics) AttachmentView(ctx context.Context, req *basics.AttachmentViewReq) (res *basics.AttachmentViewRes, err error) {
	data, err := service.BasicsAttachment().View(ctx, &req.AttachmentViewInp)
	if err != nil {
		return
	}

	res = new(basics.AttachmentViewRes)
	res.AttachmentViewModel = data
	return
}

func (c *ControllerBasics) AttachmentDelete(ctx context.Context, req *basics.AttachmentDeleteReq) (res *basics.AttachmentDeleteRes, err error) {
	err = service.BasicsAttachment().Delete(ctx, &req.AttachmentDeleteInp)
	return
}
func (c *ControllerBasics) AttachmentClearKind(ctx context.Context, req *basics.AttachmentClearKindReq) (res *basics.AttachmentClearKindRes, err error) {
	err = service.BasicsAttachment().ClearKind(ctx, &req.AttachmentClearKindInp)
	return
}
func (c *ControllerBasics) AttachmentChooserOption(ctx context.Context, _ *basics.AttachmentChooserOptionReq) (res *basics.AttachmentChooserOptionRes, err error) {
	var (
		kinds []entity.SysAttachmentKind
	)
	res = new(basics.AttachmentChooserOptionRes)

	res.Drive, err = service.BasicsDictData().Select(ctx, &input_basics.DataSelectInp{Type: "config_upload_drive"})
	if err != nil {
		return
	}

	res.Kind = append(res.Kind, entity.SysAttachmentKind{
		Id:    0,
		Key:   "",
		Label: "全部",
		Value: "",
	})
	if err = dao.SysAttachmentKind.Ctx(ctx).Scan(&kinds); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	res.Kind = append(res.Kind, kinds...)
	return
}
