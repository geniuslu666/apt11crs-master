// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ArticleCate is the golang structure for table article_cate.
type ArticleCate struct {
	Id      int    `json:"id"      orm:"id"       description:""`
	CatName string `json:"catName" orm:"cat_name" description:"自动回复分类名称"`
	UserId  string `json:"userId"  orm:"user_id"  description:"客服账户"`
	EntId   string `json:"entId"   orm:"ent_id"   description:"客服企业ID"`
	IsTop   int    `json:"isTop"   orm:"is_top"   description:"是否置顶展示，1置顶"`
}
