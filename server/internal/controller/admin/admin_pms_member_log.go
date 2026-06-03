package admin

import (
	"APT/internal/model/input/input_app_member"
	"APT/internal/service"
	"context"

	"APT/api/admin/pms"
)

func (c *ControllerPms) MemberLogList(ctx context.Context, req *pms.MemberLogListReq) (res *pms.MemberLogListRes, err error) {
	list, totalCount, err := service.AppMember().MemberLogList(ctx, &req.PmsMemberLogListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_app_member.PmsMemberLogListModel{}
	}

	res = new(pms.MemberLogListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerPms) MemberLogExport(ctx context.Context, req *pms.MemberLogExportReq) (res *pms.MemberLogExportRes, err error) {
	err = service.AppMember().MemberLogExport(ctx, &req.PmsMemberLogListInp)
	return
}
func (c *ControllerPms) MemberLogView(ctx context.Context, req *pms.MemberLogViewReq) (res *pms.MemberLogViewRes, err error) {
	data, err := service.AppMember().MemberLogView(ctx, &req.PmsMemberLogViewInp)
	if err != nil {
		return
	}

	res = new(pms.MemberLogViewRes)
	res.PmsMemberLogViewModel = data
	return
}
