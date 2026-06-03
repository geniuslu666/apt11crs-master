package admin

import (
	"APT/internal/model/input/input_app_member"
	"APT/internal/service"
	"context"

	"APT/api/admin/pms"
)

func (c *ControllerPms) ExpList(ctx context.Context, req *pms.ExpListReq) (res *pms.ExpListRes, err error) {
	list, totalCount, err := service.AppMember().ExpChangeList(ctx, &req.PmsExpChangeListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_app_member.PmsExpChangeListModel{}
	}

	res = new(pms.ExpListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
