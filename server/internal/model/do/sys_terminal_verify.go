// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysTerminalVerify is the golang structure of table hg_sys_terminal_verify for DAO operations like Where/Data.
type SysTerminalVerify struct {
	g.Meta         `orm:"table:hg_sys_terminal_verify, do:true"`
	Id             interface{} //
	TerminalId     interface{} // 终端ID
	VerifyType     interface{} // 核销类型
	MchId          interface{} // 商户ID
	VerifyMemberId interface{} // 核销用户ID
	StoreId        interface{} // 门店ID
	RestaurantId   interface{} // 餐厅ID
	FoodOrderId    interface{} // 餐厅订单ID
	MemberCouponId interface{} // 用户礼品券ID
	CouponMchName  interface{} // 核销商品名
	VerifyTime     *gtime.Time // 核销时间
	CreateAt       *gtime.Time // 创建时间
	UpdateAt       *gtime.Time // 更新时间
	DeletedAt      *gtime.Time //
}
