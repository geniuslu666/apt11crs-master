// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/input/input_homepage_article"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IBasicsHomepageArticles interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_homepage_article.HomepageArticlesListInp) (list []*input_homepage_article.HomepageArticlesListModel, totalCount int, err error)
		Edit(ctx context.Context, in *input_homepage_article.HomepageArticlesEditInp) (err error)
		Delete(ctx context.Context, in *input_homepage_article.HomepageArticlesDeleteInp) (err error)
		// Status 更新活动状态
		Status(ctx context.Context, in *input_homepage_article.HomepageArticlesStatusInp) (err error)
		View(ctx context.Context, in *input_homepage_article.HomepageArticlesViewInp) (res *input_homepage_article.HomepageArticlesViewModel, err error)
		AppView(ctx context.Context, in *input_homepage_article.HomepageArticlesAppViewInp) (res *input_homepage_article.HomepageArticlesAppViewModel, err error)
		AppList(ctx context.Context, in *input_homepage_article.HomepageArticlesAppListInp) (list []*input_homepage_article.HomepageArticlesAppListModel, totalCount int, err error)
		Thumb(ctx context.Context, in *input_homepage_article.HomepageArticlesThumbInp) (err error)
	}
)

var (
	localBasicsHomepageArticles IBasicsHomepageArticles
)

func BasicsHomepageArticles() IBasicsHomepageArticles {
	if localBasicsHomepageArticles == nil {
		panic("implement not found for interface IBasicsHomepageArticles, forgot register?")
	}
	return localBasicsHomepageArticles
}

func RegisterBasicsHomepageArticles(i IBasicsHomepageArticles) {
	localBasicsHomepageArticles = i
}
