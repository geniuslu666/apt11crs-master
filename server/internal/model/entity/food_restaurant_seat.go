// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// FoodRestaurantSeat is the golang structure for table food_restaurant_seat.
type FoodRestaurantSeat struct {
	RestaurantId int64 `json:"restaurantId" orm:"restaurant_id" description:"餐厅ID"`
	SeatId       int64 `json:"seatId"       orm:"seat_id"       description:"座位ID"`
	Num          int   `json:"num"          orm:"num"           description:"同时段可预约数量"`
}
