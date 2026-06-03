// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// FoodRestaurantTerminal is the golang structure for table food_restaurant_terminal.
type FoodRestaurantTerminal struct {
	RestaurantId int64 `json:"restaurantId" orm:"restaurant_id" description:"餐厅ID"`
	TerminalId   int64 `json:"terminalId"   orm:"terminal_id"   description:"终端ID"`
	PrintTimes   uint  `json:"printTimes"   orm:"print_times"   description:"打印联数(次数)"`
}
