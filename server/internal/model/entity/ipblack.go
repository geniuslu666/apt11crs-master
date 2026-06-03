// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Ipblack is the golang structure for table ipblack.
type Ipblack struct {
	Id       int         `json:"id"       orm:"id"        description:""`
	Ip       string      `json:"ip"       orm:"ip"        description:"IP地址"`
	Name     string      `json:"name"     orm:"name"      description:"名称"`
	CreateAt *gtime.Time `json:"createAt" orm:"create_at" description:"创建时间"`
	KefuId   string      `json:"kefuId"   orm:"kefu_id"   description:"操作的客服账户"`
	EntId    string      `json:"entId"    orm:"ent_id"    description:"客服企业ID"`
}
