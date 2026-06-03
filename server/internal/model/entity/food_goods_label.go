// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// FoodGoodsLabel is the golang structure for table food_goods_label.
type FoodGoodsLabel struct {
	GoodsId int64 `json:"goodsId" orm:"goods_id" description:"餐厅ID"`
	LabelId int64 `json:"labelId" orm:"label_id" description:"标签ID"`
}
