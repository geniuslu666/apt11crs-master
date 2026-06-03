// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Welcome is the golang structure for table welcome.
type Welcome struct {
	Id          uint        `json:"id"          orm:"id"           description:""`
	UserId      string      `json:"userId"      orm:"user_id"      description:"客服账户"`
	Keyword     string      `json:"keyword"     orm:"keyword"      description:"关键词，welcome为默认欢迎；wechat为公众号欢迎语"`
	Content     string      `json:"content"     orm:"content"      description:"欢迎消息内容"`
	IsDefault   uint        `json:"isDefault"   orm:"is_default"   description:"是否默认，未启用"`
	DelaySecond uint        `json:"delaySecond" orm:"delay_second" description:"延迟秒数"`
	Ctime       *gtime.Time `json:"ctime"       orm:"ctime"        description:"创建时间"`
}
