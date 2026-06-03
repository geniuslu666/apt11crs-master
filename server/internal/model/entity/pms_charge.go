// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsCharge is the golang structure for table pms_charge.
type PmsCharge struct {
	Id             int         `json:"id"             orm:"id"              description:"主键"`
	Uid            string      `json:"uid"            orm:"uid"             description:"三方系统 ID"`
	AirUid         string      `json:"airUid"         orm:"air_uid"         description:"airhost uid"`
	Date           *gtime.Time `json:"date"           orm:"date"            description:"费用日期，可以是发生日期或记账日期"`
	Name           string      `json:"name"           orm:"name"            description:"名称"`
	FeeType        string      `json:"feeType"        orm:"fee_type"        description:"费用类型（预订费、餐饮费、清洁费、取消费、其他费用）"`
	Amount         float64     `json:"amount"         orm:"amount"          description:"金额"`
	Currency       string      `json:"currency"       orm:"currency"        description:"货币"`
	OriginalAmount float64     `json:"originalAmount" orm:"original_amount" description:"原价"`
	Description    string      `json:"description"    orm:"description"     description:"费用详情"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:""`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:""`
	DeletedAt      *gtime.Time `json:"deletedAt"      orm:"deleted_at"      description:""`
}
