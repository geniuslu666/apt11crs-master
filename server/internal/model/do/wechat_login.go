// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// WechatLogin is the golang structure of table wechat_login for DAO operations like Where/Data.
type WechatLogin struct {
	g.Meta     `orm:"table:wechat_login, do:true"`
	Id         interface{} //
	KefuName   interface{} // 客服账户
	OpenId     interface{} // 微信公众号openid
	TempKefuId interface{} // 临时客服ID
	Status     interface{} // 当前状态
	EntId      interface{} // 企业ID
	LoginIp    interface{} // 登录IP
	CreatedAt  *gtime.Time // 创建时间
}
