// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsLanguageConfig is the golang structure of table hg_pms_language_config for DAO operations like Where/Data.
type PmsLanguageConfig struct {
	g.Meta   `orm:"table:hg_pms_language_config, do:true"`
	Id       interface{} //
	Tag      interface{} // 语言标签
	Name     interface{} // 语言名称
	BaiduTag interface{} // 百度翻译语种
	Flag     interface{} // 国旗
	CreateAt *gtime.Time //
	UpdateAt *gtime.Time //
}
