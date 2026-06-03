// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AiFile is the golang structure for table ai_file.
type AiFile struct {
	Id          int         `json:"id"          orm:"id"           description:""`
	FileName    string      `json:"fileName"    orm:"file_name"    description:"文件名"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	FileSize    string      `json:"fileSize"    orm:"file_size"    description:"字符数"`
	CollectName string      `json:"collectName" orm:"collect_name" description:"集合名称"`
}
