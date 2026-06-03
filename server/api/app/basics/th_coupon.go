package basics

import (
	"APT/internal/model/input/input_th"

	"github.com/gogf/gf/v2/frame/g"
)

type ThCouponListReq struct {
	g.Meta `path:"/th/couponList" method:"post" tags:"APP_BASICS" summary:"[礼品券]列表"`
	input_th.ThMemberCouponAppListInp
}

type ThCouponListRes struct {
	List  []*input_th.ThMemberCouponAppListModel `json:"list"   dc:"数据列表"`
	Count int                                    `json:"count"   dc:"数据总数"`
}

type ThCouponViewReq struct {
	g.Meta `path:"/th/couponDetail" method:"post" tags:"APP_BASICS" summary:"[礼品券]详情"`
	input_th.ThMemberCouponAppViewInp
}

type ThCouponViewRes struct {
	*input_th.ThMemberCouponAppViewModel
}

type ThMchViewReq struct {
	g.Meta `path:"/th/mchDetail" method:"post" tags:"APP_BASICS" summary:"[商户]详情"`
	input_th.ThMchAppViewInp
}

type ThMchViewRes struct {
	*input_th.ThMchAppViewModel
}

type ThCouponWaitViewReq struct {
	g.Meta `path:"/th/couponWaitDetail" method:"post" tags:"APP_BASICS" summary:"[礼品券]未领取详情"`
	input_th.ThCouponAppViewInp
}

type ThCouponWaitViewRes struct {
	*input_th.ThCouponAppViewModel
}

type ThCouponReceiveReq struct {
	g.Meta `path:"/th/couponReceive" method:"post" tags:"APP_BASICS" summary:"[礼品券]领取"`
	input_th.ThCouponAppReceiveInp
}

type ThCouponReceiveRes struct{}
