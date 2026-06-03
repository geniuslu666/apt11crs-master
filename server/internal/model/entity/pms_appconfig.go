// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsAppconfig is the golang structure for table pms_appconfig.
type PmsAppconfig struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	Name      string      `json:"name"      orm:"name"       description:"配置名称"`
	Key       string      `json:"key"       orm:"key"        description:"配置项"`
	Value     string      `json:"value"     orm:"value"      description:"配置值"`
	Language  string      `json:"language"  orm:"language"   description:"语言"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:""`
}
