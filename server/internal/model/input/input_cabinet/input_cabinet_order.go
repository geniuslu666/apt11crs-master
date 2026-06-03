package input_cabinet

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/util/gmeta"

	"github.com/gogf/gf/v2/os/gtime"
)

type UpdateOrderStatusInp struct {
	OrderSn   string `json:"orderSn" dc:"订单号"`
	EventType string `json:"eventType" dc:"事件类型"`
}

type OrderRefundInp struct {
	OrderSn string `json:"orderSn" dc:"订单号"`
}

// OrderAppListInp 获取订单列表
type OrderAppListInp struct {
	input_form.PageReq
	MemberId    int    `json:"memberId"      dc:"用户"`
	OrderStatus string `p:"orderStatus" v:"in:WAIT_PAY,ING,DONE,CANCEL#order_status_format_error" dc:"订单状态【WAIT_PAY、待支付 ING、进行中  DONE、已完成  CANCEL、已取消】"`
}

func (in *OrderAppListInp) Filter(ctx context.Context) (err error) {
	return
}

type OrderAppListModel struct {
	Id             int         `json:"id" dc:"订单ID"`
	OrderStatus    string      `json:"orderStatus" dc:"订单状态 WAIT_PAY、待支付 ING、进行中 CANCEL、已取消  DONE、已完成 OVERTIME、已超时"`
	OrderSn        string      `json:"orderSn" dc:"订单号"`
	Pin            string      `json:"pin" dc:"pin码"`
	BoxNo          string      `json:"boxNo" dc:"格口编号"`
	BoxAlias       string      `json:"boxAlias" dc:"格口别称"`
	Address        string      `json:"address" dc:"地址"`
	OrderAmount    float64     `json:"orderAmount" dc:"订单金额（包含基础费用+超时费用）"`
	OvertimeAmount float64     `json:"overtimeAmount" dc:"超时费用"`
	UsedTime       int         `json:"useTime" dc:"使用时长（秒）"`
	GraceSeconds   int         `json:"graceSeconds"          dc:"宽限期设置时长（秒）"`
	GraceCountdown int         `json:"graceCountdown" dc:"宽限期剩余秒数"`
	GraceEndTime   *gtime.Time `json:"graceEndTime"          dc:"宽限期结束时间"`
	StartTime      *gtime.Time `json:"startTime" dc:"租赁的开始时间"`
	EndTime        *gtime.Time `json:"endTime" dc:"租赁的结束时间"`
	CabinetName    string      `json:"cabinetName"           dc:"储物柜名称"`
	MchBranchName  string      `json:"mchBranchName"         dc:"网点名称"`
	IsFx           string      `json:"isFx" dc:"是否是分销订单   Y   是    N   不是"`
}

type OrderAppViewInp struct {
	OrderSn string `json:"orderSn" dc:"订单号"`
}

func (in *OrderAppViewInp) Filter(ctx context.Context) (err error) {
	return
}

type OrderAppViewModel struct {
	Id             int                  `json:"id"           dc:"订单ID"`
	OrderSn        string               `json:"orderSn"      dc:"订单号"`
	OrderStatus    string               `json:"orderStatus" dc:"订单状态 WAIT_PAY、待支付 ING、进行中 CANCEL、已取消  DONE、已完成 OVERTIME、已超时"`
	Pin            string               `json:"pin" dc:"pin码"`
	BoxNo          string               `json:"boxNo" dc:"格口编号"`
	BoxAlias       string               `json:"boxAlias" dc:"格口别称"`
	BoxTypeName    string               `json:"boxTypeName"           dc:"格口类型名称"`
	Address        string               `json:"address" dc:"地址"`
	CabinetName    string               `json:"cabinetName" dc:"储物柜名称"`
	StartTime      *gtime.Time          `json:"startTime" dc:"开始时间"`
	BuyHours       int                  `json:"buyHours"              dc:"租赁时长（小时）"`
	UsedTime       int                  `json:"useTime" dc:"使用时长（秒）"`
	OvertimeTime   int                  `json:"overtimeTime" dc:"超时时长（秒）"`
	GraceSeconds   int                  `json:"graceSeconds"          dc:"宽限期设置时长（秒）"`
	GraceCountdown int                  `json:"graceCountdown" dc:"宽限期剩余秒数"`
	GraceEndTime   *gtime.Time          `json:"graceEndTime"          dc:"宽限期结束时间"`
	BaseAmount     float64              `json:"baseAmount"           dc:"租赁时间内金额（基础费用）"`
	OvertimeAmount float64              `json:"overtimeAmount" dc:"超时费用"`
	BoxTypePrice   int                  `json:"boxTypePrice"          dc:"当前选择的格口单价（每小时）"`
	OpenLog        []*OrderOpenLogModel `json:"openLog"          dc:"开门记录"`
	MchBranchName  string               `json:"mchBranchName" dc:"网点名称"`
	MchBranchLat   string               `json:"mchBranchLat" dc:"网点纬度"`
	MchBranchLgt   string               `json:"mchBranchLgt" dc:"网点经度"`
	IsFx           string               `json:"isFx" dc:"是否是分销订单   Y   是    N   否"`
}

type OrderOpenLogModel struct {
	Type      int    `json:"type"           dc:"类型"`
	Timestamp int    `json:"timestamp"      dc:"时间戳"`
	Datetime  string `json:"datetime" dc:"日期时间"`
}

type PayOvertimeInfoInp struct {
	OrderSn   string `json:"orderSn" v:"required#订单号未知" dc:"订单号"`
	IsBalance int    `json:"isBalance"     dc:"是否使用余额支付【1启用，2禁用】"`
}

type PayOvertimeInfoModel struct {
	OrderSn      string             `json:"orderSn"   dc:"订单号"`
	MinHours     int                `json:"minHours" dc:"最小预定小时数"`
	BoxNo        string             `json:"boxNo" dc:"格口编号"`
	BoxAlias     string             `json:"boxAlias" dc:"格口别称"`
	Address      string             `json:"address" dc:"地址"`
	CabinetName  string             `json:"cabinetName" dc:"储物柜名称"`
	StartTime    *gtime.Time        `json:"startTime" dc:"开始时间"`
	EndTime      *gtime.Time        `json:"endTime" dc:"结束时间"`
	UsedTime     int                `json:"useTime" dc:"使用时长（秒）"`
	OvertimeTime int                `json:"overtimeTime" dc:"超时时长（秒）"`
	BoxTypePrice int                `json:"boxTypePrice"          dc:"当前选择的格口单价（每小时）"`
	BoxTypeInfo  []*BoxTypeInfoItem `json:"boxTypeInfo"          dc:"收费标准"`
	MchBranchLat string             `json:"mchBranchLat" dc:"网点纬度"`
	MchBranchLgt string             `json:"mchBranchLgt" dc:"网点经度"`
	PayInfo      struct {
		PayModel       int     `json:"payModel"      dc:"支付方式-1纯余额支付-2余额加外部支付-3纯外部支付"`
		MemberBalance  float64 `json:"memberBalance" dc:"用户当前余额"`
		OvertimeAmount float64 `json:"allAmount"     dc:"超时金额"`
		Score          float64 `json:"score"         dc:"订单可用积分上限"`
		Balance        struct {
			BalanceAmount     float64        `json:"amount"        dc:"余额支付金额"`
			BalancePayOrderSn string         `json:"payOrderSn"    dc:"余额支付订单号"`
			BalanceConfig     *BalanceConfig `json:"balanceConfig" dc:"余额计算配置"`
		}
		ThirdPay struct {
			ThirdAmount     float64 `json:"amount"     dc:"余额支付金额"`
			ThirdPayOrderSn string  `json:"payOrderSn" dc:"余额支付订单号"`
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
	} `json:"PayInfo"     dc:"支付信息"`
}

type PayOvertimeInp struct {
	OrderSn   string `json:"orderSn" v:"required#订单号未知" dc:"订单号"`
	IsBalance int    `json:"isBalance"     dc:"是否使用余额支付【1启用，2禁用】"`
}

type PayOvertimeModel struct {
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
	CreateOrderTime string `json:"createOrderTime" dc:"创建订单时间"`
	Countdown       int    `json:"countdown" dc:"订单支付倒计时秒"`
}

// OrderListInp 获取订单列表-admin
type OrderListInp struct {
	input_form.PageReq
	OrderSn       string        `json:"orderSn"       dc:"订单编号"`
	MchBranchName string        `json:"mchBranchName" dc:"网点名称"`
	CabinetName   string        `json:"cabinetName"   dc:"储物柜名称"`
	CreatedAt     []*gtime.Time `json:"createdAt"     dc:"创建时间"`
	OrderStatus   string        `json:"orderStatus" v:"in:WAIT_PAY,ING,DONE,CANCEL,GRACE,OVERTIME,HAVE_PAID#order_status_format_error" dc:"订单状态【WAIT_PAY-待支付, ING-进行中,DONE-已完成,CANCEL-已取消,GRACE-宽限期,OVERTIME-超时,HAVE_PAID-已支付】"`
	BoxNo         string        `json:"boxNo"         dc:"格口编号"`
	MemberSearch  string        `json:"memberSearch"  dc:"会员信息"`
	CreateSort    string        `json:"createSort"    dc:"下单时间排序"`
	MemberId      uint          `json:"memberId"      dc:"用户ID"`
}

func (in *OrderListInp) Filter(ctx context.Context) (err error) {
	return
}

type OrderListModel struct {
	Id                  int64       `json:"id"                     dc:""`
	OrderSn             string      `json:"orderSn"                dc:"订单编号"`
	OutOrderSn          string      `json:"outOrderSn"             dc:"三方订单号"`
	CabinetId           int         `json:"cabinetId"              dc:"储物柜ID"`
	CabinetName         string      `json:"cabinetName"            dc:"储物柜名称"`
	CityId              int         `json:"cityId"                 dc:"所属城市ID"`
	CityName            string      `json:"cityName"               dc:"所属城市名"`
	MchId               int         `json:"mchId"                  dc:"运营商ID"`
	MchName             string      `json:"mchName"                dc:"运营商名称"`
	MchBranchId         int         `json:"mchBranchId"            dc:"网点ID"`
	MchBranchName       string      `json:"mchBranchName"          dc:"网点名称"`
	Address             string      `json:"address"                dc:"地址"`
	BoxTypeJson         *gjson.Json `json:"boxTypeJson"            dc:"格口类型列表数据"`
	BoxTypeId           int         `json:"boxTypeId"              dc:"格口类型ID"`
	BoxTypeName         string      `json:"boxTypeName"            dc:"格口类型名称"`
	BoxTypePrice        int         `json:"boxTypePrice"           dc:"格口类型单价"`
	BoxId               int         `json:"boxId"                  dc:"格口ID"`
	BoxNo               string      `json:"boxNo"                  dc:"格口编号"`
	BoxAlias            string      `json:"boxAlias"               dc:"格口别名"`
	OrderFirstFeeRate   float64     `json:"orderFirstFeeRate"      dc:"首次下单优惠"`
	BuyHours            int         `json:"buyHours"               dc:"购买的小时数"`
	Pin                 string      `json:"pin"                    dc:"取件码(4位数字)"`
	MemberId            uint        `json:"memberId"               dc:"用户ID"`
	OrderAmount         float64     `json:"orderAmount"            dc:"订单金额"`
	BaseAmount          float64     `json:"baseAmount"             dc:"租赁时间内金额"`
	CouponAmount        float64     `json:"couponAmount"           dc:"优惠券抵扣金额"`
	BalAmount           float64     `json:"balAmount"              dc:"积分抵扣金额"`
	OvertimeSecs        int         `json:"overtimeSecs"           dc:"超时时长(秒)"`
	OvertimeHours       int         `json:"overtimeHours"          dc:"超时时长(小时)"`
	OvertimeFee         int         `json:"overtimeFee"            dc:"超时费用(日元)"`
	OvertimePayTime     *gtime.Time `json:"overtimePayTime"        dc:"超时费支付时间"`
	OvertimePayStatus   string      `json:"overtimePayStatus"      dc:"超时费付款状态"`
	PayStep             string      `json:"payStep"                dc:"支付流程"`
	PayModel            int         `json:"payModel"               dc:"1、余额支付 2、组合支付 3、纯外部支付"`
	PayTime             *gtime.Time `json:"payTime"                dc:"支付时间"`
	PayStatus           string      `json:"payStatus"              dc:"订单付款状态"`
	OrderStatus         string      `json:"orderStatus"            dc:"订单状态"`
	ExpirationTime      int         `json:"expirationTime"         dc:"订单过期时间"`
	CancelTime          *gtime.Time `json:"cancelTime"             dc:"取消时间"`
	StartTime           *gtime.Time `json:"startTime"              dc:"订单开始时间"`
	EndTime             *gtime.Time `json:"endTime"                dc:"订单结束时间"`
	FinishTime          *gtime.Time `json:"finishTime"             dc:"完成时间"`
	RefundStatus        string      `json:"refundStatus"           dc:"退款状态   WAIT 未退款   PART  部分退款 DONE 全部退款"`
	RefundFee           float64     `json:"refundFee"              dc:"退款手续费"`
	RefundTime          *gtime.Time `json:"refundTime"             dc:"退款时间"`
	RefundAmount        float64     `json:"refundAmount"           dc:"已退款总金额"`
	RefundBalAmount     float64     `json:"refundBalAmount"        dc:"已退款积分"`
	RefundCouponAmount  float64     `json:"refundCouponAmount"     dc:"已退款优惠券"`
	RefundReason        string      `json:"refundReason"           dc:"退款原因"`
	IsAbnormal          int         `json:"isAbnormal"             dc:"请求下单接口是否异常"`
	PayOvertimeAbnormal int         `json:"payOvertimeAbnormal"    dc:"请求支付超时费接口是否异常"`
	CreatedAt           *gtime.Time `json:"createdAt"              dc:""`
	UpdatedAt           *gtime.Time `json:"updatedAt"              dc:""`
	UsedTime            int         `json:"useTime" dc:"使用时长（秒）"`
	MemberDeleted       bool        `json:"memberDeleted" dc:"会员是否已删除"`
	MemberInfo          *struct {
		gmeta.Meta `orm:"table:hg_pms_member"`
		Id         int    `json:"id"             dc:""`
		MemberNo   string `json:"memberNo"       dc:"会员号"`
		Avatar     string `json:"avatar"         dc:"会员头像"`
		FullName   string `json:"fullName"       dc:"会员姓名"`
	} `json:"memberInfo" orm:"with:id=member_id" dc:"会员信息"`
	TransactionDetail []*struct {
		gmeta.Meta       `orm:"table:hg_pms_transaction"`
		Id               int         `json:"id"               orm:"id"                 description:"主键"`
		OrderSn          string      `json:"orderSn"          orm:"order_sn"           description:"订单号"`
		TransactionSn    string      `json:"transactionSn"    orm:"transaction_sn"     description:"支付流水号"`
		PaymentRequestId string      `json:"paymentRequestId" orm:"payment_request_id" description:"第三方支付流水号"`
		PayChannel       string      `json:"payChannel"       orm:"pay_channel"        description:"SYSTEM 系统积分  PAYCLOUD   paycloud第三方支付平台"`
		PayType          string      `json:"payType"          orm:"pay_type"           description:"支付方式   BAL 余额"`
		Amount           float64     `json:"amount"           orm:"amount"             description:"总金额"`
		PayParams        string      `json:"payParams"        orm:"pay_params"         description:"支付参数"`
		PriceCurrency    string      `json:"priceCurrency"    orm:"price_currency"     description:"币种"`
		PayAmount        float64     `json:"payAmount"        orm:"pay_amount"         description:"支付金额"`
		PayStatus        string      `json:"payStatus"        orm:"pay_status"         description:"支付状态  WAIT 等待支付、DONE 完成支付、CANCEL 取消支付"`
		PayTime          *gtime.Time `json:"payTime"          orm:"pay_time"           description:"支付时间"`
		ExpiredTime      *gtime.Time `json:"expiredTime"      orm:"expired_time"       description:"过期时间"`
		RefundAmount     float64     `json:"refundAmount"     orm:"refund_amount"      description:"退款金额"`
		RefundStatus     string      `json:"refundStatus"     orm:"refund_status"      description:"退款状态   WAIT 未退款   PART  部分退款 DONE 全部退款"`
		CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"         description:"创建时间"`
		UpdatedAt        *gtime.Time `json:"updatedAt"        orm:"updated_at"         description:"更新时间"`
		DeletedAt        *gtime.Time `json:"deletedAt"        orm:"deleted_at"         description:"删除时间"`
	} `json:"transactionDetail" orm:"with:order_sn=order_sn" dc:"支付明细"`
	TransactionRefundDetail []*struct {
		gmeta.Meta    `orm:"table:hg_pms_transaction_refund"`
		Id            int         `json:"id"            orm:"id"             description:"主键"`
		OrderSn       string      `json:"orderSn"       orm:"order_sn"       description:"订单号"`
		TransactionSn string      `json:"transactionSn" orm:"transaction_sn" description:"支付流水号"`
		RefundType    string      `json:"refundType"    orm:"refund_type"    description:"'支付方式   BAL 余额'"`
		RefundSn      string      `json:"refundSn"      orm:"refund_sn"      description:"退款流水号"`
		TransNo       string      `json:"transNo"       orm:"trans_no"       description:"退款交易号"`
		RefundAmount  float64     `json:"refundAmount"  orm:"refund_amount"  description:"退款金额"`
		RefundTime    *gtime.Time `json:"refundTime"    orm:"refund_time"    description:"退款时间"`
		RefundStatus  string      `json:"refundStatus"  orm:"refund_status"  description:"退款状态"`
		CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:"创建时间"`
		UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:"更新时间"`
	} `json:"transactionRefundDetail" orm:"with:order_sn=order_sn, where:refund_status='DONE'" dc:"退款明细"`
}

type OrderViewInp struct {
	Id      int64  `json:"id" dc:"id"`
	OrderSn string `json:"orderSn" dc:"系统订单号"`
}

func (in *OrderViewInp) Filter(ctx context.Context) (err error) {
	return
}

type OrderViewModel struct {
	entity.CabinetOrder
	UsedTime          int  `json:"useTime" dc:"使用时长（秒）"`
	MemberDeleted     bool `json:"memberDeleted" dc:"会员是否已删除"`
	TransactionDetail []*struct {
		gmeta.Meta       `orm:"table:hg_pms_transaction"`
		Id               int         `json:"id"               orm:"id"                 description:"主键"`
		OrderSn          string      `json:"orderSn"          orm:"order_sn"           description:"订单号"`
		TransactionSn    string      `json:"transactionSn"    orm:"transaction_sn"     description:"支付流水号"`
		PaymentRequestId string      `json:"paymentRequestId" orm:"payment_request_id" description:"第三方支付流水号"`
		PayChannel       string      `json:"payChannel"       orm:"pay_channel"        description:"SYSTEM 系统积分  PAYCLOUD   paycloud第三方支付平台"`
		PayType          string      `json:"payType"          orm:"pay_type"           description:"支付方式   BAL 余额"`
		Amount           float64     `json:"amount"           orm:"amount"             description:"总金额"`
		PayParams        string      `json:"payParams"        orm:"pay_params"         description:"支付参数"`
		PriceCurrency    string      `json:"priceCurrency"    orm:"price_currency"     description:"币种"`
		PayAmount        float64     `json:"payAmount"        orm:"pay_amount"         description:"支付金额"`
		PayStatus        string      `json:"payStatus"        orm:"pay_status"         description:"支付状态  WAIT 等待支付、DONE 完成支付、CANCEL 取消支付"`
		PayTime          *gtime.Time `json:"payTime"          orm:"pay_time"           description:"支付时间"`
		ExpiredTime      *gtime.Time `json:"expiredTime"      orm:"expired_time"       description:"过期时间"`
		RefundAmount     float64     `json:"refundAmount"     orm:"refund_amount"      description:"退款金额"`
		RefundStatus     string      `json:"refundStatus"     orm:"refund_status"      description:"退款状态   WAIT 未退款   PART  部分退款 DONE 全部退款"`
		CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"         description:"创建时间"`
		UpdatedAt        *gtime.Time `json:"updatedAt"        orm:"updated_at"         description:"更新时间"`
		DeletedAt        *gtime.Time `json:"deletedAt"        orm:"deleted_at"         description:"删除时间"`
	} `json:"transactionDetail" orm:"with:order_sn=order_sn" dc:"支付明细"`
	TransactionRefundDetail []*struct {
		gmeta.Meta    `orm:"table:hg_pms_transaction_refund"`
		Id            int         `json:"id"            orm:"id"             description:"主键"`
		OrderSn       string      `json:"orderSn"       orm:"order_sn"       description:"订单号"`
		TransactionSn string      `json:"transactionSn" orm:"transaction_sn" description:"支付流水号"`
		RefundType    string      `json:"refundType"    orm:"refund_type"    description:"'支付方式   BAL 余额'"`
		RefundSn      string      `json:"refundSn"      orm:"refund_sn"      description:"退款流水号"`
		TransNo       string      `json:"transNo"       orm:"trans_no"       description:"退款交易号"`
		RefundAmount  float64     `json:"refundAmount"  orm:"refund_amount"  description:"退款金额"`
		RefundTime    *gtime.Time `json:"refundTime"    orm:"refund_time"    description:"退款时间"`
		RefundStatus  string      `json:"refundStatus"  orm:"refund_status"  description:"退款状态"`
		CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:"创建时间"`
		UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:"更新时间"`
	} `json:"transactionRefundDetail" orm:"with:order_sn=order_sn, where:refund_status='DONE'" dc:"退款明细"`
	LogList []*struct {
		gmeta.Meta  `orm:"table:hg_cabinet_order_log"`
		OrderId     int         `json:"orderId"     description:"订单ID"`
		ActionWay   string      `json:"actionWay"   description:"操作名"`
		Remark      string      `json:"remark"      description:"备注"`
		Images      string      `json:"images"      description:"图集"`
		OperateType string      `json:"operateType" description:"操作员类型"`
		OperateId   int         `json:"operateId"   description:"操作员ID"`
		OperateName string      `json:"operateName"      dc:"操作人姓名"`
		CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	} `json:"logList" orm:"with:order_id=id" dc:"订单日志"`
	MemberInfo *struct {
		gmeta.Meta `orm:"table:hg_pms_member"`
		Id         int    `json:"id"             dc:""`
		MemberNo   string `json:"memberNo"       dc:"会员号"`
		FullName   string `json:"fullName"       dc:"会员姓名"`
		Phone      string `json:"phone"          dc:"手机号"`
		PhoneArea  string `json:"phoneArea"      dc:"区号"`
		Mail       string `json:"mail"           dc:"邮箱"`
	} `json:"memberInfo" orm:"with:id=member_id" dc:"会员信息"`
	OpenLogList []*OrderOpenLogModel `json:"openLogList"          dc:"开门记录"`
}

// OrderExportInp 导出-admin
type OrderExportInp struct {
	OrderSn       string        `json:"orderSn"       dc:"订单编号"`
	MchBranchName string        `json:"mchBranchName" dc:"网点名称"`
	CreatedAt     []*gtime.Time `json:"createdAt"     dc:"创建时间"`
	OrderStatus   string        `json:"orderStatus" v:"in:WAIT_PAY,ING,DONE,CANCEL#order_status_format_error" dc:"订单状态【WAIT_PAY、待支付 ING、进行中  DONE、已完成  CANCEL、已取消】"`
	BoxNo         string        `json:"boxNo"         dc:"格口编号"`
	MemberSearch  string        `json:"memberSearch"  dc:"会员信息"`
}

func (in *OrderExportInp) Filter(ctx context.Context) (err error) {
	return
}

type OrderExportModel struct {
	OrderSn                 string      `json:"orderSn"                dc:"订单编号"`
	OutOrderSn              string      `json:"outOrderSn"             dc:"三方订单号"`
	MemberNo                string      `json:"memberNo"               dc:"下单会员ID"`
	MemberFullName          string      `json:"memberFullName"         dc:"会员姓名"`
	CabinetName             string      `json:"cabinetName"            dc:"储物柜名称"`
	CityName                string      `json:"cityName"               dc:"所属城市名"`
	MchBranchName           string      `json:"mchBranchName"          dc:"网点名称"`
	Address                 string      `json:"address"                dc:"地址"`
	BoxTypeName             string      `json:"boxTypeName"            dc:"格口类型名称"`
	BoxTypePrice            int         `json:"boxTypePrice"           dc:"格口类型单价"`
	BoxNo                   string      `json:"boxNo"                  dc:"格口编号"`
	BoxAlias                string      `json:"boxAlias"               dc:"格口别名"`
	Pin                     string      `json:"pin"                    dc:"取件码(4位数字)"`
	OrderAmount             float64     `json:"orderAmount"            dc:"订单金额"`
	BaseAmount              float64     `json:"baseAmount"             dc:"租赁时间内金额"`
	PayTime                 *gtime.Time `json:"payTime"                dc:"支付时间"`
	OvertimeSecs            int         `json:"overtimeSecs"           dc:"超时时长(秒)"`
	OvertimeHours           int         `json:"overtimeHours"          dc:"超时时长(小时)"`
	OvertimeFee             int         `json:"overtimeFee"            dc:"超时费用(日元)"`
	OvertimePayTime         *gtime.Time `json:"overtimePayTime"        dc:"超时费支付时间"`
	OvertimePayStatus       string      `json:"overtimePayStatus"      dc:"超时费付款状态"`
	PayStatus               string      `json:"payStatus"              dc:"订单付款状态"`
	OrderStatus             string      `json:"orderStatus"            dc:"订单状态"`
	StartTime               *gtime.Time `json:"startTime"              dc:"订单开始时间"`
	EndTime                 *gtime.Time `json:"endTime"                dc:"订单结束时间"`
	FinishTime              *gtime.Time `json:"finishTime"             dc:"完成时间"`
	RefundAmount            float64     `json:"refundAmount"           dc:"已退款总金额"`
	RefundBalAmount         float64     `json:"refundBalAmount"        dc:"已退款积分"`
	RefundCouponAmount      float64     `json:"refundCouponAmount"     dc:"已退款优惠券"`
	OrderTime               string      `json:"createdAt"              dc:"下单时间"`
	RefundTotalAmount       float64     `json:"refundTotalAmount" dc:"退款总金额"`
	PointsPayment           float64     `json:"pointsPayment" dc:"积分支付"`
	CouponPayment           float64     `json:"couponPayment" dc:"优惠券支付"`
	PaycloudWechatPay       float64     `json:"paycloudWechatPay" dc:"PaycloudWechat支付"`
	PaycloudAlipayPay       float64     `json:"paycloudAlipayPay" dc:"PaycloudAlipay+支付"`
	PaycloudCreditPay       float64     `json:"paycloudCreditPay" dc:"Paycloud信用卡支付"`
	PaypelPay               float64     `json:"paypelPay" dc:"Paypel支付"`
	PaypelCreditPay         float64     `json:"paypelCreditPay" dc:"Paypel信用卡支付"`
	StripeCreditPay         float64     `json:"stripeCreditPay" dc:"Stripe信用卡支付"`
	MlilifeWeChatMiniPay    float64     `json:"mlilifeWeChatMiniPay" dc:"WeChatMini支付"`
	PointsRefund            float64     `json:"pointsRefund" dc:"积分退款"`
	CouponRefund            float64     `json:"couponRefund" dc:"优惠券退款"`
	PaycloudWechatRefund    float64     `json:"paycloudWechatRefund" dc:"PaycloudWechat退款"`
	PaycloudAlipayRefund    float64     `json:"paycloudAlipayRefund" dc:"PaycloudAlipay+退款"`
	PaycloudCreditRefund    float64     `json:"paycloudCreditRefund" dc:"Paycloud信用卡退款"`
	PaypelCreditRefund      float64     `json:"paypelCreditRefund" dc:"Paypel信用卡退款"`
	PaypelRefund            float64     `json:"paypelRefund" dc:"Paypel信用卡退款"`
	StripeCreditRefund      float64     `json:"stripeCreditRefund" dc:"Stripe信用卡退款"`
	MlilifeWeChatMiniRefund float64     `json:"mlilifeWeChatMiniRefund" dc:"WeChatMini退款"`
	RefundTime              string      `json:"refundTime" dc:"退款/取消时间"`
}

// OrderExportListInp 导出列表
type OrderExportListInp struct {
	input_form.PageReq
}

func (in *OrderExportListInp) Filter(ctx context.Context) (err error) {
	return
}

type OrderExportListModel struct {
	entity.OrderExport
}

// OrderCompleteInp 订单完成
type OrderCompleteInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *OrderCompleteInp) Filter(ctx context.Context) (err error) {
	return
}

type OrderCompleteModel struct{}

// CabinetOrderRefundInp 订单退款
type CabinetOrderRefundInp struct {
	Id                int64   `json:"id" v:"required#id不能为空" dc:"id"`
	RefundMoney       float64 `json:"refundMoney"              dc:"退款金额"`
	AdminCancelReason string  `json:"adminCancelReason"        dc:"退款原因"`
}

func (in *CabinetOrderRefundInp) Filter(ctx context.Context) (err error) {
	return
}

type CabinetOrderRefundModel struct{}
