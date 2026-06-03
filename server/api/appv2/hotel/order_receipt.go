package hotel

import (
	"APT/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type OrderReceiptReq struct {
	g.Meta   `path:"/v2/order/receipt" method:"post" tags:"APP_HOTEL_V2" summary:"酒店订单_酒店订单收据"`
	UserName string `json:"userName" v:"required#please_fill_in_the_username" dc:"用户姓名"`
	OrderSn  string `json:"orderSn"  v:"required#please_fill_in_the_order_number" dc:"订单号"`
	Mail     string `json:"mail"     v:"required#please_fill_in_the_email" dc:"邮箱"`
}

type OrderReceiptRes struct {
}

type OrderReceiptViewReq struct {
	g.Meta  `path:"/v2/order/receiptView" method:"post" tags:"APP_HOTEL_V2" summary:"酒店订单_酒店订单收据详情"`
	OrderSn string `json:"orderSn"  v:"required#please_fill_in_the_order_number" dc:"订单号"`
}

type OrderReceiptViewRes struct {
	*entity.Receipt
}
