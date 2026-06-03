// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsIndexNav is the golang structure of table hg_pms_index_nav for DAO operations like Where/Data.
type PmsIndexNav struct {
	g.Meta       `orm:"table:hg_pms_index_nav, do:true"`
	Id           interface{} //
	Name         interface{} // 导航名称(多语言)
	Tag          interface{} // 标签
	Image        interface{} // 图标
	AppLink      interface{} // app跳转链接
	WxLink       interface{} // 微信跳转链接
	Sort         interface{} // 排序(越大越靠前)
	Status       interface{} // 状态1、启用 2、禁用
	MinappStatus interface{} // 是否排除小程序显示 1-不排除 2-排除
	Chain        interface{} // IN 内链  OUT 外联
	LinkOpenType interface{} // 外部链接跳转方式 1-webview 2-浏览器
	CreateAt     *gtime.Time // 创建时间
	UpdateAt     *gtime.Time // 更新时间
	DeletedAt    *gtime.Time //
}
