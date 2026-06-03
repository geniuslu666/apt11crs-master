// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DouyinWebhook is the golang structure of table douyin_webhook for DAO operations like Where/Data.
type DouyinWebhook struct {
	g.Meta     `orm:"table:douyin_webhook, do:true"`
	Id         interface{} //
	KefuName   interface{} // 客服账户
	Event      interface{} // event
	FromUserId interface{} // from_user_id
	ToUserId   interface{} // to_user_id
	ClientKey  interface{} // client_key
	Content    interface{} // content
	EntId      interface{} // 企业ID
	CreatedAt  *gtime.Time // 创建时间
}
