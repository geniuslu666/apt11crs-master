// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserClient is the golang structure of table user_client for DAO operations like Where/Data.
type UserClient struct {
	g.Meta    `orm:"table:user_client, do:true"`
	Id        interface{} //
	Kefu      interface{} // 客服账户
	ClientId  interface{} // 设备ID
	CreatedAt *gtime.Time // 创建时间
}
