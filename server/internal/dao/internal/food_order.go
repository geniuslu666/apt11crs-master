// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FoodOrderDao is the data access object for table hg_food_order.
type FoodOrderDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns FoodOrderColumns // columns contains all the column names of Table for convenient usage.
}

// FoodOrderColumns defines and stores column names for table hg_food_order.
type FoodOrderColumns struct {
	Id                       string //
	OrderType                string // 订单类型
	OrderSn                  string // 订单编号
	OutOrderSn               string // 三方订单号
	MemberId                 string // 用户ID
	RestaurantId             string // 餐厅ID
	OrderAmount              string // 订单金额
	CouponAmount             string // 优惠券抵扣金额
	BalAmount                string // 积分抵扣金额
	GoodsId                  string // 套餐ID
	BookingName              string // 预定人姓名
	FirstName                string // 订单预定人姓
	LastName                 string // 订单预定人名
	GoodsNum                 string // 套餐数量
	PhoneArea                string // 手机区号
	BookingMobile            string // 预定人手机
	BookingEmail             string // 预定人邮箱
	BookDate                 string // 预定日期
	BookTime                 string // 预定时间
	BookDatetime             string // 预定日期时间
	SeatId                   string // 座位ID
	PayModel                 string // 1、余额支付 2、组合支付 3、纯外部支付
	PayTime                  string // 支付时间
	OrderStatus              string // 订单付款状态
	BookingStatus            string // 订单预定状态
	BookingTime              string // 订单确认时间
	ConfirmRefuseReason      string // 审核拒绝原因
	VerifyStatus             string // 订单核销状态
	VerifyCode               string // 核销码
	VerifyTime               string // 核销时间
	ExpirationTime           string // 订单过期时间
	SettlementRate           string // 结算比例
	SettlementStatus         string // WAIT 等待结算   SUCCESS  结算处理成功    FAIL  结算处理失败
	SettlementAmount         string // 结算金额
	SettlementTime           string // 结算时间
	SettlementOrderId        string // 结算单ID
	SettlementType           string // 结算方式  1无需结算 2按周期自动结算  3手动申请结算
	SettlementCycle          string // 结算周期  1每日结算  2每周结算  3每月结算
	MemberMessage            string // 购买人留言信息
	MemberMessageJa          string // 购买人留言信息日语版
	RestaurantMessage        string // 餐厅留言信息
	BookingCount             string // 预定人数
	RefundStatus             string // 退款状态   WAIT 未退款   PART  部分退款 DONE 全部退款
	RefundRate               string // 退款比例（废弃）
	RefundFee                string // 退款手续费
	RefundTime               string // 退款时间
	RefundAmount             string // 已退款总金额
	RefundBalAmount          string // 已退款积分
	RefundCouponAmount       string // 已退款优惠券
	AdminRefundAmount        string // 后台已退款总金额
	AdminRefundBalAmount     string // 后台已退款积分
	AdminRefundCouponAmount  string // 后台已退款优惠券
	AdminCancelReason        string // 后台取消原因
	ActivityId               string // 活动ID
	AdminCancelNum           string // 后台取消次数
	Referrer                 string // 推荐人
	RebateRate               string // 分佣比例
	RebateStatus             string // WAIT 等待处理佣金   SUCCESS   佣金处理成功    FAIL  佣金处理失败
	RebateAmount             string // 分佣结算金额
	RebateTime               string // 分佣结算时间
	IsGetOpen                string // 是否开启积分获取
	IsPayOpen                string // 是否开启积分抵扣
	FoodGetRateVip           string // 餐饮结算积分比例
	FoodGetRateScene         string // 场景结算积分比例
	FoodGetScoreStatus       string // 'WAIT','SUCCESS','FAIL'
	FoodGetAmount            string // 结算积分金额
	ExpValue                 string // 结算的经验值
	ExpTime                  string // 经验结算时间
	DepositRate              string // 定金比例
	DepositAmount            string // 定金
	DepositPayStatus         string // 定金付款状态
	DepositPayTime           string // 定金支付时间
	DepositCancelTime        string // 定金取消时间
	DepositRefundTime        string // 定金退款时间
	DepositExpirationTime    string // 定金支付过期时间
	RemainPayStatus          string // 尾款付款状态
	RemainCancelTime         string // 尾款取消时间
	RemainRefundTime         string // 尾款退款时间
	DepositCancelSource      string // 定金取消来源
	DepositCancelReason      string // 定金取消原因
	RemainCancelSource       string // 尾款取消来源
	RemainCancelReason       string // 尾款取消原因
	PayStep                  string //
	OldBookDate              string // 原预定日期
	OldBookTime              string // 原预定时间
	OldBookDatetime          string // 原预定日期时间
	OldBookingCount          string // 原预定人数
	ToretaReservationNo      string // Toreta预约号
	ToretaReservationId      string // Toreta预约ID
	ToretaReservationStatus  string // Toreta预约状态（0：未来店，1：到店，2：预约取消，3：未到，4：部分到店，5：网页取消，6：已用餐，7：已完成结账，8：已重置）
	ToretaReservationEndtime string // Toreta预约结束时间
	GoodsIsNoPay             string // 套餐是否无需支付：0-否，1-是
	ToretaHasTimeLimit       string // Toreta是否有时间限制  1-true  2-false
	ToretaEndTime            string // Toreta用餐结束时间
	CreatedAt                string // 创建时间
	UpdatedAt                string // 更新时间
	IsFx                     string // 是否是分销订单
}

// foodOrderColumns holds the columns for table hg_food_order.
var foodOrderColumns = FoodOrderColumns{
	Id:                       "id",
	OrderType:                "order_type",
	OrderSn:                  "order_sn",
	OutOrderSn:               "out_order_sn",
	MemberId:                 "member_id",
	RestaurantId:             "restaurant_id",
	OrderAmount:              "order_amount",
	CouponAmount:             "coupon_amount",
	BalAmount:                "bal_amount",
	GoodsId:                  "goods_id",
	BookingName:              "booking_name",
	FirstName:                "first_name",
	LastName:                 "last_name",
	GoodsNum:                 "goods_num",
	PhoneArea:                "phone_area",
	BookingMobile:            "booking_mobile",
	BookingEmail:             "booking_email",
	BookDate:                 "book_date",
	BookTime:                 "book_time",
	BookDatetime:             "book_datetime",
	SeatId:                   "seat_id",
	PayModel:                 "pay_model",
	PayTime:                  "pay_time",
	OrderStatus:              "order_status",
	BookingStatus:            "booking_status",
	BookingTime:              "booking_time",
	ConfirmRefuseReason:      "confirm_refuse_reason",
	VerifyStatus:             "verify_status",
	VerifyCode:               "verify_code",
	VerifyTime:               "verify_time",
	ExpirationTime:           "expiration_time",
	SettlementRate:           "settlement_rate",
	SettlementStatus:         "settlement_status",
	SettlementAmount:         "settlement_amount",
	SettlementTime:           "settlement_time",
	SettlementOrderId:        "settlement_order_id",
	SettlementType:           "settlement_type",
	SettlementCycle:          "settlement_cycle",
	MemberMessage:            "member_message",
	MemberMessageJa:          "member_message_ja",
	RestaurantMessage:        "restaurant_message",
	BookingCount:             "booking_count",
	RefundStatus:             "refund_status",
	RefundRate:               "refund_rate",
	RefundFee:                "refund_fee",
	RefundTime:               "refund_time",
	RefundAmount:             "refund_amount",
	RefundBalAmount:          "refund_bal_amount",
	RefundCouponAmount:       "refund_coupon_amount",
	AdminRefundAmount:        "admin_refund_amount",
	AdminRefundBalAmount:     "admin_refund_bal_amount",
	AdminRefundCouponAmount:  "admin_refund_coupon_amount",
	AdminCancelReason:        "admin_cancel_reason",
	ActivityId:               "activity_id",
	AdminCancelNum:           "admin_cancel_num",
	Referrer:                 "referrer",
	RebateRate:               "rebate_rate",
	RebateStatus:             "rebate_status",
	RebateAmount:             "rebate_amount",
	RebateTime:               "rebate_time",
	IsGetOpen:                "is_get_open",
	IsPayOpen:                "is_pay_open",
	FoodGetRateVip:           "food_get_rate_vip",
	FoodGetRateScene:         "food_get_rate_scene",
	FoodGetScoreStatus:       "food_get_score_status",
	FoodGetAmount:            "food_get_amount",
	ExpValue:                 "exp_value",
	ExpTime:                  "exp_time",
	DepositRate:              "deposit_rate",
	DepositAmount:            "deposit_amount",
	DepositPayStatus:         "deposit_pay_status",
	DepositPayTime:           "deposit_pay_time",
	DepositCancelTime:        "deposit_cancel_time",
	DepositRefundTime:        "deposit_refund_time",
	DepositExpirationTime:    "deposit_expiration_time",
	RemainPayStatus:          "remain_pay_status",
	RemainCancelTime:         "remain_cancel_time",
	RemainRefundTime:         "remain_refund_time",
	DepositCancelSource:      "deposit_cancel_source",
	DepositCancelReason:      "deposit_cancel_reason",
	RemainCancelSource:       "remain_cancel_source",
	RemainCancelReason:       "remain_cancel_reason",
	PayStep:                  "pay_step",
	OldBookDate:              "old_book_date",
	OldBookTime:              "old_book_time",
	OldBookDatetime:          "old_book_datetime",
	OldBookingCount:          "old_booking_count",
	ToretaReservationNo:      "toreta_reservation_no",
	ToretaReservationId:      "toreta_reservation_id",
	ToretaReservationStatus:  "toreta_reservation_status",
	ToretaReservationEndtime: "toreta_reservation_endtime",
	GoodsIsNoPay:             "goods_is_no_pay",
	ToretaHasTimeLimit:       "toreta_has_time_limit",
	ToretaEndTime:            "toreta_end_time",
	CreatedAt:                "created_at",
	UpdatedAt:                "updated_at",
	IsFx:                     "is_fx",
}

// NewFoodOrderDao creates and returns a new DAO object for table data access.
func NewFoodOrderDao() *FoodOrderDao {
	return &FoodOrderDao{
		group:   "default",
		table:   "hg_food_order",
		columns: foodOrderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FoodOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FoodOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FoodOrderDao) Columns() FoodOrderColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FoodOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FoodOrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FoodOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
