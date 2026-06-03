package hook

import "github.com/gogf/gf/v2/frame/g"

type CabinetReq struct {
	g.Meta        `path:"/cabinet/callback" method:"post" tags:"NOTIFY_HOOK" summary:"储物柜订单通知"`
	Appid         string `json:"appid" dc:"应用ID"`
	OrderId       int    `json:"orderId" dc:"mch订单ID"`
	OrderNo       string `json:"orderNo" dc:"mch订单号"`
	OutTradeNo    string `json:"outTradeNo" dc:"本系统订单ID"`
	OutTradeVipid string `json:"outTradeVipid" dc:"本系统下订单会员ID"`
	Status        int    `json:"status" dc:"订单状态"`
	EventType     string `json:"eventType" dc:"事件类型"`
	Timestamp     int    `json:"timestamp" dc:"时间戳"`
	Sign          string `json:"sign" dc:"签名"`
}

type CabinetRes struct{}
