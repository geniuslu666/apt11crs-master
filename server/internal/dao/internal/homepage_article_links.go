// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HomepageArticleLinksDao is the data access object for table hg_homepage_article_links.
type HomepageArticleLinksDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of current DAO.
	columns HomepageArticleLinksColumns // columns contains all the column names of Table for convenient usage.
}

// HomepageArticleLinksColumns defines and stores column names for table hg_homepage_article_links.
type HomepageArticleLinksColumns struct {
	Id           string //
	ArticleId    string // 文章ID
	Title        string // 链接标题
	LinkText     string // 链接显示的文字
	AppLink      string // APP跳转链接
	WxLink       string // 微信跳转链接
	UrlParam     string // APP跳转链接参数，JSON格式存储
	LinkType     string // coupon-优惠券 thCoupon-礼品券，hotelDetail-民宿详情，foodIndex-餐厅首页，foodDetail-餐厅详情，spaIndex-按摩首页，spaDetail-按摩详情，carIndex-接送机首页, outLink-外部链接
	LinkId       string // 对应跳转链接的ID
	LinkOpenType string // 外部链接跳转方式 1-webview 2-浏览器
	CreatedAt    string //
	UpdatedAt    string //
}

// homepageArticleLinksColumns holds the columns for table hg_homepage_article_links.
var homepageArticleLinksColumns = HomepageArticleLinksColumns{
	Id:           "id",
	ArticleId:    "article_id",
	Title:        "title",
	LinkText:     "link_text",
	AppLink:      "app_link",
	WxLink:       "wx_link",
	UrlParam:     "url_param",
	LinkType:     "link_type",
	LinkId:       "link_id",
	LinkOpenType: "link_open_type",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewHomepageArticleLinksDao creates and returns a new DAO object for table data access.
func NewHomepageArticleLinksDao() *HomepageArticleLinksDao {
	return &HomepageArticleLinksDao{
		group:   "default",
		table:   "hg_homepage_article_links",
		columns: homepageArticleLinksColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HomepageArticleLinksDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HomepageArticleLinksDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HomepageArticleLinksDao) Columns() HomepageArticleLinksColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HomepageArticleLinksDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HomepageArticleLinksDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HomepageArticleLinksDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
