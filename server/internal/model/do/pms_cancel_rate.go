// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsCancelRate is the golang structure of table hg_pms_cancel_rate for DAO operations like Where/Data.
type PmsCancelRate struct {
	g.Meta    `orm:"table:hg_pms_cancel_rate, do:true"`
	Id        interface{} //
	Mode      interface{} // 规则模式
	StartDays interface{} // 开始天数
	EndDays   interface{} // 结束天数
	Rate      interface{} // 费率
	Name      interface{} // 规则名
	Sort      interface{} // 排序规则  从小到大
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
