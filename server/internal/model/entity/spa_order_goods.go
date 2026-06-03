// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SpaOrderGoods is the golang structure for table spa_order_goods.
type SpaOrderGoods struct {
	Id                 int64       `json:"id"                 orm:"id"                   description:""`
	IspId              int         `json:"ispId"              orm:"isp_id"               description:"服务商ID"`
	OrderId            uint        `json:"orderId"            orm:"order_id"             description:"订单ID"`
	ServiceType        int         `json:"serviceType"        orm:"service_type"         description:"1到店  2上门"`
	ServiceId          int         `json:"serviceId"          orm:"service_id"           description:"服务ID"`
	TechnicianGoTime   *gtime.Time `json:"technicianGoTime"   orm:"technician_go_time"   description:"技师上门出发时间"`
	ActualStartTime    *gtime.Time `json:"actualStartTime"    orm:"actual_start_time"    description:"实际开始日期时间"`
	ActualEndTime      *gtime.Time `json:"actualEndTime"      orm:"actual_end_time"      description:"实际结束日期时间"`
	GoodsId            uint        `json:"goodsId"            orm:"goods_id"             description:"项目ID"`
	GoodsNum           int         `json:"goodsNum"           orm:"goods_num"            description:"项目数量"`
	OrderGoodsAmount   float64     `json:"orderGoodsAmount"   orm:"order_goods_amount"   description:"子订单金额"`
	CouponAmount       float64     `json:"couponAmount"       orm:"coupon_amount"        description:"子订单优惠券抵扣金额"`
	BalAmount          float64     `json:"balAmount"          orm:"bal_amount"           description:"子订单积分抵扣金额"`
	TechnicianIds      string      `json:"technicianIds"      orm:"technician_ids"       description:"技师ID ,分隔"`
	PayStatus          string      `json:"payStatus"          orm:"pay_status"           description:"订单付款状态"`
	StoreId            uint        `json:"storeId"            orm:"store_id"             description:"到店门店ID"`
	PropertyId         uint        `json:"propertyId"         orm:"property_id"          description:"上门物业ID"`
	RoomNo             string      `json:"roomNo"             orm:"room_no"              description:"上门房间号"`
	OrderStatus        string      `json:"orderStatus"        orm:"order_status"         description:"订单状态"`
	DispatchStatus     string      `json:"dispatchStatus"     orm:"dispatch_status"      description:"订单调度状态"`
	DispatchTime       *gtime.Time `json:"dispatchTime"       orm:"dispatch_time"        description:"订单调度时间"`
	DispatchDesc       string      `json:"dispatchDesc"       orm:"dispatch_desc"        description:"订单调度备注"`
	DispatchOperatorId uint        `json:"dispatchOperatorId" orm:"dispatch_operator_id" description:"调度操作人ID"`
	StartServeImages   string      `json:"startServeImages"   orm:"start_serve_images"   description:"开始服务图集"`
	EndServeImages     string      `json:"endServeImages"     orm:"end_serve_images"     description:"服务结束图集"`
	RefundStatus       string      `json:"refundStatus"       orm:"refund_status"        description:"退款状态   WAIT 未退款   PART  部分退款 DONE 全部退款"`
	RefundFee          float64     `json:"refundFee"          orm:"refund_fee"           description:"退款手续费"`
	RefundTime         *gtime.Time `json:"refundTime"         orm:"refund_time"          description:"退款时间"`
	RefundAmount       float64     `json:"refundAmount"       orm:"refund_amount"        description:"已退款总金额"`
	RefundBalAmount    float64     `json:"refundBalAmount"    orm:"refund_bal_amount"    description:"已退款积分"`
	RefundCouponAmount float64     `json:"refundCouponAmount" orm:"refund_coupon_amount" description:"已退款优惠券"`
	CreatedAt          *gtime.Time `json:"createdAt"          orm:"created_at"           description:"创建时间"`
	UpdatedAt          *gtime.Time `json:"updatedAt"          orm:"updated_at"           description:"更新时间"`
}
