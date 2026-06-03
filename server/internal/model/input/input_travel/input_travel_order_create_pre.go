package input_travel

type PreCreateOrderInp struct {
	ProductId int    `json:"productId" dc:"产品ID"`
	SkuId     int    `json:"skuId" dc:"车型ID"`
	BookDate  string `json:"bookDate" dc:"预定日期"`
	BookNum   int    `json:"bookNum" dc:"预定人数"`
}

type PreCreateOrderModel struct {
	PreOrderSn string `json:"preOrderSn"     dc:"预定订单号"`
}
