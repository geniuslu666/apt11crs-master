// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SpaTechnicianDao is the data access object for the table hg_spa_technician.
type SpaTechnicianDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns SpaTechnicianColumns // columns contains all the column names of Table for convenient usage.
}

// SpaTechnicianColumns defines and stores column names for the table hg_spa_technician.
type SpaTechnicianColumns struct {
	Id                    string //
	IspId                 string // 服务商ID
	IsLeader              string // 是否是店长 1是  2否
	Name                  string // 技师真实姓名
	Nickname              string // 技师昵称
	Sex                   string // 1、男 2、女
	Phone                 string // 手机号
	PhoneArea             string // 手机区号
	Photo                 string // 照片
	Age                   string // 年龄
	WorkYears             string // 从业年数
	Status                string // 状态1、启用 2、禁用
	WorkStatus            string // 工作状态
	MemberId              string // 会员ID
	WechatNo              string // 微信号
	QualityMaterials      string // 资质信息(多图)
	SettlementType        string // 服务分成类型 1跟随系统  2自定义
	SettlementId          string // 结算模式ID
	SettlementRate        string // 服务分成%
	OrderTimeType         string // 预约日期  1每天 2自定义
	OrderTimeWeek         string // 预约日期自定义周数
	RestTimeWeek          string // 定休日周数
	TotalOrderNum         string // 预约单总数量（包含退款）
	TotalOrderAmount      string // 预约单总金额（包含退款）
	PayOrderNum           string // 预约单支付数量（不包含退款）
	PayOrderAmount        string // 预约单支付金额（不包含退款）
	SettlementOrderNum    string // 预约单已结算数量
	SettlementOrderAmount string // 已结算预约单订单金额
	TotalSettlementAmount string // 已结算金额
	Balance               string // 余额
	VerifyMoney           string // 已核账金额
	CreateAt              string // 创建时间
	UpdateAt              string // 更新时间
	DeletedAt             string //
}

// spaTechnicianColumns holds the columns for the table hg_spa_technician.
var spaTechnicianColumns = SpaTechnicianColumns{
	Id:                    "id",
	IspId:                 "isp_id",
	IsLeader:              "is_leader",
	Name:                  "name",
	Nickname:              "nickname",
	Sex:                   "sex",
	Phone:                 "phone",
	PhoneArea:             "phone_area",
	Photo:                 "photo",
	Age:                   "age",
	WorkYears:             "work_years",
	Status:                "status",
	WorkStatus:            "work_status",
	MemberId:              "member_id",
	WechatNo:              "wechat_no",
	QualityMaterials:      "quality_materials",
	SettlementType:        "settlement_type",
	SettlementId:          "settlement_id",
	SettlementRate:        "settlement_rate",
	OrderTimeType:         "order_time_type",
	OrderTimeWeek:         "order_time_week",
	RestTimeWeek:          "rest_time_week",
	TotalOrderNum:         "total_order_num",
	TotalOrderAmount:      "total_order_amount",
	PayOrderNum:           "pay_order_num",
	PayOrderAmount:        "pay_order_amount",
	SettlementOrderNum:    "settlement_order_num",
	SettlementOrderAmount: "settlement_order_amount",
	TotalSettlementAmount: "total_settlement_amount",
	Balance:               "balance",
	VerifyMoney:           "verify_money",
	CreateAt:              "create_at",
	UpdateAt:              "update_at",
	DeletedAt:             "deleted_at",
}

// NewSpaTechnicianDao creates and returns a new DAO object for table data access.
func NewSpaTechnicianDao() *SpaTechnicianDao {
	return &SpaTechnicianDao{
		group:   "default",
		table:   "hg_spa_technician",
		columns: spaTechnicianColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SpaTechnicianDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SpaTechnicianDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SpaTechnicianDao) Columns() SpaTechnicianColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SpaTechnicianDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SpaTechnicianDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SpaTechnicianDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
