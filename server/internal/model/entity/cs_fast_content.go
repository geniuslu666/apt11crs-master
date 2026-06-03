// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CsFastContent is the golang structure for table cs_fast_content.
type CsFastContent struct {
	Id          int         `json:"id"          orm:"id"            description:""`
	Type        string      `json:"type"        orm:"type"          description:"类型    welcome    欢迎语    option"`
	ZhContent   string      `json:"zhContent"   orm:"zh_content"    description:"中文简体内容"`
	ZhCnContent string      `json:"zhCnContent" orm:"zh_cn_content" description:"中文繁体内容"`
	EnContent   string      `json:"enContent"   orm:"en_content"    description:"英文内容"`
	JaContent   string      `json:"jaContent"   orm:"ja_content"    description:"日文内容"`
	KoContent   string      `json:"koContent"   orm:"ko_content"    description:"韩文内容"`
	Sort        int         `json:"sort"        orm:"sort"          description:"排序  从大到小排序"`
	Status      string      `json:"status"      orm:"status"        description:"启禁用"`
	CreateAt    *gtime.Time `json:"createAt"    orm:"create_at"     description:""`
	UpdateAt    *gtime.Time `json:"updateAt"    orm:"update_at"     description:""`
}
