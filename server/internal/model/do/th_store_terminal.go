// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ThStoreTerminal is the golang structure of table hg_th_store_terminal for DAO operations like Where/Data.
type ThStoreTerminal struct {
	g.Meta     `orm:"table:hg_th_store_terminal, do:true"`
	StoreId    interface{} // 门店ID
	TerminalId interface{} // 终端ID
	PrintTimes interface{} // 打印联数(次数)
}
