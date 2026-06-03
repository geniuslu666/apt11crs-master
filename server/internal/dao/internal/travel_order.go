// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TravelOrderDao is the data access object for table hg_travel_order.
type TravelOrderDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns TravelOrderColumns // columns contains all the column names of Table for convenient usage.
}

// TravelOrderColumns defines and stores column names for table hg_travel_order.
type TravelOrderColumns struct {
	Id                      string //
	OrderSn                 string // 预约单号
	ProductId               string // 产品ID
	SkuId                   string // SKUID
	MemberId                string // 会员ID
	BookingName             string // 预订人姓名
	FirstName               string // 订单预定人名
	LastName                string // 订单预定人姓
	PhoneArea               string // 手机区号
	BookingMobile           string // 预订人电话
	BookingEmail            string // 预定人邮箱
	BookingNum              string // 预约人数
	BookDate                string // 预约日期
	OrderAmount             string // 订单金额（元）
	CouponAmount            string // 优惠券抵扣金额
	BalAmount               string // 积分抵扣金额
	OrderStatus             string // 订单状态
	PayModel                string // 1、余额支付 2、组合支付 3、纯外部支付
	PayStatus               string // 订单付款状态
	PayTime                 string // 支付时间
	VerifyStaffId           string // 核销人员ID
	VerifyTime              string // 核销时间
	ExpirationTime          string // 订单过期时间
	CancelTime              string // 取消时间
	CancelFee               string // 取消手续费（元）
	RefundStatus            string // 退款状态   WAIT 未退款   PART  部分退款 DONE 全部退款
	RefundAmount            string // 退款金额（元）
	RefundTime              string // 退款时间
	RefundBalAmount         string // 已退款积分
	RefundCouponAmount      string // 已退款优惠券
	RefundReason            string // 退款原因
	CreatedAt               string // 创建时间
	UpdatedAt               string // 更新时间
	Referrer                string // 推荐人
	RebateRate              string // 分佣比例
	RebateStatus            string // WAIT 等待处理佣金   SUCCESS   佣金处理成功    FAIL  佣金处理失败
	RebateAmount            string // 分佣结算金额
	RebateTime              string // 分佣结算时间
	IsGetOpen               string // 是否开启积分获取
	IsPayOpen               string // 是否开启积分抵扣
	TravelGetRateVip        string // 结算积分比例
	TravelGetRateScene      string // 场景结算积分比例
	TravelGetScoreStatus    string // 'WAIT','SUCCESS','FAIL'
	TravelGetAmount         string // 结算积分金额
	ExpValue                string // 结算的经验值
	ExpTime                 string // 经验结算时间
	AdminRefundAmount       string // 后台已退款总金额
	AdminRefundBalAmount    string // 后台已退款积分
	AdminRefundCouponAmount string // 后台已退款优惠券
	AdminCancelReason       string // 后台取消原因
	AdminCancelNum          string // 后台取消次数
	IsFx                    string // 是否是分销订单
}

// travelOrderColumns holds the columns for table hg_travel_order.
var travelOrderColumns = TravelOrderColumns{
	Id:                      "id",
	OrderSn:                 "order_sn",
	ProductId:               "product_id",
	SkuId:                   "sku_id",
	MemberId:                "member_id",
	BookingName:             "booking_name",
	FirstName:               "first_name",
	LastName:                "last_name",
	PhoneArea:               "phone_area",
	BookingMobile:           "booking_mobile",
	BookingEmail:            "booking_email",
	BookingNum:              "booking_num",
	BookDate:                "book_date",
	OrderAmount:             "order_amount",
	CouponAmount:            "coupon_amount",
	BalAmount:               "bal_amount",
	OrderStatus:             "order_status",
	PayModel:                "pay_model",
	PayStatus:               "pay_status",
	PayTime:                 "pay_time",
	VerifyStaffId:           "verify_staff_id",
	VerifyTime:              "verify_time",
	ExpirationTime:          "expiration_time",
	CancelTime:              "cancel_time",
	CancelFee:               "cancel_fee",
	RefundStatus:            "refund_status",
	RefundAmount:            "refund_amount",
	RefundTime:              "refund_time",
	RefundBalAmount:         "refund_bal_amount",
	RefundCouponAmount:      "refund_coupon_amount",
	RefundReason:            "refund_reason",
	CreatedAt:               "created_at",
	UpdatedAt:               "updated_at",
	Referrer:                "referrer",
	RebateRate:              "rebate_rate",
	RebateStatus:            "rebate_status",
	RebateAmount:            "rebate_amount",
	RebateTime:              "rebate_time",
	IsGetOpen:               "is_get_open",
	IsPayOpen:               "is_pay_open",
	TravelGetRateVip:        "travel_get_rate_vip",
	TravelGetRateScene:      "travel_get_rate_scene",
	TravelGetScoreStatus:    "travel_get_score_status",
	TravelGetAmount:         "travel_get_amount",
	ExpValue:                "exp_value",
	ExpTime:                 "exp_time",
	AdminRefundAmount:       "admin_refund_amount",
	AdminRefundBalAmount:    "admin_refund_bal_amount",
	AdminRefundCouponAmount: "admin_refund_coupon_amount",
	AdminCancelReason:       "admin_cancel_reason",
	AdminCancelNum:          "admin_cancel_num",
	IsFx:                    "is_fx",
}

// NewTravelOrderDao creates and returns a new DAO object for table data access.
func NewTravelOrderDao() *TravelOrderDao {
	return &TravelOrderDao{
		group:   "default",
		table:   "hg_travel_order",
		columns: travelOrderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TravelOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TravelOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TravelOrderDao) Columns() TravelOrderColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TravelOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TravelOrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TravelOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
