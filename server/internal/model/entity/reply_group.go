// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ReplyGroup is the golang structure for table reply_group.
type ReplyGroup struct {
	Id        int    `json:"id"        orm:"id"         description:""`
	GroupName string `json:"groupName" orm:"group_name" description:"组名"`
	UserId    string `json:"userId"    orm:"user_id"    description:"客服账户"`
	EntId     string `json:"entId"     orm:"ent_id"     description:"客服企业ID"`
	IsTeam    int    `json:"isTeam"    orm:"is_team"    description:"1个人,2团队"`
}
