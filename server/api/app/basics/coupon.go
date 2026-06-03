package basics

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"

	"github.com/gogf/gf/v2/frame/g"
)

type CouponListReq struct {
	g.Meta `path:"/couponList" method:"post" tags:"APP_BASICS" summary:"[优惠券]列表"`
	input_form.PageReq
	State int `json:"state" dc:"状态  1 已领用（未使用） 2 已使用 3 已过期"`
}

type CouponListRes struct {
	List  []*input_basics.PmsCouponListModel `json:"list"   dc:"数据列表"`
	Count int                                `json:"count"   dc:"数据总数"`
}

type CouponViewReq struct {
	g.Meta `path:"/couponDetail" method:"post" tags:"APP_BASICS" summary:"[优惠券]详情"`
	input_basics.PmsCouponViewInp
}

type CouponViewRes struct {
	*input_basics.PmsCouponViewModel
}

type CouponOrderListReq struct {
	g.Meta `path:"/orderCouponList" method:"post" tags:"APP_BASICS" summary:"[优惠券]订单优惠券列表"`
	input_form.PageReq
	PreOrderSn string `json:"preOrderSn"    dc:"预订单号"`
	Scene      int    `json:"scene"         dc:"场景 1 住宿场景 2 餐饮场景 3 按摩  4 接送机  5 储物柜"`
}

type CouponOrderListRes struct {
	List  []*input_basics.PmsCouponListModel `json:"list"   dc:"数据列表"`
	Count int                                `json:"count"   dc:"数据总数"`
}
