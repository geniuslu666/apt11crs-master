// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsMemberAuth is the golang structure of table hg_pms_member_auth for DAO operations like Where/Data.
type PmsMemberAuth struct {
	g.Meta    `orm:"table:hg_pms_member_auth, do:true"`
	Id        interface{} //
	AuthId    interface{} // 第三方用户ID标识
	WxUnionid interface{} // 微信unionid
	Channel   interface{} // WX、APPLE
	MemberId  interface{} // 关联的会员ID
	Email     interface{} // 邮箱
	Phone     interface{} // 手机号
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
	IsFxAuth  interface{} // 是否是分销授权用户
}
