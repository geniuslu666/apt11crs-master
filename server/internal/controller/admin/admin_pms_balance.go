package admin

import (
	"APT/internal/model/input/input_app_member"
	"APT/internal/service"
	"context"

	"APT/api/admin/pms"
)

func (c *ControllerPms) BalanceList(ctx context.Context, req *pms.BalanceListReq) (res *pms.BalanceListRes, err error) {
	list, totalCount, err := service.AppMember().BalanceChangeList(ctx, &req.PmsBalanceChangeListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_app_member.PmsBalanceChangeListModel{}
	}

	res = new(pms.BalanceListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerPms) BalanceStat(ctx context.Context, req *pms.BalanceStatReq) (res *pms.BalanceStatRes, err error) {
	data, err := service.AppMember().BalanceChangeStat(ctx, &req.PmsBalanceChangeStatInp)

	res = new(pms.BalanceStatRes)
	res.PmsBalanceChangeStatModel = data
	return
}
