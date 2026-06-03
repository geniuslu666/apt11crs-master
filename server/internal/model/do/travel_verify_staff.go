// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TravelVerifyStaff is the golang structure of table hg_travel_verify_staff for DAO operations like Where/Data.
type TravelVerifyStaff struct {
	g.Meta             `orm:"table:hg_travel_verify_staff, do:true"`
	Id                 interface{} //
	Name               interface{} // 姓名
	Mobile             interface{} // 电话
	Username           interface{} // 登录账号
	Status             interface{} // 状态（1启用 2禁用）
	PasswordHash       interface{} // 密码
	Salt               interface{} // 密码盐
	PasswordResetToken interface{} // 密码重置令牌
	CreatedAt          *gtime.Time // 创建时间
	UpdatedAt          *gtime.Time // 更新时间
	DeletedAt          *gtime.Time // 软删除时间（NULL=正常）
}
