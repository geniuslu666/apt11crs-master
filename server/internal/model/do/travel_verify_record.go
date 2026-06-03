// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TravelVerifyRecord is the golang structure of table hg_travel_verify_record for DAO operations like Where/Data.
type TravelVerifyRecord struct {
	g.Meta        `orm:"table:hg_travel_verify_record, do:true"`
	Id            interface{} //
	OrderId       interface{} // 订单ID
	OrderSn       interface{} // 预约单号
	ProductId     interface{} // 产品ID
	MemberId      interface{} // 会员ID
	BookDate      *gtime.Time // 预约日期
	VerifyStaffId interface{} // 核销人员ID
	VerifyTime    *gtime.Time // 核销时间
	CreatedAt     *gtime.Time // 创建时间
}
