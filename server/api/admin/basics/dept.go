package basics

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"APT/utility/tree"
	"github.com/gogf/gf/v2/frame/g"
)

type DeptListReq struct {
	g.Meta `path:"/dept/list" method:"get" tags:"ADMIN" summary:"部门_获取部门列表"`
	input_basics.DeptListInp
}

type DeptListRes *input_basics.DeptListModel

type DeptViewReq struct {
	g.Meta `path:"/dept/view" method:"get" tags:"ADMIN" summary:"部门_获取指定信息"`
	input_basics.DeptViewInp
}

type DeptViewRes struct {
	*input_basics.DeptViewModel
}

type DeptEditReq struct {
	g.Meta `path:"/dept/edit" method:"post" tags:"ADMIN" summary:"部门_修改/新增部门"`
	input_basics.DeptEditInp
}

type DeptEditRes struct{}

type DeptDeleteReq struct {
	g.Meta `path:"/dept/delete" method:"post" tags:"ADMIN" summary:"部门_删除部门"`
	input_basics.DeptDeleteInp
}

type DeptDeleteRes struct{}

type DeptMaxSortReq struct {
	g.Meta `path:"/dept/maxSort" method:"get" tags:"ADMIN" summary:"部门_部门最大排序"`
	input_basics.DeptMaxSortInp
}

type DeptMaxSortRes struct {
	*input_basics.DeptMaxSortModel
}

type DeptOptionReq struct {
	g.Meta `path:"/dept/option" method:"get" tags:"ADMIN" summary:"部门_获取当前登录用户可选的部门选项"`
	input_basics.DeptOptionInp
}

type DeptOptionRes struct {
	*input_basics.DeptOptionModel
	input_form.PageRes
}

type DeptTreeOptionReq struct {
	g.Meta `path:"/dept/treeOption" method:"get" tags:"ADMIN" summary:"部门_获取部门关系树选项"`
}

type DeptTreeOptionRes []tree.Node
