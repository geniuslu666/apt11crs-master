package admin

import (
	"APT/internal/model/input/input_homepage_article"
	"APT/internal/service"
	"context"

	"APT/api/admin/basics"
)

func (c *ControllerBasics) ArticleList(ctx context.Context, req *basics.ArticleListReq) (res *basics.ArticleListRes, err error) {
	list, totalCount, err := service.BasicsHomepageArticles().List(ctx, &req.HomepageArticlesListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_homepage_article.HomepageArticlesListModel{}
	}

	res = new(basics.ArticleListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerBasics) ArticleView(ctx context.Context, req *basics.ArticleViewReq) (res *basics.ArticleViewRes, err error) {
	data, err := service.BasicsHomepageArticles().View(ctx, &req.HomepageArticlesViewInp)
	if err != nil {
		return
	}

	res = new(basics.ArticleViewRes)
	res.HomepageArticlesViewModel = data
	return
}
func (c *ControllerBasics) ArticleEdit(ctx context.Context, req *basics.ArticleEditReq) (res *basics.ArticleEditRes, err error) {
	err = service.BasicsHomepageArticles().Edit(ctx, &req.HomepageArticlesEditInp)
	return
}
func (c *ControllerBasics) ArticleDelete(ctx context.Context, req *basics.ArticleDeleteReq) (res *basics.ArticleDeleteRes, err error) {
	err = service.BasicsHomepageArticles().Delete(ctx, &req.HomepageArticlesDeleteInp)
	return
}
func (c *ControllerBasics) ArticleStatus(ctx context.Context, req *basics.ArticleStatusReq) (res *basics.ArticleStatusRes, err error) {
	err = service.BasicsHomepageArticles().Status(ctx, &req.HomepageArticlesStatusInp)
	return
}
