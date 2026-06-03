// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FoodRestaurantDao is the data access object for the table hg_food_restaurant.
type FoodRestaurantDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of the current DAO.
	columns FoodRestaurantColumns // columns contains all the column names of Table for convenient usage.
}

// FoodRestaurantColumns defines and stores column names for the table hg_food_restaurant.
type FoodRestaurantColumns struct {
	Id                      string //
	OrderMode               string // 预定模式
	ToretaId                string // Toreta的餐厅ID
	Name                    string // 名称
	CuisineIds              string // 菜系（多选）
	LabelIds                string // 标签（多选）
	CooperateTypeId         string // 合作类型ID
	Logo                    string // LOGO
	Images                  string // 图集
	Content                 string // 简介
	Phone                   string // 联系电话
	OpenTime                string // 营业时间
	AreaPid                 string // 地区省级ID
	AreaId                  string // 地区市级ID
	DetailAddress           string // 详细地址
	GgLat                   string // 谷歌纬度
	GgLng                   string // 谷歌经度
	Lat                     string // 纬度
	Lng                     string // 经度
	OpenStatus              string // 营业状态
	OrderTimeType           string // 预约日期  1每天 2自定义
	OrderTimeWeek           string // 预约日期自定义周数
	RestTimeWeek            string // 定休日周数
	RestTimeDate            string // 定休日固定日期
	OrderConfirmType        string // 预约确定模式 1手动确认 2自动确认(废除)
	OrderConfirmDayType     string // 1 手动确认在前   2自动确认在前
	OrderConfirmBeforeDays  string // 多少天内手动确认
	OrderConfirmTime        string // 预约确认自定义分钟数
	OpenTimeType            string // 1 按时段  2 按时间点
	AmOpenTime              string // 早市开始时间
	AmCloseTime             string // 早市结束时间
	PmOpenTime              string // 晚市开始时间
	PmCloseTime             string // 晚市结束时间
	TimePoints              string // 时间点
	TimeDuration            string // 时间间隔
	OrderMaxType            string // 最大容纳数  1按时段  2按时间点
	TimeDurationMax         string // 时段最大容纳预定数量  0不限制
	DayTimeLimit            string // 当天预约处理截止时间
	AdvanceOrderDay         string // 至少提前预约天数  0当日可约
	MaxOrderDay             string // 最长预约天数
	CancelPolicyOpen        string // 是否允许取消  1允许  2不允许
	BeforeConfirmCancelRate string // 订单确认前取消费用为订单的%，0则免费
	AfterConfirmCancelDay   string // 订单确认后距离到店多少天
	AfterConfirmCancelRate1 string // 订单确认后距离到店天数前取消费用为订单的%，0则免费
	AfterConfirmCancelRate2 string // 订单确认后距离到店天数后取消费用为订单的%，0则免费
	MaxSeat                 string // 最大席位数
	CanSmoking              string // 是否允许抽烟  1允许  2不允许
	SettlementType          string // 门店抽成类型 1跟随系统  2自定义
	SettlementId            string // 结算模式ID
	SettlementRate          string // 门店结算比例
	Desc                    string // 开业/停业原因
	TotalOrderNum           string // 预约单总数量（包含退款）
	TotalOrderAmount        string // 预约单总金额（包含退款）
	PayOrderNum             string // 预约单支付数量（不包含退款）
	PayOrderAmount          string // 预约单支付金额（不包含退款）
	WaitConfirmOrderNum     string // 预约单待确认数量（已支付待确认）
	SettlementOrderNum      string // 预约单已结算数量
	SettlementOrderAmount   string // 已结算预约单订单金额
	TotalSettlementAmount   string // 已结算金额
	VerifyCode              string // 核销码
	CanOrder                string // 是否开放预定  1开放  2关闭
	Sort                    string // 排序(越大越靠前)
	DepositRate             string // 定金比例
	Account                 string // 账号
	PasswordHash            string // 密码
	Salt                    string // 密码盐
	PasswordResetToken      string // 密码重置令牌
	ToretaCancelEnable      string // Toreta是否允许取消 1 允许  2 不允许
	ToretaCancelLimitDay    string // Toreta允许取消几天前
	ToretaCancelLimitTime   string // Toreta允许取消时间前
	MaxDatetimeOrderOpen    string // 日期限制 1 开启  2关闭
	MaxDatetimeOrderDate    string // 限制的日期 逗号分隔
	CreateAt                string // 创建时间
	UpdateAt                string // 更新时间
	DeletedAt               string //
}

// foodRestaurantColumns holds the columns for the table hg_food_restaurant.
var foodRestaurantColumns = FoodRestaurantColumns{
	Id:                      "id",
	OrderMode:               "order_mode",
	ToretaId:                "toreta_id",
	Name:                    "name",
	CuisineIds:              "cuisine_ids",
	LabelIds:                "label_ids",
	CooperateTypeId:         "cooperate_type_id",
	Logo:                    "logo",
	Images:                  "images",
	Content:                 "content",
	Phone:                   "phone",
	OpenTime:                "open_time",
	AreaPid:                 "area_pid",
	AreaId:                  "area_id",
	DetailAddress:           "detail_address",
	GgLat:                   "gg_lat",
	GgLng:                   "gg_lng",
	Lat:                     "lat",
	Lng:                     "lng",
	OpenStatus:              "open_status",
	OrderTimeType:           "order_time_type",
	OrderTimeWeek:           "order_time_week",
	RestTimeWeek:            "rest_time_week",
	RestTimeDate:            "rest_time_date",
	OrderConfirmType:        "order_confirm_type",
	OrderConfirmDayType:     "order_confirm_day_type",
	OrderConfirmBeforeDays:  "order_confirm_before_days",
	OrderConfirmTime:        "order_confirm_time",
	OpenTimeType:            "open_time_type",
	AmOpenTime:              "am_open_time",
	AmCloseTime:             "am_close_time",
	PmOpenTime:              "pm_open_time",
	PmCloseTime:             "pm_close_time",
	TimePoints:              "time_points",
	TimeDuration:            "time_duration",
	OrderMaxType:            "order_max_type",
	TimeDurationMax:         "time_duration_max",
	DayTimeLimit:            "day_time_limit",
	AdvanceOrderDay:         "advance_order_day",
	MaxOrderDay:             "max_order_day",
	CancelPolicyOpen:        "cancel_policy_open",
	BeforeConfirmCancelRate: "before_confirm_cancel_rate",
	AfterConfirmCancelDay:   "after_confirm_cancel_day",
	AfterConfirmCancelRate1: "after_confirm_cancel_rate1",
	AfterConfirmCancelRate2: "after_confirm_cancel_rate2",
	MaxSeat:                 "max_seat",
	CanSmoking:              "can_smoking",
	SettlementType:          "settlement_type",
	SettlementId:            "settlement_id",
	SettlementRate:          "settlement_rate",
	Desc:                    "desc",
	TotalOrderNum:           "total_order_num",
	TotalOrderAmount:        "total_order_amount",
	PayOrderNum:             "pay_order_num",
	PayOrderAmount:          "pay_order_amount",
	WaitConfirmOrderNum:     "wait_confirm_order_num",
	SettlementOrderNum:      "settlement_order_num",
	SettlementOrderAmount:   "settlement_order_amount",
	TotalSettlementAmount:   "total_settlement_amount",
	VerifyCode:              "verify_code",
	CanOrder:                "can_order",
	Sort:                    "sort",
	DepositRate:             "deposit_rate",
	Account:                 "account",
	PasswordHash:            "password_hash",
	Salt:                    "salt",
	PasswordResetToken:      "password_reset_token",
	ToretaCancelEnable:      "toreta_cancel_enable",
	ToretaCancelLimitDay:    "toreta_cancel_limit_day",
	ToretaCancelLimitTime:   "toreta_cancel_limit_time",
	MaxDatetimeOrderOpen:    "max_datetime_order_open",
	MaxDatetimeOrderDate:    "max_datetime_order_date",
	CreateAt:                "create_at",
	UpdateAt:                "update_at",
	DeletedAt:               "deleted_at",
}

// NewFoodRestaurantDao creates and returns a new DAO object for table data access.
func NewFoodRestaurantDao() *FoodRestaurantDao {
	return &FoodRestaurantDao{
		group:   "default",
		table:   "hg_food_restaurant",
		columns: foodRestaurantColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FoodRestaurantDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FoodRestaurantDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FoodRestaurantDao) Columns() FoodRestaurantColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FoodRestaurantDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FoodRestaurantDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *FoodRestaurantDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
