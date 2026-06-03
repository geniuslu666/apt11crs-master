// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// FoodRestaurantSeat is the golang structure of table hg_food_restaurant_seat for DAO operations like Where/Data.
type FoodRestaurantSeat struct {
	g.Meta       `orm:"table:hg_food_restaurant_seat, do:true"`
	RestaurantId interface{} // 餐厅ID
	SeatId       interface{} // 座位ID
	Num          interface{} // 同时段可预约数量
}
