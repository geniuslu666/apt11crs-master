// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsBanner is the golang structure of table hg_pms_banner for DAO operations like Where/Data.
type PmsBanner struct {
	g.Meta       `orm:"table:hg_pms_banner, do:true"`
	Id           interface{} //
	Language     interface{} // 语言
	BannerName   interface{} // banner名称
	BannerImage  interface{} // 轮播图
	Model        interface{} // 模块   BANNER  横幅   EVENTS  事件
	Chain        interface{} // IN 内链  OUT 外联
	Path         interface{} // 链接内容
	BannerStatus interface{} // 1、启用 2、禁用
	CreateAt     *gtime.Time // 创建时间
	UpdateAt     *gtime.Time // 更新时间
}
