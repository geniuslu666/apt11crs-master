// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsBalanceChange is the golang structure for table pms_balance_change.
type PmsBalanceChange struct {
	Id          int         `json:"id"          orm:"id"           description:"主键"`
	MemberId    int         `json:"memberId"    orm:"member_id"    description:"会员ID"`
	Scene       string      `json:"scene"       orm:"scene"        description:"场景值   HOTEL    酒店   SYSTEM 系统"`
	Type        string      `json:"type"        orm:"type"         description:"金额变动方式   CONSUME 消费   REFUND   退款    AWARD  奖励    BROKERAGE   佣金    SYS 系统调整"`
	ChangePrice float64     `json:"changePrice" orm:"change_price" description:"变更金额"`
	OrderSn     string      `json:"orderSn"     orm:"order_sn"     description:"订单号"`
	Des         string      `json:"des"         orm:"des"          description:"消费描述（后端展示）"`
	Reason      string      `json:"reason"      orm:"reason"       description:"原因（前端展示）"`
	OperatorId  int         `json:"operatorId"  orm:"operator_id"  description:"操作员ID"`
	MdCode      string      `json:"mdCode"      orm:"md_code"      description:"注册设备码"`
	MpModel     string      `json:"mpModel"     orm:"mp_model"     description:"注册设备型号"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"更新时间"`
}
