package basics

import (
	"APT/internal/model/input/input_basics"

	"github.com/gogf/gf/v2/frame/g"
)

type CouponTypeViewReq struct {
	g.Meta `path:"/couponTypeDetail" method:"post" tags:"APP_BASICS" summary:"[优惠券]未领取详情"`
	input_basics.PmsCouponTypeAppViewInp
}

type CouponTypeViewRes struct {
	*input_basics.PmsCouponTypeAppViewModel
}

type CouponTypeReceiveReq struct {
	g.Meta `path:"/couponTypeReceive" method:"post" tags:"APP_BASICS" summary:"[优惠券]领取"`
	input_basics.PmsCouponTypeAppReceiveInp
}

type CouponTypeReceiveRes struct{}
