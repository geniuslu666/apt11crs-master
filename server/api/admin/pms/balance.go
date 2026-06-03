package pms

import (
	"APT/internal/model/input/input_app_member"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type BalanceListReq struct {
	g.Meta `path:"/pmsBalanceChange/list" method:"get" tags:"ADMIN_PMS" summary:"积分明细_列表"`
	input_app_member.PmsBalanceChangeListInp
}

type BalanceListRes struct {
	input_form.PageRes
	List []*input_app_member.PmsBalanceChangeListModel `json:"list"   dc:"数据列表"`
}

type BalanceStatReq struct {
	g.Meta `path:"/pmsBalanceChange/stat" method:"get" tags:"ADMIN_PMS" summary:"积分明细_概况"`
	input_app_member.PmsBalanceChangeStatInp
}

type BalanceStatRes struct {
	*input_app_member.PmsBalanceChangeStatModel
}
