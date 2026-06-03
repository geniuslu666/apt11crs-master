// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FlyLog is the golang structure for table fly_log.
type FlyLog struct {
	Id         int         `json:"id"         orm:"id"          description:""`
	EntId      string      `json:"entId"      orm:"ent_id"      description:"企业ID"`
	LogType    string      `json:"logType"    orm:"log_type"    description:"日志类型"`
	LogContent string      `json:"logContent" orm:"log_content" description:"日志内容"`
	IpAddress  string      `json:"ipAddress"  orm:"ip_address"  description:"用户IP地址"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"`
}
