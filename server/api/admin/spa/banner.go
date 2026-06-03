package spa

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_spa"

	"github.com/gogf/gf/v2/frame/g"
)

// BannerListReq 查询Banner列表
type BannerListReq struct {
	g.Meta `path:"/spaBanner/list" method:"get" tags:"ADMIN_CAR" summary:"获取Banner列表"`
	input_spa.SpaBannerListInp
}

type BannerListRes struct {
	input_form.PageRes
	List []*input_spa.SpaBannerListModel `json:"list"   dc:"数据列表"`
}

// BannerViewReq 获取Banner指定信息
type BannerViewReq struct {
	g.Meta `path:"/spaBanner/view" method:"get" tags:"ADMIN_CAR" summary:"获取Banner指定信息"`
	input_spa.SpaBannerViewInp
}

type BannerViewRes struct {
	*input_spa.SpaBannerViewModel
}

// BannerEditReq 修改/新增Banner
type BannerEditReq struct {
	g.Meta `path:"/spaBanner/edit" method:"post" tags:"ADMIN_CAR" summary:"修改/新增Banner"`
	input_spa.SpaBannerEditInp
}

type BannerEditRes struct{}

// BannerDeleteReq 删除Banner
type BannerDeleteReq struct {
	g.Meta `path:"/spaBanner/delete" method:"post" tags:"ADMIN_CAR" summary:"删除Banner"`
	input_spa.SpaBannerDeleteInp
}

type BannerDeleteRes struct{}

// BannerMaxSortReq 获取Banner最大排序
type BannerMaxSortReq struct {
	g.Meta `path:"/spaBanner/maxSort" method:"get" tags:"ADMIN_CAR" summary:"获取Banner最大排序"`
	input_spa.SpaBannerMaxSortInp
}

type BannerMaxSortRes struct {
	*input_spa.SpaBannerMaxSortModel
}

// BannerStatusReq 更新Banner状态
type BannerStatusReq struct {
	g.Meta `path:"/spaBanner/status" method:"post" tags:"ADMIN_CAR" summary:"更新Banner状态"`
	input_spa.SpaBannerStatusInp
}

type BannerStatusRes struct{}
