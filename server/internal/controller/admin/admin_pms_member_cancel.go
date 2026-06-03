package admin

import (
	"APT/api/admin/pms"
	"APT/internal/model/input/input_app_member"
	"APT/internal/service"
	"context"
)

func (c *ControllerPms) MemberCancelList(ctx context.Context, req *pms.MemberCancelListReq) (res *pms.MemberCancelListRes, err error) {
	list, totalCount, err := service.AppMember().MemberCancelList(ctx, &req.PmsMemberCancelListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_app_member.PmsMemberCancelListModel{}
	}

	res = new(pms.MemberCancelListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerPms) MemberCancelAgree(ctx context.Context, req *pms.MemberCancelAgreeReq) (res *pms.MemberCancelAgreeRes, err error) {
	err = service.AppMember().MemberCancelAgree(ctx, &req.PmsMemberCancelAgreeInp)
	return
}
func (c *ControllerPms) MemberCancelDisagree(ctx context.Context, req *pms.MemberCancelDisagreeReq) (res *pms.MemberCancelDisagreeRes, err error) {
	err = service.AppMember().MemberCancelDisagree(ctx, &req.PmsMemberCancelDisagreeInp)
	return
}
