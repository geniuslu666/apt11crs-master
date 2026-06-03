// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsFolio is the golang structure for table pms_folio.
type PmsFolio struct {
	Id                 int         `json:"id"                 orm:"id"                   description:"主键"`
	Uid                string      `json:"uid"                orm:"uid"                  description:"第三方系统的ID"`
	Charges            *gjson.Json `json:"charges"            orm:"charges"              description:"费用数组，参考Charges"`
	Payments           *gjson.Json `json:"payments"           orm:"payments"             description:"付款数组，参考Payments"`
	TotalChargeAmount  float64     `json:"totalChargeAmount"  orm:"total_charge_amount"  description:"总收费金额"`
	TotalPaymentAmount float64     `json:"totalPaymentAmount" orm:"total_payment_amount" description:"总付款金额"`
	OutstandingBalance float64     `json:"outstandingBalance" orm:"outstanding_balance"  description:"未清余额"`
	CreatedAt          *gtime.Time `json:"createdAt"          orm:"created_at"           description:""`
	UpdatedAt          *gtime.Time `json:"updatedAt"          orm:"updated_at"           description:""`
}
