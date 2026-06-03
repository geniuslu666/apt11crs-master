// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ToretaNotificationFailedLog is the golang structure of table hg_toreta_notification_failed_log for DAO operations like Where/Data.
type ToretaNotificationFailedLog struct {
	g.Meta         `orm:"table:hg_toreta_notification_failed_log, do:true"`
	Id             interface{} // 主键ID
	ResourceKey    interface{} // 资源密钥
	ResourceType   interface{} // 资源类型
	ResourceAction interface{} // 资源操作
	RestaurantKey  interface{} // 餐厅密钥
	StatusCode     interface{} // HTTP状态码
	RetryCount     interface{} // 重试次数
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
