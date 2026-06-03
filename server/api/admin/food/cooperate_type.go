package food

import (
	"APT/internal/model/input/input_food"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type CooperateTypeListReq struct {
	g.Meta `path:"/foodCooperateType/list" method:"get" tags:"ADMIN_FOOD" summary:"获取订餐合作类型列表"`
	input_food.FoodCooperateTypeListInp
}

type CooperateTypeListRes struct {
	input_form.PageRes
	List []*input_food.FoodCooperateTypeListModel `json:"list"   dc:"数据列表"`
}

type CooperateTypeExportReq struct {
	g.Meta `path:"/foodCooperateType/export" method:"get" tags:"ADMIN_FOOD" summary:"导出订餐合作类型列表"`
	input_food.FoodCooperateTypeListInp
}

type CooperateTypeExportRes struct{}

type CooperateTypeViewReq struct {
	g.Meta `path:"/foodCooperateType/view" method:"get" tags:"ADMIN_FOOD" summary:"获取订餐合作类型指定信息"`
	input_food.FoodCooperateTypeViewInp
}

type CooperateTypeViewRes struct {
	*input_food.FoodCooperateTypeViewModel
}

type CooperateTypeEditReq struct {
	g.Meta `path:"/foodCooperateType/edit" method:"post" tags:"ADMIN_FOOD" summary:"修改/新增订餐合作类型"`
	input_food.FoodCooperateTypeEditInp
}

type CooperateTypeEditRes struct{}

type CooperateTypeDeleteReq struct {
	g.Meta `path:"/foodCooperateType/delete" method:"post" tags:"ADMIN_FOOD" summary:"删除订餐合作类型"`
	input_food.FoodCooperateTypeDeleteInp
}

type CooperateTypeDeleteRes struct{}

type CooperateTypeStatusReq struct {
	g.Meta `path:"/foodCooperateType/status" method:"post" tags:"ADMIN_FOOD" summary:"更新订餐合作类型状态"`
	input_food.FoodCooperateTypeStatusInp
}

type CooperateTypeStatusRes struct{}
