// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Receipt is the golang structure of table hg_receipt for DAO operations like Where/Data.
type Receipt struct {
	g.Meta      `orm:"table:hg_receipt, do:true"`
	Id          interface{} //
	UserName    interface{} // 发票抬头
	OrderSn     interface{} // 订单号
	ReceiptPath interface{} // 领售书文件地址
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 删除时间
}
