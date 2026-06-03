package car

import (
	"APT/internal/model/input/input_car"
	"APT/internal/model/input/input_form"

	"github.com/gogf/gf/v2/frame/g"
)

// HelpListReq 查询Help列表
type HelpListReq struct {
	g.Meta `path:"/carHelp/list" method:"get" tags:"ADMIN_CAR" summary:"获取Help列表"`
	input_car.CarHelpListInp
}

type HelpListRes struct {
	input_form.PageRes
	List []*input_car.CarHelpListModel `json:"list"   dc:"数据列表"`
}

// HelpViewReq 获取Help指定信息
type HelpViewReq struct {
	g.Meta `path:"/carHelp/view" method:"get" tags:"ADMIN_CAR" summary:"获取Help指定信息"`
	input_car.CarHelpViewInp
}

type HelpViewRes struct {
	*input_car.CarHelpViewModel
}

// HelpEditReq 修改/新增Help
type HelpEditReq struct {
	g.Meta `path:"/carHelp/edit" method:"post" tags:"ADMIN_CAR" summary:"修改/新增Help"`
	input_car.CarHelpEditInp
}

type HelpEditRes struct{}

// HelpDeleteReq 删除Help
type HelpDeleteReq struct {
	g.Meta `path:"/carHelp/delete" method:"post" tags:"ADMIN_CAR" summary:"删除Help"`
	input_car.CarHelpDeleteInp
}

type HelpDeleteRes struct{}

// HelpLanguageListReq 获取Help语言
type HelpLanguageListReq struct {
	g.Meta `path:"/carHelp/languageList" method:"get" tags:"ADMIN_CAR" summary:"获取Help语言"`
	input_car.CarHelpLanguageListInp
}

type HelpLanguageListRes struct {
	List []*input_car.CarHelpLanguageListModel `json:"list"   dc:"数据列表"`
}
