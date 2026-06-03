package basics

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type ReferrerInfoReq struct {
	g.Meta `path:"/referrer/info" method:"post" tags:"APP_BASICS" summary:"[分销]推荐人信息"`
	input_form.PageReq
}

type ReferrerInfoRes struct {
	List []*struct {
		*entity.PmsReferrerLog
		MemberInfo *struct {
			g.Meta `orm:"table:hg_pms_member"`
			*entity.PmsMember
		} `json:"memberInfo" orm:"with:id=memberId"`
	} `json:"list" dc:"推荐人列表"`
	Count           int `json:"count" dc:"分页总数"`
	InviteCount     int `json:"inviteCount" dc:"邀请总数"`
	TodayInvite     int `json:"todayInvite" dc:"今日邀请"`
	WeekInvite      int `json:"weekInvite" dc:"本周邀请"`
	YesterdayInvite int `json:"yesterdayInvite" dc:"昨日邀请"`
}

type ReferrerListReq struct {
	g.Meta `path:"/referrer/list" method:"post" tags:"APP_BASICS" summary:"[分销]推荐人列表"`
	input_form.PageReq
}

type ReferrerListRes struct {
	List []*struct {
		*entity.PmsReferrerLog
		MemberInfo *struct {
			g.Meta `orm:"table:hg_pms_member"`
			*entity.PmsMember
		} `json:"memberInfo" orm:"with:id=memberId"`
	} `json:"list" dc:"推荐人列表"`
	Count int `json:"count" dc:"分页总数"`
}
