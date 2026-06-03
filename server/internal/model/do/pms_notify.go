// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsNotify is the golang structure of table hg_pms_notify for DAO operations like Where/Data.
type PmsNotify struct {
	g.Meta        `orm:"table:hg_pms_notify, do:true"`
	Id            interface{} //
	MemberId      interface{} // 会员ID
	NotifyTitle   interface{} // 通知标题
	NotifyType    interface{} // 通知类型 SYSTEM、系统消息    BOOKING、预定消息
	NotifyContent interface{} // 通知内容
	NotifyData    *gjson.Json // 通知数据
	IsRead        interface{} // Y 已读  N 未读
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
	DeletedAt     *gtime.Time // 删除时间
}
