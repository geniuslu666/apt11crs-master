package app

import (
	"APT/api/app/basics"
	"APT/internal/library/contexts"
	"APT/internal/model/input/input_th"
	"APT/internal/service"
	"context"
)

func (c *ControllerBasics) ThCouponList(ctx context.Context, req *basics.ThCouponListReq) (res *basics.ThCouponListRes, err error) {
	var (
		MemberInfo = contexts.GetMemberUser(ctx)
	)
	res = new(basics.ThCouponListRes)
	if res.List, res.Count, err = service.ThMemberCoupon().AppList(ctx, &input_th.ThMemberCouponAppListInp{
		PageReq:  req.PageReq,
		State:    req.State,
		Source:   req.Source,
		MemberId: MemberInfo.Id,
	}); err != nil {
		return
	}
	return
}
func (c *ControllerBasics) ThCouponView(ctx context.Context, req *basics.ThCouponViewReq) (res *basics.ThCouponViewRes, err error) {
	res = new(basics.ThCouponViewRes)
	if res.ThMemberCouponAppViewModel, err = service.ThMemberCoupon().AppView(ctx, &req.ThMemberCouponAppViewInp); err != nil {
		return
	}
	return
}
func (c *ControllerBasics) ThMchView(ctx context.Context, req *basics.ThMchViewReq) (res *basics.ThMchViewRes, err error) {
	res = new(basics.ThMchViewRes)
	if res.ThMchAppViewModel, err = service.ThMch().AppView(ctx, &req.ThMchAppViewInp); err != nil {
		return
	}
	return
}
func (c *ControllerBasics) ThCouponWaitView(ctx context.Context, req *basics.ThCouponWaitViewReq) (res *basics.ThCouponWaitViewRes, err error) {
	res = new(basics.ThCouponWaitViewRes)
	if res.ThCouponAppViewModel, err = service.ThCoupon().AppView(ctx, &req.ThCouponAppViewInp); err != nil {
		return
	}
	return
}
func (c *ControllerBasics) ThCouponReceive(ctx context.Context, req *basics.ThCouponReceiveReq) (res *basics.ThCouponReceiveRes, err error) {
	err = service.ThCoupon().AppReceiveCoupon(ctx, &req.ThCouponAppReceiveInp)
	return
}
