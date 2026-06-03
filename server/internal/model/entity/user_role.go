// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UserRole is the golang structure for table user_role.
type UserRole struct {
	Id     int `json:"id"     orm:"id"      description:""`
	UserId int `json:"userId" orm:"user_id" description:"客服ID"`
	RoleId int `json:"roleId" orm:"role_id" description:"角色ID"`
}
