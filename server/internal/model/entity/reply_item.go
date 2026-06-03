// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ReplyItem is the golang structure for table reply_item.
type ReplyItem struct {
	Id       int    `json:"id"       orm:"id"        description:""`
	Content  string `json:"content"  orm:"content"   description:"快捷回复内容"`
	GroupId  int    `json:"groupId"  orm:"group_id"  description:"快捷回复分组ID"`
	UserId   string `json:"userId"   orm:"user_id"   description:"客服账户"`
	ItemName string `json:"itemName" orm:"item_name" description:"快捷回复标题"`
	EntId    string `json:"entId"    orm:"ent_id"    description:"客服企业ID"`
	IsTeam   int    `json:"isTeam"   orm:"is_team"   description:"1个人,2团队"`
}
