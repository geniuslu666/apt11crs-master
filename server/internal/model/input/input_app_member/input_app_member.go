package input_app_member

type MemberBalanceInp struct {
	MemberId      int     `json:"member_id" v:"required" dc:"用户ID"`
	Scene         string  `json:"scene" v:"required" dc:"场景"`
	Type          string  `json:"type" v:"required" dc:"类型"`
	ChangeBalance float64 `json:"change_balance" v:"required" dc:"变动金额"`
	OrderSn       string  `json:"order_sn" v:"required" dc:"订单号"`
	Des           string  `json:"des" dc:"消费描述"`
	Reason        string  `json:"reason" dc:"原因"`
	MdCode        string  `json:"mdCode"      dc:"注册设备码"`
	MpModel       string  `json:"mpModel"     dc:"注册设备型号"`
}
