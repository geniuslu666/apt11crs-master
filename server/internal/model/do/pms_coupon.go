// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsCoupon is the golang structure of table hg_pms_coupon for DAO operations like Where/Data.
type PmsCoupon struct {
	g.Meta          `orm:"table:hg_pms_coupon, do:true"`
	Id              interface{} // 主键ID
	Type            interface{} // 优惠券类型 reward-满减 discount-折扣 random-随机
	CouponName      interface{} // 优惠券名称
	CouponTypeId    interface{} // 优惠券类型id
	MemberId        interface{} // 领用人
	Scene           interface{} // 场景 1-住宿 2-餐饮  3-按摩 4-接送机/包车   5-储物柜
	PropertyIds     interface{} //
	RestaurantIds   interface{} // 餐厅ID
	ServiceIds      interface{} // 按摩服务ID
	CarServiceTypes interface{} // 汽车服务类型
	AtLeast         interface{} // 满多少元使用 0代表无限制
	Money           interface{} // 发放面额 当type为reward时需要添加
	Discount        interface{} // 1 =< 折扣 <= 9.9 当type为discount时需要添加
	DiscountLimit   interface{} // 最多折扣金额 当type为discount时可选择性添加
	State           interface{} // 优惠券状态 1已领用（未使用） 2已使用 3已过期 4已关闭 5已回收
	FetchTime       *gtime.Time // 领取时间
	UseTime         *gtime.Time // 使用时间
	StartTime       *gtime.Time // 可使用的开始时间
	EndTime         *gtime.Time // 有效期结束时间
	Source          interface{} // 来源：1-自主领取 2-系统发放 3-注册奖励 4-邀请奖励  5-首页活动
	SourceOrderId   interface{} // 来源订单ID
	OperatorId      interface{} // 操作员ID（回收）
	IndexActivityId interface{} // 活动ID（首页活动）
	RecoveryTime    *gtime.Time // 回收时间
	CreateAt        *gtime.Time // 创建时间
	UpdateAt        *gtime.Time // 修改时间
	DeletedAt       *gtime.Time // 删除时间
}
