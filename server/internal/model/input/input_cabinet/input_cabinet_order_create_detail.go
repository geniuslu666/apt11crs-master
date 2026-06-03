package input_cabinet

import (
	"APT/internal/library/cabinetApi"

	"github.com/gogf/gf/v2/container/gvar"
)

// PreOrderDetailInp 预订单详情入参
type PreOrderDetailInp struct {
	PreOrderSn string `json:"preOrderSn" dc:"预定订单号"`
}

// PreOrderDetailModel 预订单详情内容
type PreOrderDetailModel struct {
	PreOrderSn  string                              `json:"preOrderSn"    dc:"预订单号"`
	IsFx        string                              `json:"isFx"       dc:"是否是分销"`
	CabinetInfo *cabinetApi.CabinetInfoResponseItem `json:"cabinetInfo" dc:"储物柜信息"`
	BoxTypeInfo *BoxTypeInfoItem                    `json:"boxTypeInfo" dc:"所选择的格口类型信息"`
	ChooseHours int                                 `json:"chooseHours"    dc:"选择的小时数"`
	MemberId    int                                 `json:"memberId"    dc:"会员ID"`
	PayInfo     *OrderPayInfoModel                  `json:"PayInfo"     dc:"支付信息"`
}

type BoxTypeInfoItem struct {
	Id     int    `json:"id"               dc:"ID"`
	Name   string `json:"name"             dc:"名称"`
	NameZh string `json:"nameZh"             dc:"名称_简体中文"`
	NameEn string `json:"nameEn"             dc:"名称_英文"`
	NameJa string `json:"nameJa"             dc:"名称_日语"`
	NameKo string `json:"nameKo"             dc:"名称_韩语"`
	NameTw string `json:"nameTw"             dc:"名称_繁体中文"`
	Price  int    `json:"price"  dc:"价格"`
}

type PrePayInfoInp struct {
	PreOrderSn string `json:"preOrderSn"    dc:"预订单号"`
	BoxTypeId  int    `json:"boxTypeId"      dc:"格口类型ID"`
	Hours      int    `json:"hours"      dc:"租赁时长（小时）"`
	IsBalance  int    `json:"isBalance"     dc:"是否使用余额支付【1启用，2禁用】"`
	CouponId   int    `json:"couponId"      dc:"使用的优惠券ID"`
}

type OrderPayInfoModel struct {
	PreOrderSn    string  `json:"preOrderSn"    dc:"预订单号"`
	PayModel      int     `json:"payModel"      dc:"支付方式-1纯余额支付-2余额加外部支付-3纯外部支付"`
	MemberBalance float64 `json:"memberBalance" dc:"用户当前余额"`
	AllAmount     float64 `json:"allAmount"     dc:"订单总金额"`
	Score         float64 `json:"score"         dc:"订单可用积分上限"`
	Balance       struct {
		BalanceAmount     float64 `json:"amount"        dc:"余额支付金额"`
		BalancePayOrderSn string  `json:"payOrderSn"    dc:"余额支付订单号"`
		BalanceConfig     *BalanceConfig
	}
	ThirdPay struct {
		ThirdAmount     float64   `json:"amount"     dc:"外部支付金额"`
		ThirdPayOrderSn string    `json:"payOrderSn" dc:"外部支付订单号"`
		ThirdConfig     *gvar.Var `json:"config"     dc:"外部支付配置"`
	}
	Coupon struct {
		CouponId         int     `json:"couponId"      dc:"优惠券ID"`
		CouponAmount     float64 `json:"couponAmount"  dc:"优惠券抵用金额"`
		CouponName       string  `json:"couponName"    dc:"优惠券名称"`
		CouponPayOrderSn string  `json:"couponPayOrderSn" dc:"优惠券支付订单号"`
	}
}

type BalanceConfig struct {
	ScenePayRate float64 `json:"scenePayRate"  dc:"支付场景费率"`
	Level        int     `json:"level"         dc:"会员等级"`
	LevelName    string  `json:"levelName"     dc:"会员等级名"`
	ExchangeRate float64 `json:"exchangeRate"  dc:"汇率"`
	IsPayOpen    string  `json:"isPayOpen"     dc:"是否积分抵扣"`
}
