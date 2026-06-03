// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ImMessage is the golang structure of table hg_im_message for DAO operations like Where/Data.
type ImMessage struct {
	g.Meta       `orm:"table:hg_im_message, do:true"`
	Id           interface{} //
	Source       interface{} // 消息发起方
	SourceId     interface{} // 身份对应角色ID
	SourceName   interface{} // 身份对应角色昵称
	SourcePhoto  interface{} // 身份对应角色照片
	OrderSn      interface{} // 订单号
	FromMemberId interface{} // 发起消息方的用户ID
	ToMemberId   interface{} // 接收消息方的用户ID
	Content      *gjson.Json // 消息内容
	CreatedAt    *gtime.Time // 创建消息时间
	UpdatedAt    *gtime.Time // 更新消息时间
}
