// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderExport is the golang structure of table hg_order_export for DAO operations like Where/Data.
type OrderExport struct {
	g.Meta    `orm:"table:hg_order_export, do:true"`
	Id        interface{} //
	Scene     interface{} // 场景 1-住宿 2-餐饮  3-按摩 4-接送机/包车 5-储物柜
	Condition interface{} // 查询条件
	Status    interface{} // 0：导出中   1：导出成功  2：导出失败
	Path      interface{} // 路径
	CreateAt  *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 修改时间
}
