// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CabinetOrderDao is the data access object for the table hg_cabinet_order.
type CabinetOrderDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of the current DAO.
	columns CabinetOrderColumns // columns contains all the column names of Table for convenient usage.
}

// CabinetOrderColumns defines and stores column names for the table hg_cabinet_order.
type CabinetOrderColumns struct {
	Id                      string //
	OrderSn                 string // 订单编号
	OutOrderSn              string // 三方订单号
	MinHours                string // 最小小时数
	CabinetId               string // 储物柜ID
	CabinetName             string // 储物柜名称
	CabinetNameJson         string // 储物柜名称多语言JSON
	CityId                  string // 所属城市ID
	CityName                string // 所属城市名
	CityNameJson            string // 所属城市名多语言JSON
	MchId                   string // 运营商ID
	MchName                 string // 运营商名称
	MchNameJson             string // 运营商名称多语言JSON
	MchBranchId             string // 网点ID
	MchBranchName           string // 网点名称
	MchBranchNameJson       string // 网点名称多语言JSON
	MchBranchLat            string // 网点lat
	MchBranchLgt            string // 网点lgt
	Address                 string // 地址
	AddressJson             string // 地址多语言JSON
	BoxTypeJson             string // 格口类型列表数据
	BoxTypeId               string // 格口类型ID
	BoxTypeName             string // 格口类型名称
	BoxTypeNameJson         string // 格口类型名称多语言JSON
	BoxTypePrice            string // 格口类型单价
	BoxId                   string // 格口ID
	BoxNo                   string // 格口编号
	BoxAlias                string // 格口别名
	OrderFirstFeeRate       string // 首次下单优惠
	BuyHours                string // 购买的小时数
	Pin                     string // 取件码(4位数字)
	MemberId                string // 用户ID
	OrderAmount             string // 订单金额
	BaseAmount              string // 租赁时间内金额
	CouponAmount            string // 优惠券抵扣金额
	BalAmount               string // 积分抵扣金额
	OvertimeSecs            string // 超时时长(秒)
	OvertimeHours           string // 超时时长(小时)
	OvertimeFee             string // 超时费用(日元)
	OvertimePayTime         string // 超时费支付时间
	OvertimePayStatus       string // 超时费付款状态
	OvertimeBalAmount       string // 超时费积分抵扣金额
	OvertimePayModel        string // 1、余额支付 2、组合支付 3、纯外部支付
	GraceSeconds            string // 宽限期设置时长（秒）
	GraceEndTime            string // 宽限期结束时间
	PayStep                 string // 支付流程
	PayModel                string // 1、余额支付 2、组合支付 3、纯外部支付
	PayTime                 string // 支付时间
	PayStatus               string // 订单付款状态
	OrderStatus             string // 订单状态
	ExpirationTime          string // 订单过期时间
	CancelTime              string // 取消时间
	StartTime               string // 订单开始时间
	EndTime                 string // 订单结束时间
	FinishTime              string // 完成时间
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
	CabinetGetRateVip       string // 储物柜结算积分比例
	CabinetGetRateScene     string // 场景结算积分比例
	CabinetGetScoreStatus   string // 'WAIT','SUCCESS','FAIL'
	CabinetGetAmount        string // 结算积分金额
	ExpValue                string // 结算的经验值
	ExpTime                 string // 经验结算时间
	IsAbnormal              string // 请求下单接口是否异常
	PayOvertimeAbnormal     string // 请求支付超时费接口是否异常
	IsAdminComplete         string // 是否是后台强制完成
	AdminCompleteOperatorId string // 强制完成处理操作人ID
	CreatedAt               string //
	UpdatedAt               string //
	IsFx                    string // 是否是分销订单
}

// cabinetOrderColumns holds the columns for the table hg_cabinet_order.
var cabinetOrderColumns = CabinetOrderColumns{
	Id:                      "id",
	OrderSn:                 "order_sn",
	OutOrderSn:              "out_order_sn",
	MinHours:                "min_hours",
	CabinetId:               "cabinet_id",
	CabinetName:             "cabinet_name",
	CabinetNameJson:         "cabinet_name_json",
	CityId:                  "city_id",
	CityName:                "city_name",
	CityNameJson:            "city_name_json",
	MchId:                   "mch_id",
	MchName:                 "mch_name",
	MchNameJson:             "mch_name_json",
	MchBranchId:             "mch_branch_id",
	MchBranchName:           "mch_branch_name",
	MchBranchNameJson:       "mch_branch_name_json",
	MchBranchLat:            "mch_branch_lat",
	MchBranchLgt:            "mch_branch_lgt",
	Address:                 "address",
	AddressJson:             "address_json",
	BoxTypeJson:             "box_type_json",
	BoxTypeId:               "box_type_id",
	BoxTypeName:             "box_type_name",
	BoxTypeNameJson:         "box_type_name_json",
	BoxTypePrice:            "box_type_price",
	BoxId:                   "box_id",
	BoxNo:                   "box_no",
	BoxAlias:                "box_alias",
	OrderFirstFeeRate:       "order_first_fee_rate",
	BuyHours:                "buy_hours",
	Pin:                     "pin",
	MemberId:                "member_id",
	OrderAmount:             "order_amount",
	BaseAmount:              "base_amount",
	CouponAmount:            "coupon_amount",
	BalAmount:               "bal_amount",
	OvertimeSecs:            "overtime_secs",
	OvertimeHours:           "overtime_hours",
	OvertimeFee:             "overtime_fee",
	OvertimePayTime:         "overtime_pay_time",
	OvertimePayStatus:       "overtime_pay_status",
	OvertimeBalAmount:       "overtime_bal_amount",
	OvertimePayModel:        "overtime_pay_model",
	GraceSeconds:            "grace_seconds",
	GraceEndTime:            "grace_end_time",
	PayStep:                 "pay_step",
	PayModel:                "pay_model",
	PayTime:                 "pay_time",
	PayStatus:               "pay_status",
	OrderStatus:             "order_status",
	ExpirationTime:          "expiration_time",
	CancelTime:              "cancel_time",
	StartTime:               "start_time",
	EndTime:                 "end_time",
	FinishTime:              "finish_time",
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
	CabinetGetRateVip:       "cabinet_get_rate_vip",
	CabinetGetRateScene:     "cabinet_get_rate_scene",
	CabinetGetScoreStatus:   "cabinet_get_score_status",
	CabinetGetAmount:        "cabinet_get_amount",
	ExpValue:                "exp_value",
	ExpTime:                 "exp_time",
	IsAbnormal:              "is_abnormal",
	PayOvertimeAbnormal:     "pay_overtime_abnormal",
	IsAdminComplete:         "is_admin_complete",
	AdminCompleteOperatorId: "admin_complete_operator_id",
	CreatedAt:               "created_at",
	UpdatedAt:               "updated_at",
	IsFx:                    "is_fx",
}

// NewCabinetOrderDao creates and returns a new DAO object for table data access.
func NewCabinetOrderDao() *CabinetOrderDao {
	return &CabinetOrderDao{
		group:   "default",
		table:   "hg_cabinet_order",
		columns: cabinetOrderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CabinetOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CabinetOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CabinetOrderDao) Columns() CabinetOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CabinetOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CabinetOrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CabinetOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
