// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsAppStayDao is the data access object for table hg_pms_app_stay.
type PmsAppStayDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns PmsAppStayColumns // columns contains all the column names of Table for convenient usage.
}

// PmsAppStayColumns defines and stores column names for table hg_pms_app_stay.
type PmsAppStayColumns struct {
	Id                  string // APP订单主键
	Uid                 string // 我方系统 ID
	Uuid                string // 三方系统 ID
	Puid                string // 物业ID
	Source              string // 订单来源
	MemberId            string // 用户ID
	OrderSn             string // 订单号
	OutOrderSn          string // 三方订单号
	Booker              string // 预定人
	OrderAmount         string // 订单金额
	PayModel            string // 1、余额支付 2、组合支付 3、纯外部支付
	OrderStatus         string // 订单付款状态
	ExpirationTime      string // 订单过期时间
	CancelTime          string // 未支付取消时间
	RefundStatus        string // 退款状态   WAIT 未退款   PART  部分退款 DONE 全部退款
	RefundTime          string // 退款时间
	RefundAmount        string // 已退款金额
	CleanFee            string // 取消费用
	CheckInDate         string // 入住时间
	CheckOutDate        string // 退房时间
	CancelRate          string // 退款政策
	Referrer            string // 推荐人
	RebateRate          string // 分佣比例
	RebateStatus        string // WAIT 等待处理佣金   SUCCESS   佣金处理成功    FAIL  佣金处理失败
	RebateAmount        string // 分佣结算金额
	RebateTime          string // 分佣结算时间
	IsGetOpen           string // 是否开启积分获取
	IsPayOpen           string // 是否开启积分抵扣
	HotelGetRateVip     string // 酒店结算积分比例
	HotelGetRateScene   string // 场景结算积分比例
	HotelGetScoreStatus string // 'WAIT','SUCCESS','FAIL'
	HotelGetAmount      string // 结算积分金额
	CreatedAt           string // 创建时间
	UpdatedAt           string // 更新时间
	PricePercent        string // 全局溢价比例
	TotalAmount         string // 订单总价
	ChangeAmount        string // 优惠金额
	IsFx                string // 是否是分销订单
}

// pmsAppStayColumns holds the columns for table hg_pms_app_stay.
var pmsAppStayColumns = PmsAppStayColumns{
	Id:                  "id",
	Uid:                 "uid",
	Uuid:                "uuid",
	Puid:                "puid",
	Source:              "source",
	MemberId:            "member_id",
	OrderSn:             "order_sn",
	OutOrderSn:          "out_order_sn",
	Booker:              "booker",
	OrderAmount:         "order_amount",
	PayModel:            "pay_model",
	OrderStatus:         "order_status",
	ExpirationTime:      "expiration_time",
	CancelTime:          "cancel_time",
	RefundStatus:        "refund_status",
	RefundTime:          "refund_time",
	RefundAmount:        "refund_amount",
	CleanFee:            "clean_fee",
	CheckInDate:         "check_in_date",
	CheckOutDate:        "check_out_date",
	CancelRate:          "cancel_rate",
	Referrer:            "referrer",
	RebateRate:          "rebate_rate",
	RebateStatus:        "rebate_status",
	RebateAmount:        "rebate_amount",
	RebateTime:          "rebate_time",
	IsGetOpen:           "is_get_open",
	IsPayOpen:           "is_pay_open",
	HotelGetRateVip:     "hotel_get_rate_vip",
	HotelGetRateScene:   "hotel_get_rate_scene",
	HotelGetScoreStatus: "hotel_get_score_status",
	HotelGetAmount:      "hotel_get_amount",
	CreatedAt:           "created_at",
	UpdatedAt:           "updated_at",
	PricePercent:        "price_percent",
	TotalAmount:         "total_amount",
	ChangeAmount:        "change_amount",
	IsFx:                "is_fx",
}

// NewPmsAppStayDao creates and returns a new DAO object for table data access.
func NewPmsAppStayDao() *PmsAppStayDao {
	return &PmsAppStayDao{
		group:   "default",
		table:   "hg_pms_app_stay",
		columns: pmsAppStayColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PmsAppStayDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PmsAppStayDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PmsAppStayDao) Columns() PmsAppStayColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PmsAppStayDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PmsAppStayDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PmsAppStayDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
