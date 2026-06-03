// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// EntConfig is the golang structure of table ent_config for DAO operations like Where/Data.
type EntConfig struct {
	g.Meta    `orm:"table:ent_config, do:true"`
	Id        interface{} //
	ConfName  interface{} // 配置描述
	ConfKey   interface{} // 配置key
	ConfValue interface{} // 配置值
	Language  interface{} // 语言
	EntId     interface{} // 客服企业ID
}
