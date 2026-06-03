// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Config is the golang structure for table config.
type Config struct {
	Id        uint   `json:"id"        orm:"id"         description:""`
	ConfName  string `json:"confName"  orm:"conf_name"  description:"配置项描述"`
	ConfKey   string `json:"confKey"   orm:"conf_key"   description:"配置项key"`
	ConfValue string `json:"confValue" orm:"conf_value" description:"配置项值"`
}
