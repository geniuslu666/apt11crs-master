// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Welcome is the golang structure of table welcome for DAO operations like Where/Data.
type Welcome struct {
	g.Meta      `orm:"table:welcome, do:true"`
	Id          interface{} //
	UserId      interface{} // 客服账户
	Keyword     interface{} // 关键词，welcome为默认欢迎；wechat为公众号欢迎语
	Content     interface{} // 欢迎消息内容
	IsDefault   interface{} // 是否默认，未启用
	DelaySecond interface{} // 延迟秒数
	Ctime       *gtime.Time // 创建时间
}
