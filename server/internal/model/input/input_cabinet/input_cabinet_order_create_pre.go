package input_cabinet

type PreCreateOrderInp struct {
	CabinetId int `json:"cabinetId" dc:"储物柜ID"`
}

type PreCreateOrderModel struct {
	PreOrderSn string `json:"preOrderSn" dc:"预定订单号"`
}
