package food

import (
	"APT/internal/model/input/input_food"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type SeatListReq struct {
	g.Meta `path:"/foodSeat/list" method:"get" tags:"ADMIN_FOOD" summary:"获取订餐-座位表列表"`
	input_food.FoodSeatListInp
}

type SeatListRes struct {
	input_form.PageRes
	List []*input_food.FoodSeatListModel `json:"list"   dc:"数据列表"`
}

type SeatViewReq struct {
	g.Meta `path:"/foodSeat/view" method:"get" tags:"ADMIN_FOOD" summary:"获取订餐-座位表指定信息"`
	input_food.FoodSeatViewInp
}

type SeatViewRes struct {
	*input_food.FoodSeatViewModel
}

type SeatEditReq struct {
	g.Meta `path:"/foodSeat/edit" method:"post" tags:"ADMIN_FOOD" summary:"修改/新增订餐-座位表"`
	input_food.FoodSeatEditInp
}

type SeatEditRes struct{}

type SeatDeleteReq struct {
	g.Meta `path:"/foodSeat/delete" method:"post" tags:"ADMIN_FOOD" summary:"删除订餐-座位表"`
	input_food.FoodSeatDeleteInp
}

type SeatDeleteRes struct{}

type SeatStatusReq struct {
	g.Meta `path:"/foodSeat/status" method:"post" tags:"ADMIN_FOOD" summary:"更新订餐-座位表状态"`
	input_food.FoodSeatStatusInp
}

type SeatStatusRes struct{}
