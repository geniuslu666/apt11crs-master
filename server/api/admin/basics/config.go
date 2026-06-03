package basics

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"

	"github.com/gogf/gf/v2/frame/g"
)

type ConfigGetReq struct {
	g.Meta `path:"/config/get" method:"get" tags:"ADMIN" summary:"配置_获取指定分组的配置"`
	input_basics.GetConfigInp
}

type ConfigGetRes struct {
	*input_basics.GetConfigModel
}

type ConfigUpdateReq struct {
	g.Meta `path:"/config/update" method:"post" tags:"ADMIN" summary:"配置_获取指定分组的配置"`
	input_basics.UpdateConfigInp
}

type ConfigUpdateRes struct {
}

type ConfigTypeSelectReq struct {
	g.Meta `path:"/config/typeSelect" method:"get" tags:"ADMIN" summary:"配置_数据类型选项"`
}

type ConfigTypeSelectRes struct {
	List []*input_form.Select `json:"list" response:"arr"`
}

type ConfigGetCashReq struct {
	g.Meta `path:"/config/getCash" method:"get" tags:"ADMIN" summary:"配置_获取提现配置"`
}

type ConfigGetCashRes struct {
	*input_basics.GetConfigModel
}

type SpaSmsConfigUpdateReq struct {
	g.Meta `path:"/config/spaSmsConfigUpdate" method:"post" tags:"ADMIN" summary:"短信_配置按摩短信"`
	input_basics.UpdateSpaSmsConfigInp
}

type SpaSmsConfigUpdateRes struct {
}

type ConfigUpdateOrderMemberReq struct {
	g.Meta `path:"/config/updateOrderMember" method:"post" tags:"ADMIN" summary:"配置_更新订单用户"`
}

type ConfigUpdateOrderMemberRes struct {
}
