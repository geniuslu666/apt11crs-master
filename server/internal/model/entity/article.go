// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Article is the golang structure for table article.
type Article struct {
	Id         int    `json:"id"         orm:"id"          description:""`
	Title      string `json:"title"      orm:"title"       description:"自动回复关键词"`
	Content    string `json:"content"    orm:"content"     description:"自动回复内容"`
	CatId      int    `json:"catId"      orm:"cat_id"      description:"自动回复分类ID"`
	UserId     string `json:"userId"     orm:"user_id"     description:"客服账户"`
	EntId      string `json:"entId"      orm:"ent_id"      description:"客服企业ID"`
	ApiUrl     string `json:"apiUrl"     orm:"api_url"     description:"第三方接口地址"`
	SearchType int    `json:"searchType" orm:"search_type" description:"1包含匹配,2精准匹配"`
	Score      int    `json:"score"      orm:"score"       description:"命中次数"`
}
