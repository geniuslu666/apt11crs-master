// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// FoodRestaurantCuisine is the golang structure for table food_restaurant_cuisine.
type FoodRestaurantCuisine struct {
	RestaurantId int64 `json:"restaurantId" orm:"restaurant_id" description:"餐厅ID"`
	CuisineId    int64 `json:"cuisineId"    orm:"cuisine_id"    description:"菜系ID"`
}
