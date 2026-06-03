// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// FoodRestaurantCuisine is the golang structure of table hg_food_restaurant_cuisine for DAO operations like Where/Data.
type FoodRestaurantCuisine struct {
	g.Meta       `orm:"table:hg_food_restaurant_cuisine, do:true"`
	RestaurantId interface{} // 餐厅ID
	CuisineId    interface{} // 菜系ID
}
