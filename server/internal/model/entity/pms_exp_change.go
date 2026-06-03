// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsExpChange is the golang structure for table pms_exp_change.
type PmsExpChange struct {
	Id         int         `json:"id"         orm:"id"          description:""`
	Scene      string      `json:"scene"      orm:"scene"       description:"场景值   HOTEL 酒店  SYSTEM 系统"`
	Exp        float64     `json:"exp"        orm:"exp"         description:"变动金额"`
	OrderSn    string      `json:"orderSn"    orm:"order_sn"    description:"订单号"`
	Des        string      `json:"des"        orm:"des"         description:"描述"`
	MemberId   int         `json:"memberId"   orm:"member_id"   description:"会员ID"`
	OperatorId int         `json:"operatorId" orm:"operator_id" description:"操作员ID"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:""`
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:""`
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at"  description:""`
}
