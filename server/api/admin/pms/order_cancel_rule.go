package pms

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_hotel"
	"github.com/gogf/gf/v2/frame/g"
)

type CancelOrderRuleListReq struct {
	g.Meta `path:"/pmsCancelRate/list" method:"get" tags:"ADMIN_PMS" summary:"取消政策_列表"`
}

type CancelOrderRuleListRes struct {
	input_form.PageRes
	List []*input_hotel.PmsCancelRateListModel `json:"list"   dc:"数据列表"`
}

type CancelOrderRuleViewReq struct {
	g.Meta `path:"/pmsCancelRate/view" method:"get" tags:"ADMIN_PMS" summary:"取消政策_详情"`
	input_hotel.PmsCancelRateViewInp
}

type CancelOrderRuleViewRes struct {
	*input_hotel.PmsCancelRateViewModel
}

type CancelOrderRuleEditReq struct {
	g.Meta `path:"/pmsCancelRate/edit" method:"post" tags:"ADMIN_PMS" summary:"取消政策_修改/新增"`
	input_hotel.PmsCancelRateEditInp
}

type CancelOrderRuleEditRes struct{}

type CancelOrderRuleDeleteReq struct {
	g.Meta `path:"/pmsCancelRate/delete" method:"post" tags:"ADMIN_PMS" summary:"取消政策_删除"`
	input_hotel.PmsCancelRateDeleteInp
}

type CancelOrderRuleDeleteRes struct{}

type CancelOrderRuleMaxSortReq struct {
	g.Meta `path:"/pmsCancelRate/maxSort" method:"get" tags:"ADMIN_PMS" summary:"取消政策_最大排序"`
	input_hotel.PmsCancelRateMaxSortInp
}

type CancelOrderRuleMaxSortRes struct {
	*input_hotel.PmsCancelRateMaxSortModel
}
