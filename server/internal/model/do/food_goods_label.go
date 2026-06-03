// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// FoodGoodsLabel is the golang structure of table hg_food_goods_label for DAO operations like Where/Data.
type FoodGoodsLabel struct {
	g.Meta  `orm:"table:hg_food_goods_label, do:true"`
	GoodsId interface{} // 餐厅ID
	LabelId interface{} // 标签ID
}
