package pms

import (
	"APT/internal/model/input/input_app_member"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type ExpListReq struct {
	g.Meta `path:"/pmsExpChange/list" method:"get" tags:"ADMIN_PMS" summary:"成长值明细_列表"`
	input_app_member.PmsExpChangeListInp
}

type ExpListRes struct {
	input_form.PageRes
	List []*input_app_member.PmsExpChangeListModel `json:"list"   dc:"数据列表"`
}
