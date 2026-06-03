// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Visitor is the golang structure for table visitor.
type Visitor struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	Name      string      `json:"name"      orm:"name"       description:"访客显示名称"`
	RealName  string      `json:"realName"  orm:"real_name"  description:"访客真实姓名"`
	Avator    string      `json:"avator"    orm:"avator"     description:"访客头像"`
	SourceIp  string      `json:"sourceIp"  orm:"source_ip"  description:"访客来源IP"`
	ToId      string      `json:"toId"      orm:"to_id"      description:"对接客服账户"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`
	VisitorId string      `json:"visitorId" orm:"visitor_id" description:"访客唯一ID"`
	Status    int         `json:"status"    orm:"status"     description:"访客状态"`
	State     string      `json:"state"     orm:"state"      description:"访客状态位"`
	Refer     string      `json:"refer"     orm:"refer"      description:"访客来源"`
	City      string      `json:"city"      orm:"city"       description:"访客城市"`
	ClientIp  string      `json:"clientIp"  orm:"client_ip"  description:"访客IP"`
	Extra     string      `json:"extra"     orm:"extra"      description:"访客扩展信息"`
	EntId     uint        `json:"entId"     orm:"ent_id"     description:"对接的企业ID"`
	VisitNum  uint        `json:"visitNum"  orm:"visit_num"  description:"访客访问次数"`
}
