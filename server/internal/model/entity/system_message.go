// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemMessage is the golang structure for table system_message.
type SystemMessage struct {
	Id                uint64      `json:"id"                orm:"id"                  description:"消息ID"`
	Scene             string      `json:"scene"             orm:"scene"               description:"场景：system-系统，hotel-酒店，food-餐饮，car-接送机，spa-按摩"`
	Type              string      `json:"type"              orm:"type"                description:"类型：order-订单流转，im-聊天"`
	MemberId          uint64      `json:"memberId"          orm:"member_id"           description:"会员ID (为0则是所有会员)"`
	OrderSn           string      `json:"orderSn"           orm:"order_sn"            description:"订单号"`
	GroupIds          string      `json:"groupIds"          orm:"group_ids"           description:"会员分组IDS"`
	LevelIds          string      `json:"levelIds"          orm:"level_ids"           description:"会员等级IDS"`
	EmployeeIds       string      `json:"employeeIds"       orm:"employee_ids"        description:"员工IDS"`
	SendObject        string      `json:"sendObject"        orm:"send_object"         description:"发送对象"`
	Title             *gjson.Json `json:"title"             orm:"title"               description:"消息标题，多语言JSON格式：{\"zh\":\"中文标题\",\"en\":\"English Title\"}"`
	Content           *gjson.Json `json:"content"           orm:"content"             description:"消息内容，多语言JSON格式：{\"zh\":\"中文副标题\",\"en\":\"English Subtitle\"}"`
	Image             string      `json:"image"             orm:"image"               description:"消息图片URL"`
	AppLink           string      `json:"appLink"           orm:"app_link"            description:"APP跳转链接"`
	WxLink            string      `json:"wxLink"            orm:"wx_link"             description:"微信跳转链接"`
	UrlParam          *gjson.Json `json:"urlParam"          orm:"url_param"           description:"APP跳转链接参数，JSON格式存储"`
	MobilePushSuccess int         `json:"mobilePushSuccess" orm:"mobile_push_success" description:"手机消息推送是否成功：0-失败，1-成功，NULL-未推送"`
	MobileSmsSuccess  int         `json:"mobileSmsSuccess"  orm:"mobile_sms_success"  description:"手机短信是否发送成功：0-失败，1-成功，NULL-未发送"`
	OperatorRole      string      `json:"operatorRole"      orm:"operator_role"       description:"操作员角色"`
	OperatorId        uint64      `json:"operatorId"        orm:"operator_id"         description:"操作员ID"`
	ShowIndex         int         `json:"showIndex"         orm:"show_index"          description:"是否显示在首页"`
	CreatedAt         *gtime.Time `json:"createdAt"         orm:"created_at"          description:"创建时间"`
	UpdatedAt         *gtime.Time `json:"updatedAt"         orm:"updated_at"          description:"更新时间"`
}
