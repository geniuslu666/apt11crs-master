// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Visitor is the golang structure of table visitor for DAO operations like Where/Data.
type Visitor struct {
	g.Meta    `orm:"table:visitor, do:true"`
	Id        interface{} //
	Name      interface{} // 访客显示名称
	RealName  interface{} // 访客真实姓名
	Avator    interface{} // 访客头像
	SourceIp  interface{} // 访客来源IP
	ToId      interface{} // 对接客服账户
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
	VisitorId interface{} // 访客唯一ID
	Status    interface{} // 访客状态
	State     interface{} // 访客状态位
	Refer     interface{} // 访客来源
	City      interface{} // 访客城市
	ClientIp  interface{} // 访客IP
	Extra     interface{} // 访客扩展信息
	EntId     interface{} // 对接的企业ID
	VisitNum  interface{} // 访客访问次数
}
