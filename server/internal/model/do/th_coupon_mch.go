// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ThCouponMch is the golang structure of table hg_th_coupon_mch for DAO operations like Where/Data.
type ThCouponMch struct {
	g.Meta   `orm:"table:hg_th_coupon_mch, do:true"`
	CouponId interface{} // 券ID
	MchId    interface{} // 商户ID
	Name     interface{} // 核销商品名
}
