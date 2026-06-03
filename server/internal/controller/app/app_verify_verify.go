package app

import (
	"APT/internal/library/contexts"
	"APT/internal/model/input/input_food"
	"APT/internal/model/input/input_th"
	"APT/internal/service"
	"context"
	"strings"

	"APT/api/app/verify"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerVerify) MemberVerify(ctx context.Context, req *verify.MemberVerifyReq) (res *verify.MemberVerifyRes, err error) {
	res = new(verify.MemberVerifyRes)

	valArr := strings.Split(req.Val, "|")
	if len(valArr) == 2 && valArr[1] == "FOOD" {
		// |FOOD区分
		if err = service.FoodOrder().Verify(ctx, &req.ThMemberCouponVerifyInp); err != nil {
			return
		}
	} else if len(valArr) == 2 && valArr[1] == "MEMBER" {
		// 会员码 - 调用新的会员码验证逻辑（发券+核销）
		if err = service.AppMember().VerifyMemberCode(ctx, &req.ThMemberCouponVerifyInp); err != nil {
			return
		}

		ghttp.RequestFromCtx(ctx).Response.WriteJsonExit(g.Map{
			"data": "OK",
		})
	} else {
		if err = service.ThMemberCoupon().Verify(ctx, &req.ThMemberCouponVerifyInp); err != nil {
			return
		}
	}
	return
}
func (c *ControllerVerify) MemberVerifyCodeRefresh(ctx context.Context, req *verify.MemberVerifyCodeRefreshReq) (res *verify.MemberVerifyCodeRefreshRes, err error) {
	res = new(verify.MemberVerifyCodeRefreshRes)
	var (
		MemberInfo = contexts.GetMemberUser(ctx)
	)
	if req.Scene == "food" {
		if res.Code, res.VerifyStatus, err = service.FoodOrder().RefreshCode(ctx, &input_food.FoodsOrderRefreshCodeInp{
			Id:       req.Id,
			MemberId: MemberInfo.Id,
		}); err != nil {
			return
		}
	} else {
		res.Code, res.VerifyStatus, err = service.ThMemberCoupon().RefreshCode(ctx, &input_th.ThMemberCouponRefreshCodeInp{
			Id:       req.Id,
			MemberId: MemberInfo.Id,
		})
		if err != nil {
			return
		}
	}
	return
}
