// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TravelVerifyStaff is the golang structure for table travel_verify_staff.
type TravelVerifyStaff struct {
	Id                 uint64      `json:"id"                 orm:"id"                   description:""`
	Name               string      `json:"name"               orm:"name"                 description:"姓名"`
	Mobile             string      `json:"mobile"             orm:"mobile"               description:"电话"`
	Username           string      `json:"username"           orm:"username"             description:"登录账号"`
	Status             int         `json:"status"             orm:"status"               description:"状态（1启用 2禁用）"`
	PasswordHash       string      `json:"passwordHash"       orm:"password_hash"        description:"密码"`
	Salt               string      `json:"salt"               orm:"salt"                 description:"密码盐"`
	PasswordResetToken string      `json:"passwordResetToken" orm:"password_reset_token" description:"密码重置令牌"`
	CreatedAt          *gtime.Time `json:"createdAt"          orm:"created_at"           description:"创建时间"`
	UpdatedAt          *gtime.Time `json:"updatedAt"          orm:"updated_at"           description:"更新时间"`
	DeletedAt          *gtime.Time `json:"deletedAt"          orm:"deleted_at"           description:"软删除时间（NULL=正常）"`
}
