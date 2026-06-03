package food

import (
	"APT/internal/model/input/input_food"
	"APT/internal/model/input/input_form"

	"github.com/gogf/gf/v2/frame/g"
)

// MaintenanceListReq 获取餐厅维护模式列表
type MaintenanceListReq struct {
	g.Meta `path:"/foodMaintenance/list" method:"get" tags:"ADMIN_FOOD" summary:"获取餐厅维护模式列表"`
	input_food.FoodMaintenanceListInp
}

type MaintenanceListRes struct {
	input_form.PageRes
	List []*input_food.FoodMaintenanceListModel `json:"list"   dc:"数据列表"`
}

// MaintenanceViewReq 获取餐厅维护模式指定信息
type MaintenanceViewReq struct {
	g.Meta `path:"/foodMaintenance/view" method:"get" tags:"ADMIN_FOOD" summary:"获取餐厅维护模式指定信息"`
	input_food.FoodMaintenanceViewInp
}

type MaintenanceViewRes struct {
	*input_food.FoodMaintenanceViewModel
}

// MaintenanceEditReq 修改/新增餐厅维护模式
type MaintenanceEditReq struct {
	g.Meta `path:"/foodMaintenance/edit" method:"post" tags:"ADMIN_FOOD" summary:"修改/新增餐厅维护模式"`
	input_food.FoodMaintenanceEditInp
}

type MaintenanceEditRes struct{}

// MaintenanceDeleteReq 删除餐厅维护模式
type MaintenanceDeleteReq struct {
	g.Meta `path:"/foodMaintenance/delete" method:"post" tags:"ADMIN_FOOD" summary:"删除餐厅维护模式"`
	input_food.FoodMaintenanceDeleteInp
}

type MaintenanceDeleteRes struct{}

// MaintenanceLanguageListReq 获取餐厅维护模式语言
type MaintenanceLanguageListReq struct {
	g.Meta `path:"/foodMaintenance/languageList" method:"get" tags:"ADMIN_FOOD" summary:"获取餐厅维护模式语言"`
	input_food.FoodMaintenanceLanguageListInp
}

type MaintenanceLanguageListRes struct {
	List []*input_food.FoodMaintenanceLanguageListModel `json:"list"   dc:"数据列表"`
}
