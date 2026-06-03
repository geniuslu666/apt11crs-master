package basics

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type GenCodesListReq struct {
	g.Meta `path:"/genCodes/list" method:"get" tags:"ADMIN" summary:"生成代码_获取生成代码列表"`
	input_basics.GenCodesListInp
}

type GenCodesListRes struct {
	List []*input_basics.GenCodesListModel `json:"list"   dc:"数据列表"`
	input_form.PageRes
}

type GenCodesViewReq struct {
	g.Meta `path:"/genCodes/view" method:"get" tags:"ADMIN" summary:"生成代码_获取指定信息"`
	input_basics.GenCodesViewInp
}

type GenCodesViewRes struct {
	*input_basics.GenCodesViewModel
}

type GenCodesEditReq struct {
	g.Meta `path:"/genCodes/edit" method:"post" tags:"ADMIN" summary:"生成代码_修改/新增生成代码"`
	input_basics.GenCodesEditInp
}

type GenCodesEditRes struct {
	*input_basics.GenCodesEditModel
}

type GenCodesDeleteReq struct {
	g.Meta `path:"/genCodes/delete" method:"post" tags:"ADMIN" summary:"生成代码_删除生成代码"`
	input_basics.GenCodesDeleteInp
}

type GenCodesDeleteRes struct{}

type GenCodesMaxSortReq struct {
	g.Meta `path:"/genCodes/maxSort" method:"get" tags:"ADMIN" summary:"生成代码_生成代码最大排序"`
	input_basics.GenCodesMaxSortInp
}

type GenCodesMaxSortRes struct {
	*input_basics.GenCodesMaxSortModel
}

type GenCodesStatusReq struct {
	g.Meta `path:"/genCodes/status" method:"post" tags:"ADMIN" summary:"生成代码_更新生成代码状态"`
	input_basics.GenCodesStatusInp
}

type GenCodesStatusRes struct{}

type GenCodesSelectsReq struct {
	g.Meta `path:"/genCodes/selects" method:"get" tags:"ADMIN" summary:"生成代码_生成入口选项"`
	input_basics.GenCodesSelectsInp
}

type GenCodesSelectsRes struct {
	*input_basics.GenCodesSelectsModel
}

type GenCodesTableSelectReq struct {
	g.Meta `path:"/genCodes/tableSelect" method:"get" tags:"ADMIN" summary:"生成代码_数据库表选项"`
	input_basics.GenCodesTableSelectInp
}

type GenCodesTableSelectRes []*input_basics.GenCodesTableSelectModel

type GenCodesColumnSelectReq struct {
	g.Meta `path:"/genCodes/columnSelect" method:"get" tags:"ADMIN" summary:"生成代码_表字段选项"`
	input_basics.GenCodesColumnSelectInp
}

type GenCodesColumnSelectRes []*input_basics.GenCodesColumnSelectModel

type GenCodesColumnListReq struct {
	g.Meta `path:"/genCodes/columnList" method:"get" tags:"ADMIN" summary:"生成代码_表字段列表"`
	input_basics.GenCodesColumnListInp
}

type GenCodesColumnListRes []*input_basics.GenCodesColumnListModel

type GenCodesPreviewReq struct {
	g.Meta `path:"/genCodes/preview" method:"post" tags:"ADMIN" summary:"生成代码_生成预览"`
	input_basics.GenCodesPreviewInp
}

type GenCodesPreviewRes struct {
	*input_basics.GenCodesPreviewModel
}

type GenCodesBuildReq struct {
	g.Meta `path:"/genCodes/build" method:"post" tags:"ADMIN" summary:"生成代码_提交生成"`
	input_basics.GenCodesBuildInp
}

type GenCodesBuildRes struct {
}
