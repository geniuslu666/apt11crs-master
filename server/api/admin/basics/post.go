package basics

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type PostEditReq struct {
	g.Meta `path:"/post/edit" method:"post" tags:"ADMIN" summary:"岗位_修改/新增岗位"`
	input_basics.PostEditInp
}

type PostEditRes struct{}

type PostDeleteReq struct {
	g.Meta `path:"/post/delete" method:"post" tags:"ADMIN" summary:"岗位_删除岗位"`
	input_basics.PostDeleteInp
}

type PostDeleteRes struct{}

type PostMaxSortReq struct {
	g.Meta `path:"/post/maxSort" method:"get" tags:"ADMIN" summary:"岗位_岗位最大排序"`
	input_basics.PostMaxSortInp
}

type PostMaxSortRes struct {
	*input_basics.PostMaxSortModel
}

type PostListReq struct {
	g.Meta `path:"/post/list" method:"get" tags:"ADMIN" summary:"岗位_获取岗位列表"`
	input_basics.PostListInp
}

type PostListRes struct {
	List []*input_basics.PostListModel `json:"list"   description:"数据列表"`
	input_form.PageRes
}

type PostViewReq struct {
	g.Meta `path:"/post/view" method:"get" tags:"ADMIN" summary:"岗位_获取指定信息"`
	input_basics.PostViewInp
}

type PostViewRes struct {
	*input_basics.PostViewModel
}

type PostStatusReq struct {
	g.Meta `path:"/post/status" method:"post" tags:"ADMIN" summary:"岗位_更新岗位状态"`
	input_basics.PostStatusInp
}

type PostStatusRes struct{}
