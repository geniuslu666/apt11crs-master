// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CarBanner is the golang structure of table hg_car_banner for DAO operations like Where/Data.
type CarBanner struct {
	g.Meta      `orm:"table:hg_car_banner, do:true"`
	Id          interface{} //
	Language    interface{} // 语言
	BannerImage interface{} // 轮播图
	Status      interface{} // 状态1、启用 2、禁用
	Sort        interface{} // 排序(越大越靠前)
	CreateAt    *gtime.Time // 创建时间
	UpdateAt    *gtime.Time // 更新时间
}
