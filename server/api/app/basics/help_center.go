package basics

import (
	"APT/internal/model/input/input_basics"
	"github.com/gogf/gf/v2/frame/g"
)

type HelpCenterCategoryListReq struct {
	g.Meta `path:"/pmsHelpcenterCategory/list" method:"post" tags:"APP_BASICS" summary:"[帮助中心]获取分类列表"`
}

type HelpCenterCategoryListRes struct {
	List []*input_basics.PmsHelpcenterCategoryListModel `json:"list"   dc:"数据列表"`
}

type HelpCenterListReq struct {
	g.Meta     `path:"/pmsHelpcenter/list" method:"post" tags:"APP_BASICS" summary:"[帮助中心]获取列表"`
	CategoryId int    `json:"categoryId" dc:"分类ID"`
	Keyword    string `json:"keyword" dc:"关键字"`
	Title      string `json:"title" dc:"标题"`
}

type HelpCenterListRes struct {
	List []*input_basics.PmsHelpcenterListModel `json:"list"   dc:"数据列表"`
}

type HelpCenterViewReq struct {
	g.Meta `path:"/pmsHelpcenter/view" method:"post" tags:"APP_BASICS" summary:"[帮助中心]获取详情"`
	input_basics.PmsHelpcenterViewInp
}

type HelpCenterViewRes struct {
	*input_basics.PmsHelpcenterViewModel
}
