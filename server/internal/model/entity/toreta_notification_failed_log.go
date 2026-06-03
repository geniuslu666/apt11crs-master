// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ToretaNotificationFailedLog is the golang structure for table toreta_notification_failed_log.
type ToretaNotificationFailedLog struct {
	Id             uint64      `json:"id"             orm:"id"              description:"主键ID"`
	ResourceKey    string      `json:"resourceKey"    orm:"resource_key"    description:"资源密钥"`
	ResourceType   string      `json:"resourceType"   orm:"resource_type"   description:"资源类型"`
	ResourceAction string      `json:"resourceAction" orm:"resource_action" description:"资源操作"`
	RestaurantKey  string      `json:"restaurantKey"  orm:"restaurant_key"  description:"餐厅密钥"`
	StatusCode     int         `json:"statusCode"     orm:"status_code"     description:"HTTP状态码"`
	RetryCount     int         `json:"retryCount"     orm:"retry_count"     description:"重试次数"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:"更新时间"`
}
