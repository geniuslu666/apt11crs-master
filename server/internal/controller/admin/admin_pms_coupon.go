package admin

import (
	"APT/api/admin/pms"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"
)

func (c *ControllerPms) CouponList(ctx context.Context, req *pms.CouponListReq) (res *pms.CouponListRes, err error) {
	list, totalCount, err := service.BasicsCoupon().List(ctx, &req.PmsCouponListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_basics.PmsCouponListModel{}
	}

	res = new(pms.CouponListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

func (c *ControllerPms) CouponStat(ctx context.Context, req *pms.CouponStatReq) (res *pms.CouponStatRes, err error) {
	data, err := service.BasicsCoupon().Stat(ctx, &req.PmsCouponStatInp)

	res = new(pms.CouponStatRes)
	res.PmsCouponStatModel = data
	return
}
func (c *ControllerPms) CouponRecycle(ctx context.Context, req *pms.CouponRecycleReq) (res *pms.CouponRecycleRes, err error) {
	err = service.BasicsCoupon().Recycle(ctx, &req.PmsCouponRecycleInp)
	return
}
