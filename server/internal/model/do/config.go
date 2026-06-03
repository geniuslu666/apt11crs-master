// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Config is the golang structure of table config for DAO operations like Where/Data.
type Config struct {
	g.Meta    `orm:"table:config, do:true"`
	Id        interface{} //
	ConfName  interface{} // 配置项描述
	ConfKey   interface{} // 配置项key
	ConfValue interface{} // 配置项值
}
