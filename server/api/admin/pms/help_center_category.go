package pms

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type HelpCenterCategoryListReq struct {
	g.Meta `path:"/pmsHelpcenterCategory/list" method:"get" tags:"ADMIN_PMS" summary:"帮助中心_分类列表"`
	input_basics.PmsHelpcenterCategoryListInp
}

type HelpCenterCategoryListRes struct {
	input_form.PageRes
	List []*input_basics.PmsHelpcenterCategoryListModel `json:"list"   dc:"数据列表"`
}

type HelpCenterCategoryAllReq struct {
	g.Meta `path:"/pmsHelpcenterCategory/all" method:"get" tags:"ADMIN_PMS" summary:"帮助中心_分类全部列表"`
	input_basics.PmsHelpcenterCategoryAllInp
}

type HelpCenterCategoryAllRes struct {
	List []*input_basics.PmsHelpcenterCategoryAllModel `json:"list"   dc:"数据列表"`
}

type HelpCenterCategoryViewReq struct {
	g.Meta `path:"/pmsHelpcenterCategory/view" method:"get" tags:"ADMIN_PMS" summary:"帮助中心_详情"`
	input_basics.PmsHelpcenterCategoryViewInp
}

type HelpCenterCategoryViewRes struct {
	*input_basics.PmsHelpcenterCategoryViewModel
}

type HelpCenterCategoryEditReq struct {
	g.Meta `path:"/pmsHelpcenterCategory/edit" method:"post" tags:"ADMIN_PMS" summary:"帮助中心_修改/新增"`
	input_basics.PmsHelpcenterCategoryEditInp
}

type HelpCenterCategoryEditRes struct{}

type HelpCenterCategoryDeleteReq struct {
	g.Meta `path:"/pmsHelpcenterCategory/delete" method:"post" tags:"ADMIN_PMS" summary:"帮助中心_删除"`
	input_basics.PmsHelpcenterCategoryDeleteInp
}

type HelpCenterCategoryDeleteRes struct{}
