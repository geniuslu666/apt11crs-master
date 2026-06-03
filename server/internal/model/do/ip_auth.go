// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// IpAuth is the golang structure of table ip_auth for DAO operations like Where/Data.
type IpAuth struct {
	g.Meta     `orm:"table:ip_auth, do:true"`
	Id         interface{} //
	Content    interface{} // 备注
	IpAddress  interface{} // IP地址
	ExpireTime interface{} // 过期时间，未启用
	Phone      interface{} // 客服账户手机号
	CreatedAt  *gtime.Time // 创建时间
	Status     interface{} // 开启状态，1正常，2关闭
	Code       interface{} // 授权码
}
