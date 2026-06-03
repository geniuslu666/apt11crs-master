// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FoodGoodsDao is the data access object for table hg_food_goods.
type FoodGoodsDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns FoodGoodsColumns // columns contains all the column names of Table for convenient usage.
}

// FoodGoodsColumns defines and stores column names for table hg_food_goods.
type FoodGoodsColumns struct {
	Id                  string //
	RestaurantId        string // 餐厅ID
	ToretaCourseId      string // Toreta的课程ID
	GoodsName           string // 套餐名称
	Introduction        string // 促销语
	LabelIds            string // 标签（多选）
	Images              string // 图集
	Price               string // 套餐售价
	MarketPrice         string // 套餐原价
	MaxOrderNum         string // 单次最大可预定数
	TimeDuration        string // 用餐停留时间
	GoodsContent        string // 套餐详情
	Notice              string // 注意事项
	GoodsState          string // 状态（1.正常2下架）
	TotalOrderNum       string // 预约单购买套餐的总数量（包含退款）
	TotalOrderAmount    string // 预约单总金额（包含退款）
	PayOrderNum         string // 预约单支付购买套餐的数量（不包含退款）
	PayOrderAmount      string // 预约单支付金额（不包含退款）
	IsThCouponExclusive string // 是否是礼品券兑换专属：0-否，1-是
	ThCouponId          string // 绑定的礼品券ID，用于礼品券兑换专属商品
	IsNoPay             string // 是否无需支付：0-否，1-是
	MaxTimeOrderOpen    string // 是否开启时间限制 1-开始  2-关闭
	OrderTimeForm       string // 限制时段
	CreateAt            string // 创建时间
	UpdateAt            string // 更新时间
	DeletedAt           string //
}

// foodGoodsColumns holds the columns for table hg_food_goods.
var foodGoodsColumns = FoodGoodsColumns{
	Id:                  "id",
	RestaurantId:        "restaurant_id",
	ToretaCourseId:      "toreta_course_id",
	GoodsName:           "goods_name",
	Introduction:        "introduction",
	LabelIds:            "label_ids",
	Images:              "images",
	Price:               "price",
	MarketPrice:         "market_price",
	MaxOrderNum:         "max_order_num",
	TimeDuration:        "time_duration",
	GoodsContent:        "goods_content",
	Notice:              "notice",
	GoodsState:          "goods_state",
	TotalOrderNum:       "total_order_num",
	TotalOrderAmount:    "total_order_amount",
	PayOrderNum:         "pay_order_num",
	PayOrderAmount:      "pay_order_amount",
	IsThCouponExclusive: "is_th_coupon_exclusive",
	ThCouponId:          "th_coupon_id",
	IsNoPay:             "is_no_pay",
	MaxTimeOrderOpen:    "max_time_order_open",
	OrderTimeForm:       "order_time_form",
	CreateAt:            "create_at",
	UpdateAt:            "update_at",
	DeletedAt:           "deleted_at",
}

// NewFoodGoodsDao creates and returns a new DAO object for table data access.
func NewFoodGoodsDao() *FoodGoodsDao {
	return &FoodGoodsDao{
		group:   "default",
		table:   "hg_food_goods",
		columns: foodGoodsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FoodGoodsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FoodGoodsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FoodGoodsDao) Columns() FoodGoodsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FoodGoodsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FoodGoodsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FoodGoodsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
