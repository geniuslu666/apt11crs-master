// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AigcSessionMessage is the golang structure of table aigc_session_message for DAO operations like Where/Data.
type AigcSessionMessage struct {
	g.Meta     `orm:"table:aigc_session_message, do:true"`
	Id         interface{} //
	CollectId  interface{} // 集合ID
	CreatedAt  *gtime.Time // 创建时间
	KefuAvatar interface{} // 客服头像
	AiAvatar   interface{} // AI头像
	Content    interface{} // 内容
	KefuName   interface{} // 客服名称
	EntId      interface{} // 企业ID
	MsgType    interface{} // 消息类型
}
