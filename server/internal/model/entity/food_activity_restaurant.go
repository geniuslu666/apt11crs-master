// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// FoodActivityRestaurant is the golang structure for table food_activity_restaurant.
type FoodActivityRestaurant struct {
	ActivityId   int64 `json:"activityId"   orm:"activity_id"   description:"活动ID"`
	RestaurantId int64 `json:"restaurantId" orm:"restaurant_id" description:"餐厅ID"`
}
