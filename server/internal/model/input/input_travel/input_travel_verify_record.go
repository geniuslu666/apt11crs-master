package input_travel

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// TravelVerifyRecordListInp 获取核销记录列表
type TravelVerifyRecordListInp struct {
	input_form.PageReq
	OrderSn       string   `json:"orderSn"       dc:"预约单号"`
	ProductName   string   `json:"productName"   dc:"活动名称"`
	VerifyStaffId int64    `json:"verifyStaffId" dc:"核销人员ID"`
	BookDate      []string `json:"bookDate"      dc:"预约日期范围 [开始, 结束]"`
	VerifyTime    []string `json:"verifyTime"    dc:"核销时间范围 [开始, 结束]"`
}

// TravelVerifyRecordListModel 核销记录列表模型
type TravelVerifyRecordListModel struct {
	entity.TravelVerifyRecord
	PmsMemberMemberNo string `json:"pmsMemberMemberNo"        dc:"会员号"`
	MemberDeleted     bool   `json:"memberDeleted"            dc:"会员是否已删除"`
	OrderInfo         *struct {
		gmeta.Meta    `orm:"table:hg_travel_order"`
		Id            uint64 `json:"id"              dc:""`
		OrderSn       string `json:"orderSn"         dc:"预约单号"`
		BookingName   string `json:"bookingName"     dc:"预订人姓名"`
		PhoneArea     string `json:"phoneArea"       dc:"手机区号"`
		BookingMobile string `json:"bookingMobile"   dc:"预订人电话"`
		BookingNum    uint   `json:"bookingNum"      dc:"预约人数"`
		BookDate      string `json:"bookDate"      dc:"预约日期"`
		OrderStatus   string `json:"orderStatus"     dc:"订单状态_WAIT_PAY-待支付 WAIT_VERIFY-已支付待核销 DONE-已完成 CANCEL-取消 REFUND-退款"`
	} `json:"orderInfo" orm:"with:id=order_id"  dc:"订单信息"`
	ProductInfo *struct {
		gmeta.Meta `orm:"table:hg_travel_product"`
		*entity.TravelProduct
	} `json:"productInfo" orm:"with:id=product_id"  dc:"产品"`
	VerifyStaffInfo *struct {
		gmeta.Meta `orm:"table:hg_travel_verify_staff"`
		*entity.TravelVerifyStaff
	} `json:"verifyStaffInfo" orm:"with:id=verify_staff_id"  dc:"核销人员"`
}

// VerifyListInp 获取核销记录列表
type VerifyListInp struct {
	input_form.PageReq
	VerifyStaffId int64 `json:"verifyStaffId" dc:"核销人员ID"`
}

func (in *VerifyListInp) Filter(ctx context.Context) (err error) {
	return
}

// VerifyListModel 核销记录列表模型
type VerifyListModel struct {
	Id            uint64      `json:"id"            dc:""`
	OrderId       uint64      `json:"orderId"       dc:"订单ID"`
	OrderSn       string      `json:"orderSn"       dc:"预约单号"`
	ProductId     uint64      `json:"productId"     dc:"产品ID"`
	MemberId      uint64      `json:"memberId"      dc:"会员ID"`
	BookDate      string      `json:"bookDate"      dc:"预约日期"`
	VerifyStaffId uint64      `json:"verifyStaffId" dc:"核销人员ID"`
	VerifyTime    *gtime.Time `json:"verifyTime"    dc:"核销时间"`
	OrderInfo     *struct {
		gmeta.Meta    `orm:"table:hg_travel_order"`
		Id            uint64 `json:"id"              dc:""`
		OrderSn       string `json:"orderSn"         dc:"预约单号"`
		BookingName   string `json:"bookingName"     dc:"预订人姓名"`
		PhoneArea     string `json:"phoneArea"       dc:"手机区号"`
		BookingMobile string `json:"bookingMobile"   dc:"预订人电话"`
		BookingNum    uint   `json:"bookingNum"      dc:"预约人数"`
		OrderStatus   string `json:"orderStatus"     dc:"订单状态_WAIT_PAY-待支付 WAIT_VERIFY-已支付待核销 DONE-已完成 CANCEL-取消 REFUND-退款"`
	} `json:"orderInfo" orm:"with:id=order_id"  dc:"订单信息"`
	ProductInfo *struct {
		gmeta.Meta `orm:"table:hg_travel_product"`
		Id         uint64 `json:"id"              dc:""`
		Title      string `json:"title"           dc:"标题（默认语言；多语言存 hg_pms_language）"`
		SubTitle   string `json:"subTitle"        dc:"副标题（默认语言；多语言存 hg_pms_language）"`
	} `json:"productInfo" orm:"with:id=product_id"  dc:"产品"`
}

type VerifyLogViewInp struct {
	Id int `json:"id" v:"required#核销记录ID不能为空" dc:"核销记录ID"`
}

func (in *VerifyLogViewInp) Filter(ctx context.Context) (err error) {
	return
}

type VerifyLogViewModel struct {
	Id            uint64      `json:"id"            dc:""`
	OrderId       uint64      `json:"orderId"       dc:"订单ID"`
	OrderSn       string      `json:"orderSn"       dc:"预约单号"`
	ProductId     uint64      `json:"productId"     dc:"产品ID"`
	MemberId      uint64      `json:"memberId"      dc:"会员ID"`
	BookDate      *gtime.Time `json:"bookDate"      dc:"预约日期"`
	VerifyStaffId uint64      `json:"verifyStaffId" dc:"核销人员ID"`
	StaffName     string      `json:"staffName" dc:"核销人员"`
	VerifyTime    *gtime.Time `json:"verifyTime"    dc:"核销时间"`
	OrderInfo     *struct {
		gmeta.Meta    `orm:"table:hg_travel_order"`
		Id            uint64 `json:"id"              dc:""`
		OrderSn       string `json:"orderSn"         dc:"预约单号"`
		BookingName   string `json:"bookingName"     dc:"预订人姓名"`
		PhoneArea     string `json:"phoneArea"       dc:"手机区号"`
		BookingMobile string `json:"bookingMobile"   dc:"预订人电话"`
		BookingNum    uint   `json:"bookingNum"      dc:"预约人数"`
	} `json:"orderInfo" orm:"with:id=order_id"  dc:"订单信息"`
	ProductInfo *struct {
		gmeta.Meta `orm:"table:hg_travel_product"`
		Id         uint64 `json:"id"              dc:""`
		Title      string `json:"title"           dc:"标题（默认语言；多语言存 hg_pms_language）"`
		SubTitle   string `json:"subTitle"        dc:"副标题（默认语言；多语言存 hg_pms_language）"`
	} `json:"productInfo" orm:"with:id=product_id"  dc:"产品"`
}

// CodeViewInp 获取二维码详情（核销端）
type CodeViewInp struct {
	Code string `json:"code" v:"required#券码不能为空" dc:"券码"`
}

func (in *CodeViewInp) Filter(ctx context.Context) (err error) {
	return
}

type CodeViewModel struct {
	Id            uint64      `json:"id"                 dc:""`
	OrderSn       string      `json:"orderSn"            dc:"订单编号"`
	ProductId     uint64      `json:"productId"          dc:"产品ID"`
	MemberId      uint64      `json:"memberId"           dc:"会员ID"`
	BookingName   string      `json:"bookingName"        dc:"预订人姓名"`
	PhoneArea     string      `json:"phoneArea"          dc:"手机区号"`
	BookingMobile string      `json:"bookingMobile"      dc:"预订人电话"`
	BookingEmail  string      `json:"bookingEmail"       dc:"预定人邮箱"`
	BookingNum    uint        `json:"bookingNum"         dc:"预约人数"`
	BookDate      string      `json:"bookDate"           dc:"预约日期"`
	OrderAmount   float64     `json:"orderAmount"        dc:"订单金额（元）"`
	CreatedAt     *gtime.Time `json:"createdAt"          dc:"下单时间"`
	OrderStatus   string      `json:"orderStatus"        dc:"订单状态_WAIT_PAY-待支付 WAIT_VERIFY-已支付待核销 DONE-已完成 CANCEL-已取消 REFUND-已退款 OVERDUE-已逾期"`
	ProductInfo   *struct {
		gmeta.Meta   `orm:"table:hg_travel_product"`
		Id           uint64 `json:"id"              dc:""`
		Title        string `json:"title"           dc:"标题（默认语言；多语言存 hg_pms_language）"`
		SubTitle     string `json:"subTitle"        dc:"副标题（默认语言；多语言存 hg_pms_language）"`
		MeetingPlace string `json:"meetingPlace"    dc:"集合地点"`
		MeetingTime  string `json:"meetingTime"     dc:"集合时间（格式：HH:MM）"`
		GgLat        string `json:"ggLat"           dc:"谷歌纬度"`
		GgLng        string `json:"ggLng"           dc:"谷歌经度"`
	} `json:"productInfo" orm:"with:id=product_id"  dc:"产品"`
}

type CodeVerifyInp struct {
	OrderSn string `json:"orderSn" v:"required#order_number_miss" dc:"订单号"`
}

func (in *CodeVerifyInp) Filter(ctx context.Context) (err error) {
	return
}

type CodeVerifyModel struct {
	VerifyTime *gtime.Time `json:"verifyTime"     dc:"核销时间"`
	OrderSn    string      `json:"orderSn"        dc:"订单编号"`
}
