// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsMemberDao is the data access object for the table hg_pms_member.
type PmsMemberDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns PmsMemberColumns // columns contains all the column names of Table for convenient usage.
}

// PmsMemberColumns defines and stores column names for the table hg_pms_member.
type PmsMemberColumns struct {
	Id              string // 主键
	MemberNo        string // 会员号
	GroupId         string // 会员分组ID
	Avatar          string // 头像
	Sex             string // 1、男 2、女
	FirstName       string // 名
	LastName        string // 姓
	FullName        string // 全名
	Level           string // 等级
	Phone           string // 手机号
	PhoneArea       string // 手机区号
	Mail            string // 邮箱
	Password        string // 密码
	Birthday        string // 生日
	Balance         string // 积分
	Exp             string // 经验
	Referrer        string // 推荐人
	LastReferrer    string // 最后推荐人
	Source          string // 注册来源  IOS,Andriod,H5
	LoginMode       string // password,phone,email,googleOauth
	LastLoginIp     string // 上次登录IP
	LastLogin       string // 上次登录时间
	RegisterIp      string // 注册IP
	RegisterTime    string // 注册时间
	RegisterMdCode  string // 注册设备码
	RegisterMpModel string // 注册设备型号
	Address         string // 地址
	Nationality     string // 国籍
	Replenish       string // 是否完善过用户信息    N 未完善  Y 已完善
	RebateMode      string // MEMBER 会员     STAFF    员工 CHANNEL  渠道
	MemberId        string // 会员ID
	StaffId         string // 员工ID
	ChannelId       string // 渠道ID
	IsNewPrice      string // 是否获取过新人奖  1、获得  0、未获得  2、没有邀新奖励
	IsInviteRewards string //
	Status          string // 状态1、启用 2、禁用
	CreatedAt       string // 创建时间
	UpdatedAt       string // 更新时间
	DeletedAt       string // 删除时间
}

// pmsMemberColumns holds the columns for the table hg_pms_member.
var pmsMemberColumns = PmsMemberColumns{
	Id:              "id",
	MemberNo:        "member_no",
	GroupId:         "group_id",
	Avatar:          "avatar",
	Sex:             "sex",
	FirstName:       "first_name",
	LastName:        "last_name",
	FullName:        "full_name",
	Level:           "level",
	Phone:           "phone",
	PhoneArea:       "phone_area",
	Mail:            "mail",
	Password:        "password",
	Birthday:        "birthday",
	Balance:         "balance",
	Exp:             "exp",
	Referrer:        "referrer",
	LastReferrer:    "last_referrer",
	Source:          "source",
	LoginMode:       "login_mode",
	LastLoginIp:     "last_login_ip",
	LastLogin:       "last_login",
	RegisterIp:      "register_ip",
	RegisterTime:    "register_time",
	RegisterMdCode:  "register_md_code",
	RegisterMpModel: "register_mp_model",
	Address:         "address",
	Nationality:     "nationality",
	Replenish:       "replenish",
	RebateMode:      "rebate_mode",
	MemberId:        "member_id",
	StaffId:         "staff_id",
	ChannelId:       "channel_id",
	IsNewPrice:      "is_new_price",
	IsInviteRewards: "is_Invite_rewards",
	Status:          "status",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	DeletedAt:       "deleted_at",
}

// NewPmsMemberDao creates and returns a new DAO object for table data access.
func NewPmsMemberDao() *PmsMemberDao {
	return &PmsMemberDao{
		group:   "default",
		table:   "hg_pms_member",
		columns: pmsMemberColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsMemberDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsMemberDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsMemberDao) Columns() PmsMemberColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsMemberDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsMemberDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsMemberDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
