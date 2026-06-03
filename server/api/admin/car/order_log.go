package car

import (
	"APT/internal/model/input/input_car"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type OrderLogListReq struct {
	g.Meta `path:"/carOrderLog/list" method:"get" tags:"ADMIN_CAR" summary:"获取接送机预订单日志列表"`
	input_car.CarOrderLogListInp
}

type OrderLogListRes struct {
	input_form.PageRes
	List []*input_car.CarOrderLogListModel `json:"list"   dc:"数据列表"`
}
