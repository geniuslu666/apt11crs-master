// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsFolio is the golang structure of table hg_pms_folio for DAO operations like Where/Data.
type PmsFolio struct {
	g.Meta             `orm:"table:hg_pms_folio, do:true"`
	Id                 interface{} // 主键
	Uid                interface{} // 第三方系统的ID
	Charges            *gjson.Json // 费用数组，参考Charges
	Payments           *gjson.Json // 付款数组，参考Payments
	TotalChargeAmount  interface{} // 总收费金额
	TotalPaymentAmount interface{} // 总付款金额
	OutstandingBalance interface{} // 未清余额
	CreatedAt          *gtime.Time //
	UpdatedAt          *gtime.Time //
}
