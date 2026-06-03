// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsLanguageConfig is the golang structure for table pms_language_config.
type PmsLanguageConfig struct {
	Id       int         `json:"id"       orm:"id"        description:""`
	Tag      string      `json:"tag"      orm:"tag"       description:"语言标签"`
	Name     string      `json:"name"     orm:"name"      description:"语言名称"`
	BaiduTag string      `json:"baiduTag" orm:"baidu_tag" description:"百度翻译语种"`
	Flag     string      `json:"flag"     orm:"flag"      description:"国旗"`
	CreateAt *gtime.Time `json:"createAt" orm:"create_at" description:""`
	UpdateAt *gtime.Time `json:"updateAt" orm:"update_at" description:""`
}
