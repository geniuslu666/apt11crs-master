package pms

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_hotel"
	"github.com/gogf/gf/v2/frame/g"
)

type RegionListReq struct {
	g.Meta `path:"/propertyRegion/list" method:"get" tags:"ADMIN_PMS" summary:"物业区域列表"`
	input_hotel.PropertyRegionListInp
}

type RegionListRes struct {
	input_form.PageRes
	List []*input_hotel.PropertyRegionListModel `json:"list"   dc:"数据列表"`
}
