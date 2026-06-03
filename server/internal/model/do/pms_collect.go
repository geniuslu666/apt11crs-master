// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsCollect is the golang structure of table hg_pms_collect for DAO operations like Where/Data.
type PmsCollect struct {
	g.Meta      `orm:"table:hg_pms_collect, do:true"`
	Id          interface{} //
	MemberId    interface{} //
	CollectType interface{} // 收藏类型   HOST 、 酒店  REPAST 餐饮
	CollectId   interface{} // 收藏数据ID
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 删除时间
}
