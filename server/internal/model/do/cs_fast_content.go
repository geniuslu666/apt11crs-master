// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CsFastContent is the golang structure of table hg_cs_fast_content for DAO operations like Where/Data.
type CsFastContent struct {
	g.Meta      `orm:"table:hg_cs_fast_content, do:true"`
	Id          interface{} //
	Type        interface{} // 类型    welcome    欢迎语    option
	ZhContent   interface{} // 中文简体内容
	ZhCnContent interface{} // 中文繁体内容
	EnContent   interface{} // 英文内容
	JaContent   interface{} // 日文内容
	KoContent   interface{} // 韩文内容
	Sort        interface{} // 排序  从大到小排序
	Status      interface{} // 启禁用
	CreateAt    *gtime.Time //
	UpdateAt    *gtime.Time //
}
