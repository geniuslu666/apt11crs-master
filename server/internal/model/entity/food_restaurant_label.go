// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// FoodRestaurantLabel is the golang structure for table food_restaurant_label.
type FoodRestaurantLabel struct {
	RestaurantId int64 `json:"restaurantId" orm:"restaurant_id" description:"餐厅ID"`
	LabelId      int64 `json:"labelId"      orm:"label_id"      description:"标签ID"`
}
