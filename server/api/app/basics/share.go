package basics

import "github.com/gogf/gf/v2/frame/g"

type ShareInfoReq struct {
	g.Meta    `path:"/home/shareImage" method:"post" tags:"APP_BASICS" summary:"[分享]分享"`
	ShareType string `p:"shareType" v:"required|in:member,host#share_type_cannot_be_empty|share_type_error"`
	Pid       int    `p:"pid" v:"required-if:shareType,host#share_id_cannot_be_empty"`
}

type ShareInfoRes struct {
	Icon        string `json:"icon"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Url         string `json:"url"`
	ShareDomain string `json:"shareDomain"`
}

type ShareInvitePageInfoReq struct {
	g.Meta `path:"/home/invitePageInfo" method:"post" tags:"APP_BASICS" summary:"[分享]邀请分享"`
}

type ShareInvitePageInfoRes struct {
	AwardScore      float64 `json:"awardScore"        dc:"邀请奖励积分"`
	AwardCommission float64 `json:"awardCommission"   dc:"佣金比例"`
	GetScore        float64 `json:"getScore"          dc:"获得奖励积分"`
	GetCommission   float64 `json:"getCommission"     dc:"获得总获得佣金"`
	SuccessMember   int     `json:"successMember"     dc:"邀请成功人数"`
	WithdrawPrice   float64 `json:"withdrawPrice"     dc:"可提现金额"`
}
