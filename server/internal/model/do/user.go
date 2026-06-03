// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta         `orm:"table:user, do:true"`
	Id             interface{} //
	Pid            interface{} // 父级ID
	Name           interface{} // 账户
	Password       interface{} // 密码，md5加密
	Nickname       interface{} // 显示名称
	CreatedAt      *gtime.Time // 创建时间
	ExpiredAt      *gtime.Time // 过期时间
	UpdatedAt      *gtime.Time // 更新时间
	DeletedAt      *gtime.Time // 删除时间
	Avator         interface{} // 客服头像
	RecNum         interface{} // 正在接待数量
	OnlineStatus   interface{} // 在线状态，1在线，2离线
	Status         interface{} // 开启状态，0正常，1关闭
	AgentNum       interface{} // 子账号个数
	Email          interface{} // 绑定邮箱
	CompanyPic     interface{} // 宣传图
	Tel            interface{} // 绑定手机
	Uuid           interface{} // 企业uuid
	ReceptionValue interface{} // 接待固定值
	Reception      interface{} // 接待模式
}
