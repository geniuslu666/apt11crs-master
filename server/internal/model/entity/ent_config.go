// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// EntConfig is the golang structure for table ent_config.
type EntConfig struct {
	Id        uint   `json:"id"        orm:"id"         description:""`
	ConfName  string `json:"confName"  orm:"conf_name"  description:"配置描述"`
	ConfKey   string `json:"confKey"   orm:"conf_key"   description:"配置key"`
	ConfValue string `json:"confValue" orm:"conf_value" description:"配置值"`
	Language  string `json:"language"  orm:"language"   description:"语言"`
	EntId     string `json:"entId"     orm:"ent_id"     description:"客服企业ID"`
}
