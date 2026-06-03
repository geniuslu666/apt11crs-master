package pms

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"

	"github.com/gogf/gf/v2/frame/g"
)

type AppConfigListReq struct {
	g.Meta `path:"/pmsAppconfig/list" method:"get" tags:"ADMIN_PMS" summary:"APP配置信息_列表"`
	input_basics.PmsAppconfigListInp
}

type AppConfigListRes struct {
	input_form.PageRes
	List []*input_basics.PmsAppconfigListModel `json:"list"   dc:"数据列表"`
}

type AppConfigViewReq struct {
	g.Meta `path:"/pmsAppconfig/view" method:"get" tags:"ADMIN_PMS" summary:"APP配置信息_详情"`
	input_basics.PmsAppconfigViewInp
}

type AppConfigViewRes struct {
	*input_basics.PmsAppconfigViewModel
}

type AppConfigEditReq struct {
	g.Meta `path:"/pmsAppconfig/edit" method:"post" tags:"ADMIN_PMS" summary:"APP配置信息_编辑"`
	input_basics.PmsAppconfigEditInp
}

type AppConfigEditRes struct{}

// AppConfigDeleteReq 删除APP配置
type AppConfigDeleteReq struct {
	g.Meta `path:"/pmsAppconfig/delete" method:"post" tags:"ADMIN_PMS" summary:"APP配置信息_删除"`
	input_basics.PmsAppconfigDeleteInp
}

type AppConfigDeleteRes struct{}
