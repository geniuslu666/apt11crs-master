// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CarDriverDao is the data access object for the table hg_car_driver.
type CarDriverDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns CarDriverColumns // columns contains all the column names of Table for convenient usage.
}

// CarDriverColumns defines and stores column names for the table hg_car_driver.
type CarDriverColumns struct {
	Id                    string //
	CarId                 string // 车辆ID
	IsLeader              string // 是否是车队长 1是  2否
	Name                  string // 司机真实姓名
	Nickname              string // 司机昵称
	CooperateTypeId       string // 合作类型ID
	Sex                   string // 1、男 2、女
	Phone                 string // 手机号
	PhoneArea             string // 手机区号
	Photo                 string // 照片
	Age                   string // 年龄
	WorkYears             string // 从业年数
	Language              string // 语言能力
	Status                string // 状态1、启用 2、禁用
	WorkStatus            string // 工作状态
	MemberId              string // 会员ID
	QualityMaterials      string // 资质信息(多图)
	SettlementType        string // 服务分成类型 1跟随系统  2自定义
	SettlementId          string // 结算模式ID
	SettlementRate        string // 服务分成%
	TotalOrderNum         string // 预约单总数量（包含退款）
	TotalOrderAmount      string // 预约单总金额（包含退款）
	PayOrderNum           string // 预约单支付数量（不包含退款）
	PayOrderAmount        string // 预约单支付金额（不包含退款）
	SettlementOrderNum    string // 预约单已结算数量
	SettlementOrderAmount string // 已结算预约单订单金额
	TotalSettlementAmount string // 已结算金额
	Balance               string // 余额
	VerifyMoney           string // 已核账金额
	ApplyWithdrawBalance  string // 提现中余额
	WithdrawBalance       string // 已提现余额
	CreateAt              string // 创建时间
	UpdateAt              string // 更新时间
	DeletedAt             string //
}

// carDriverColumns holds the columns for the table hg_car_driver.
var carDriverColumns = CarDriverColumns{
	Id:                    "id",
	CarId:                 "car_id",
	IsLeader:              "is_leader",
	Name:                  "name",
	Nickname:              "nickname",
	CooperateTypeId:       "cooperate_type_id",
	Sex:                   "sex",
	Phone:                 "phone",
	PhoneArea:             "phone_area",
	Photo:                 "photo",
	Age:                   "age",
	WorkYears:             "work_years",
	Language:              "language",
	Status:                "status",
	WorkStatus:            "work_status",
	MemberId:              "member_id",
	QualityMaterials:      "quality_materials",
	SettlementType:        "settlement_type",
	SettlementId:          "settlement_id",
	SettlementRate:        "settlement_rate",
	TotalOrderNum:         "total_order_num",
	TotalOrderAmount:      "total_order_amount",
	PayOrderNum:           "pay_order_num",
	PayOrderAmount:        "pay_order_amount",
	SettlementOrderNum:    "settlement_order_num",
	SettlementOrderAmount: "settlement_order_amount",
	TotalSettlementAmount: "total_settlement_amount",
	Balance:               "balance",
	VerifyMoney:           "verify_money",
	ApplyWithdrawBalance:  "apply_withdraw_balance",
	WithdrawBalance:       "withdraw_balance",
	CreateAt:              "create_at",
	UpdateAt:              "update_at",
	DeletedAt:             "deleted_at",
}

// NewCarDriverDao creates and returns a new DAO object for table data access.
func NewCarDriverDao() *CarDriverDao {
	return &CarDriverDao{
		group:   "default",
		table:   "hg_car_driver",
		columns: carDriverColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CarDriverDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CarDriverDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CarDriverDao) Columns() CarDriverColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CarDriverDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CarDriverDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CarDriverDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
