// Package pmsAdminApi

package pms

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

// CouponListReq 查询已领取优惠券列表
type CouponListReq struct {
	g.Meta `path:"/pmsCoupon/list" method:"get" tags:"ADMIN_PMS" summary:"会员优惠券_获取已领取列表"`
	input_basics.PmsCouponListInp
}

type CouponListRes struct {
	input_form.PageRes
	List []*input_basics.PmsCouponListModel `json:"list"   dc:"数据列表"`
}

// CouponStatReq 会员优惠券概况
type CouponStatReq struct {
	g.Meta `path:"/pmsCoupon/stat" method:"get" tags:"ADMIN_PMS" summary:"会员优惠券_概况"`
	input_basics.PmsCouponStatInp
}

type CouponStatRes struct {
	*input_basics.PmsCouponStatModel
}

// CouponRecycleReq 回收优惠券
type CouponRecycleReq struct {
	g.Meta `path:"/pmsCoupon/recycle" method:"post" tags:"ADMIN_PMS" summary:"会员优惠券_回收"`
	input_basics.PmsCouponRecycleInp
}

type CouponRecycleRes struct{}
