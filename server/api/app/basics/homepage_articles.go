package basics

import (
	"APT/internal/model/input/input_homepage_article"

	"github.com/gogf/gf/v2/frame/g"
)

type HomepageArticleListReq struct {
	g.Meta `path:"/article/appList" method:"post" tags:"APP_ARTICLE" summary:"[首页活动]列表"`
	input_homepage_article.HomepageArticlesAppListInp
}

type HomepageArticleListRes struct {
	List  []*input_homepage_article.HomepageArticlesAppListModel `json:"list"   dc:"数据列表"`
	Count int                                                    `json:"count"   dc:"数据总数"`
}

type HomepageArticleThumbReq struct {
	g.Meta `path:"/article/thumb" method:"post" tags:"APP_ARTICLE" summary:"[首页活动]点赞"`
	input_homepage_article.HomepageArticlesThumbInp
}

type HomepageArticleThumbRes struct{}

type HomepageArticleViewReq struct {
	g.Meta `path:"/article/appView" method:"post" tags:"APP_ARTICLE" summary:"[首页活动]详情"`
	input_homepage_article.HomepageArticlesAppViewInp
}

type HomepageArticleViewRes struct {
	*input_homepage_article.HomepageArticlesAppViewModel
}
