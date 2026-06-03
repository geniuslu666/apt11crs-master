// Package pmsAdminApi

package pms

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

// CouponTypeListReq CouponTypeListReq ListReq 查询优惠券列表
type CouponTypeListReq struct {
	g.Meta `path:"/pmsCouponType/list" method:"get" tags:"ADMIN_PMS" summary:"优惠券_列表"`
	input_basics.PmsCouponTypeListInp
}

type CouponTypeListRes struct {
	input_form.PageRes
	List []*input_basics.PmsCouponTypeListModel `json:"list"   dc:"数据列表"`
}

// CouponTypeAllListReq 查询优惠券列表
type CouponTypeAllListReq struct {
	g.Meta `path:"/pmsCouponType/all" method:"get" tags:"ADMIN_PMS" summary:"优惠券_列表ALL"`
	input_basics.PmsCouponTypeListInp
}

type CouponTypeAllListRes struct {
	List []*input_basics.PmsCouponTypeAllListModel `json:"list"   dc:"数据列表"`
}

// CouponTypeViewReq CouponTypeViewReq ViewReq 获取优惠券指定信息
type CouponTypeViewReq struct {
	g.Meta `path:"/pmsCouponType/view" method:"get" tags:"ADMIN_PMS" summary:"优惠券_详情"`
	input_basics.PmsCouponTypeViewInp
}

type CouponTypeViewRes struct {
	*input_basics.PmsCouponTypeViewModel
}

// CouponTypeEditReq CouponTypeEditReq EditReq 修改/新增优惠券
type CouponTypeEditReq struct {
	g.Meta `path:"/pmsCouponType/edit" method:"post" tags:"ADMIN_PMS" summary:"优惠券_修改/新增"`
	input_basics.PmsCouponTypeEditInp
}

type CouponTypeEditRes struct{}

// CouponTypeDeleteReq CouponTypeDeleteReq DeleteReq 删除优惠券
type CouponTypeDeleteReq struct {
	g.Meta `path:"/pmsCouponType/delete" method:"post" tags:"ADMIN_PMS" summary:"优惠券_删除"`
	input_basics.PmsCouponTypeDeleteInp
}

type CouponTypeDeleteRes struct{}

// CouponTypeMaxSortReq MaxSortReq 获取优惠券最大排序
type CouponTypeMaxSortReq struct {
	g.Meta `path:"/pmsCouponType/maxSort" method:"get" tags:"ADMIN_PMS" summary:"优惠券_最大排序"`
	input_basics.PmsCouponTypeMaxSortInp
}

type CouponTypeMaxSortRes struct {
	*input_basics.PmsCouponTypeMaxSortModel
}

// CouponTypeStatusReq StatusReq 更新优惠券状态
type CouponTypeStatusReq struct {
	g.Meta `path:"/pmsCouponType/status" method:"post" tags:"ADMIN_PMS" summary:"优惠券_更新状态"`
	input_basics.PmsCouponTypeStatusInp
}

type CouponTypeStatusRes struct{}

// SendMemberCouponReq 发放会员优惠券
type SendMemberCouponReq struct {
	g.Meta `path:"/pmsCouponType/sendCoupon" method:"post" tags:"ADMIN_PMS" summary:"优惠券_发放会员优惠券"`
	input_basics.PmsSendMemberCouponInp
}

type SendMemberCouponRes struct{}
