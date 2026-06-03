// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ThCouponMch is the golang structure for table th_coupon_mch.
type ThCouponMch struct {
	CouponId int64  `json:"couponId" orm:"coupon_id" description:"券ID"`
	MchId    int64  `json:"mchId"    orm:"mch_id"    description:"商户ID"`
	Name     string `json:"name"     orm:"name"      description:"核销商品名"`
}
