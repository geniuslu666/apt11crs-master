// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysTerminalModel is the golang structure for table sys_terminal_model.
type SysTerminalModel struct {
	Id           int         `json:"id"           orm:"id"            description:""`
	BrandModel   string      `json:"brandModel"   orm:"brand_model"   description:"品牌型号"`
	ClientId     string      `json:"clientId"     orm:"client_id"     description:"开发者ID"`
	ClientSecret string      `json:"clientSecret" orm:"client_secret" description:"开发者秘钥"`
	CreateAt     *gtime.Time `json:"createAt"     orm:"create_at"     description:"创建时间"`
	UpdateAt     *gtime.Time `json:"updateAt"     orm:"update_at"     description:"更新时间"`
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"    description:""`
}
