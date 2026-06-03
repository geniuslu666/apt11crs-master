// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserClient is the golang structure for table user_client.
type UserClient struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	Kefu      string      `json:"kefu"      orm:"kefu"       description:"客服账户"`
	ClientId  string      `json:"clientId"  orm:"client_id"  description:"设备ID"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
}
