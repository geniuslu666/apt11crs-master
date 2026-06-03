package pms

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_hotel"
	"github.com/gogf/gf/v2/frame/g"
)

type OrderLogListReq struct {
	g.Meta `path:"/hotelOrderLog/list" method:"get" tags:"ADMIN_PMS" summary:"获取酒店预订单日志列表"`
	input_hotel.AppStayLogListInp
}

type OrderLogListRes struct {
	input_form.PageRes
	List []*input_hotel.AppStayLogListModel `json:"list"   dc:"数据列表"`
}
