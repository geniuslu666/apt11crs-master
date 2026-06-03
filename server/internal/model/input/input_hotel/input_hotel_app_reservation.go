// Package sysin

package input_hotel

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/utility/validate"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// PmsAppReservationUpdateFields 修改系统入住订单字段过滤
type PmsAppReservationUpdateFields struct {
	Uid           string      `json:"uid"           dc:"房间预订ID"`
	MemberId      int         `json:"memberId"      dc:"用户ID"`
	Puid          string      `json:"puid"          dc:"物业ID"`
	OrderSn       string      `json:"orderSn"       dc:"系统订单号"`
	TicketId      string      `json:"ticketId"      dc:"票号"`
	RoomType      string      `json:"roomType"      dc:"房型信息"`
	RoomUnit      string      `json:"roomUnit"      dc:"房间单元"`
	GuestProfile  string      `json:"guestProfile"  dc:"预订人信息"`
	CheckinDate   *gtime.Time `json:"checkinDate"   dc:"入住日期"`
	CheckoutDate  *gtime.Time `json:"checkoutDate"  dc:"退房日期"`
	CheckinTime   string      `json:"checkinTime"   dc:"入住时间"`
	CheckoutTime  string      `json:"checkoutTime"  dc:"退房时间"`
	Status        string      `json:"status"        dc:"预订状态"`
	CheckinStatus string      `json:"checkinStatus" dc:"入住状态"`
	OrderStatus   string      `json:"orderStatus"   dc:"支付状态"`
	AdultCount    int         `json:"adultCount"    dc:"成人数量"`
	ChildCount    int         `json:"childCount"    dc:"儿童数量"`
	BookingFee    float64     `json:"bookingFee"    dc:"预订费"`
	ChannelFee    float64     `json:"channelFee"    dc:"渠道费"`
}

// PmsAppReservationInsertFields 新增系统入住订单字段过滤
type PmsAppReservationInsertFields struct {
	Uid           string      `json:"uid"           dc:"房间预订ID"`
	MemberId      int         `json:"memberId"      dc:"用户ID"`
	Puid          string      `json:"puid"          dc:"物业ID"`
	OrderSn       string      `json:"orderSn"       dc:"系统订单号"`
	TicketId      string      `json:"ticketId"      dc:"票号"`
	RoomType      string      `json:"roomType"      dc:"房型信息"`
	RoomUnit      string      `json:"roomUnit"      dc:"房间单元"`
	GuestProfile  string      `json:"guestProfile"  dc:"预订人信息"`
	CheckinDate   *gtime.Time `json:"checkinDate"   dc:"入住日期"`
	CheckoutDate  *gtime.Time `json:"checkoutDate"  dc:"退房日期"`
	CheckinTime   string      `json:"checkinTime"   dc:"入住时间"`
	CheckoutTime  string      `json:"checkoutTime"  dc:"退房时间"`
	Status        string      `json:"status"        dc:"预订状态"`
	CheckinStatus string      `json:"checkinStatus" dc:"入住状态"`
	OrderStatus   string      `json:"orderStatus"   dc:"支付状态"`
	AdultCount    int         `json:"adultCount"    dc:"成人数量"`
	ChildCount    int         `json:"childCount"    dc:"儿童数量"`
	BookingFee    float64     `json:"bookingFee"    dc:"预订费"`
	ChannelFee    float64     `json:"channelFee"    dc:"渠道费"`
}

// PmsAppReservationEditInp 修改/新增系统入住订单
type PmsAppReservationEditInp struct {
	entity.PmsAppReservation
}

func (in *PmsAppReservationEditInp) Filter(ctx context.Context) (err error) {
	// 验证支付状态
	if err := g.Validator().Rules("required").Data(in.OrderStatus).Messages("支付状态不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:WAIT_PAY,HAVE_PAID,CANCEL").Data(in.OrderStatus).Messages("支付状态值不正确").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type PmsAppReservationEditModel struct{}

// PmsAppReservationDeleteInp 删除系统入住订单
type PmsAppReservationDeleteInp struct {
	Id interface{} `json:"id" v:"required#主键不能为空" dc:"主键"`
}

func (in *PmsAppReservationDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsAppReservationDeleteModel struct{}

// PmsAppReservationViewInp 获取指定系统入住订单信息
type PmsAppReservationViewInp struct {
	Id         int    `json:"id" dc:"主键"`
	OrderSn    string `json:"ordersn" dc:"系统订单号"`
	OutOrderSn string `json:"outOrderSn" dc:"外部订单号"`
}

func (in *PmsAppReservationViewInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsAppReservationViewModel struct {
	gmeta.Meta `orm:"table:hg_pms_app_stay"`
	*entity.PmsAppStay
	MemberDeleted   bool `json:"memberDeleted" dc:"会员是否已删除"`
	ReferrerDeleted bool `json:"referrerDeleted" dc:"推荐人是否已删除"`
	PropertyDeleted bool `json:"propertyDeleted" dc:"物业是否已删除"`
	MemberDetail    *struct {
		gmeta.Meta `orm:"table:hg_pms_member"`
		Id         int    `json:"id"              orm:"id"                description:"主键"`
		MemberNo   string `json:"memberNo"        orm:"member_no"         description:"会员号"`
		FirstName  string `json:"firstName"       orm:"first_name"        description:"名"`
		LastName   string `json:"lastName"        orm:"last_name"         description:"姓"`
		FullName   string `json:"fullName"        orm:"full_name"         description:"全名"`
		Phone      string `json:"phone"           orm:"phone"             description:"手机号"`
		PhoneArea  string `json:"phoneArea"       orm:"phone_area"        description:"手机区号"`
		Mail       string `json:"mail"            orm:"mail"              description:"邮箱"`
	} `json:"memberDetail" orm:"with:id=member_id"`
	ReferrerDetail *struct {
		gmeta.Meta `orm:"table:hg_pms_member"`
		*entity.PmsMember
		Staff *struct {
			gmeta.Meta `orm:"table:hg_pms_staff"`
			*entity.PmsStaff
		} `json:"staff" orm:"with:id=staff_id"`
		Channel *struct {
			gmeta.Meta `orm:"table:hg_pms_channel"`
			*entity.PmsChannel
		} `json:"channel" orm:"with:id=channel_id"`
	} `json:"referrerDetail" orm:"with:id=referrer"`
	PropertyDetail *struct {
		gmeta.Meta `orm:"table:hg_pms_property"`
		*entity.PmsProperty
	} `json:"propertyDetail" orm:"with:uid=puid"`
	AppReservation []*struct {
		gmeta.Meta `orm:"table:hg_pms_app_reservation"`
		*entity.PmsAppReservation
		RoomTypeDetail *struct {
			gmeta.Meta `orm:"table:hg_pms_room_type"`
			*entity.PmsRoomType
		} `json:"roomTypeDetail" orm:"with:uid=room_type"  dc:"房型信息"`
		RoomUnitDetail *struct {
			gmeta.Meta `orm:"table:hg_pms_room_unit"`
			*entity.PmsRoomUnit
		} `json:"roomUnitDetail" orm:"with:uid=room_unit" dc:"房间信息"`
		ChargesDetail         []*entity.PmsCharge `json:"chargesDetail" orm:"with:uid=charges" dc:"费用明细"`
		ReservationChangeList []*struct {
			gmeta.Meta `orm:"table:hg_pms_app_reservation_change"`
			*entity.PmsAppReservationChange
		} `json:"reservationChangeList" orm:"with:order_id=id" dc:"入住订单变更信息"`
	} `json:"appReservation" orm:"with:out_order_sn=out_order_sn"`
	GuestProfileDetail *struct {
		gmeta.Meta `orm:"table:hg_pms_guest_profile"`
		*entity.PmsGuestProfile
	} `json:"guestProfileDetail" orm:"with:uid=booker"`
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
	} `json:"transactionRefundDetail" orm:"with:order_sn=order_sn" dc:"退款明细"`
	CancelOrder *entity.PmsAppCancelOrder `json:"cancelOrder" orm:"with:order_sn=order_sn" dc:"取消订单"`
	LogList     []*struct {
		gmeta.Meta  `orm:"table:hg_pms_app_stay_log"`
		OrderId     int         `json:"orderId"     description:"订单ID"`
		ActionWay   string      `json:"actionWay"   description:"操作名"`
		Remark      string      `json:"remark"      description:"备注"`
		Images      string      `json:"images"      description:"图集"`
		OperateType string      `json:"operateType" description:"操作员类型"`
		OperateId   int         `json:"operateId"   description:"操作员ID"`
		OperateName string      `json:"operateName"      dc:"操作人姓名"`
		CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	} `json:"logList" orm:"with:order_id=id" dc:"服务流程"`
}

// PmsAppReservationListInp 获取系统入住订单列表
type PmsAppReservationListInp struct {
	input_form.PageReq
	OrderSn      string   `json:"orderSn"       dc:"系统订单号"`
	OutOrderSn   string   `json:"outOrderSn"    dc:"外部订单号"`
	CheckinDate  []string `json:"checkinDate"  dc:"入住日期"`
	CheckoutDate []string `json:"checkoutDate"  dc:"退房日期"`
	RefundStatus string   `json:"status"        dc:"退款状态"`
	OrderStatus  string   `json:"orderStatus"   dc:"订单状态"`
	MemberId     int      `json:"member_id" dc:"会员ID"`
	MemberNo     string   `json:"memberNo"       dc:"会员编号"`
	CreatedAt    []string `json:"createdAt" dc:"创建时间"`
	Puid         string   `json:"puid" dc:"物业ID"`
	SnKeyword    string   `json:"snKeyword"       dc:"系统订单号|外部订单号"`
}

func (in *PmsAppReservationListInp) Filter(ctx context.Context) (err error) {
	return
}

// PmsAppReservationExportListInp 获取系统入住订单列表
type PmsAppReservationExportListInp struct {
	input_form.PageReq
	OrderSn      string `json:"orderSn"       dc:"系统订单号"`
	OutOrderSn   string `json:"outOrderSn"    dc:"外部订单号"`
	CheckinDate  string `json:"checkinDate"  dc:"入住日期"`
	CheckoutDate string `json:"checkoutDate"  dc:"退房日期"`
	RefundStatus string `json:"status"        dc:"退款状态"`
	OrderStatus  string `json:"orderStatus"   dc:"订单状态"`
	MemberId     int    `json:"member_id" dc:"会员ID"`
	MemberNo     string `json:"memberNo"       dc:"会员编号"`
	CreatedAt    string `json:"createdAt" dc:"创建时间"`
	Puid         string `json:"puid" dc:"物业ID"`
	SnKeyword    string `json:"snKeyword"       dc:"系统订单号|外部订单号"`
}

func (in *PmsAppReservationExportListInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsAppReservationListModel struct {
	*entity.PmsAppStay
	MemberDeleted bool `json:"memberDeleted" dc:"会员是否已删除"`
	MemberDetail  *struct {
		gmeta.Meta `orm:"table:hg_pms_member"`
		Id         int    `json:"id"              orm:"id"                description:"主键"`
		MemberNo   string `json:"memberNo"        orm:"member_no"         description:"会员号"`
		FirstName  string `json:"firstName"       orm:"first_name"        description:"名"`
		LastName   string `json:"lastName"        orm:"last_name"         description:"姓"`
		FullName   string `json:"fullName"        orm:"full_name"         description:"全名"`
		Phone      string `json:"phone"           orm:"phone"             description:"手机号"`
		PhoneArea  string `json:"phoneArea"       orm:"phone_area"        description:"手机区号"`
		Mail       string `json:"mail"            orm:"mail"              description:"邮箱"`
	} `json:"memberDetail" orm:"with:id=member_id"`
	PropertyDeleted bool `json:"propertyDeleted" dc:"物业是否已删除"`
	PropertyDetail  *struct {
		gmeta.Meta `orm:"table:hg_pms_property"`
		Uid        string `json:"uid"                  orm:"uid"                     description:"在API合作伙伴系统中的物业ID"`
		Cover      string `json:"cover"                orm:"cover"                   description:"封面"`
		Name       string `json:"name"                 orm:"name"                    description:"物业名称   多语言"`
	} `json:"propertyDetail" orm:"with:uid=puid" dc:"物业信息"`
	GuestProfileDetail struct {
		gmeta.Meta `orm:"table:hg_pms_guest_profile"`
		*entity.PmsGuestProfile
	} `json:"guestProfileDetail"  orm:"with:uid=booker" dc:"预订人信息"`
	TransactionDetail []*struct {
		g.Meta `orm:"table:hg_pms_transaction"`
		*entity.PmsTransaction
	} `json:"transactionDetail" orm:"with:order_sn=order_sn" dc:"支付明细"`
	IsRefund                bool `json:"isRefund"          orm:"is_refund"          description:"是否能退款"`
	TransactionRefundDetail []*struct {
		g.Meta `orm:"table:hg_pms_transaction_refund"`
		*entity.PmsTransactionRefund
	} `json:"transactionRefundDetail" orm:"with:order_sn=order_sn" dc:"退款明细"`
}

// PmsAppReservationRoomListInp 获取系统入住订单列表
type PmsAppReservationRoomListInp struct {
	input_form.PageReq
	Source     string `json:"source"       dc:"来源"`
	OrderSn    string `json:"orderSn"       dc:"系统订单号"`
	OutOrderSn string `json:"outOrderSn"    dc:"外部订单号"`
	//CheckinDate   []*gtime.Time `json:"checkinDate"  dc:"入住日期"`
	//CheckoutDate  []*gtime.Time `json:"checkoutDate"  dc:"退房日期"`
	CheckinDate   []string      `json:"checkinDate"  dc:"入住日期"`
	CheckoutDate  []string      `json:"checkoutDate"  dc:"退房日期"`
	Status        string        `json:"status"        dc:"预订状态"`
	CheckinStatus string        `json:"checkinStatus" dc:"入住状态"`
	OrderStatus   string        `json:"orderStatus"   dc:"支付状态"`
	MemberId      string        `json:"member_id" dc:"会员ID"`
	IsNew         string        `json:"isNew" dc:"是否新订单"`
	PropertyName  string        `json:"propertyName" dc:"物业名称"`
	CreatedAt     []*gtime.Time `json:"createdAt" dc:"创建时间"`
	Puid          string        `json:"puid" dc:"物业ID"`
	SelectStatus  string        `json:"selectStatus" dc:"tab栏状态筛选 CONFIRM-已确认 CANCEL-已取消 ALL-全部"`
}

func (in *PmsAppReservationRoomListInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsAppReservationRoomListModel struct {
	Id              string  `json:"id"             orm:"id"                description:"订单ID"`
	Puid            string  `json:"puid"           orm:"puid"              description:"物业ID"`
	PropertyDeleted bool    `json:"propertyDeleted" dc:"物业是否已删除"`
	OrderSn         string  `json:"orderSn"        orm:"order_sn"          description:"订单号"`
	CreatedAt       string  `json:"createdAt"      orm:"created_at"        description:"订单创建时间"`
	OutOrderSn      string  `json:"outOrderSn"     orm:"out_order_sn"      description:"外部订单号"`
	Status          string  `json:"status"          orm:"status"           description:"预订状态（确认/confirmed、取消/cancelled）"`
	CheckinStatus   string  `json:"checkinStatus"   orm:"checkin_status"   description:"入住状态  before_checkin  在入住之前  checked_in  已入住  checked_out  已退房"`
	OrderStatus     string  `json:"orderStatus"     orm:"order_status"     description:"WAIT_PAY、待支付 CANCEL、支付过期 HAVE_PAID、支付成功"`
	CheckinDate     string  `json:"checkinDate"    orm:"checkin_date"      description:"入住时间"`
	CheckoutDate    string  `json:"checkoutDate"   orm:"checkout_date"     description:"退房时间"`
	CheckinTime     string  `json:"checkinTime"     orm:"checkin_time"     description:"入住时间，24小时格式"`
	CheckoutTime    string  `json:"checkoutTime"    orm:"checkout_time"    description:"退房时间，24小时格式"`
	AdultCount      int     `json:"adultCount"      orm:"adult_count"      description:"成人数量"`
	ChildCount      int     `json:"childCount"      orm:"child_count"      description:"儿童数量"`
	InfantCount     int     `json:"infantCount"     orm:"infant_count"     description:"婴儿数量"`
	BookingFee      float64 `json:"bookingFee"      orm:"booking_fee"      description:"预订费"`
	CancellationFee float64 `json:"cancellationFee" orm:"cancellation_fee" description:"取消费，仅在取消时适用"`
	MainGuest       string  `json:"mainGuest"       orm:"main_guest"       description:"住宿人编号"`
	RoomType        string  `json:"roomType"        orm:"room_type"        description:"房型信息，参考房型uid"`
	RoomUnit        string  `json:"roomUnit"        orm:"room_unit"        description:"房间单元的id, uid或组合"`
	SourceName      string  `json:"sourceName"      orm:"source_name"      description:"渠道名称"`
	GuestRemarks    string  `json:"guestRemark"    orm:"guest_remarks"     description:"备注"`
	PropertyDetail  *struct {
		gmeta.Meta `orm:"table:hg_pms_property"`
		Uid        string `json:"uid"                  orm:"uid"                     description:"在API合作伙伴系统中的物业ID"`
		Cover      string `json:"cover"                orm:"cover"                   description:"封面"`
		Name       string `json:"name"                 orm:"name"                    description:"物业名称   多语言"`
	} `json:"propertyDetail" orm:"with:uid=puid" dc:"物业信息"`
	GuestProfileDetail struct {
		//entity.PmsGuestProfile
		gmeta.Meta  `orm:"table:hg_pms_guest_profile"`
		Uid         string `json:"uid"           orm:"uid"             description:"订单号"`
		FirstName   string `json:"firstName"     orm:"first_name"      description:"名"`
		LastName    string `json:"lastName"      orm:"last_name"       description:"姓"`
		FullName    string `json:"fullName"      orm:"full_name"       description:"全名"`
		Language    string `json:"language"      orm:"language"        description:"语言"`
		Email       string `json:"email"         orm:"email"           description:"电子邮件"`
		Phone       string `json:"phone"         orm:"phone"           description:"电话"`
		AreaNo      string `json:"areaNo"        orm:"area_no"         description:"电话国际区号"`
		Nationality string `json:"nationality"   orm:"nationality"     description:"国籍"`
		Address     string `json:"address"       orm:"address"         description:"地址"`
	} `json:"guestProfileDetail"  orm:"with:uid=main_guest" dc:"预订人信息"`
	RoomTypeDetail *struct {
		gmeta.Meta `orm:"table:hg_pms_room_type"`
		*entity.PmsRoomType
	} `json:"roomTypeDetail" orm:"with:uid=room_type"  dc:"房型信息"`
	RoomUnitDetail *struct {
		gmeta.Meta `orm:"table:hg_pms_room_unit"`
		*entity.PmsRoomUnit
	} `json:"roomUnitDetail" orm:"with:uid=room_unit" dc:"房间信息"`
	TransactionDetail []*struct {
		g.Meta `orm:"table:hg_pms_transaction"`
		//*entity.PmsTransaction
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
	} `json:"transactionDetail" orm:"with:order_sn=order_sn" dc:"支付明细"`
	IsRefund                bool `json:"isRefund"          orm:"is_refund"          description:"是否能退款"`
	TransactionRefundDetail []*struct {
		g.Meta `orm:"table:hg_pms_transaction_refund"`
		//*entity.PmsTransactionRefund
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
	} `json:"transactionRefundDetail" orm:"with:order_sn=order_sn" dc:"退款明细"`
}

// PmsAppReservationRebateListInp 获取系统入住结算订单列表
type PmsAppReservationRebateListInp struct {
	input_form.PageReq
	OrderSn       string `json:"orderSn"       dc:"系统订单号"`
	OutOrderSn    string `json:"outOrderSn"    dc:"外部订单号"`
	CheckinDate   string `json:"checkinDate"   v:"date#日期格式错误" dc:"入住日期"`
	CheckoutDate  string `json:"checkoutDate"  v:"date#日期格式错误" dc:"退房日期"`
	Status        string `json:"status"        dc:"预订状态"`
	CheckinStatus string `json:"checkinStatus" dc:"入住状态"`
	OrderStatus   string `json:"orderStatus"   dc:"支付状态"`
	MemberId      string `json:"member_id" dc:"会员ID"`
	FxType        string `json:"fx_type" dc:"staff：员工  channel：渠道"`
	FxId          string `json:"fx_id" dc:"ID"`
}

func (in *PmsAppReservationRebateListInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsAppReservationRebateListModel struct {
	Id                string      `json:"id"             orm:"id"              description:"订单ID"`
	Puid              string      `json:"puid"           orm:"puid"            description:"物业ID"`
	OrderSn           string      `json:"orderSn"        orm:"order_sn"        description:"订单号"`
	CreatedAt         string      `json:"createdAt"      orm:"created_at"      description:"订单创建时间"`
	OutOrderSn        string      `json:"outOrderSn"     orm:"out_order_sn"    description:"外部订单号"`
	OrderAmount       float64     `json:"orderAmount"    orm:"order_amount"    description:"订单金额"`
	PayModel          int         `json:"payModel"       orm:"pay_model"       description:"1、余额支付 2、组合支付 3、纯外部支付"`
	OrderStatus       string      `json:"orderStatus"    orm:"order_status"    description:"订单付款状态"`
	ExpirationTime    *gtime.Time `json:"expirationTime" orm:"expiration_time" description:"订单过期时间"`
	CheckInDate       string      `json:"checkInDate"    orm:"check_in_date"    description:"入住日期"`
	CheckOutDate      string      `json:"checkOutDate"   orm:"check_out_date"   description:"退房日期"`
	RoomNum           int         `json:"roomNum"        orm:"roomNum"         description:"预定房间总数"`
	GuestNum          int         `json:"guestNum"       orm:"guestNum"        description:"入住人数"`
	Booker            string      `json:"booker"         orm:"booker"          description:"预定人"`
	RebateRate        float64     `json:"rebateRate"     orm:"rebate_rate"     description:"分佣比例"`
	RebateStatus      string      `json:"rebateStatus"   orm:"rebate_status"   description:"'WAIT','SUCCESS','FAIL'"`
	RebateAmount      float64     `json:"rebateAmount"   orm:"rebate_amount"   description:"分佣结算金额"`
	TransactionDetail []*struct {
		g.Meta `orm:"table:hg_pms_transaction"`
		//*entity.PmsTransaction
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
	} `json:"transactionDetail" orm:"with:order_sn=order_sn" dc:"支付明细"`
	IsRefund                bool    `json:"isRefund"          orm:"is_refund"          description:"是否能退款"`
	TotalRefundAmount       float64 `json:"refundAmount"     orm:"total_refund_amount"      description:"总退款金额"`
	TransactionRefundDetail []*struct {
		g.Meta `orm:"table:hg_pms_transaction_refund"`
		//*entity.PmsTransactionRefund
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
	} `json:"transactionRefundDetail" orm:"with:order_sn=order_sn" dc:"退款明细"`
}

type PmsAppReservationExportModel struct {
	PropertyName            string  `json:"propertyName" dc:"物业名称"`
	MemberNo                string  `json:"memberNo" dc:"下单会员ID"`
	GuestName               string  `json:"guestName" dc:"预定人姓名"`
	SystemOrderID           string  `json:"systemOrderID" dc:"系统单号"`
	ChannelOrderID          string  `json:"channelOrderID" dc:"渠道订单号"`
	PaymentStatus           string  `json:"paymentStatus" dc:"支付状态"`
	CheckInDate             string  `json:"checkInDate" dc:"入住日期"`
	CheckOutDate            string  `json:"checkOutDate" dc:"离店日期"`
	NumberOfGuests          int     `json:"numberOfGuests" dc:"人数"`
	NumberOfNights          int     `json:"numberOfNights" dc:"间夜数"`
	OrderTime               string  `json:"orderTime" dc:"下单时间"`
	TotalAmount             float64 `json:"totalAmount" dc:"订单总金额"`
	RefundTotalAmount       float64 `json:"refundTotalAmount" dc:"退款总金额"`
	NowPayAmount            float64 `json:"nowPayAmount" dc:"实付金额"`
	PointsPayment           float64 `json:"pointsPayment" dc:"积分支付"`
	CouponPayment           float64 `json:"couponPayment" dc:"优惠券支付"`
	PaycloudWechatPay       float64 `json:"paycloudWechatPay" dc:"PaycloudWechat支付"`
	PaycloudAlipayPay       float64 `json:"paycloudAlipayPay" dc:"PaycloudAlipay+支付"`
	PaycloudCreditPay       float64 `json:"paycloudCreditPay" dc:"Paycloud信用卡支付"`
	PaypelPay               float64 `json:"paypelPay" dc:"Paypel支付"`
	PaypelCreditPay         float64 `json:"paypelCreditPay" dc:"Paypel信用卡支付"`
	StripeCreditPay         float64 `json:"stripeCreditPay" dc:"Stripe信用卡支付"`
	MlilifeWeChatMiniPay    float64 `json:"mlilifeWeChatMiniPay" dc:"WeChatMini支付"`
	PointsRefund            float64 `json:"pointsRefund" dc:"积分退款"`
	CouponRefund            float64 `json:"couponRefund" dc:"优惠券退款"`
	PaycloudWechatRefund    float64 `json:"paycloudWechatRefund" dc:"PaycloudWechat退款"`
	PaycloudAlipayRefund    float64 `json:"paycloudAlipayRefund" dc:"PaycloudAlipay+退款"`
	PaycloudCreditRefund    float64 `json:"paycloudCreditRefund" dc:"Paycloud信用卡退款"`
	PaypelCreditRefund      float64 `json:"paypelCreditRefund" dc:"Paypel信用卡退款"`
	PaypelRefund            float64 `json:"paypelRefund" dc:"Paypel信用卡退款"`
	StripeCreditRefund      float64 `json:"stripeCreditRefund" dc:"Stripe信用卡退款"`
	MlilifeWeChatMiniRefund float64 `json:"mlilifeWeChatMiniRefund" dc:"WeChatMini退款"`
	RefundTime              string  `json:"refundTime" dc:"退款/取消时间"`
}

// PmsAppReservationStatusInp 更新系统入住订单状态
type PmsAppReservationStatusInp struct {
	Id     int `json:"id" v:"required#主键不能为空" dc:"主键"`
	Status int `json:"status" dc:"状态"`
}

type PmsAppReservationSyncOrderInp struct {
	OrderSn    string `json:"orderSn" v:"required#订单号不能为空" dc:"订单号"`
	OutOrderSn string `json:"outOrderSn" v:"required#第三方订单号不能为空" dc:"第三方订单号"`
}

func (in *PmsAppReservationStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("主键不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}
	return
}

type PmsAppReservationStatusModel struct{}

// PmsAppReservationReferrerListInp 获取系统入住分销订单列表
type PmsAppReservationReferrerListInp struct {
	input_form.PageReq
	OrderSn      string        `json:"orderSn"       dc:"系统订单号"`
	OutOrderSn   string        `json:"outOrderSn"    dc:"外部订单号"`
	CheckinDate  string        `json:"checkinDate"   v:"date#日期格式错误" dc:"入住日期"`
	CheckoutDate string        `json:"checkoutDate"  v:"date#日期格式错误" dc:"退房日期"`
	RefundStatus string        `json:"status"        dc:"退款状态"`
	OrderStatus  string        `json:"orderStatus"   dc:"订单状态"`
	MemberId     int           `json:"member_id" dc:"会员ID"`
	CreatedAt    []*gtime.Time `json:"createdAt" dc:"创建时间"`
	Puid         string        `json:"puid" dc:"物业ID"`
}

func (in *PmsAppReservationReferrerListInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsAppReservationReferrerListModel struct {
	*entity.PmsAppStay
	PropertyDetail *struct {
		gmeta.Meta `orm:"table:hg_pms_property"`
		Uid        string `json:"uid"                  orm:"uid"                     description:"在API合作伙伴系统中的物业ID"`
		Cover      string `json:"cover"                orm:"cover"                   description:"封面"`
		Name       string `json:"name"                 orm:"name"                    description:"物业名称   多语言"`
	} `json:"propertyDetail" orm:"with:uid=puid" dc:"物业信息"`
	GuestProfileDetail struct {
		gmeta.Meta `orm:"table:hg_pms_guest_profile"`
		*entity.PmsGuestProfile
	} `json:"guestProfileDetail"  orm:"with:uid=booker" dc:"预订人信息"`
	TransactionDetail []*struct {
		g.Meta `orm:"table:hg_pms_transaction"`
		*entity.PmsTransaction
	} `json:"transactionDetail" orm:"with:order_sn=order_sn" dc:"支付明细"`

	TransactionRefundDetail []*struct {
		g.Meta `orm:"table:hg_pms_transaction_refund"`
		*entity.PmsTransactionRefund
	} `json:"transactionRefundDetail" orm:"with:order_sn=order_sn" dc:"退款明细"`
	ReferrerDetail *struct {
		gmeta.Meta `orm:"table:hg_pms_member"`
		*entity.PmsMember
		Staff *struct {
			gmeta.Meta `orm:"table:hg_pms_staff"`
			*entity.PmsStaff
		} `json:"staff" orm:"with:id=staff_id"`
		Channel *struct {
			gmeta.Meta `orm:"table:hg_pms_channel"`
			*entity.PmsChannel
		} `json:"channel" orm:"with:id=channel_id"`
	} `json:"referrerDetail" orm:"with:id=referrer"`
}

type PmsAppStayInfoModel struct {
	gmeta.Meta `orm:"table:hg_pms_app_stay"`
	*entity.PmsAppStay
	MemberDetail *struct {
		gmeta.Meta `orm:"table:hg_pms_member"`
		Id         int    `json:"id"              orm:"id"                description:"主键"`
		MemberNo   string `json:"memberNo"        orm:"member_no"         description:"会员号"`
		FirstName  string `json:"firstName"       orm:"first_name"        description:"名"`
		LastName   string `json:"lastName"        orm:"last_name"         description:"姓"`
		FullName   string `json:"fullName"        orm:"full_name"         description:"全名"`
		Phone      string `json:"phone"           orm:"phone"             description:"手机号"`
		PhoneArea  string `json:"phoneArea"       orm:"phone_area"        description:"手机区号"`
		Mail       string `json:"mail"            orm:"mail"              description:"邮箱"`
	} `json:"memberDetail" orm:"with:id=member_id"`
	PropertyDetail *struct {
		gmeta.Meta `orm:"table:hg_pms_property"`
		Id         int    `json:"id"        orm:"id"         description:"主键"`
		Uid        string `json:"uid"       orm:"uid"        description:"在API合作伙伴系统中的物业ID"`
		GroupIds   string `json:"groupIds"  orm:"group_ids"  description:"开放显示的会员分组IDS"`
		Name       string `json:"name"      orm:"name"       description:"物业名称   多语言"`
	} `json:"propertyDetail" orm:"with:uid=puid"`
}
