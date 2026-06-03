package input_terminal

import (
	"APT/internal/model/input/input_form"
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// VerifyListInp 获取核销记录列表
type VerifyListInp struct {
	input_form.PageReq
	TerminalId   int `json:"terminalId"     dc:"终端ID"`
	MchId        int `json:"mchId"         dc:"商户ID"`
	StoreId      int `json:"storeId"        dc:"门店ID"`
	RestaurantId int `json:"restaurantId"   dc:"餐厅ID"`
}

func (in *VerifyListInp) Filter(ctx context.Context) (err error) {
	return
}

type VerifyListModel struct {
	Id             int    `json:"id"             dc:"核销记录ID"`
	VerifyType     string `json:"verifyType"     dc:"核销类型"`
	VerifyMemberId int    `json:"verifyMemberId" dc:"核销用户ID"`
	TerminalId     int    `json:"terminalId"     dc:"终端ID"`
	FoodOrderId    int    `json:"foodOrderId"    dc:"餐厅订单ID"`
	MemberCouponId int    `json:"memberCouponId" dc:"用户礼品券ID"`
	CouponMchName  string `json:"couponMchName"  dc:"核销商品名"`
	MemberInfo     *struct {
		gmeta.Meta `orm:"table:hg_pms_member"`
		Id         int    `json:"id"             dc:""`
		MemberNo   string `json:"memberNo"        dc:"会员号"`
	} `json:"memberInfo" orm:"with:id=verify_member_id" dc:"会员信息"`
	TerminalInfo *struct {
		gmeta.Meta   `orm:"table:hg_sys_terminal"`
		Id           int    `json:"id"             dc:""`
		TerminalName string `json:"terminalName" dc:"终端名称"`
		Sn           string `json:"sn"           dc:"终端编号"`
	} `json:"terminalInfo" orm:"with:id=terminal_id" dc:"终端信息"`
	MemberCouponInfo *struct {
		gmeta.Meta `orm:"table:hg_th_member_coupon"`
		Id         int  `json:"id"             dc:""`
		CouponId   uint `json:"couponId"                 dc:"券ID"`
		CouponInfo *struct {
			gmeta.Meta `orm:"table:hg_th_coupon"`
			Id         int    `json:"id"             dc:""`
			CouponName string `json:"couponName"       dc:"券名称"`
		} `json:"couponInfo" orm:"with:id=coupon_id" dc:"券信息"`
	} `json:"memberCouponInfo" orm:"with:id=member_coupon_id" dc:"会员券领取信息"`
	FoodOrderInfo *struct {
		gmeta.Meta         `orm:"table:hg_food_order"`
		Id                 int    `json:"id"             dc:""`
		OrderSn            string `json:"orderSn"                dc:"订单编号"`
		BookingName        string `json:"bookingName"             dc:"预定人姓名"`
		PhoneArea          string `json:"phoneArea"              dc:"手机区号"`
		BookingMobile      string `json:"bookingMobile"          dc:"预定人手机"`
		GoodsId            uint   `json:"goodsId"                 dc:"套餐ID"`
		FoodOrderGoodsInfo *struct {
			gmeta.Meta `orm:"table:hg_food_goods"`
			Id         int    `json:"id"             dc:""`
			GoodsName  string `json:"goodsName"       dc:"套餐名称"`
		} `json:"foodOrderGoodsInfo" orm:"with:id=goods_id" dc:"套餐信息"`
	} `json:"foodOrderInfo" orm:"with:id=food_order_id" dc:"餐厅订单信息"`
}

type VerifyLogViewInp struct {
	Id int `json:"id" v:"required#核销记录ID不能为空" dc:"核销记录ID"`
}

func (in *VerifyLogViewInp) Filter(ctx context.Context) (err error) {
	return
}

type VerifyLogViewModel struct {
	Id            int            `json:"id"             dc:"核销记录ID"`
	VerifyType    string         `json:"verifyType"     dc:"核销类型"`
	VerifyInfo    *VerifyInfo    `json:"verifyInfo"     dc:"核销信息"`
	CouponInfo    *CouponInfo    `json:"couponInfo"     dc:"券信息"`
	FoodOrderInfo *FoodOrderInfo `json:"foodOrderInfo"     dc:"餐厅订单信息"`
}

type VerifyInfo struct {
	VerifyTime  *gtime.Time `json:"verifyTime"    dc:"核销时间"`
	VerifyStore string      `json:"verifyStore"    dc:"核销商家"`
	Account     string      `json:"account"     dc:"核销人员"`
}

type CouponInfo struct {
	CouponMchName string      `json:"couponMchName"  dc:"核销商品名"`
	CouponName    string      `json:"couponName"       dc:"券名称"`
	StartTime     *gtime.Time `json:"startTime"     dc:"有效期开始时间"`
	EndTime       *gtime.Time `json:"endTime"       dc:"有效期结束时间"`
	MemberNo      string      `json:"memberNo"        dc:"会员号"`
	StoreAddress  string      `json:"storeAddress"        dc:"门店地址"`
}

type FoodOrderInfo struct {
	OrderSn           string      `json:"orderSn"                dc:"订单编号"`
	BookingName       string      `json:"bookingName"             dc:"预定人姓名"`
	PhoneArea         string      `json:"phoneArea"              dc:"手机区号"`
	BookingMobile     string      `json:"bookingMobile"          dc:"预定人手机"`
	GoodsNum          int         `json:"goodsNum"                dc:"套餐数量"`
	BookDate          string      `json:"bookDate"                dc:"预定日期"`
	BookTime          string      `json:"bookTime"                dc:"预定时间"`
	RestaurantName    string      `json:"restaurantName"       dc:"餐厅名称"`
	RestaurantAddress string      `json:"restaurantAddress"       dc:"餐厅地址"`
	MemberMessage     string      `json:"memberMessage"       dc:"备注"`
	MemberMessageJa   string      `json:"memberMessageJa"       dc:"备注-日语"`
	GoodsName         string      `json:"goodsName"       dc:"套餐名称"`
	OrderAmount       float64     `json:"orderAmount"            dc:"订单金额"`
	CreatedTime       *gtime.Time `json:"createdTime"             dc:"下单时间"`
}

type VerifyFoodOrderInfo struct {
	gmeta.Meta              `orm:"table:hg_food_order"`
	Id                      int         `json:"id"             dc:""`
	OrderSn                 string      `json:"orderSn"                dc:"订单编号"`
	BookingName             string      `json:"bookingName"             dc:"预定人姓名"`
	PhoneArea               string      `json:"phoneArea"              dc:"手机区号"`
	BookingMobile           string      `json:"bookingMobile"          dc:"预定人手机"`
	RestaurantId            int         `json:"restaurantId"            dc:"餐厅ID"`
	GoodsId                 uint        `json:"goodsId"                 dc:"套餐ID"`
	BookDate                string      `json:"bookDate"                dc:"预定日期"`
	BookTime                string      `json:"bookTime"                dc:"预定时间"`
	GoodsNum                int         `json:"goodsNum"                dc:"套餐数量"`
	MemberMessage           string      `json:"memberMessage"          dc:"购买人留言信息"`
	MemberMessageJa         string      `json:"memberMessageJa"         dc:"购买人留言信息日语版"`
	OrderAmount             float64     `json:"orderAmount"            dc:"订单金额"`
	CreatedAt               *gtime.Time `json:"createdAt"             dc:"下单时间"`
	VerifyStatus            string      `json:"verifyStatus"           dc:"订单核销状态"`
	BookDatetime            *gtime.Time `json:"bookDatetime"            dc:"预定日期时间"`
	OrderStatus             string      `json:"orderStatus"            dc:"订单付款状态"`
	BookingStatus           string      `json:"bookingStatus"          dc:"订单预定状态"`
	FoodOrderRestaurantInfo *struct {
		gmeta.Meta    `orm:"table:hg_food_restaurant"`
		Id            int    `json:"id"             dc:""`
		Name          string `json:"name"       dc:"餐厅名称"`
		DetailAddress string `json:"detailAddress"         dc:"详细地址"`
	} `json:"foodOrderRestaurantInfo" orm:"with:id=restaurant_id" dc:"餐厅信息"`
	FoodOrderGoodsInfo *struct {
		gmeta.Meta `orm:"table:hg_food_goods"`
		Id         int    `json:"id"             dc:""`
		GoodsName  string `json:"goodsName"       dc:"套餐名称"`
	} `json:"foodOrderGoodsInfo" orm:"with:id=goods_id" dc:"套餐信息"`
}

type VerifyMemberCouponInfo struct {
	gmeta.Meta `orm:"table:hg_th_member_coupon"`
	Id         int         `json:"id"             dc:""`
	MemberId   uint        `json:"memberId"                 dc:"会员ID"`
	CouponId   uint        `json:"couponId"                 dc:"券ID"`
	StartTime  *gtime.Time `json:"startTime"     dc:"有效期开始时间"`
	EndTime    *gtime.Time `json:"endTime"       dc:"有效期结束时间"`
	State      int         `json:"state"        dc:"状态 1待生效 2未使用 3已核销 4已过期  5已失效"`
	MemberInfo *struct {
		gmeta.Meta `orm:"table:hg_pms_member"`
		Id         int    `json:"id"             dc:""`
		MemberNo   string `json:"memberNo"        dc:"会员号"`
	} `json:"memberInfo" orm:"with:id=member_id" dc:"会员信息"`
	CouponInfo *struct {
		gmeta.Meta `orm:"table:hg_th_coupon"`
		Id         int    `json:"id"             dc:""`
		CouponName string `json:"couponName"       dc:"券名称"`
		UseStatus  int    `json:"useStatus"      dc:"使用状态（1开始使用  2停止使用）"`
	} `json:"couponInfo" orm:"with:id=coupon_id" dc:"券信息"`
}

type CodeViewInp struct {
	Code string `json:"code" v:"required#券码不能为空" dc:"券码"`
}

func (in *CodeViewInp) Filter(ctx context.Context) (err error) {
	return
}

type CodeViewModel struct {
	VerifyType    string          `json:"verifyType"     dc:"核销类型"`
	CouponInfo    *CodeCouponInfo `json:"couponInfo"     dc:"券信息"`
	FoodOrderInfo *FoodOrderInfo  `json:"foodOrderInfo"     dc:"餐厅订单信息"`
}

type CodeCouponInfo struct {
	CouponMchName string      `json:"couponMchName"  dc:"核销商品名"`
	CouponName    string      `json:"couponName"       dc:"券名称"`
	StartTime     *gtime.Time `json:"startTime"     dc:"有效期开始时间"`
	EndTime       *gtime.Time `json:"endTime"       dc:"有效期结束时间"`
	MemberNo      string      `json:"memberNo"        dc:"会员号"`
	MchName       string      `json:"mchName"        dc:"商户名称"`
	StoreName     string      `json:"storeName"        dc:"门店名称"`
	StoreAddress  string      `json:"storeAddress"        dc:"门店地址"`
}

type CodeVerifyInp struct {
	Code string `json:"code" v:"required#coupon_code_cannot_be_empty" dc:"券码"`
}

func (in *CodeVerifyInp) Filter(ctx context.Context) (err error) {
	return
}

type CodeVerifyModel struct {
	VerifyType    string          `json:"verifyType"     dc:"核销类型"`
	VerifyTime    *gtime.Time     `json:"verifyTime"     dc:"核销时间"`
	CouponInfo    *CodeCouponInfo `json:"couponInfo"     dc:"券信息"`
	FoodOrderInfo *FoodOrderInfo  `json:"foodOrderInfo"     dc:"餐厅订单信息"`
}

type FoodOrderVerifyInp struct {
	OrderId    int `json:"orderId" v:"required#order_id_cannot_be_empty" dc:"订单ID"`
	TerminalId int `json:"terminalId" v:"required#terminal_id_cannot_be_empty" dc:"终端ID"`
}

type MemberCouponVerifyInp struct {
	MemberCouponId int    `json:"memberCouponId" v:"required#member_coupon_id_cannot_be_empty" dc:"券ID"`
	TerminalId     int    `json:"terminalId" v:"required#terminal_id_cannot_be_empty" dc:"终端ID"`
	CouponMchName  string `json:"couponMchName" dc:"核销商品名"`
}
