// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderExport is the golang structure for table order_export.
type OrderExport struct {
	Id        int64       `json:"id"        orm:"id"        description:""`
	Scene     int         `json:"scene"     orm:"scene"     description:"场景 1-住宿 2-餐饮  3-按摩 4-接送机/包车 5-储物柜"`
	Condition string      `json:"condition" orm:"condition" description:"查询条件"`
	Status    uint        `json:"status"    orm:"status"    description:"0：导出中   1：导出成功  2：导出失败"`
	Path      string      `json:"path"      orm:"path"      description:"路径"`
	CreateAt  *gtime.Time `json:"createAt"  orm:"create_at" description:"创建时间"`
	UpdateAt  *gtime.Time `json:"updateAt"  orm:"update_at" description:"修改时间"`
}
