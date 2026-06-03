// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UpDownLine is the golang structure of table up_down_line for DAO operations like Where/Data.
type UpDownLine struct {
	g.Meta       `orm:"table:up_down_line, do:true"`
	Id           interface{} //
	KefuName     interface{} // 客服账户
	EntId        interface{} // 企业ID
	OnlineStatus interface{} // 在线状态，1在线，2离线
	ClientIp     interface{} // ip地址
	CreatedAt    *gtime.Time // 创建时间
}
