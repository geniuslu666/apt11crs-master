// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Learn is the golang structure of table learn for DAO operations like Where/Data.
type Learn struct {
	g.Meta    `orm:"table:learn, do:true"`
	Id        interface{} //
	Content   interface{} // 问题内容
	Score     interface{} // 次数
	CreatedAt *gtime.Time // 创建时间
	EntId     interface{} // 企业ID
	KefuName  interface{} // kefu名称
	Finshed   interface{} // 是否解决，1未解决，2已解决
}
