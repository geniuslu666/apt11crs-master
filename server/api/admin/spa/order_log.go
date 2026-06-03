package spa

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_spa"
	"github.com/gogf/gf/v2/frame/g"
)

type OrderLogListReq struct {
	g.Meta `path:"/spaOrderLog/list" method:"get" tags:"ADMIN_SPA" summary:"获取按摩预订单日志列表"`
	input_spa.SpaOrderLogListInp
}

type OrderLogListRes struct {
	input_form.PageRes
	List []*input_spa.SpaOrderLogListModel `json:"list"   dc:"数据列表"`
}
