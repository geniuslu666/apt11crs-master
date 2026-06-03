package admin

import (
	"APT/internal/model/input/input_app_member"
	"APT/internal/service"
	"context"

	"APT/api/admin/pms"
)

func (c *ControllerPms) MemberGroupList(ctx context.Context, req *pms.MemberGroupListReq) (res *pms.MemberGroupListRes, err error) {
	list, totalCount, err := service.AppMember().MemberGroupList(ctx, &req.PmsMemberGroupListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_app_member.PmsMemberGroupListModel{}
	}

	res = new(pms.MemberGroupListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerPms) MemberGroupAll(ctx context.Context, req *pms.MemberGroupAllReq) (res *pms.MemberGroupAllRes, err error) {
	list, err := service.AppMember().MemberGroupAllList(ctx, &req.PmsMemberGroupAllInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_app_member.PmsMemberGroupAllModel{}
	}

	res = new(pms.MemberGroupAllRes)
	res.List = list
	return
}
func (c *ControllerPms) MemberGroupExport(ctx context.Context, req *pms.MemberGroupExportReq) (res *pms.MemberGroupExportRes, err error) {
	err = service.AppMember().MemberGroupExport(ctx, &req.PmsMemberGroupListInp)
	return
}
func (c *ControllerPms) MemberGroupView(ctx context.Context, req *pms.MemberGroupViewReq) (res *pms.MemberGroupViewRes, err error) {
	data, err := service.AppMember().MemberGroupView(ctx, &req.PmsMemberGroupViewInp)
	if err != nil {
		return
	}

	res = new(pms.MemberGroupViewRes)
	res.PmsMemberGroupViewModel = data
	return
}
func (c *ControllerPms) MemberGroupEdit(ctx context.Context, req *pms.MemberGroupEditReq) (res *pms.MemberGroupEditRes, err error) {
	err = service.AppMember().MemberGroupEdit(ctx, &req.PmsMemberGroupEditInp)
	return
}
func (c *ControllerPms) MemberGroupDelete(ctx context.Context, req *pms.MemberGroupDeleteReq) (res *pms.MemberGroupDeleteRes, err error) {
	err = service.AppMember().MemberGroupDelete(ctx, &req.PmsMemberGroupDeleteInp)
	return
}
