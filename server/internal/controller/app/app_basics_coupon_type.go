package app

import (
	"context"

	"APT/api/app/basics"
	"APT/internal/service"
)

func (c *ControllerBasics) CouponTypeView(ctx context.Context, req *basics.CouponTypeViewReq) (res *basics.CouponTypeViewRes, err error) {
	res = new(basics.CouponTypeViewRes)
	if res.PmsCouponTypeAppViewModel, err = service.BasicsCouponType().AppView(ctx, &req.PmsCouponTypeAppViewInp); err != nil {
		return
	}
	return
}
func (c *ControllerBasics) CouponTypeReceive(ctx context.Context, req *basics.CouponTypeReceiveReq) (res *basics.CouponTypeReceiveRes, err error) {
	err = service.BasicsCouponType().AppReceiveCouponType(ctx, &req.PmsCouponTypeAppReceiveInp)
	return
}
