// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsAppReservationDao is the data access object for the table hg_pms_app_reservation.
type PmsAppReservationDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of the current DAO.
	columns PmsAppReservationColumns // columns contains all the column names of Table for convenient usage.
}

// PmsAppReservationColumns defines and stores column names for the table hg_pms_app_reservation.
type PmsAppReservationColumns struct {
	Id               string // 主键
	Uid              string // 我方系统 ID
	Uuid             string // airhost   ID
	Source           string // 来源
	SourceCode       string // 来源渠道
	SourceName       string // 渠道名称
	MemberId         string // 用户ID
	Puid             string // 物业ID
	OrderSn          string // 系统订单号
	OrderIndex       string // 订单索引
	OutOrderSn       string // 三方订单号
	RoomType         string // 房型信息，参考房型uid
	RoomUnit         string // 房间单元的id, uid或组合
	RatePlanId       string // 费率ID
	CheckinDate      string // 入住日期
	CheckoutDate     string // 退房日期
	CheckinTime      string // 入住时间，24小时格式
	CheckoutTime     string // 退房时间，24小时格式
	Status           string // 预订状态（确认/confirmed、取消/cancelled）
	CheckinStatus    string // 入住状态  before_checkin  在入住之前  checked_in  已入住  checked_out  已退房
	OrderStatus      string // WAIT_PAY、待支付 CANCEL、支付过期 HAVE_PAID、支付成功
	MainGuest        string // 住宿人编号
	AdultCount       string // 成人数量
	ChildCount       string // 儿童数量
	InfantCount      string // 婴儿数量
	PricePlanId      string // 价格plan ID
	ChangeAmount     string // 变动金额
	PricePlanInfo    string // 价格plan内容
	BookingFee       string // 预订费
	ChannelFee       string // 渠道费
	CleaningFee      string // 清洁费
	CancellationFee  string // 取消费，仅在取消时适用
	Charges          string // 费用详情，参考Charge对象
	GuestRemarks     string // 备注
	CancelRemake     string // 取消原因
	ExpValue         string // 结算的经验值
	ExpTime          string // 经验结算时间
	CreatedAt        string //
	UpdatedAt        string //
	IsChangeGuest    string // 是否变更入住人信息 Y 是 N 否
	IsChangeGuestId  string // 入住人信息变更ID
	IsChangePeople   string // 是否变更入住人数 Y 是 N 否
	IsChangePeopleId string // 入住人数信息变更ID
	IsChangeDate     string // 是否变更入住日期 Y 是 N 否
	IsChangeDateId   string // 入住日期变更ID
	IsFx             string // 是否是分销订单
}

// pmsAppReservationColumns holds the columns for the table hg_pms_app_reservation.
var pmsAppReservationColumns = PmsAppReservationColumns{
	Id:               "id",
	Uid:              "uid",
	Uuid:             "uuid",
	Source:           "source",
	SourceCode:       "source_code",
	SourceName:       "source_name",
	MemberId:         "member_id",
	Puid:             "puid",
	OrderSn:          "order_sn",
	OrderIndex:       "order_index",
	OutOrderSn:       "out_order_sn",
	RoomType:         "room_type",
	RoomUnit:         "room_unit",
	RatePlanId:       "rate_plan_id",
	CheckinDate:      "checkin_date",
	CheckoutDate:     "checkout_date",
	CheckinTime:      "checkin_time",
	CheckoutTime:     "checkout_time",
	Status:           "status",
	CheckinStatus:    "checkin_status",
	OrderStatus:      "order_status",
	MainGuest:        "main_guest",
	AdultCount:       "adult_count",
	ChildCount:       "child_count",
	InfantCount:      "infant_count",
	PricePlanId:      "price_plan_id",
	ChangeAmount:     "change_amount",
	PricePlanInfo:    "price_plan_info",
	BookingFee:       "booking_fee",
	ChannelFee:       "channel_fee",
	CleaningFee:      "cleaning_fee",
	CancellationFee:  "cancellation_fee",
	Charges:          "charges",
	GuestRemarks:     "guest_remarks",
	CancelRemake:     "cancel_remake",
	ExpValue:         "exp_value",
	ExpTime:          "exp_time",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	IsChangeGuest:    "is_change_guest",
	IsChangeGuestId:  "is_change_guest_id",
	IsChangePeople:   "is_change_people",
	IsChangePeopleId: "is_change_people_id",
	IsChangeDate:     "is_change_date",
	IsChangeDateId:   "is_change_date_id",
	IsFx:             "is_fx",
}

// NewPmsAppReservationDao creates and returns a new DAO object for table data access.
func NewPmsAppReservationDao() *PmsAppReservationDao {
	return &PmsAppReservationDao{
		group:   "default",
		table:   "hg_pms_app_reservation",
		columns: pmsAppReservationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsAppReservationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsAppReservationDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsAppReservationDao) Columns() PmsAppReservationColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsAppReservationDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsAppReservationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsAppReservationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
