package car

import (
	"APT/internal/model/input/input_car"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type AddressTypeListReq struct {
	g.Meta `path:"/carAddressType/list" method:"get" tags:"ADMIN_CAR" summary:"获取接送机地址类型列表"`
	input_car.CarAddressTypeListInp
}

type AddressTypeListRes struct {
	input_form.PageRes
	List []*input_car.CarAddressTypeListModel `json:"list"   dc:"数据列表"`
}
