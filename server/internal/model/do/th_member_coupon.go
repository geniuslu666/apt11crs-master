// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ThMemberCoupon is the golang structure of table hg_th_member_coupon for DAO operations like Where/Data.
type ThMemberCoupon struct {
	g.Meta          `orm:"table:hg_th_member_coupon, do:true"`
	Id              interface{} // 主键ID
	CouponNo        interface{} // 券号
	CouponId        interface{} // 券id
	MemberId        interface{} // 领用人
	State           interface{} // 状态 1待生效 2未使用 3已核销 4已过期  5已失效  6已回收
	InvalidTime     *gtime.Time // 失效时间
	VerifyTime      *gtime.Time // 核销时间
	VerifyMchId     interface{} // 核销商户ID
	VerifyStoreId   interface{} // 核销门店ID
	StartTime       *gtime.Time // 有效期开始时间
	EndTime         *gtime.Time // 有效期结束时间
	Source          interface{} // 来源：1-手动发放 2-自动发放(下单奖励)  3-员工福利  4-首页活动
	SourceOrderId   interface{} // 来源订单ID
	ActivityId      interface{} // 活动ID（员工福利活动）
	EmployeeId      interface{} // 员工ID（员工福利发放）
	IndexActivityId interface{} // 活动ID（首页活动）
	CountDown       interface{} // 倒计时
	OperatorId      interface{} // 操作员ID（回收）
	RecoveryTime    *gtime.Time // 回收时间
	CreateAt        *gtime.Time // 创建时间
	UpdateAt        *gtime.Time // 修改时间
	DeletedAt       *gtime.Time // 删除时间
}
