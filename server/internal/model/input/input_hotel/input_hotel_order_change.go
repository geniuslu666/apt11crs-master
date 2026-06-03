package input_hotel

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderChangeUpdateFields 修改入住订单变更信息表字段过滤
type OrderChangeUpdateFields struct {
	ChangeOrderSn   string      `json:"changeOrderSn"   dc:"变更订单号"`
	OrderId         int         `json:"orderId"         dc:"订单ID"`
	OrderSn         string      `json:"orderSn"         dc:"预订订单号"`
	OutOrderSn      string      `json:"outOrderSn"      dc:"预订外部订单号"`
	ChangeType      string      `json:"changeType"      dc:"变更内容  GUEST 预定人信息变更  PEOPLE   入住人数变更  DATE   日期变更"`
	OldCheckinDate  *gtime.Time `json:"oldCheckinDate"  dc:"入住日期"`
	OldCheckoutDate *gtime.Time `json:"oldCheckoutDate" dc:"退房日期"`
	NewCheckinDate  *gtime.Time `json:"newCheckinDate"  dc:"入住日期"`
	NewCheckoutDate *gtime.Time `json:"newCheckoutDate" dc:"退房日期"`
	OldMainGuest    *gjson.Json `json:"oldMainGuest"    dc:"住宿人编号"`
	NewMainGuest    *gjson.Json `json:"newMainGuest"    dc:"住宿人编号"`
	OldAdultCount   int         `json:"oldAdultCount"   dc:"成人数量"`
	NewAdultCount   int         `json:"newAdultCount"   dc:"成人数量"`
	OldChildCount   int         `json:"oldChildCount"   dc:"儿童数量"`
	NewChildCount   int         `json:"newChildCount"   dc:"儿童数量"`
	OldInfantCount  int         `json:"oldInfantCount"  dc:"婴儿数量"`
	NewInfantCount  int         `json:"newInfantCount"  dc:"婴儿数量"`
	ChangeStatus    string      `json:"changeStatus"    dc:"变动状态 ING   处理中   DONE   变更完成   FAIL   变更失败"`
	ChangeAmount    float64     `json:"changeAmount"    dc:"变动金额"`
	SubmitDate      *gtime.Time `json:"submitDate"      dc:"提交变更时间"`
	DoneDate        *gtime.Time `json:"doneDate"        dc:"变更成功时间"`
	OldOrderPrice   float64     `json:"oldOrderPrice"   dc:"原订单价格"`
	NewOrderPrice   float64     `json:"newOrderPrice"   dc:"变更后订单价格"`
	ExpirationTime  int         `json:"expirationTime"  dc:"订单过期时间"`
}

// OrderChangeInsertFields 新增入住订单变更信息表字段过滤
type OrderChangeInsertFields struct {
	ChangeOrderSn   string      `json:"changeOrderSn"   dc:"变更订单号"`
	OrderId         int         `json:"orderId"         dc:"订单ID"`
	OrderSn         string      `json:"orderSn"         dc:"预订订单号"`
	OutOrderSn      string      `json:"outOrderSn"      dc:"预订外部订单号"`
	ChangeType      string      `json:"changeType"      dc:"变更内容  GUEST 预定人信息变更  PEOPLE   入住人数变更  DATE   日期变更"`
	OldCheckinDate  *gtime.Time `json:"oldCheckinDate"  dc:"入住日期"`
	OldCheckoutDate *gtime.Time `json:"oldCheckoutDate" dc:"退房日期"`
	NewCheckinDate  *gtime.Time `json:"newCheckinDate"  dc:"入住日期"`
	NewCheckoutDate *gtime.Time `json:"newCheckoutDate" dc:"退房日期"`
	OldMainGuest    *gjson.Json `json:"oldMainGuest"    dc:"住宿人编号"`
	NewMainGuest    *gjson.Json `json:"newMainGuest"    dc:"住宿人编号"`
	OldAdultCount   int         `json:"oldAdultCount"   dc:"成人数量"`
	NewAdultCount   int         `json:"newAdultCount"   dc:"成人数量"`
	OldChildCount   int         `json:"oldChildCount"   dc:"儿童数量"`
	NewChildCount   int         `json:"newChildCount"   dc:"儿童数量"`
	OldInfantCount  int         `json:"oldInfantCount"  dc:"婴儿数量"`
	NewInfantCount  int         `json:"newInfantCount"  dc:"婴儿数量"`
	ChangeStatus    string      `json:"changeStatus"    dc:"变动状态 ING   处理中   DONE   变更完成   FAIL   变更失败"`
	ChangeAmount    float64     `json:"changeAmount"    dc:"变动金额"`
	SubmitDate      *gtime.Time `json:"submitDate"      dc:"提交变更时间"`
	DoneDate        *gtime.Time `json:"doneDate"        dc:"变更成功时间"`
	OldOrderPrice   float64     `json:"oldOrderPrice"   dc:"原订单价格"`
	NewOrderPrice   float64     `json:"newOrderPrice"   dc:"变更后订单价格"`
	ExpirationTime  int         `json:"expirationTime"  dc:"订单过期时间"`
}

// OrderChangeEditInp 修改/新增入住订单变更信息表
type OrderChangeEditInp struct {
	entity.PmsAppReservationChange
}

type OrderChangeEditModel struct{}

// OrderChangeDeleteInp 删除入住订单变更信息表
type OrderChangeDeleteInp struct {
	Id interface{} `json:"id" v:"required#主键不能为空" dc:"主键"`
}

type OrderChangeDeleteModel struct{}

// OrderChangeViewInp 获取指定入住订单变更信息表信息
type OrderChangeViewInp struct {
	Id int `json:"id" v:"required#主键不能为空" dc:"主键"`
}

type OrderChangeViewModel struct {
	entity.PmsAppReservationChange
}

// OrderChangeListInp 获取入住订单变更信息表列表
type OrderChangeListInp struct {
	input_form.PageReq
	Id        int           `json:"id"        dc:"主键"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"created_at"`
}

type OrderChangeListModel struct {
	Id              int         `json:"id"              dc:"主键"`
	ChangeOrderSn   string      `json:"changeOrderSn"   dc:"变更订单号"`
	OrderId         int         `json:"orderId"         dc:"订单ID"`
	OrderSn         string      `json:"orderSn"         dc:"预订订单号"`
	OutOrderSn      string      `json:"outOrderSn"      dc:"预订外部订单号"`
	ChangeType      string      `json:"changeType"      dc:"变更内容  GUEST 预定人信息变更  PEOPLE   入住人数变更  DATE   日期变更"`
	OldCheckinDate  *gtime.Time `json:"oldCheckinDate"  dc:"入住日期"`
	OldCheckoutDate *gtime.Time `json:"oldCheckoutDate" dc:"退房日期"`
	NewCheckinDate  *gtime.Time `json:"newCheckinDate"  dc:"入住日期"`
	NewCheckoutDate *gtime.Time `json:"newCheckoutDate" dc:"退房日期"`
	OldAdultCount   int         `json:"oldAdultCount"   dc:"成人数量"`
	NewAdultCount   int         `json:"newAdultCount"   dc:"成人数量"`
	OldChildCount   int         `json:"oldChildCount"   dc:"儿童数量"`
	NewChildCount   int         `json:"newChildCount"   dc:"儿童数量"`
	OldInfantCount  int         `json:"oldInfantCount"  dc:"婴儿数量"`
	NewInfantCount  int         `json:"newInfantCount"  dc:"婴儿数量"`
	ChangeStatus    string      `json:"changeStatus"    dc:"变动状态 ING   处理中   DONE   变更完成   FAIL   变更失败"`
	ChangeAmount    float64     `json:"changeAmount"    dc:"变动金额"`
	SubmitDate      *gtime.Time `json:"submitDate"      dc:"提交变更时间"`
	DoneDate        *gtime.Time `json:"doneDate"        dc:"变更成功时间"`
	OldOrderPrice   float64     `json:"oldOrderPrice"   dc:"原订单价格"`
	NewOrderPrice   float64     `json:"newOrderPrice"   dc:"变更后订单价格"`
	ExpirationTime  int         `json:"expirationTime"  dc:"订单过期时间"`
	CreatedAt       *gtime.Time `json:"createdAt"       dc:"created_at"`
	UpdatedAt       *gtime.Time `json:"updatedAt"       dc:"updated_at"`
}

// OrderChangeExportModel 导出入住订单变更信息表
type OrderChangeExportModel struct {
	Id              int         `json:"id"              dc:"主键"`
	ChangeOrderSn   string      `json:"changeOrderSn"   dc:"变更订单号"`
	OrderId         int         `json:"orderId"         dc:"订单ID"`
	OrderSn         string      `json:"orderSn"         dc:"预订订单号"`
	OutOrderSn      string      `json:"outOrderSn"      dc:"预订外部订单号"`
	ChangeType      string      `json:"changeType"      dc:"变更内容  GUEST 预定人信息变更  PEOPLE   入住人数变更  DATE   日期变更"`
	OldCheckinDate  *gtime.Time `json:"oldCheckinDate"  dc:"入住日期"`
	OldCheckoutDate *gtime.Time `json:"oldCheckoutDate" dc:"退房日期"`
	NewCheckinDate  *gtime.Time `json:"newCheckinDate"  dc:"入住日期"`
	NewCheckoutDate *gtime.Time `json:"newCheckoutDate" dc:"退房日期"`
	OldAdultCount   int         `json:"oldAdultCount"   dc:"成人数量"`
	NewAdultCount   int         `json:"newAdultCount"   dc:"成人数量"`
	OldChildCount   int         `json:"oldChildCount"   dc:"儿童数量"`
	NewChildCount   int         `json:"newChildCount"   dc:"儿童数量"`
	OldInfantCount  int         `json:"oldInfantCount"  dc:"婴儿数量"`
	NewInfantCount  int         `json:"newInfantCount"  dc:"婴儿数量"`
	ChangeStatus    string      `json:"changeStatus"    dc:"变动状态 ING   处理中   DONE   变更完成   FAIL   变更失败"`
	ChangeAmount    float64     `json:"changeAmount"    dc:"变动金额"`
	SubmitDate      *gtime.Time `json:"submitDate"      dc:"提交变更时间"`
	DoneDate        *gtime.Time `json:"doneDate"        dc:"变更成功时间"`
	OldOrderPrice   float64     `json:"oldOrderPrice"   dc:"原订单价格"`
	NewOrderPrice   float64     `json:"newOrderPrice"   dc:"变更后订单价格"`
	ExpirationTime  int         `json:"expirationTime"  dc:"订单过期时间"`
	CreatedAt       *gtime.Time `json:"createdAt"       dc:"created_at"`
	UpdatedAt       *gtime.Time `json:"updatedAt"       dc:"updated_at"`
}

type OrderChangeGuestReq struct {
	AirhostOrderUuid string `json:"airhostOrderUuid" v:"required#airhost_order_id_unknown" dc:"airhost订单ID"`
	OrderId          int    `json:"orderId" v:"required#order_id_missing" dc:"订单ID"`
	FirstName        string `json:"firstName"     dc:"名"`
	LastName         string `json:"lastName"      dc:"姓"`
	Email            string `json:"email"         dc:"电子邮件"`
	Phone            string `json:"phone"         dc:"电话"`
	AreaNo           string `json:"areaNo"        dc:"电话国际区号"`
}
