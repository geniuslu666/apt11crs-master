package admin

import (
	"APT/internal/model/input/input_app_member"
	"APT/internal/service"
	"context"

	"APT/api/admin/pms"
)

func (c *ControllerPms) MemberSceneList(ctx context.Context, req *pms.MemberSceneListReq) (res *pms.MemberSceneListRes, err error) {
	list, totalCount, err := service.AppMember().MemberSceneList(ctx, &req.PmsMemberSceneListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_app_member.PmsMemberSceneListModel{}
	}

	res = new(pms.MemberSceneListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerPms) MemberSceneView(ctx context.Context, req *pms.MemberSceneViewReq) (res *pms.MemberSceneViewRes, err error) {
	data, err := service.AppMember().MemberSceneView(ctx, &req.PmsMemberSceneViewInp)
	if err != nil {
		return
	}

	res = new(pms.MemberSceneViewRes)
	res.PmsMemberSceneViewModel = data
	return
}
func (c *ControllerPms) MemberSceneEdit(ctx context.Context, req *pms.MemberSceneEditReq) (res *pms.MemberSceneEditRes, err error) {
	err = service.AppMember().MemberSceneEdit(ctx, &req.PmsMemberSceneEditInp)
	return
}
func (c *ControllerPms) MemberSceneDelete(ctx context.Context, req *pms.MemberSceneDeleteReq) (res *pms.MemberSceneDeleteRes, err error) {
	err = service.AppMember().MemberSceneDelete(ctx, &req.PmsMemberSceneDeleteInp)
	return
}
func (c *ControllerPms) MemberSceneSwitch(ctx context.Context, req *pms.MemberSceneSwitchReq) (res *pms.MemberSceneSwitchRes, err error) {
	err = service.AppMember().MemberSceneSwitch(ctx, &req.PmsMemberSceneSwitchInp)
	return
}
