// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// FoodRestaurantTerminal is the golang structure of table hg_food_restaurant_terminal for DAO operations like Where/Data.
type FoodRestaurantTerminal struct {
	g.Meta       `orm:"table:hg_food_restaurant_terminal, do:true"`
	RestaurantId interface{} // 餐厅ID
	TerminalId   interface{} // 终端ID
	PrintTimes   interface{} // 打印联数(次数)
}
