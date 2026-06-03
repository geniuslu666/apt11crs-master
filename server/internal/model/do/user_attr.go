// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserAttr is the golang structure of table user_attr for DAO operations like Where/Data.
type UserAttr struct {
	g.Meta           `orm:"table:user_attr, do:true"`
	Id               interface{} //
	KefuName         interface{} // 客服账户
	AuthKey          interface{} // authKey
	IdCard           interface{} // 身份证号码
	EntId            interface{} // 企业ID
	GoodNum          interface{} // 评价好评数量
	NormalNum        interface{} // 评价中评数量
	BadNum           interface{} // 评价差评数量
	AigcSessionScore interface{} // aigc积分
	Money            interface{} // 总金额
	CreatedAt        *gtime.Time // 创建时间
}
