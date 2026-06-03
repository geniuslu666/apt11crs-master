package basics

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_homepage_article"
	"github.com/gogf/gf/v2/frame/g"
)

type ArticleListReq struct {
	g.Meta `path:"/article/list" method:"get" tags:"ADMIN_ARTICLE" summary:"首页活动_列表"`
	input_homepage_article.HomepageArticlesListInp
}

type ArticleListRes struct {
	input_form.PageRes
	List []*input_homepage_article.HomepageArticlesListModel `json:"list"   dc:"数据列表"`
}

type ArticleViewReq struct {
	g.Meta `path:"/article/view" method:"get" tags:"ADMIN_ARTICLE" summary:"首页活动_详情"`
	input_homepage_article.HomepageArticlesViewInp
}

type ArticleViewRes struct {
	*input_homepage_article.HomepageArticlesViewModel
}

type ArticleEditReq struct {
	g.Meta `path:"/article/edit" method:"post" tags:"ADMIN_ARTICLE" summary:"首页活动_修改/新增"`
	input_homepage_article.HomepageArticlesEditInp
}

type ArticleEditRes struct{}

type ArticleDeleteReq struct {
	g.Meta `path:"/article/delete" method:"post" tags:"ADMIN_ARTICLE" summary:"首页活动_删除"`
	input_homepage_article.HomepageArticlesDeleteInp
}

type ArticleDeleteRes struct{}

// ArticleStatusReq 更新员工状态
type ArticleStatusReq struct {
	g.Meta `path:"/article/status" method:"post" tags:"ADMIN_ARTICLE" summary:"首页活动_更新状态"`
	input_homepage_article.HomepageArticlesStatusInp
}

type ArticleStatusRes struct{}
