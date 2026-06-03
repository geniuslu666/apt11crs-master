package pms

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type IndexNavListReq struct {
	g.Meta `path:"/pmsIndexNav/list" method:"get" tags:"ADMIN_PMS" summary:"获取导航列表"`
	input_basics.PmsIndexNavListInp
}

type IndexNavListRes struct {
	input_form.PageRes
	List []*input_basics.PmsIndexNavListModel `json:"list"   dc:"数据列表"`
}

type IndexNavAllReq struct {
	g.Meta `path:"/pmsIndexNav/all" method:"get" tags:"ADMIN_PMS" summary:"导航全部列表"`
	input_basics.PmsIndexNavAllInp
}

type IndexNavAllRes struct {
	List []*input_basics.PmsIndexNavAllModel `json:"list"   dc:"数据列表"`
}

type IndexNavViewReq struct {
	g.Meta `path:"/pmsIndexNav/view" method:"get" tags:"ADMIN_PMS" summary:"获取导航指定信息"`
	input_basics.PmsIndexNavViewInp
}

type IndexNavViewRes struct {
	*input_basics.PmsIndexNavViewModel
}

type IndexNavEditReq struct {
	g.Meta `path:"/pmsIndexNav/edit" method:"post" tags:"ADMIN_PMS" summary:"修改/新增导航"`
	input_basics.PmsIndexNavEditInp
}

type IndexNavEditRes struct{}

type IndexNavDeleteReq struct {
	g.Meta `path:"/pmsIndexNav/delete" method:"post" tags:"ADMIN_PMS" summary:"删除导航"`
	input_basics.PmsIndexNavDeleteInp
}

type IndexNavDeleteRes struct{}

type IndexNavStatusReq struct {
	g.Meta `path:"/pmsIndexNav/status" method:"post" tags:"ADMIN_PMS" summary:"更新导航状态"`
	input_basics.PmsIndexNavStatusInp
}

type IndexNavStatusRes struct{}

type IndexNavSwitchReq struct {
	g.Meta `path:"/pmsIndexNav/switch" method:"post" tags:"ADMIN_PMS" summary:"导航_更新状态"`
	input_basics.PmsIndexNavSwitchInp
}

type IndexNavSwitchRes struct{}

// IndexNavSortReq 排序
type IndexNavSortReq struct {
	g.Meta `path:"/pmsIndexNav/sortUpdate" method:"post" tags:"ADMIN_PMS" summary:"导航_排序"`
	input_basics.PmsIndexNavSortInp
}

type IndexNavSortRes struct{}

type IndexNavMinappStatusReq struct {
	g.Meta `path:"/pmsIndexNav/minappStatus" method:"post" tags:"ADMIN_PMS" summary:"导航_更新状态"`
	input_basics.PmsIndexNavSwitchInp
}

type IndexNavMinappStatusRes struct{}
