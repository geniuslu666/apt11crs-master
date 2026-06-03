package food

import (
	"APT/internal/model/input/input_food"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type GoodsListReq struct {
	g.Meta `path:"/foodGoods/list" method:"get" tags:"ADMIN_FOOD" summary:"获取套餐管理列表"`
	input_food.FoodGoodsListInp
}

type GoodsListRes struct {
	input_form.PageRes
	List []*input_food.FoodGoodsListModel `json:"list"   dc:"数据列表"`
}

type GoodsExportReq struct {
	g.Meta `path:"/foodGoods/export" method:"get" tags:"ADMIN_FOOD" summary:"导出套餐管理列表"`
	input_food.FoodGoodsListInp
}

type GoodsExportRes struct{}

type GoodsViewReq struct {
	g.Meta `path:"/foodGoods/view" method:"get" tags:"ADMIN_FOOD" summary:"获取套餐管理指定信息"`
	input_food.FoodGoodsViewInp
}

type GoodsViewRes struct {
	*input_food.FoodGoodsViewModel
}

type GoodsEditReq struct {
	g.Meta `path:"/foodGoods/edit" method:"post" tags:"ADMIN_FOOD" summary:"修改/新增套餐管理"`
	input_food.FoodGoodsEditInp
}

type GoodsEditRes struct{}

type GoodsDeleteReq struct {
	g.Meta `path:"/foodGoods/delete" method:"post" tags:"ADMIN_FOOD" summary:"删除套餐管理"`
	input_food.FoodGoodsDeleteInp
}

type GoodsDeleteRes struct{}

type GoodsStatusReq struct {
	g.Meta `path:"/foodGoods/status" method:"post" tags:"ADMIN_FOOD" summary:"更新套餐状态"`
	input_food.FoodGoodsStatusInp
}

type GoodsStatusRes struct{}
