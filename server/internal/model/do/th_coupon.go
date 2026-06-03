// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ThCoupon is the golang structure of table hg_th_coupon for DAO operations like Where/Data.
type ThCoupon struct {
	g.Meta                   `orm:"table:hg_th_coupon, do:true"`
	Id                       interface{} // ID
	CouponName               interface{} // 提货券名称
	CouponSubName            interface{} // 提货券副标题
	IdentityName             interface{} // 券识别名称
	CouponNoPrefix           interface{} // 编号前缀
	Logo                     interface{} // LOGO
	Status                   interface{} // 发放状态（1立即启用  2暂不启用）
	FixedTerm                interface{} // 激活后几天内有效
	Desc                     interface{} // 券说明
	Count                    interface{} // 发放数量
	UseStatus                interface{} // 使用状态（1开始使用  2停止使用）
	UseMode                  interface{} // 使用模式
	CategoryId               interface{} // 分类ID
	UsedCount                interface{} // 已使用数量
	Sort                     interface{} // 排序
	NeedReservation          interface{} // 是否需要预约：0-不需要，1-需要
	ReservationRestaurantIds interface{} // 需要预约的餐厅IDs，多个用逗号分隔，如：1,2,3
	CreateAt                 *gtime.Time // 创建时间
	UpdateAt                 *gtime.Time // 修改时间
	DeletedAt                *gtime.Time // 删除时间
}
