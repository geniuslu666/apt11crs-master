package th

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_th"

	"github.com/gogf/gf/v2/frame/g"
)

// MemberCouponListReq 查询已发放礼品券列表
type MemberCouponListReq struct {
	g.Meta `path:"/thMemberCoupon/list" method:"get" tags:"ADMIN_TH" summary:"会员礼品券_获取已发放列表"`
	input_th.ThMemberCouponListInp
}

type MemberCouponListRes struct {
	input_form.PageRes
	List []*input_th.ThMemberCouponListModel `json:"list"   dc:"数据列表"`
}

type MemberCouponExportReq struct {
	g.Meta `path:"/thMemberCoupon/export" method:"get" tags:"ADMIN_TH" summary:"会员礼品券_导出"`
	input_th.ThMemberCouponListInp
}

type MemberCouponExportRes struct{}

// MemberCouponViewReq 获取会员礼品券详情
type MemberCouponViewReq struct {
	g.Meta `path:"/thMemberCoupon/view" method:"get" tags:"ADMIN_TH" summary:"会员礼品券_详情"`
	input_th.ThMemberCouponViewInp
}

type MemberCouponViewRes struct {
	*input_th.ThMemberCouponAdminViewModel
}

// MemberCouponRecycleReq 回收礼品券
type MemberCouponRecycleReq struct {
	g.Meta `path:"/thMemberCoupon/recycle" method:"post" tags:"ADMIN_TH" summary:"会员礼品券_回收"`
	input_th.ThMemberCouponRecycleInp
}

type MemberCouponRecycleRes struct{}

// ManualVerifyMemberCouponReq 手动核销会员礼品券
type ManualVerifyMemberCouponReq struct {
	g.Meta `path:"/thMemberCoupon/manualVerify" method:"post" tags:"ADMIN_TH" summary:"礼品券_手动核销"`
	input_th.ThMemberCouponManualVerifyInp
}

type ManualVerifyMemberCouponRes struct{}
