// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FlyLog is the golang structure of table fly_log for DAO operations like Where/Data.
type FlyLog struct {
	g.Meta     `orm:"table:fly_log, do:true"`
	Id         interface{} //
	EntId      interface{} // 企业ID
	LogType    interface{} // 日志类型
	LogContent interface{} // 日志内容
	IpAddress  interface{} // 用户IP地址
	CreatedAt  *gtime.Time // 创建时间
}
