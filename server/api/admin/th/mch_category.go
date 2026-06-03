package th

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_th"
	"github.com/gogf/gf/v2/frame/g"
)

type MchCategoryListReq struct {
	g.Meta `path:"/thMchCategory/list" method:"get" tags:"ADMIN_TH" summary:"商户_分类列表"`
	input_th.ThMchCategoryListInp
}

type MchCategoryListRes struct {
	input_form.PageRes
	List []*input_th.ThMchCategoryListModel `json:"list"   dc:"数据列表"`
}

type MchCategoryAllReq struct {
	g.Meta `path:"/thMchCategory/all" method:"get" tags:"ADMIN_TH" summary:"商户_分类全部列表"`
	input_th.ThMchCategoryAllInp
}

type MchCategoryAllRes struct {
	List []*input_th.ThMchCategoryAllModel `json:"list"   dc:"数据列表"`
}

type MchCategoryViewReq struct {
	g.Meta `path:"/thMchCategory/view" method:"get" tags:"ADMIN_TH" summary:"商户_分类详情"`
	input_th.ThMchCategoryViewInp
}

type MchCategoryViewRes struct {
	*input_th.ThMchCategoryViewModel
}

type MchCategoryEditReq struct {
	g.Meta `path:"/thMchCategory/edit" method:"post" tags:"ADMIN_TH" summary:"商户_分类修改/新增"`
	input_th.ThMchCategoryEditInp
}

type MchCategoryEditRes struct{}

type MchCategoryDeleteReq struct {
	g.Meta `path:"/thMchCategory/delete" method:"post" tags:"ADMIN_TH" summary:"商户_分类删除"`
	input_th.ThMchCategoryDeleteInp
}

type MchCategoryDeleteRes struct{}

type MchCategorySwitchReq struct {
	g.Meta `path:"/thMchCategory/switch" method:"post" tags:"ADMIN_TH" summary:"商户_分类更新状态"`
	input_th.ThMchCategorySwitchInp
}

type MchCategorySwitchRes struct{}
