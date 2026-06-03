// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemMessage is the golang structure of table hg_system_message for DAO operations like Where/Data.
type SystemMessage struct {
	g.Meta            `orm:"table:hg_system_message, do:true"`
	Id                interface{} // 消息ID
	Scene             interface{} // 场景：system-系统，hotel-酒店，food-餐饮，car-接送机，spa-按摩
	Type              interface{} // 类型：order-订单流转，im-聊天
	MemberId          interface{} // 会员ID (为0则是所有会员)
	OrderSn           interface{} // 订单号
	GroupIds          interface{} // 会员分组IDS
	LevelIds          interface{} // 会员等级IDS
	EmployeeIds       interface{} // 员工IDS
	SendObject        interface{} // 发送对象
	Title             *gjson.Json // 消息标题，多语言JSON格式：{"zh":"中文标题","en":"English Title"}
	Content           *gjson.Json // 消息内容，多语言JSON格式：{"zh":"中文副标题","en":"English Subtitle"}
	Image             interface{} // 消息图片URL
	AppLink           interface{} // APP跳转链接
	WxLink            interface{} // 微信跳转链接
	UrlParam          *gjson.Json // APP跳转链接参数，JSON格式存储
	MobilePushSuccess interface{} // 手机消息推送是否成功：0-失败，1-成功，NULL-未推送
	MobileSmsSuccess  interface{} // 手机短信是否发送成功：0-失败，1-成功，NULL-未发送
	OperatorRole      interface{} // 操作员角色
	OperatorId        interface{} // 操作员ID
	ShowIndex         interface{} // 是否显示在首页
	CreatedAt         *gtime.Time // 创建时间
	UpdatedAt         *gtime.Time // 更新时间
}
