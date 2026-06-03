package pms

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"

	"github.com/gogf/gf/v2/frame/g"
)

type TestNavListReq struct {
	g.Meta `path:"/pmsTestNav/list" method:"get" tags:"ADMIN_PMS" summary:"获取导航列表"`
	input_basics.PmsTestNavListInp
}

type TestNavListRes struct {
	input_form.PageRes
	List []*input_basics.PmsTestNavListModel `json:"list"   dc:"数据列表"`
}

type TestNavAllReq struct {
	g.Meta `path:"/pmsTestNav/all" method:"get" tags:"ADMIN_PMS" summary:"导航全部列表"`
	input_basics.PmsTestNavAllInp
}

type TestNavAllRes struct {
	List []*input_basics.PmsTestNavAllModel `json:"list"   dc:"数据列表"`
}

type TestNavViewReq struct {
	g.Meta `path:"/pmsTestNav/view" method:"get" tags:"ADMIN_PMS" summary:"获取导航指定信息"`
	input_basics.PmsTestNavViewInp
}

type TestNavViewRes struct {
	*input_basics.PmsTestNavViewModel
}

type TestNavEditReq struct {
	g.Meta `path:"/pmsTestNav/edit" method:"post" tags:"ADMIN_PMS" summary:"修改/新增导航"`
	input_basics.PmsTestNavEditInp
}

type TestNavEditRes struct{}

type TestNavDeleteReq struct {
	g.Meta `path:"/pmsTestNav/delete" method:"post" tags:"ADMIN_PMS" summary:"删除导航"`
	input_basics.PmsTestNavDeleteInp
}

type TestNavDeleteRes struct{}

type TestNavStatusReq struct {
	g.Meta `path:"/pmsTestNav/status" method:"post" tags:"ADMIN_PMS" summary:"更新导航状态"`
	input_basics.PmsTestNavStatusInp
}

type TestNavStatusRes struct{}

type TestNavSwitchReq struct {
	g.Meta `path:"/pmsTestNav/switch" method:"post" tags:"ADMIN_PMS" summary:"导航_更新状态"`
	input_basics.PmsTestNavSwitchInp
}

type TestNavSwitchRes struct{}

// TestNavSortReq 排序
type TestNavSortReq struct {
	g.Meta `path:"/pmsTestNav/sortUpdate" method:"post" tags:"ADMIN_PMS" summary:"导航_排序"`
	input_basics.PmsTestNavSortInp
}

type TestNavSortRes struct{}
