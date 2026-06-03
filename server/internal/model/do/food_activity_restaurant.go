// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// FoodActivityRestaurant is the golang structure of table hg_food_activity_restaurant for DAO operations like Where/Data.
type FoodActivityRestaurant struct {
	g.Meta       `orm:"table:hg_food_activity_restaurant, do:true"`
	ActivityId   interface{} // 活动ID
	RestaurantId interface{} // 餐厅ID
}
