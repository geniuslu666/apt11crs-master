package food

import (
	"APT/internal/model/input/input_food"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type RestaurantNoticeListReq struct {
	g.Meta `path:"/foodRestaurantNotice/list" method:"get" tags:"ADMIN_FOOD" summary:"获取商家通知列表"`
	input_food.FoodRestaurantNoticeListInp
}

type RestaurantNoticeListRes struct {
	input_form.PageRes
	List []*input_food.FoodRestaurantNoticeListModel `json:"list"   dc:"数据列表"`
}

type RestaurantNoticeExportReq struct {
	g.Meta `path:"/foodRestaurantNotice/export" method:"get" tags:"ADMIN_FOOD" summary:"导出通知管理列表"`
	input_food.FoodRestaurantNoticeListInp
}

type RestaurantNoticeExportRes struct{}

type RestaurantNoticeViewReq struct {
	g.Meta `path:"/foodRestaurantNotice/view" method:"get" tags:"ADMIN_FOOD" summary:"获取商家通知指定信息"`
	input_food.FoodRestaurantNoticeViewInp
}

type RestaurantNoticeViewRes struct {
	*input_food.FoodRestaurantNoticeViewModel
}

type RestaurantNoticeEditReq struct {
	g.Meta `path:"/foodRestaurantNotice/edit" method:"post" tags:"ADMIN_FOOD" summary:"修改/新增商家通知"`
	input_food.FoodRestaurantNoticeEditInp
}

type RestaurantNoticeEditRes struct{}

type RestaurantNoticeDeleteReq struct {
	g.Meta `path:"/foodRestaurantNotice/delete" method:"post" tags:"ADMIN_FOOD" summary:"删除商家通知"`
	input_food.FoodRestaurantNoticeDeleteInp
}

type RestaurantNoticeDeleteRes struct{}

type RestaurantNoticeMaxSortReq struct {
	g.Meta `path:"/foodRestaurantNotice/maxSort" method:"get" tags:"ADMIN_FOOD" summary:"获取商家通知最大排序"`
	input_food.FoodRestaurantNoticeMaxSortInp
}

type RestaurantNoticeMaxSortRes struct {
	*input_food.FoodRestaurantNoticeMaxSortModel
}

type RestaurantNoticeStatusReq struct {
	g.Meta `path:"/foodRestaurantNotice/status" method:"post" tags:"ADMIN_FOOD" summary:"更新商家通知状态"`
	input_food.FoodRestaurantNoticeStatusInp
}

type RestaurantNoticeStatusRes struct{}
