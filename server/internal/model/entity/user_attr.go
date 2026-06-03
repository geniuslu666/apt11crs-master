// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserAttr is the golang structure for table user_attr.
type UserAttr struct {
	Id               int         `json:"id"               orm:"id"                 description:""`
	KefuName         string      `json:"kefuName"         orm:"kefu_name"          description:"客服账户"`
	AuthKey          string      `json:"authKey"          orm:"auth_key"           description:"authKey"`
	IdCard           string      `json:"idCard"           orm:"id_card"            description:"身份证号码"`
	EntId            string      `json:"entId"            orm:"ent_id"             description:"企业ID"`
	GoodNum          uint        `json:"goodNum"          orm:"good_num"           description:"评价好评数量"`
	NormalNum        uint        `json:"normalNum"        orm:"normal_num"         description:"评价中评数量"`
	BadNum           uint        `json:"badNum"           orm:"bad_num"            description:"评价差评数量"`
	AigcSessionScore uint        `json:"aigcSessionScore" orm:"aigc_session_score" description:"aigc积分"`
	Money            uint        `json:"money"            orm:"money"              description:"总金额"`
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"         description:"创建时间"`
}
