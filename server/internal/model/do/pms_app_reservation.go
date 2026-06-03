// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsAppReservation is the golang structure of table hg_pms_app_reservation for DAO operations like Where/Data.
type PmsAppReservation struct {
	g.Meta           `orm:"table:hg_pms_app_reservation, do:true"`
	Id               interface{} // 主键
	Uid              interface{} // 我方系统 ID
	Uuid             interface{} // airhost   ID
	Source           interface{} // 来源
	SourceCode       interface{} // 来源渠道
	SourceName       interface{} // 渠道名称
	MemberId         interface{} // 用户ID
	Puid             interface{} // 物业ID
	OrderSn          interface{} // 系统订单号
	OrderIndex       interface{} // 订单索引
	OutOrderSn       interface{} // 三方订单号
	RoomType         interface{} // 房型信息，参考房型uid
	RoomUnit         interface{} // 房间单元的id, uid或组合
	RatePlanId       interface{} // 费率ID
	CheckinDate      *gtime.Time // 入住日期
	CheckoutDate     *gtime.Time // 退房日期
	CheckinTime      interface{} // 入住时间，24小时格式
	CheckoutTime     interface{} // 退房时间，24小时格式
	Status           interface{} // 预订状态（确认/confirmed、取消/cancelled）
	CheckinStatus    interface{} // 入住状态  before_checkin  在入住之前  checked_in  已入住  checked_out  已退房
	OrderStatus      interface{} // WAIT_PAY、待支付 CANCEL、支付过期 HAVE_PAID、支付成功
	MainGuest        interface{} // 住宿人编号
	AdultCount       interface{} // 成人数量
	ChildCount       interface{} // 儿童数量
	InfantCount      interface{} // 婴儿数量
	PricePlanId      interface{} // 价格plan ID
	ChangeAmount     interface{} // 变动金额
	PricePlanInfo    *gjson.Json // 价格plan内容
	BookingFee       interface{} // 预订费
	ChannelFee       interface{} // 渠道费
	CleaningFee      interface{} // 清洁费
	CancellationFee  interface{} // 取消费，仅在取消时适用
	Charges          interface{} // 费用详情，参考Charge对象
	GuestRemarks     interface{} // 备注
	CancelRemake     interface{} // 取消原因
	ExpValue         interface{} // 结算的经验值
	ExpTime          *gtime.Time // 经验结算时间
	CreatedAt        *gtime.Time //
	UpdatedAt        *gtime.Time //
	IsChangeGuest    interface{} // 是否变更入住人信息 Y 是 N 否
	IsChangeGuestId  interface{} // 入住人信息变更ID
	IsChangePeople   interface{} // 是否变更入住人数 Y 是 N 否
	IsChangePeopleId interface{} // 入住人数信息变更ID
	IsChangeDate     interface{} // 是否变更入住日期 Y 是 N 否
	IsChangeDateId   interface{} // 入住日期变更ID
	IsFx             interface{} // 是否是分销订单
}
