// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id             int         `json:"id"             orm:"id"              description:""`
	Pid            uint        `json:"pid"            orm:"pid"             description:"父级ID"`
	Name           string      `json:"name"           orm:"name"            description:"账户"`
	Password       string      `json:"password"       orm:"password"        description:"密码，md5加密"`
	Nickname       string      `json:"nickname"       orm:"nickname"        description:"显示名称"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:"创建时间"`
	ExpiredAt      *gtime.Time `json:"expiredAt"      orm:"expired_at"      description:"过期时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:"更新时间"`
	DeletedAt      *gtime.Time `json:"deletedAt"      orm:"deleted_at"      description:"删除时间"`
	Avator         string      `json:"avator"         orm:"avator"          description:"客服头像"`
	RecNum         uint        `json:"recNum"         orm:"rec_num"         description:"正在接待数量"`
	OnlineStatus   int         `json:"onlineStatus"   orm:"online_status"   description:"在线状态，1在线，2离线"`
	Status         int         `json:"status"         orm:"status"          description:"开启状态，0正常，1关闭"`
	AgentNum       uint        `json:"agentNum"       orm:"agent_num"       description:"子账号个数"`
	Email          string      `json:"email"          orm:"email"           description:"绑定邮箱"`
	CompanyPic     string      `json:"companyPic"     orm:"company_pic"     description:"宣传图"`
	Tel            string      `json:"tel"            orm:"tel"             description:"绑定手机"`
	Uuid           string      `json:"uuid"           orm:"uuid"            description:"企业uuid"`
	ReceptionValue string      `json:"receptionValue" orm:"reception_value" description:"接待固定值"`
	Reception      string      `json:"reception"      orm:"reception"       description:"接待模式"`
}
