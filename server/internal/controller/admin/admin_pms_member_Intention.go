package admin

import (
	"APT/internal/model/input/input_app_member"
	"APT/internal/service"
	"context"

	"APT/api/admin/pms"
)

func (c *ControllerPms) MemberIntentionList(ctx context.Context, req *pms.MemberIntentionListReq) (res *pms.MemberIntentionListRes, err error) {
	list, totalCount, err := service.AppMember().MemberIntentionList(ctx, &req.PmsMemberIntentionListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_app_member.PmsMemberIntentionListModel{}
	}

	res = new(pms.MemberIntentionListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerPms) MemberIntentionExport(ctx context.Context, req *pms.MemberIntentionExportReq) (res *pms.MemberIntentionExportRes, err error) {
	err = service.AppMember().MemberIntentionExport(ctx, &req.PmsMemberIntentionListInp)
	return
}
func (c *ControllerPms) MemberIntentionView(ctx context.Context, req *pms.MemberIntentionViewReq) (res *pms.MemberIntentionViewRes, err error) {
	data, err := service.AppMember().MemberIntentionView(ctx, &req.PmsMemberIntentionViewInp)
	if err != nil {
		return
	}

	res = new(pms.MemberIntentionViewRes)
	res.PmsMemberIntentionViewModel = data
	return
}
