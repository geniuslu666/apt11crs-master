package kefu

import (
	"APT/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type FindMemberInfoReq struct {
	g.Meta `path:"/findMemberInfo" method:"get" tags:"NOTIFY_KEFU" summary:"查询会员信息" `
	Id     int `json:"id" v:"required#member_id_cannot_be_empty" dc:"用户编号"`
}

type FindMemberInfoRes struct {
	*FindMemberInfo
}

type FindMemberInfosReq struct {
	g.Meta   `path:"/findMemberInfos" method:"get" tags:"NOTIFY_KEFU" summary:"查询会员信息" `
	Ids      string `json:"ids" dc:"用户编号"`
	KeyWords string `json:"keyWords" dc:"搜索关键字"`
}

type FindMemberInfosRes struct {
	List []*FindMemberInfo
}

type FindMemberInfo struct {
	*entity.PmsMember
	LevelInfo *struct {
		g.Meta `orm:"table:hg_pms_member_level"`
		*entity.PmsMemberLevel
	} `json:"levelInfo" dc:"会员等级信息" orm:"with:id=level"`
}
