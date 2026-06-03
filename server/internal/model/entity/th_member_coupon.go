// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ThMemberCoupon is the golang structure for table th_member_coupon.
type ThMemberCoupon struct {
	Id              int         `json:"id"              orm:"id"                description:"主键ID"`
	CouponNo        string      `json:"couponNo"        orm:"coupon_no"         description:"券号"`
	CouponId        int         `json:"couponId"        orm:"coupon_id"         description:"券id"`
	MemberId        int         `json:"memberId"        orm:"member_id"         description:"领用人"`
	State           int         `json:"state"           orm:"state"             description:"状态 1待生效 2未使用 3已核销 4已过期  5已失效  6已回收"`
	InvalidTime     *gtime.Time `json:"invalidTime"     orm:"invalid_time"      description:"失效时间"`
	VerifyTime      *gtime.Time `json:"verifyTime"      orm:"verify_time"       description:"核销时间"`
	VerifyMchId     int         `json:"verifyMchId"     orm:"verify_mch_id"     description:"核销商户ID"`
	VerifyStoreId   int         `json:"verifyStoreId"   orm:"verify_store_id"   description:"核销门店ID"`
	StartTime       *gtime.Time `json:"startTime"       orm:"start_time"        description:"有效期开始时间"`
	EndTime         *gtime.Time `json:"endTime"         orm:"end_time"          description:"有效期结束时间"`
	Source          uint        `json:"source"          orm:"source"            description:"来源：1-手动发放 2-自动发放(下单奖励)  3-员工福利  4-首页活动"`
	SourceOrderId   int         `json:"sourceOrderId"   orm:"source_order_id"   description:"来源订单ID"`
	ActivityId      uint64      `json:"activityId"      orm:"activity_id"       description:"活动ID（员工福利活动）"`
	EmployeeId      uint64      `json:"employeeId"      orm:"employee_id"       description:"员工ID（员工福利发放）"`
	IndexActivityId uint64      `json:"indexActivityId" orm:"index_activity_id" description:"活动ID（首页活动）"`
	CountDown       int         `json:"countDown"       orm:"count_down"        description:"倒计时"`
	OperatorId      int         `json:"operatorId"      orm:"operator_id"       description:"操作员ID（回收）"`
	RecoveryTime    *gtime.Time `json:"recoveryTime"    orm:"recovery_time"     description:"回收时间"`
	CreateAt        *gtime.Time `json:"createAt"        orm:"create_at"         description:"创建时间"`
	UpdateAt        *gtime.Time `json:"updateAt"        orm:"update_at"         description:"修改时间"`
	DeletedAt       *gtime.Time `json:"deletedAt"       orm:"deleted_at"        description:"删除时间"`
}
