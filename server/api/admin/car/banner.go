package car

import (
	"APT/internal/model/input/input_car"
	"APT/internal/model/input/input_form"

	"github.com/gogf/gf/v2/frame/g"
)

// BannerListReq 查询Banner列表
type BannerListReq struct {
	g.Meta `path:"/carBanner/list" method:"get" tags:"ADMIN_CAR" summary:"获取Banner列表"`
	input_car.CarBannerListInp
}

type BannerListRes struct {
	input_form.PageRes
	List []*input_car.CarBannerListModel `json:"list"   dc:"数据列表"`
}

// BannerViewReq 获取Banner指定信息
type BannerViewReq struct {
	g.Meta `path:"/carBanner/view" method:"get" tags:"ADMIN_CAR" summary:"获取Banner指定信息"`
	input_car.CarBannerViewInp
}

type BannerViewRes struct {
	*input_car.CarBannerViewModel
}

// BannerEditReq 修改/新增Banner
type BannerEditReq struct {
	g.Meta `path:"/carBanner/edit" method:"post" tags:"ADMIN_CAR" summary:"修改/新增Banner"`
	input_car.CarBannerEditInp
}

type BannerEditRes struct{}

// BannerDeleteReq 删除Banner
type BannerDeleteReq struct {
	g.Meta `path:"/carBanner/delete" method:"post" tags:"ADMIN_CAR" summary:"删除Banner"`
	input_car.CarBannerDeleteInp
}

type BannerDeleteRes struct{}

// BannerMaxSortReq 获取Banner最大排序
type BannerMaxSortReq struct {
	g.Meta `path:"/carBanner/maxSort" method:"get" tags:"ADMIN_CAR" summary:"获取Banner最大排序"`
	input_car.CarBannerMaxSortInp
}

type BannerMaxSortRes struct {
	*input_car.CarBannerMaxSortModel
}

// BannerStatusReq 更新Banner状态
type BannerStatusReq struct {
	g.Meta `path:"/carBanner/status" method:"post" tags:"ADMIN_CAR" summary:"更新Banner状态"`
	input_car.CarBannerStatusInp
}

type BannerStatusRes struct{}
