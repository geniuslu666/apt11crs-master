package food

import (
	"APT/internal/model/input/input_food"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type ActivityListReq struct {
	g.Meta `path:"/foodActivity/list" method:"get" tags:"ADMIN_FOOD" summary:"订餐活动_列表"`
	input_food.FoodActivityListInp
}

type ActivityListRes struct {
	input_form.PageRes
	List []*input_food.FoodActivityListModel `json:"list"   dc:"数据列表"`
}

type ActivityExportReq struct {
	g.Meta `path:"/foodActivity/export" method:"get" tags:"ADMIN_FOOD" summary:"订餐活动_导出"`
	input_food.FoodActivityListInp
}

type ActivityExportRes struct{}

type ActivityViewReq struct {
	g.Meta `path:"/foodActivity/view" method:"get" tags:"ADMIN_FOOD" summary:"订餐活动_详情"`
	input_food.FoodActivityViewInp
}

type ActivityViewRes struct {
	*input_food.FoodActivityViewModel
}

type ActivityEditReq struct {
	g.Meta `path:"/foodActivity/edit" method:"post" tags:"ADMIN_FOOD" summary:"订餐活动_修改/新增"`
	input_food.FoodActivityEditInp
}

type ActivityEditRes struct{}

type ActivityDeleteReq struct {
	g.Meta `path:"/foodActivity/delete" method:"post" tags:"ADMIN_FOOD" summary:"订餐活动_删除"`
	input_food.FoodActivityDeleteInp
}

type ActivityDeleteRes struct{}

type ActivityStatusReq struct {
	g.Meta `path:"/foodActivity/status" method:"post" tags:"ADMIN_FOOD" summary:"订餐活动_更新"`
	input_food.FoodActivityStatusInp
}

type ActivityStatusRes struct{}
