// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// FoodRestaurantMember is the golang structure for table food_restaurant_member.
type FoodRestaurantMember struct {
	RestaurantId int64 `json:"restaurantId" orm:"restaurant_id" description:"餐厅ID"`
	MemberId     int64 `json:"memberId"     orm:"member_id"     description:"用户ID"`
}
