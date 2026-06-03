package travel

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_travel"
	"github.com/gogf/gf/v2/frame/g"
)

// VerifyRecordListReq 获取核销记录列表
type VerifyRecordListReq struct {
	g.Meta `path:"/travel/verifyRecord/list" method:"get" tags:"ADMIN_TRAVEL" summary:"获取一日游核销记录列表"`
	input_travel.TravelVerifyRecordListInp
}

type VerifyRecordListRes struct {
	input_form.PageRes
	List []*input_travel.TravelVerifyRecordListModel `json:"list" dc:"数据列表"`
}
