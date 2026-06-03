package admin

import (
	"APT/internal/model/input/input_app_member"
	"APT/internal/service"
	"context"

	"APT/api/admin/pms"
)

func (c *ControllerPms) MemberLevelList(ctx context.Context, req *pms.MemberLevelListReq) (res *pms.MemberLevelListRes, err error) {
	list, totalCount, err := service.AppMember().MemberLevelList(ctx, &req.PmsMemberLevelListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_app_member.PmsMemberLevelListModel{}
	}

	res = new(pms.MemberLevelListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerPms) MemberLevelAll(ctx context.Context, req *pms.MemberLevelAllReq) (res *pms.MemberLevelAllRes, err error) {
	list, err := service.AppMember().MemberLevelAllList(ctx, &req.PmsMemberLevelAllInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_app_member.PmsMemberLevelAllModel{}
	}

	res = new(pms.MemberLevelAllRes)
	res.List = list
	return
}
func (c *ControllerPms) MemberLevelView(ctx context.Context, req *pms.MemberLevelViewReq) (res *pms.MemberLevelViewRes, err error) {
	data, err := service.AppMember().MemberLevelView(ctx, &req.PmsMemberLevelViewInp)
	if err != nil {
		return
	}

	res = new(pms.MemberLevelViewRes)
	res.PmsMemberLevelViewModel = data
	return
}
func (c *ControllerPms) MemberLevelEdit(ctx context.Context, req *pms.MemberLevelEditReq) (res *pms.MemberLevelEditRes, err error) {
	err = service.AppMember().MemberLevelEdit(ctx, &req.PmsMemberLevelEditInp)
	return
}
func (c *ControllerPms) MemberLevelDelete(ctx context.Context, req *pms.MemberLevelDeleteReq) (res *pms.MemberLevelDeleteRes, err error) {
	err = service.AppMember().MemberLevelDelete(ctx, &req.PmsMemberLevelDeleteInp)
	return
}
