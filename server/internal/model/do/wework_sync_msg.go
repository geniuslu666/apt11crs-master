// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// WeworkSyncMsg is the golang structure of table wework_sync_msg for DAO operations like Where/Data.
type WeworkSyncMsg struct {
	g.Meta     `orm:"table:wework_sync_msg, do:true"`
	Id         interface{} //
	SyncCursor interface{} //
	JsonTxt    interface{} //
	VisitorId  interface{} //
	KefuId     interface{} //
	CreatedAt  *gtime.Time //
	EntId      interface{} //
}
