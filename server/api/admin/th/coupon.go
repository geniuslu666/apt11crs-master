package th

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_th"
	"github.com/gogf/gf/v2/frame/g"
)

// CouponListReq CouponListReq ListReq 查询礼品券列表
type CouponListReq struct {
	g.Meta `path:"/thCoupon/list" method:"get" tags:"ADMIN_TH" summary:"礼品券_列表"`
	input_th.ThCouponListInp
}

type CouponListRes struct {
	input_form.PageRes
	List []*input_th.ThCouponListModel `json:"list"   dc:"数据列表"`
}

// CouponAllListReq 查询礼品券列表
type CouponAllListReq struct {
	g.Meta `path:"/thCoupon/all" method:"get" tags:"ADMIN_TH" summary:"礼品券_列表ALL"`
	input_th.ThCouponListInp
}

type CouponAllListRes struct {
	List []*input_th.ThCouponAllListModel `json:"list"   dc:"数据列表"`
}

// CouponViewReq CouponViewReq ViewReq 获取礼品券指定信息
type CouponViewReq struct {
	g.Meta `path:"/thCoupon/view" method:"get" tags:"ADMIN_TH" summary:"礼品券_详情"`
	input_th.ThCouponViewInp
}

type CouponViewRes struct {
	*input_th.ThCouponViewModel
}

// CouponEditReq CouponEditReq EditReq 修改/新增礼品券
type CouponEditReq struct {
	g.Meta `path:"/thCoupon/edit" method:"post" tags:"ADMIN_TH" summary:"礼品券_修改/新增"`
	input_th.ThCouponEditInp
}

type CouponEditRes struct{}

// CouponDeleteReq CouponDeleteReq DeleteReq 删除礼品券
type CouponDeleteReq struct {
	g.Meta `path:"/thCoupon/delete" method:"post" tags:"ADMIN_TH" summary:"礼品券_删除"`
	input_th.ThCouponDeleteInp
}

type CouponDeleteRes struct{}

// CouponMaxSortReq MaxSortReq 获取礼品券最大排序
type CouponMaxSortReq struct {
	g.Meta `path:"/thCoupon/maxSort" method:"get" tags:"ADMIN_TH" summary:"礼品券_最大排序"`
	input_th.ThCouponMaxSortInp
}

type CouponMaxSortRes struct {
	*input_th.ThCouponMaxSortModel
}

// CouponStatusReq StatusReq 更新礼品券状态
type CouponStatusReq struct {
	g.Meta `path:"/thCoupon/status" method:"post" tags:"ADMIN_TH" summary:"礼品券_更新状态"`
	input_th.ThCouponStatusInp
}

type CouponStatusRes struct{}

// CouponUseStatusReq StatusReq 更新礼品券状态
type CouponUseStatusReq struct {
	g.Meta `path:"/thCoupon/useStatus" method:"post" tags:"ADMIN_TH" summary:"礼品券_更新使用状态"`
	input_th.ThCouponUseStatusInp
}

type CouponUseStatusRes struct{}

// SendMemberCouponReq 发放会员礼品券
type SendMemberCouponReq struct {
	g.Meta `path:"/thCoupon/sendCoupon" method:"post" tags:"ADMIN_TH" summary:"礼品券_发放会员礼品券"`
	input_th.ThSendMemberCouponInp
}

type SendMemberCouponRes struct{}
