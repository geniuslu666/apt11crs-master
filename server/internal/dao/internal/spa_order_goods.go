// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SpaOrderGoodsDao is the data access object for the table hg_spa_order_goods.
type SpaOrderGoodsDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns SpaOrderGoodsColumns // columns contains all the column names of Table for convenient usage.
}

// SpaOrderGoodsColumns defines and stores column names for the table hg_spa_order_goods.
type SpaOrderGoodsColumns struct {
	Id                 string //
	IspId              string // 服务商ID
	OrderId            string // 订单ID
	ServiceType        string // 1到店  2上门
	ServiceId          string // 服务ID
	TechnicianGoTime   string // 技师上门出发时间
	ActualStartTime    string // 实际开始日期时间
	ActualEndTime      string // 实际结束日期时间
	GoodsId            string // 项目ID
	GoodsNum           string // 项目数量
	OrderGoodsAmount   string // 子订单金额
	CouponAmount       string // 子订单优惠券抵扣金额
	BalAmount          string // 子订单积分抵扣金额
	TechnicianIds      string // 技师ID ,分隔
	PayStatus          string // 订单付款状态
	StoreId            string // 到店门店ID
	PropertyId         string // 上门物业ID
	RoomNo             string // 上门房间号
	OrderStatus        string // 订单状态
	DispatchStatus     string // 订单调度状态
	DispatchTime       string // 订单调度时间
	DispatchDesc       string // 订单调度备注
	DispatchOperatorId string // 调度操作人ID
	StartServeImages   string // 开始服务图集
	EndServeImages     string // 服务结束图集
	RefundStatus       string // 退款状态   WAIT 未退款   PART  部分退款 DONE 全部退款
	RefundFee          string // 退款手续费
	RefundTime         string // 退款时间
	RefundAmount       string // 已退款总金额
	RefundBalAmount    string // 已退款积分
	RefundCouponAmount string // 已退款优惠券
	CreatedAt          string // 创建时间
	UpdatedAt          string // 更新时间
}

// spaOrderGoodsColumns holds the columns for the table hg_spa_order_goods.
var spaOrderGoodsColumns = SpaOrderGoodsColumns{
	Id:                 "id",
	IspId:              "isp_id",
	OrderId:            "order_id",
	ServiceType:        "service_type",
	ServiceId:          "service_id",
	TechnicianGoTime:   "technician_go_time",
	ActualStartTime:    "actual_start_time",
	ActualEndTime:      "actual_end_time",
	GoodsId:            "goods_id",
	GoodsNum:           "goods_num",
	OrderGoodsAmount:   "order_goods_amount",
	CouponAmount:       "coupon_amount",
	BalAmount:          "bal_amount",
	TechnicianIds:      "technician_ids",
	PayStatus:          "pay_status",
	StoreId:            "store_id",
	PropertyId:         "property_id",
	RoomNo:             "room_no",
	OrderStatus:        "order_status",
	DispatchStatus:     "dispatch_status",
	DispatchTime:       "dispatch_time",
	DispatchDesc:       "dispatch_desc",
	DispatchOperatorId: "dispatch_operator_id",
	StartServeImages:   "start_serve_images",
	EndServeImages:     "end_serve_images",
	RefundStatus:       "refund_status",
	RefundFee:          "refund_fee",
	RefundTime:         "refund_time",
	RefundAmount:       "refund_amount",
	RefundBalAmount:    "refund_bal_amount",
	RefundCouponAmount: "refund_coupon_amount",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}

// NewSpaOrderGoodsDao creates and returns a new DAO object for table data access.
func NewSpaOrderGoodsDao() *SpaOrderGoodsDao {
	return &SpaOrderGoodsDao{
		group:   "default",
		table:   "hg_spa_order_goods",
		columns: spaOrderGoodsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SpaOrderGoodsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SpaOrderGoodsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SpaOrderGoodsDao) Columns() SpaOrderGoodsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SpaOrderGoodsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SpaOrderGoodsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SpaOrderGoodsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
