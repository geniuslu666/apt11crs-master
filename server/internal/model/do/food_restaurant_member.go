// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// FoodRestaurantMember is the golang structure of table hg_food_restaurant_member for DAO operations like Where/Data.
type FoodRestaurantMember struct {
	g.Meta       `orm:"table:hg_food_restaurant_member, do:true"`
	RestaurantId interface{} // 餐厅ID
	MemberId     interface{} // 用户ID
}
