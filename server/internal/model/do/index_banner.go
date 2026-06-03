// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// IndexBanner is the golang structure of table hg_index_banner for DAO operations like Where/Data.
type IndexBanner struct {
	g.Meta       `orm:"table:hg_index_banner, do:true"`
	Id           interface{} //
	Language     interface{} // 语言
	BannerImage  interface{} // 轮播图
	Model        interface{} // 模块   BANNER  横幅   EVENTS  事件
	Chain        interface{} // IN 内链  OUT 外联
	Path         interface{} // 链接内容
	BannerStatus interface{} // 1、启用 2、禁用
	MinappStatus interface{} // 是否排除小程序显示 1-不排除 2-排除
	Sort         interface{} // 排序(越大越靠前)
	LinkOpenType interface{} // 外部链接跳转方式 1-webview 2-浏览器
	CreateAt     *gtime.Time // 创建时间
	UpdateAt     *gtime.Time // 更新时间
}
