package basics

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"

	"github.com/gogf/gf/v2/frame/g"
)

// IndexBannerListReq 查询Banner列表
type IndexBannerListReq struct {
	g.Meta `path:"/indexBanner/list" method:"get" tags:"ADMIN" summary:"首页Banner-列表"`
	input_basics.IndexBannerListInp
}

type IndexBannerListRes struct {
	input_form.PageRes
	List []*input_basics.IndexBannerListModel `json:"list"   dc:"数据列表"`
}

// IndexBannerViewReq 获取Banner指定信息
type IndexBannerViewReq struct {
	g.Meta `path:"/indexBanner/view" method:"get" tags:"ADMIN" summary:"首页Banner-详情"`
	input_basics.IndexBannerViewInp
}

type IndexBannerViewRes struct {
	*input_basics.IndexBannerViewModel
}

// IndexBannerEditReq 修改/新增Banner
type IndexBannerEditReq struct {
	g.Meta `path:"/indexBanner/edit" method:"post" tags:"ADMIN" summary:"首页Banner-修改/新增"`
	input_basics.IndexBannerEditInp
}

type IndexBannerEditRes struct{}

// IndexBannerDeleteReq 删除Banner
type IndexBannerDeleteReq struct {
	g.Meta `path:"/indexBanner/delete" method:"post" tags:"ADMIN" summary:"首页Banner-删除"`
	input_basics.IndexBannerDeleteInp
}

type IndexBannerDeleteRes struct{}

// IndexBannerMaxSortReq 获取Banner最大排序
type IndexBannerMaxSortReq struct {
	g.Meta `path:"/indexBanner/maxSort" method:"get" tags:"ADMIN" summary:"首页Banner-获取最大排序"`
	input_basics.IndexBannerMaxSortInp
}

type IndexBannerMaxSortRes struct {
	*input_basics.IndexBannerMaxSortModel
}

// IndexBannerStatusReq 更新Banner状态
type IndexBannerStatusReq struct {
	g.Meta `path:"/indexBanner/status" method:"post" tags:"ADMIN" summary:"首页Banner-更新状态"`
	input_basics.IndexBannerStatusInp
}

type IndexBannerStatusRes struct{}

// IndexBannerSortReq 排序
type IndexBannerSortReq struct {
	g.Meta `path:"/indexBanner/sortUpdate" method:"post" tags:"ADMIN" summary:"首页Banner-排序"`
	input_basics.IndexBannerSortInp
}

type IndexBannerSortRes struct{}
