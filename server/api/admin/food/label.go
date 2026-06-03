package food

import (
	"APT/internal/model/input/input_food"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type LabelListReq struct {
	g.Meta `path:"/foodLabel/list" method:"get" tags:"ADMIN_FOOD" summary:"获取餐厅标签列表"`
	input_food.FoodLabelListInp
}

type LabelListRes struct {
	input_form.PageRes
	List []*input_food.FoodLabelListModel `json:"list"   dc:"数据列表"`
}

type LabelViewReq struct {
	g.Meta `path:"/foodLabel/view" method:"get" tags:"ADMIN_FOOD" summary:"获取餐厅标签指定信息"`
	input_food.FoodLabelViewInp
}

type LabelViewRes struct {
	*input_food.FoodLabelViewModel
}

type LabelEditReq struct {
	g.Meta `path:"/foodLabel/edit" method:"post" tags:"ADMIN_FOOD" summary:"修改/新增餐厅标签"`
	input_food.FoodLabelEditInp
}

type LabelEditRes struct{}

type LabelDeleteReq struct {
	g.Meta `path:"/foodLabel/delete" method:"post" tags:"ADMIN_FOOD" summary:"删除餐厅标签"`
	input_food.FoodLabelDeleteInp
}

type LabelDeleteRes struct{}

type LabelMaxSortReq struct {
	g.Meta `path:"/foodLabel/maxSort" method:"get" tags:"ADMIN_FOOD" summary:"获取餐厅标签最大排序"`
	input_food.FoodLabelMaxSortInp
}

type LabelMaxSortRes struct {
	*input_food.FoodLabelMaxSortModel
}

type LabelStatusReq struct {
	g.Meta `path:"/foodLabel/status" method:"post" tags:"ADMIN_FOOD" summary:"更新餐厅标签状态"`
	input_food.FoodLabelStatusInp
}

type LabelStatusRes struct{}
