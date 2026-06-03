// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CarOrderDao is the data access object for the table hg_car_order.
type CarOrderDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of the current DAO.
	columns CarOrderColumns // columns contains all the column names of Table for convenient usage.
}

// CarOrderColumns defines and stores column names for the table hg_car_order.
type CarOrderColumns struct {
	Id                      string //
	OrderType               string // 订单类型
	OrderSn                 string // 订单编号
	OutOrderSn              string // 三方订单号
	ServiceType             string // 服务类型
	ConfirmType             string // 1手动确认  2自动确认
	DispatchType            string // 1手动派单  2自动派单
	MemberId                string // 用户ID
	BookDate                string // 预定日期
	BookTime                string // 预定时间
	AdultNum                string // 成人数
	ChildNum                string // 儿童数
	BookStartTime           string // 预定开始日期时间
	BookEndTime             string // 预定结束日期时间
	DriverGoTime            string // 司机出发时间
	ActualStartTime         string // 实际开始日期时间
	ActualEndTime           string // 实际结束日期时间
	BookingName             string // 预定人姓名
	PhoneArea               string // 手机区号
	BookingMobile           string // 预定人手机
	BookingEmail            string // 预定人邮箱
	StartAddressId          string // 出发地ID
	EndAddressId            string // 目的地ID
	ServiceId               string // 服务ID
	DriverId                string // 司机ID
	CarId                   string // 车辆ID
	IsReturn                string // 是否退单中 1是  2否
	ReturnDriverId          string // 退单司机ID
	ReturnCarId             string // 退单车辆ID
	PickUpSign              string // 是否选择举牌接机
	PickUpSignAmount        string // 举牌接机价格
	ChildSeatAddNum         string // 婴儿座椅数量
	ChildSeatAddAmount      string // 婴儿座椅总价
	OrderAmount             string // 订单金额
	CouponAmount            string // 优惠券抵扣金额
	NightAmount             string // 深夜费
	BalAmount               string // 积分抵扣金额
	PayModel                string // 1、余额支付 2、组合支付 3、纯外部支付
	PayTime                 string // 支付时间
	PayStatus               string // 订单付款状态
	FlightNumber            string // 航班号
	OrderStatus             string // 订单状态
	ConfirmTime             string // 订单确认时间
	ConfirmRefuseReason     string // 审核拒绝原因
	DispatchStatus          string // 订单调度状态
	DispatchTime            string // 订单调度时间
	DispatchDesc            string // 订单调度备注
	DispatchOperatorId      string // 调度操作人ID
	ExpirationTime          string // 订单过期时间
	CancelTime              string // 取消时间
	DriverLanguage          string // 司机语言
	EmergencyName           string // 紧急联系人
	EmergencyPhoneArea      string // 紧急联系人手机区号
	EmergencyMobile         string // 紧急联系人手机
	StartServeImages        string // 开始服务图集
	EndServeImages          string // 服务结束图集
	SettlementRate          string // 结算比例
	SettlementStatus        string // WAIT 等待结算   SUCCESS  结算处理成功    FAIL  结算处理失败
	SettlementAmount        string // 结算金额
	SettlementTime          string // 结算时间
	SettlementOrderId       string // 结算单ID
	SettlementType          string // 结算方式  1无需结算 2按周期自动结算  3手动申请结算
	SettlementCycle         string // 结算周期  1每日结算  2每周结算  3每月结算
	MemberMessage           string // 购买人留言信息
	MemberMessageJa         string // 购买人留言信息日语
	RefundStatus            string // 退款状态   WAIT 未退款   PART  部分退款 DONE 全部退款
	RefundFee               string // 退款手续费
	RefundTime              string // 退款时间
	RefundAmount            string // 已退款总金额
	RefundBalAmount         string // 已退款积分
	RefundCouponAmount      string // 已退款优惠券
	RefundReason            string // 退款原因
	Referrer                string // 推荐人
	RebateRate              string // 分佣比例
	RebateStatus            string // WAIT 等待处理佣金   SUCCESS   佣金处理成功    FAIL  佣金处理失败
	RebateAmount            string // 分佣结算金额
	RebateTime              string // 分佣结算时间
	IsGetOpen               string // 是否开启积分获取
	IsPayOpen               string // 是否开启积分抵扣
	CarGetRateVip           string // 接送机结算积分比例
	CarGetRateScene         string // 场景结算积分比例
	CarGetScoreStatus       string // 'WAIT','SUCCESS','FAIL'
	CarGetAmount            string // 结算积分金额
	InnnLuggageFreeNum      string // innn免费行李数
	InnnCarPreAmount        string // innn车型单价
	InnnCarTotalAmount      string // innn车型总价
	InnnLuggagePreAmount    string // innn行李额单价
	InnnLuggageNum          string // innn额外行李数
	InnnLuggageTotalAmount  string // innn行李额总价
	InnnOrderId             string // innn订单ID
	InnnOrderNo             string // innn订单号
	InnnLuggageOrderId      string // innn行李订单ID
	InnnLuggageOrderNo      string // innn行李订单号
	InnnQrcode              string // innn二维码内容
	InnnQrcodeExpireTime    string // innn二维码过期时间戳
	InnnOrderCancelSuccess  string // innn取消接口是否请求成功  0无请求   1请求成功  2请求失败
	InnnCancelRule          string // innn取消政策
	InnnVerifyStatus        string // innn核销状态
	ExpValue                string // 结算的经验值
	ExpTime                 string // 经验结算时间
	AdminRefundAmount       string // 后台已退款总金额
	AdminRefundBalAmount    string // 后台已退款积分
	AdminRefundCouponAmount string // 后台已退款优惠券
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

// carOrderColumns holds the columns for the table hg_car_order.
var carOrderColumns = CarOrderColumns{
	Id:                      "id",
	OrderType:               "order_type",
	OrderSn:                 "order_sn",
	OutOrderSn:              "out_order_sn",
	ServiceType:             "service_type",
	ConfirmType:             "confirm_type",
	DispatchType:            "dispatch_type",
	MemberId:                "member_id",
	BookDate:                "book_date",
	BookTime:                "book_time",
	AdultNum:                "adult_num",
	ChildNum:                "child_num",
	BookStartTime:           "book_start_time",
	BookEndTime:             "book_end_time",
	DriverGoTime:            "driver_go_time",
	ActualStartTime:         "actual_start_time",
	ActualEndTime:           "actual_end_time",
	BookingName:             "booking_name",
	PhoneArea:               "phone_area",
	BookingMobile:           "booking_mobile",
	BookingEmail:            "booking_email",
	StartAddressId:          "start_address_id",
	EndAddressId:            "end_address_id",
	ServiceId:               "service_id",
	DriverId:                "driver_id",
	CarId:                   "car_id",
	IsReturn:                "is_return",
	ReturnDriverId:          "return_driver_id",
	ReturnCarId:             "return_car_id",
	PickUpSign:              "pick_up_sign",
	PickUpSignAmount:        "pick_up_sign_amount",
	ChildSeatAddNum:         "child_seat_add_num",
	ChildSeatAddAmount:      "child_seat_add_amount",
	OrderAmount:             "order_amount",
	CouponAmount:            "coupon_amount",
	NightAmount:             "night_amount",
	BalAmount:               "bal_amount",
	PayModel:                "pay_model",
	PayTime:                 "pay_time",
	PayStatus:               "pay_status",
	FlightNumber:            "flight_number",
	OrderStatus:             "order_status",
	ConfirmTime:             "confirm_time",
	ConfirmRefuseReason:     "confirm_refuse_reason",
	DispatchStatus:          "dispatch_status",
	DispatchTime:            "dispatch_time",
	DispatchDesc:            "dispatch_desc",
	DispatchOperatorId:      "dispatch_operator_id",
	ExpirationTime:          "expiration_time",
	CancelTime:              "cancel_time",
	DriverLanguage:          "driver_language",
	EmergencyName:           "emergency_name",
	EmergencyPhoneArea:      "emergency_phone_area",
	EmergencyMobile:         "emergency_mobile",
	StartServeImages:        "start_serve_images",
	EndServeImages:          "end_serve_images",
	SettlementRate:          "settlement_rate",
	SettlementStatus:        "settlement_status",
	SettlementAmount:        "settlement_amount",
	SettlementTime:          "settlement_time",
	SettlementOrderId:       "settlement_order_id",
	SettlementType:          "settlement_type",
	SettlementCycle:         "settlement_cycle",
	MemberMessage:           "member_message",
	MemberMessageJa:         "member_message_ja",
	RefundStatus:            "refund_status",
	RefundFee:               "refund_fee",
	RefundTime:              "refund_time",
	RefundAmount:            "refund_amount",
	RefundBalAmount:         "refund_bal_amount",
	RefundCouponAmount:      "refund_coupon_amount",
	RefundReason:            "refund_reason",
	Referrer:                "referrer",
	RebateRate:              "rebate_rate",
	RebateStatus:            "rebate_status",
	RebateAmount:            "rebate_amount",
	RebateTime:              "rebate_time",
	IsGetOpen:               "is_get_open",
	IsPayOpen:               "is_pay_open",
	CarGetRateVip:           "car_get_rate_vip",
	CarGetRateScene:         "car_get_rate_scene",
	CarGetScoreStatus:       "car_get_score_status",
	CarGetAmount:            "car_get_amount",
	InnnLuggageFreeNum:      "innn_luggage_free_num",
	InnnCarPreAmount:        "innn_car_pre_amount",
	InnnCarTotalAmount:      "innn_car_total_amount",
	InnnLuggagePreAmount:    "innn_luggage_pre_amount",
	InnnLuggageNum:          "innn_luggage_num",
	InnnLuggageTotalAmount:  "innn_luggage_total_amount",
	InnnOrderId:             "innn_order_id",
	InnnOrderNo:             "innn_order_no",
	InnnLuggageOrderId:      "innn_luggage_order_id",
	InnnLuggageOrderNo:      "innn_luggage_order_no",
	InnnQrcode:              "innn_qrcode",
	InnnQrcodeExpireTime:    "innn_qrcode_expire_time",
	InnnOrderCancelSuccess:  "innn_order_cancel_success",
	InnnCancelRule:          "innn_cancel_rule",
	InnnVerifyStatus:        "innn_verify_status",
	ExpValue:                "exp_value",
	ExpTime:                 "exp_time",
	AdminRefundAmount:       "admin_refund_amount",
	AdminRefundBalAmount:    "admin_refund_bal_amount",
	AdminRefundCouponAmount: "admin_refund_coupon_amount",
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

// NewCarOrderDao creates and returns a new DAO object for table data access.
func NewCarOrderDao() *CarOrderDao {
	return &CarOrderDao{
		group:   "default",
		table:   "hg_car_order",
		columns: carOrderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CarOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CarOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CarOrderDao) Columns() CarOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CarOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CarOrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CarOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
