// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HomepageArticlesDao is the data access object for the table hg_homepage_articles.
type HomepageArticlesDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of the current DAO.
	columns HomepageArticlesColumns // columns contains all the column names of Table for convenient usage.
}

// HomepageArticlesColumns defines and stores column names for the table hg_homepage_articles.
type HomepageArticlesColumns struct {
	Id                    string //
	Language              string //
	No                    string // 活动编号
	ListPic               string // 列表图URL
	Title                 string // 标题
	Author                string // 作者
	Content               string // 内容
	Sort                  string // 排序(越大越靠前)
	Views                 string // 浏览量
	ThumbNum              string // 收藏数量
	RecommendTitle        string // 推荐标题
	RecommendLinkType     string // 推荐链接类型 coupon-优惠券 thCoupon-礼品券，hotelDetail-民宿详情，foodIndex-餐厅首页，foodDetail-餐厅详情，spaIndex-按摩首页，spaDetail-按摩详情，carIndex-接送机首页
	CouponLimit           string // 每人每张券可领取数量(废弃)
	HotelIds              string // 民宿ids
	RestaurantIds         string // 餐厅ids
	SpaServiceIds         string // 按摩ids
	StartTime             string // 开始时间
	EndTime               string // 结束时间
	Status                string // 状态：1-正常 2-禁用
	MinappStatus          string // 是否排除小程序显示 1-不排除 2-排除
	FoodIndexTitle        string // 餐厅首页推荐标题
	SpaIndexTitle         string // 按摩首页推荐标题
	CarIndexTitle         string // 接送机首页推荐标题
	CouponLinkTitle       string // 券链接区域标题
	CreatedAt             string //
	UpdatedAt             string //
	DeletedAt             string //
	CouponLinkSubTitle    string // 券链接区域副标题
	RecommendLinkTitle    string // 推荐链接区域标题
	RecommendLinkSubTitle string // 推荐链接区域副标题
	ButtonTxt             string // 外链按钮文字
	OutLinkTitle          string // 外链推荐标题
	OutLinkContent        string // 外链推荐链接
	Chain                 string // 活动内外链: IN 内链  OUT 外联
	Path                  string // 外链链接内容
	LinkOpenType          string // 外部链接跳转方式 1-webview 2-浏览器
	LimitWeek             string // 不可领取限制：6-是周六不可领， 7是周日不可领，以逗号分割
}

// homepageArticlesColumns holds the columns for the table hg_homepage_articles.
var homepageArticlesColumns = HomepageArticlesColumns{
	Id:                    "id",
	Language:              "language",
	No:                    "no",
	ListPic:               "list_pic",
	Title:                 "title",
	Author:                "author",
	Content:               "content",
	Sort:                  "sort",
	Views:                 "views",
	ThumbNum:              "thumb_num",
	RecommendTitle:        "recommend_title",
	RecommendLinkType:     "recommend_link_type",
	CouponLimit:           "coupon_limit",
	HotelIds:              "hotel_ids",
	RestaurantIds:         "restaurant_ids",
	SpaServiceIds:         "spa_service_ids",
	StartTime:             "start_time",
	EndTime:               "end_time",
	Status:                "status",
	MinappStatus:          "minapp_status",
	FoodIndexTitle:        "food_index_title",
	SpaIndexTitle:         "spa_index_title",
	CarIndexTitle:         "car_index_title",
	CouponLinkTitle:       "coupon_link_title",
	CreatedAt:             "created_at",
	UpdatedAt:             "updated_at",
	DeletedAt:             "deleted_at",
	CouponLinkSubTitle:    "coupon_link_sub_title",
	RecommendLinkTitle:    "recommend_link_title",
	RecommendLinkSubTitle: "recommend_link_sub_title",
	ButtonTxt:             "button_txt",
	OutLinkTitle:          "out_link_title",
	OutLinkContent:        "out_link_content",
	Chain:                 "chain",
	Path:                  "path",
	LinkOpenType:          "link_open_type",
	LimitWeek:             "limit_week",
}

// NewHomepageArticlesDao creates and returns a new DAO object for table data access.
func NewHomepageArticlesDao() *HomepageArticlesDao {
	return &HomepageArticlesDao{
		group:   "default",
		table:   "hg_homepage_articles",
		columns: homepageArticlesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *HomepageArticlesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *HomepageArticlesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *HomepageArticlesDao) Columns() HomepageArticlesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *HomepageArticlesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *HomepageArticlesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *HomepageArticlesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
