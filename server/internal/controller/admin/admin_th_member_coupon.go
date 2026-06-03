package admin

import (
	"APT/internal/model/input/input_th"
	"APT/internal/service"
	"context"

	"APT/api/admin/th"
)

func (c *ControllerTh) MemberCouponList(ctx context.Context, req *th.MemberCouponListReq) (res *th.MemberCouponListRes, err error) {
	list, totalCount, err := service.ThMemberCoupon().List(ctx, &req.ThMemberCouponListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_th.ThMemberCouponListModel{}
	}

	res = new(th.MemberCouponListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerTh) MemberCouponExport(ctx context.Context, req *th.MemberCouponExportReq) (res *th.MemberCouponExportRes, err error) {
	err = service.ThMemberCoupon().Export(ctx, &req.ThMemberCouponListInp)
	return
}
func (c *ControllerTh) MemberCouponView(ctx context.Context, req *th.MemberCouponViewReq) (res *th.MemberCouponViewRes, err error) {
	data, err := service.ThMemberCoupon().View(ctx, &req.ThMemberCouponViewInp)
	if err != nil {
		return
	}

	res = new(th.MemberCouponViewRes)
	res.ThMemberCouponAdminViewModel = data
	return
}
func (c *ControllerTh) MemberCouponRecycle(ctx context.Context, req *th.MemberCouponRecycleReq) (res *th.MemberCouponRecycleRes, err error) {
	err = service.ThMemberCoupon().Recycle(ctx, &req.ThMemberCouponRecycleInp)
	return
}
func (c *ControllerTh) ManualVerifyMemberCoupon(ctx context.Context, req *th.ManualVerifyMemberCouponReq) (res *th.ManualVerifyMemberCouponRes, err error) {
	err = service.ThMemberCoupon().ManualVerify(ctx, &req.ThMemberCouponManualVerifyInp)
	return
}
