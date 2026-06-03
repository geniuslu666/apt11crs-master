package member

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type BalanceChangeReq struct {
	g.Meta     `path:"/balanceChange/list" method:"post" tags:"APP_MEMBER" summary:"余额_消费明细"`
	PageNum    int    `json:"pageNum" v:"required#page_number_unknown" dc:"页码"`
	PageSize   int    `json:"pageSize" v:"required#page_length_unknown" dc:"页长"`
	ChangeMode string `json:"changeMode" dc:"in、收入  out、支出"`
}

type BalanceChangeRes struct {
	List []*struct {
		ChangePrice float64     `json:"changePrice" dc:"变更金额"`
		OrderSn     string      `json:"orderSn"     dc:"订单号"`
		Des         string      `json:"des"         dc:"描述"`
		CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
	} `json:"list"`
	Count int `json:"count"`
}
