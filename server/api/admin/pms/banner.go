package pms

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

// BannerListReq 查询banner 横幅列表
type BannerListReq struct {
	g.Meta `path:"/pmsBanner/list" method:"get" tags:"ADMIN_PMS" summary:"BANNER_列表"`
	input_basics.PmsBannerListInp
}

type BannerListRes struct {
	input_form.PageRes
	List []*input_basics.PmsBannerListModel `json:"list"   dc:"数据列表"`
}

// BannerViewReq 获取banner 横幅指定信息
type BannerViewReq struct {
	g.Meta `path:"/pmsBanner/view" method:"get" tags:"ADMIN_PMS" summary:"BANNER_详情"`
	input_basics.PmsBannerViewInp
}

type BannerViewRes struct {
	*input_basics.PmsBannerViewModel
}

// BannerEditReq 修改/新增banner 横幅
type BannerEditReq struct {
	g.Meta `path:"/pmsBanner/edit" method:"post" tags:"ADMIN_PMS" summary:"BANNER_修改/新增"`
	input_basics.PmsBannerEditInp
}

type BannerEditRes struct{}

// BannerDeleteReq 删除banner 横幅
type BannerDeleteReq struct {
	g.Meta `path:"/pmsBanner/delete" method:"post" tags:"ADMIN_PMS" summary:"BANNER_删除"`
	input_basics.PmsBannerDeleteInp
}

type BannerDeleteRes struct{}
