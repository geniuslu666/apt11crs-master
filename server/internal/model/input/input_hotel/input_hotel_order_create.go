package input_hotel

type CreateOrderInp struct {
	PreOrderSn string `json:"preOrderSn" v:"required#预定订单号未知" dc:"预定订单号"`
	IsBalance  int    `json:"isBalance" v:"required#请选择是否使用余额支付" dc:"是否使用余额支付【1启用，2禁用】"`
	CouponId   int    `json:"couponId" dc:"使用优惠券ID"`
	Remark     string `json:"remark" dc:"备注"`
	User       struct {
		FirstName string `json:"first_name" v:"required#预订人firstName未知" dc:"名"`
		LastName  string `json:"last_name" v:"required#预订人lastName未知" dc:"姓"`
		Email     string `json:"email" v:"required|email#邮箱未知|邮箱格式错误" dc:"邮箱"`
		Phone     string `json:"phone" v:"required#手机号未知" dc:"手机号"`
		AreaNo    string `json:"area_no" v:"required#手机区号未知" dc:"手机国际区号"`
	} `json:"user" dc:"预订人信息"`
}

type CreateOrderModel struct {
	OrderSn  string `json:"orderSn"   dc:"订单号"`
	PayModel int    `json:"payModel"  dc:"支付方式-1纯余额支付-2余额加外部支付-3纯外部支付"`
	Balance  struct {
		BalanceAmount     float64        `json:"amount"        dc:"余额支付金额"`
		BalancePayOrderSn string         `json:"payOrderSn"    dc:"余额支付订单号"`
		BalanceConfig     *BalanceConfig `json:"balanceConfig" dc:"余额计算配置"`
	}
	ThirdPay struct {
		ThirdAmount     float64 `json:"amount"     dc:"余额支付金额"`
		ThirdPayOrderSn string  `json:"payOrderSn" dc:"余额支付订单号"`
		IsFxOrder       bool    `json:"isFxOrder" dc:"是否为分销订单"`
		ThirdConfig     struct {
			WebPayUrl          string `json:"webPayUrl"     dc:"web支付地址"`
			AppPayParams       string `json:"appPayParams"  dc:"APP支付参数包"`
			AppWechatPayParams struct {
				Appid     string `json:"appid"`
				Partnerid string `json:"partnerid"`
				Prepayid  string `json:"prepayid"`
				Package   string `json:"package"`
				Noncestr  string `json:"noncestr"`
				Timestamp string `json:"timestamp"`
				Sign      string `json:"sign"`
			} `json:"appWechatPayParams" dc:"微信支付APP支付参数包"`
		} `json:"config" dc:"第三方支付配置"`
	}
	Coupon struct {
		CouponId         int     `json:"couponId"    dc:"优惠券ID"`
		CouponName       string  `json:"couponName"  dc:"优惠券名称"`
		CouponAmount     float64 `json:"couponAmount" dc:"优惠券抵用金额"`
		CouponPayOrderSn string  `json:"couponPayOrderSn" dc:"优惠券支付订单号"`
	}
	CreateOrderTime string `json:"createOrderTime" dc:"创建订单时间"`
	Countdown       int    `json:"countdown" dc:"订单支付倒计时秒"`
	RoomItems       []*OrderRooms
}
