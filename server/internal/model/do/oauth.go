// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Oauth is the golang structure of table oauth for DAO operations like Where/Data.
type Oauth struct {
	g.Meta    `orm:"table:oauth, do:true"`
	Id        interface{} //
	UserId    interface{} // 访客/客服账户
	OauthId   interface{} // 公众号OPENID
	CreatedAt *gtime.Time // 创建时间
	Status    interface{} // 状态，未启用
}
