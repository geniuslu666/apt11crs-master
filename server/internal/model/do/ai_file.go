// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AiFile is the golang structure of table ai_file for DAO operations like Where/Data.
type AiFile struct {
	g.Meta      `orm:"table:ai_file, do:true"`
	Id          interface{} //
	FileName    interface{} // 文件名
	CreatedAt   *gtime.Time // 创建时间
	FileSize    interface{} // 字符数
	CollectName interface{} // 集合名称
}
