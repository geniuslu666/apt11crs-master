package pms

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type HelpCenterListReq struct {
	g.Meta `path:"/pmsHelpcenter/list" method:"get" tags:"ADMIN_PMS" summary:"帮助中心_列表"`
	input_basics.PmsHelpcenterListInp
}

type HelpCenterListRes struct {
	input_form.PageRes
	List []*input_basics.PmsHelpcenterListModel `json:"list"   dc:"数据列表"`
}

type HelpCenterViewReq struct {
	g.Meta `path:"/pmsHelpcenter/view" method:"get" tags:"ADMIN_PMS" summary:"帮助中心_详情"`
	input_basics.PmsHelpcenterViewInp
}

type HelpCenterViewRes struct {
	*input_basics.PmsHelpcenterViewModel
}

type HelpCenterEditReq struct {
	g.Meta `path:"/pmsHelpcenter/edit" method:"post" tags:"ADMIN_PMS" summary:"帮助中心_修改/新增"`
	input_basics.PmsHelpcenterEditInp
}

type HelpCenterEditRes struct{}

type HelpCenterDeleteReq struct {
	g.Meta `path:"/pmsHelpcenter/delete" method:"post" tags:"ADMIN_PMS" summary:"帮助中心_删除"`
	input_basics.PmsHelpcenterDeleteInp
}

type HelpCenterDeleteRes struct{}
