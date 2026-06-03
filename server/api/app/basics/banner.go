package basics

import (
	"APT/internal/model/input/input_basics"
	"github.com/gogf/gf/v2/frame/g"
)

type AdSpaceAfterPayBannerListReq struct {
	g.Meta `path:"/banner/adSpaceAfterPayList" method:"post" tags:"APP_BASICS" summary:"[轮播]获取支付后广告位列表"`
}

type AdSpaceAfterPayBannerListRes struct {
	List []*input_basics.PmsBannerAppListModel `json:"list"   dc:"数据列表"`
}
