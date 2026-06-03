// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AiFilePoints is the golang structure of table ai_file_points for DAO operations like Where/Data.
type AiFilePoints struct {
	g.Meta      `orm:"table:ai_file_points, do:true"`
	Id          interface{} //
	FileId      interface{} // 文件表自增ID
	CreatedAt   *gtime.Time // 创建时间
	CollectName interface{} // 集合名称
	PointsId    interface{} // 向量ID
}
