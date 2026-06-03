package app

import (
	"APT/internal/model/input/input_homepage_article"
	"APT/internal/service"
	"context"

	"APT/api/app/basics"
)

func (c *ControllerBasics) HomepageArticleList(ctx context.Context, req *basics.HomepageArticleListReq) (res *basics.HomepageArticleListRes, err error) {
	res = new(basics.HomepageArticleListRes)
	if res.List, res.Count, err = service.BasicsHomepageArticles().AppList(ctx, &input_homepage_article.HomepageArticlesAppListInp{
		PageReq: req.PageReq,
	}); err != nil {
		return
	}
	return
}
func (c *ControllerBasics) HomepageArticleView(ctx context.Context, req *basics.HomepageArticleViewReq) (res *basics.HomepageArticleViewRes, err error) {
	res = new(basics.HomepageArticleViewRes)
	if res.HomepageArticlesAppViewModel, err = service.BasicsHomepageArticles().AppView(ctx, &req.HomepageArticlesAppViewInp); err != nil {
		return
	}
	return
}

func (c *ControllerBasics) HomepageArticleThumb(ctx context.Context, req *basics.HomepageArticleThumbReq) (res *basics.HomepageArticleThumbRes, err error) {
	err = service.BasicsHomepageArticles().Thumb(ctx, &req.HomepageArticlesThumbInp)
	return
}
