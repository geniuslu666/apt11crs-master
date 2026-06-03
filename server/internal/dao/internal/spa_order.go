// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SpaOrderDao is the data access object for the table hg_spa_order.
type SpaOrderDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of the current DAO.
	columns SpaOrderColumns // columns contains all the column names of Table for convenient usage.
}

// SpaOrderColumns defines and stores column names for the table hg_spa_order.
type SpaOrderColumns struct {
	Id                      string //
	IspId                   string // 服务商ID
	OrderSn                 string // 订单编号
	OutOrderSn              string // 三方订单号
	ServiceType             string // 1到店  2上门
	ConfirmType             string // 1手动确认  2自动确认
	MemberId                string // 用户ID
	BookDate                string // 预定日期
	BookTime                string // 预定时间
	BookStartTime           string // 预定开始日期时间
	BookEndTime             string // 预定结束日期时间
	TechnicianGoTime        string // 技师上门出发时间
	MemberArriveTime        string // 到店时间
	ActualStartTime         string // 实际开始日期时间
	ActualEndTime           string // 实际结束日期时间
	BookingName             string // 预定人姓名
	PhoneArea               string // 手机区号
	BookingMobile           string // 预定人手机
	BookingEmail            string // 预定人邮箱
	ServiceId               string // 服务ID
	OrderAmount             string // 订单金额
	CouponAmount            string // 优惠券抵扣金额
	BalAmount               string // 积分抵扣金额
	GoodsId                 string // 项目ID
	GoodsNum                string // 项目数量
	TechnicianIds           string // 技师ID ,分隔
	IsReturn                string // 是否退单中 1是  2否
	PayModel                string // 1、余额支付 2、组合支付 3、纯外部支付
	PayTime                 string // 支付时间
	PayStatus               string // 订单付款状态
	StoreId                 string // 到店门店ID
	PropertyId              string // 上门物业ID
	RoomNo                  string // 上门房间号
	OrderStatus             string // 订单状态
	ConfirmTime             string // 订单确认时间
	ConfirmRefuseReason     string // 审核拒绝原因
	DispatchStatus          string // 订单调度状态
	DispatchTime            string // 订单调度时间
	DispatchDesc            string // 订单调度备注
	DispatchOperatorId      string // 调度操作人ID
	ExpirationTime          string // 订单过期时间
	CancelTime              string // 取消时间
	StartServeImages        string // 开始服务图集
	EndServeImages          string // 服务结束图集
	SettlementRate          string // 结算比例
	SettlementStatus        string // WAIT 等待结算   SUCCESS  结算处理成功    FAIL  结算处理失败
	SettlementAmount        string // 结算金额
	SettlementTime          string // 结算时间
	SettlementOrderId       string // 结算单ID
	MemberMessage           string // 购买人留言信息
	MemberMessageJa         string // 购买人留言信息日语
	RefundStatus            string // 退款状态   WAIT 未退款   PART  部分退款 DONE 全部退款
	RefundFee               string // 退款手续费
	RefundTime              string // 退款时间
	RefundAmount            string // 已退款总金额
	RefundBalAmount         string // 已退款积分
	RefundCouponAmount      string // 已退款优惠券
	AdminRefundAmount       string // 后台已退款总金额
	AdminRefundBalAmount    string // 后台已退款积分
	AdminRefundCouponAmount string // 后台已退款优惠券
	RefundReason            string // 退款原因
	Referrer                string // 推荐人
	RebateRate              string // 分佣比例
	RebateStatus            string // WAIT 等待处理佣金   SUCCESS   佣金处理成功    FAIL  佣金处理失败
	RebateAmount            string // 分佣结算金额
	RebateTime              string // 分佣结算时间
	IsGetOpen               string // 是否开启积分获取
	IsPayOpen               string // 是否开启积分抵扣
	SpaGetRateVip           string // 按摩结算积分比例
	SpaGetRateScene         string // 场景结算积分比例
	SpaGetScoreStatus       string // 'WAIT','SUCCESS','FAIL'
	SpaGetAmount            string // 结算积分金额
	ExpValue                string // 结算的经验值
	ExpTime                 string // 经验结算时间
	AdminCancelReason       string // 后台取消原因
	AdminCancelNum          string // 后台取消次数
	AbnormalStatus          string // 异常单状态（1：不是异常单  2异常待处理  3异常已处理）
	AbnormalReason          string // 异常处理原因
	AbnormalOperatorId      string // 异常处理操作人ID
	AbnormalTime            string // 异常处理时间
	CreatedAt               string // 创建时间
	UpdatedAt               string // 更新时间
	IsFx                    string // 是否是分销订单
}

// spaOrderColumns holds the columns for the table hg_spa_order.
var spaOrderColumns = SpaOrderColumns{
	Id:                      "id",
	IspId:                   "isp_id",
	OrderSn:                 "order_sn",
	OutOrderSn:              "out_order_sn",
	ServiceType:             "service_type",
	ConfirmType:             "confirm_type",
	MemberId:                "member_id",
	BookDate:                "book_date",
	BookTime:                "book_time",
	BookStartTime:           "book_start_time",
	BookEndTime:             "book_end_time",
	TechnicianGoTime:        "technician_go_time",
	MemberArriveTime:        "member_arrive_time",
	ActualStartTime:         "actual_start_time",
	ActualEndTime:           "actual_end_time",
	BookingName:             "booking_name",
	PhoneArea:               "phone_area",
	BookingMobile:           "booking_mobile",
	BookingEmail:            "booking_email",
	ServiceId:               "service_id",
	OrderAmount:             "order_amount",
	CouponAmount:            "coupon_amount",
	BalAmount:               "bal_amount",
	GoodsId:                 "goods_id",
	GoodsNum:                "goods_num",
	TechnicianIds:           "technician_ids",
	IsReturn:                "is_return",
	PayModel:                "pay_model",
	PayTime:                 "pay_time",
	PayStatus:               "pay_status",
	StoreId:                 "store_id",
	PropertyId:              "property_id",
	RoomNo:                  "room_no",
	OrderStatus:             "order_status",
	ConfirmTime:             "confirm_time",
	ConfirmRefuseReason:     "confirm_refuse_reason",
	DispatchStatus:          "dispatch_status",
	DispatchTime:            "dispatch_time",
	DispatchDesc:            "dispatch_desc",
	DispatchOperatorId:      "dispatch_operator_id",
	ExpirationTime:          "expiration_time",
	CancelTime:              "cancel_time",
	StartServeImages:        "start_serve_images",
	EndServeImages:          "end_serve_images",
	SettlementRate:          "settlement_rate",
	SettlementStatus:        "settlement_status",
	SettlementAmount:        "settlement_amount",
	SettlementTime:          "settlement_time",
	SettlementOrderId:       "settlement_order_id",
	MemberMessage:           "member_message",
	MemberMessageJa:         "member_message_ja",
	RefundStatus:            "refund_status",
	RefundFee:               "refund_fee",
	RefundTime:              "refund_time",
	RefundAmount:            "refund_amount",
	RefundBalAmount:         "refund_bal_amount",
	RefundCouponAmount:      "refund_coupon_amount",
	AdminRefundAmount:       "admin_refund_amount",
	AdminRefundBalAmount:    "admin_refund_bal_amount",
	AdminRefundCouponAmount: "admin_refund_coupon_amount",
	RefundReason:            "refund_reason",
	Referrer:                "referrer",
	RebateRate:              "rebate_rate",
	RebateStatus:            "rebate_status",
	RebateAmount:            "rebate_amount",
	RebateTime:              "rebate_time",
	IsGetOpen:               "is_get_open",
	IsPayOpen:               "is_pay_open",
	SpaGetRateVip:           "spa_get_rate_vip",
	SpaGetRateScene:         "spa_get_rate_scene",
	SpaGetScoreStatus:       "spa_get_score_status",
	SpaGetAmount:            "spa_get_amount",
	ExpValue:                "exp_value",
	ExpTime:                 "exp_time",
	AdminCancelReason:       "admin_cancel_reason",
	AdminCancelNum:          "admin_cancel_num",
	AbnormalStatus:          "abnormal_status",
	AbnormalReason:          "abnormal_reason",
	AbnormalOperatorId:      "abnormal_operator_id",
	AbnormalTime:            "abnormal_time",
	CreatedAt:               "created_at",
	UpdatedAt:               "updated_at",
	IsFx:                    "is_fx",
}

// NewSpaOrderDao creates and returns a new DAO object for table data access.
func NewSpaOrderDao() *SpaOrderDao {
	return &SpaOrderDao{
		group:   "default",
		table:   "hg_spa_order",
		columns: spaOrderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SpaOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SpaOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SpaOrderDao) Columns() SpaOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SpaOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SpaOrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SpaOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
