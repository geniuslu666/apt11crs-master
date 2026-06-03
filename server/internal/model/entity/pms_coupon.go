// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsCoupon is the golang structure for table pms_coupon.
type PmsCoupon struct {
	Id              int         `json:"id"              orm:"id"                description:"主键ID"`
	Type            string      `json:"type"            orm:"type"              description:"优惠券类型 reward-满减 discount-折扣 random-随机"`
	CouponName      string      `json:"couponName"      orm:"coupon_name"       description:"优惠券名称"`
	CouponTypeId    int         `json:"couponTypeId"    orm:"coupon_type_id"    description:"优惠券类型id"`
	MemberId        int         `json:"memberId"        orm:"member_id"         description:"领用人"`
	Scene           int         `json:"scene"           orm:"scene"             description:"场景 1-住宿 2-餐饮  3-按摩 4-接送机/包车   5-储物柜"`
	PropertyIds     string      `json:"propertyIds"     orm:"property_ids"      description:""`
	RestaurantIds   string      `json:"restaurantIds"   orm:"restaurant_ids"    description:"餐厅ID"`
	ServiceIds      string      `json:"serviceIds"      orm:"service_ids"       description:"按摩服务ID"`
	CarServiceTypes string      `json:"carServiceTypes" orm:"car_service_types" description:"汽车服务类型"`
	AtLeast         float64     `json:"atLeast"         orm:"at_least"          description:"满多少元使用 0代表无限制"`
	Money           float64     `json:"money"           orm:"money"             description:"发放面额 当type为reward时需要添加"`
	Discount        float64     `json:"discount"        orm:"discount"          description:"1 =< 折扣 <= 9.9 当type为discount时需要添加"`
	DiscountLimit   float64     `json:"discountLimit"   orm:"discount_limit"    description:"最多折扣金额 当type为discount时可选择性添加"`
	State           int         `json:"state"           orm:"state"             description:"优惠券状态 1已领用（未使用） 2已使用 3已过期 4已关闭 5已回收"`
	FetchTime       *gtime.Time `json:"fetchTime"       orm:"fetch_time"        description:"领取时间"`
	UseTime         *gtime.Time `json:"useTime"         orm:"use_time"          description:"使用时间"`
	StartTime       *gtime.Time `json:"startTime"       orm:"start_time"        description:"可使用的开始时间"`
	EndTime         *gtime.Time `json:"endTime"         orm:"end_time"          description:"有效期结束时间"`
	Source          uint        `json:"source"          orm:"source"            description:"来源：1-自主领取 2-系统发放 3-注册奖励 4-邀请奖励  5-首页活动"`
	SourceOrderId   int         `json:"sourceOrderId"   orm:"source_order_id"   description:"来源订单ID"`
	OperatorId      int         `json:"operatorId"      orm:"operator_id"       description:"操作员ID（回收）"`
	IndexActivityId uint64      `json:"indexActivityId" orm:"index_activity_id" description:"活动ID（首页活动）"`
	RecoveryTime    *gtime.Time `json:"recoveryTime"    orm:"recovery_time"     description:"回收时间"`
	CreateAt        *gtime.Time `json:"createAt"        orm:"create_at"         description:"创建时间"`
	UpdateAt        *gtime.Time `json:"updateAt"        orm:"update_at"         description:"修改时间"`
	DeletedAt       *gtime.Time `json:"deletedAt"       orm:"deleted_at"        description:"删除时间"`
}
