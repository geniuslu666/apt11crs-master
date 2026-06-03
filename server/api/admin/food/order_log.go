package food

import (
	"APT/internal/model/input/input_food"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type OrderLogListReq struct {
	g.Meta `path:"/foodOrderLog/list" method:"get" tags:"ADMIN_FOOD" summary:"获取餐厅预订单日志列表"`
	input_food.FoodOrderLogListInp
}

type OrderLogListRes struct {
	input_form.PageRes
	List []*input_food.FoodOrderLogListModel `json:"list"   dc:"数据列表"`
}
