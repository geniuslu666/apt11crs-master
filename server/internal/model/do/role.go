// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Role is the golang structure of table role for DAO operations like Where/Data.
type Role struct {
	g.Meta  `orm:"table:role, do:true"`
	Id      interface{} //
	Name    interface{} // 角色名称
	Method  interface{} // 允许的HTTP方法，*代表全部
	Path    interface{} // 允许的访客路径，*代表全部
	IsSuper interface{} // 是否为超管，未启用
}
