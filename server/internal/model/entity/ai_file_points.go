// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AiFilePoints is the golang structure for table ai_file_points.
type AiFilePoints struct {
	Id          int         `json:"id"          orm:"id"           description:""`
	FileId      string      `json:"fileId"      orm:"file_id"      description:"文件表自增ID"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	CollectName string      `json:"collectName" orm:"collect_name" description:"集合名称"`
	PointsId    string      `json:"pointsId"    orm:"points_id"    description:"向量ID"`
}
