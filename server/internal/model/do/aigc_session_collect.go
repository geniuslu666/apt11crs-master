// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AigcSessionCollect is the golang structure of table aigc_session_collect for DAO operations like Where/Data.
type AigcSessionCollect struct {
	g.Meta    `orm:"table:aigc_session_collect, do:true"`
	Id        interface{} //
	Title     interface{} // 集合标题
	CreatedAt *gtime.Time // 创建时间
	KefuName  interface{} // 客服名称
	EntId     interface{} // 企业ID
}
