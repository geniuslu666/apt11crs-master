// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemMessageDao is the data access object for the table hg_system_message.
type SystemMessageDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns SystemMessageColumns // columns contains all the column names of Table for convenient usage.
}

// SystemMessageColumns defines and stores column names for the table hg_system_message.
type SystemMessageColumns struct {
	Id                string // 消息ID
	Scene             string // 场景：system-系统，hotel-酒店，food-餐饮，car-接送机，spa-按摩
	Type              string // 类型：order-订单流转，im-聊天
	MemberId          string // 会员ID (为0则是所有会员)
	OrderSn           string // 订单号
	GroupIds          string // 会员分组IDS
	LevelIds          string // 会员等级IDS
	EmployeeIds       string // 员工IDS
	SendObject        string // 发送对象
	Title             string // 消息标题，多语言JSON格式：{"zh":"中文标题","en":"English Title"}
	Content           string // 消息内容，多语言JSON格式：{"zh":"中文副标题","en":"English Subtitle"}
	Image             string // 消息图片URL
	AppLink           string // APP跳转链接
	WxLink            string // 微信跳转链接
	UrlParam          string // APP跳转链接参数，JSON格式存储
	MobilePushSuccess string // 手机消息推送是否成功：0-失败，1-成功，NULL-未推送
	MobileSmsSuccess  string // 手机短信是否发送成功：0-失败，1-成功，NULL-未发送
	OperatorRole      string // 操作员角色
	OperatorId        string // 操作员ID
	ShowIndex         string // 是否显示在首页
	CreatedAt         string // 创建时间
	UpdatedAt         string // 更新时间
}

// systemMessageColumns holds the columns for the table hg_system_message.
var systemMessageColumns = SystemMessageColumns{
	Id:                "id",
	Scene:             "scene",
	Type:              "type",
	MemberId:          "member_id",
	OrderSn:           "order_sn",
	GroupIds:          "group_ids",
	LevelIds:          "level_ids",
	EmployeeIds:       "employee_ids",
	SendObject:        "send_object",
	Title:             "title",
	Content:           "content",
	Image:             "image",
	AppLink:           "app_link",
	WxLink:            "wx_link",
	UrlParam:          "url_param",
	MobilePushSuccess: "mobile_push_success",
	MobileSmsSuccess:  "mobile_sms_success",
	OperatorRole:      "operator_role",
	OperatorId:        "operator_id",
	ShowIndex:         "show_index",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
}

// NewSystemMessageDao creates and returns a new DAO object for table data access.
func NewSystemMessageDao() *SystemMessageDao {
	return &SystemMessageDao{
		group:   "default",
		table:   "hg_system_message",
		columns: systemMessageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SystemMessageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SystemMessageDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SystemMessageDao) Columns() SystemMessageColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SystemMessageDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SystemMessageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SystemMessageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
