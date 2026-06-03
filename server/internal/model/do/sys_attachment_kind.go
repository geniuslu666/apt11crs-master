// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysAttachmentKind is the golang structure of table hg_sys_attachment_kind for DAO operations like Where/Data.
type SysAttachmentKind struct {
	g.Meta `orm:"table:hg_sys_attachment_kind, do:true"`
	Id     interface{} //
	Label  interface{} // 附件分类名称
	Key    interface{} // 分类key
	Value  interface{} // 分类value
	Icon   interface{} // 分类图标
	Tag    interface{} //
}
