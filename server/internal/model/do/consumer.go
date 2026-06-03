// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Consumer is the golang structure of table consumer for DAO operations like Where/Data.
type Consumer struct {
	g.Meta     `orm:"table:consumer, do:true"`
	Id         interface{} //
	Company    interface{} // 公司名
	Realname   interface{} // 姓名
	Score      interface{} // 级别
	ConsumerSn interface{} // 客户编号
	EntId      interface{} // 企业ID
	KefuName   interface{} // kefu名称
	Tel        interface{} // 手机
	Wechat     interface{} // 微信
	Qq         interface{} // qq
	Email      interface{} // 邮箱
	Remark     interface{} // 备注
	CreatedAt  *gtime.Time // 创建时间
}
