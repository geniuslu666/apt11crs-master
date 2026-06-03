package food

import (
	"APT/internal/model/input/input_food"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type AreaListReq struct {
	g.Meta `path:"/foodArea/list" method:"get" tags:"ADMIN_FOOD" summary:"获取区域管理列表"`
	input_food.FoodAreaListInp
}

type AreaListRes struct {
	input_form.PageRes
	*input_food.FoodAreaListModel
}

type AreaViewReq struct {
	g.Meta `path:"/foodArea/view" method:"get" tags:"ADMIN_FOOD" summary:"获取区域管理指定信息"`
	input_food.FoodAreaViewInp
}

type AreaViewRes struct {
	*input_food.FoodAreaViewModel
}

type AreaEditReq struct {
	g.Meta `path:"/foodArea/edit" method:"post" tags:"ADMIN_FOOD" summary:"修改/新增区域管理"`
	input_food.FoodAreaEditInp
}

type AreaEditRes struct{}

type AreaDeleteReq struct {
	g.Meta `path:"/foodArea/delete" method:"post" tags:"ADMIN_FOOD" summary:"删除区域管理"`
	input_food.FoodAreaDeleteInp
}

type AreaDeleteRes struct{}

type AreaStatusReq struct {
	g.Meta `path:"/foodArea/status" method:"post" tags:"ADMIN_FOOD" summary:"更新区域管理状态"`
	input_food.FoodAreaStatusInp
}

type AreaStatusRes struct{}
