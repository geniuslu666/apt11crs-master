package th

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_th"
	"github.com/gogf/gf/v2/frame/g"
)

type CouponCategoryListReq struct {
	g.Meta `path:"/thCouponCategory/list" method:"get" tags:"ADMIN_TH" summary:"礼品券_分类列表"`
	input_th.ThCouponCategoryListInp
}

type CouponCategoryListRes struct {
	input_form.PageRes
	List []*input_th.ThCouponCategoryListModel `json:"list"   dc:"数据列表"`
}

type CouponCategoryAllReq struct {
	g.Meta `path:"/thCouponCategory/all" method:"get" tags:"ADMIN_TH" summary:"礼品券_分类全部列表"`
	input_th.ThCouponCategoryAllInp
}

type CouponCategoryAllRes struct {
	List []*input_th.ThCouponCategoryAllModel `json:"list"   dc:"数据列表"`
}

type CouponCategoryViewReq struct {
	g.Meta `path:"/thCouponCategory/view" method:"get" tags:"ADMIN_TH" summary:"礼品券_分类详情"`
	input_th.ThCouponCategoryViewInp
}

type CouponCategoryViewRes struct {
	*input_th.ThCouponCategoryViewModel
}

type CouponCategoryEditReq struct {
	g.Meta `path:"/thCouponCategory/edit" method:"post" tags:"ADMIN_TH" summary:"礼品券_分类修改/新增"`
	input_th.ThCouponCategoryEditInp
}

type CouponCategoryEditRes struct{}

type CouponCategoryDeleteReq struct {
	g.Meta `path:"/thCouponCategory/delete" method:"post" tags:"ADMIN_TH" summary:"礼品券_分类删除"`
	input_th.ThCouponCategoryDeleteInp
}

type CouponCategoryDeleteRes struct{}

type CouponCategorySwitchReq struct {
	g.Meta `path:"/thCouponCategory/switch" method:"post" tags:"ADMIN_TH" summary:"礼品券_分类更新状态"`
	input_th.ThCouponCategorySwitchInp
}

type CouponCategorySwitchRes struct{}
