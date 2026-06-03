// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// CabinetOrder is the golang structure for table cabinet_order.
type CabinetOrder struct {
	Id                      int64       `json:"id"                      orm:"id"                         description:""`
	OrderSn                 string      `json:"orderSn"                 orm:"order_sn"                   description:"订单编号"`
	OutOrderSn              string      `json:"outOrderSn"              orm:"out_order_sn"               description:"三方订单号"`
	MinHours                int         `json:"minHours"                orm:"min_hours"                  description:"最小小时数"`
	CabinetId               int         `json:"cabinetId"               orm:"cabinet_id"                 description:"储物柜ID"`
	CabinetName             string      `json:"cabinetName"             orm:"cabinet_name"               description:"储物柜名称"`
	CabinetNameJson         *gjson.Json `json:"cabinetNameJson"         orm:"cabinet_name_json"          description:"储物柜名称多语言JSON"`
	CityId                  int         `json:"cityId"                  orm:"city_id"                    description:"所属城市ID"`
	CityName                string      `json:"cityName"                orm:"city_name"                  description:"所属城市名"`
	CityNameJson            *gjson.Json `json:"cityNameJson"            orm:"city_name_json"             description:"所属城市名多语言JSON"`
	MchId                   int         `json:"mchId"                   orm:"mch_id"                     description:"运营商ID"`
	MchName                 string      `json:"mchName"                 orm:"mch_name"                   description:"运营商名称"`
	MchNameJson             *gjson.Json `json:"mchNameJson"             orm:"mch_name_json"              description:"运营商名称多语言JSON"`
	MchBranchId             int         `json:"mchBranchId"             orm:"mch_branch_id"              description:"网点ID"`
	MchBranchName           string      `json:"mchBranchName"           orm:"mch_branch_name"            description:"网点名称"`
	MchBranchNameJson       *gjson.Json `json:"mchBranchNameJson"       orm:"mch_branch_name_json"       description:"网点名称多语言JSON"`
	MchBranchLat            string      `json:"mchBranchLat"            orm:"mch_branch_lat"             description:"网点lat"`
	MchBranchLgt            string      `json:"mchBranchLgt"            orm:"mch_branch_lgt"             description:"网点lgt"`
	Address                 string      `json:"address"                 orm:"address"                    description:"地址"`
	AddressJson             *gjson.Json `json:"addressJson"             orm:"address_json"               description:"地址多语言JSON"`
	BoxTypeJson             *gjson.Json `json:"boxTypeJson"             orm:"box_type_json"              description:"格口类型列表数据"`
	BoxTypeId               int         `json:"boxTypeId"               orm:"box_type_id"                description:"格口类型ID"`
	BoxTypeName             string      `json:"boxTypeName"             orm:"box_type_name"              description:"格口类型名称"`
	BoxTypeNameJson         *gjson.Json `json:"boxTypeNameJson"         orm:"box_type_name_json"         description:"格口类型名称多语言JSON"`
	BoxTypePrice            int         `json:"boxTypePrice"            orm:"box_type_price"             description:"格口类型单价"`
	BoxId                   int         `json:"boxId"                   orm:"box_id"                     description:"格口ID"`
	BoxNo                   string      `json:"boxNo"                   orm:"box_no"                     description:"格口编号"`
	BoxAlias                string      `json:"boxAlias"                orm:"box_alias"                  description:"格口别名"`
	OrderFirstFeeRate       float64     `json:"orderFirstFeeRate"       orm:"order_first_fee_rate"       description:"首次下单优惠"`
	BuyHours                int         `json:"buyHours"                orm:"buy_hours"                  description:"购买的小时数"`
	Pin                     string      `json:"pin"                     orm:"pin"                        description:"取件码(4位数字)"`
	MemberId                uint        `json:"memberId"                orm:"member_id"                  description:"用户ID"`
	OrderAmount             float64     `json:"orderAmount"             orm:"order_amount"               description:"订单金额"`
	BaseAmount              float64     `json:"baseAmount"              orm:"base_amount"                description:"租赁时间内金额"`
	CouponAmount            float64     `json:"couponAmount"            orm:"coupon_amount"              description:"优惠券抵扣金额"`
	BalAmount               float64     `json:"balAmount"               orm:"bal_amount"                 description:"积分抵扣金额"`
	OvertimeSecs            int         `json:"overtimeSecs"            orm:"overtime_secs"              description:"超时时长(秒)"`
	OvertimeHours           int         `json:"overtimeHours"           orm:"overtime_hours"             description:"超时时长(小时)"`
	OvertimeFee             int         `json:"overtimeFee"             orm:"overtime_fee"               description:"超时费用(日元)"`
	OvertimePayTime         *gtime.Time `json:"overtimePayTime"         orm:"overtime_pay_time"          description:"超时费支付时间"`
	OvertimePayStatus       string      `json:"overtimePayStatus"       orm:"overtime_pay_status"        description:"超时费付款状态"`
	OvertimeBalAmount       float64     `json:"overtimeBalAmount"       orm:"overtime_bal_amount"        description:"超时费积分抵扣金额"`
	OvertimePayModel        int         `json:"overtimePayModel"        orm:"overtime_pay_model"         description:"1、余额支付 2、组合支付 3、纯外部支付"`
	GraceSeconds            int         `json:"graceSeconds"            orm:"grace_seconds"              description:"宽限期设置时长（秒）"`
	GraceEndTime            *gtime.Time `json:"graceEndTime"            orm:"grace_end_time"             description:"宽限期结束时间"`
	PayStep                 string      `json:"payStep"                 orm:"pay_step"                   description:"支付流程"`
	PayModel                int         `json:"payModel"                orm:"pay_model"                  description:"1、余额支付 2、组合支付 3、纯外部支付"`
	PayTime                 *gtime.Time `json:"payTime"                 orm:"pay_time"                   description:"支付时间"`
	PayStatus               string      `json:"payStatus"               orm:"pay_status"                 description:"订单付款状态"`
	OrderStatus             string      `json:"orderStatus"             orm:"order_status"               description:"订单状态"`
	ExpirationTime          int         `json:"expirationTime"          orm:"expiration_time"            description:"订单过期时间"`
	CancelTime              *gtime.Time `json:"cancelTime"              orm:"cancel_time"                description:"取消时间"`
	StartTime               *gtime.Time `json:"startTime"               orm:"start_time"                 description:"订单开始时间"`
	EndTime                 *gtime.Time `json:"endTime"                 orm:"end_time"                   description:"订单结束时间"`
	FinishTime              *gtime.Time `json:"finishTime"              orm:"finish_time"                description:"完成时间"`
	RefundStatus            string      `json:"refundStatus"            orm:"refund_status"              description:"退款状态   WAIT 未退款   PART  部分退款 DONE 全部退款"`
	RefundFee               float64     `json:"refundFee"               orm:"refund_fee"                 description:"退款手续费"`
	RefundTime              *gtime.Time `json:"refundTime"              orm:"refund_time"                description:"退款时间"`
	RefundAmount            float64     `json:"refundAmount"            orm:"refund_amount"              description:"已退款总金额"`
	RefundBalAmount         float64     `json:"refundBalAmount"         orm:"refund_bal_amount"          description:"已退款积分"`
	RefundCouponAmount      float64     `json:"refundCouponAmount"      orm:"refund_coupon_amount"       description:"已退款优惠券"`
	RefundReason            string      `json:"refundReason"            orm:"refund_reason"              description:"退款原因"`
	Referrer                int         `json:"referrer"                orm:"referrer"                   description:"推荐人"`
	RebateRate              float64     `json:"rebateRate"              orm:"rebate_rate"                description:"分佣比例"`
	RebateStatus            string      `json:"rebateStatus"            orm:"rebate_status"              description:"WAIT 等待处理佣金   SUCCESS   佣金处理成功    FAIL  佣金处理失败"`
	RebateAmount            float64     `json:"rebateAmount"            orm:"rebate_amount"              description:"分佣结算金额"`
	RebateTime              *gtime.Time `json:"rebateTime"              orm:"rebate_time"                description:"分佣结算时间"`
	IsGetOpen               string      `json:"isGetOpen"               orm:"is_get_open"                description:"是否开启积分获取"`
	IsPayOpen               string      `json:"isPayOpen"               orm:"is_pay_open"                description:"是否开启积分抵扣"`
	CabinetGetRateVip       float64     `json:"cabinetGetRateVip"       orm:"cabinet_get_rate_vip"       description:"储物柜结算积分比例"`
	CabinetGetRateScene     float64     `json:"cabinetGetRateScene"     orm:"cabinet_get_rate_scene"     description:"场景结算积分比例"`
	CabinetGetScoreStatus   string      `json:"cabinetGetScoreStatus"   orm:"cabinet_get_score_status"   description:"'WAIT','SUCCESS','FAIL'"`
	CabinetGetAmount        float64     `json:"cabinetGetAmount"        orm:"cabinet_get_amount"         description:"结算积分金额"`
	ExpValue                float64     `json:"expValue"                orm:"exp_value"                  description:"结算的经验值"`
	ExpTime                 *gtime.Time `json:"expTime"                 orm:"exp_time"                   description:"经验结算时间"`
	IsAbnormal              int         `json:"isAbnormal"              orm:"is_abnormal"                description:"请求下单接口是否异常"`
	PayOvertimeAbnormal     int         `json:"payOvertimeAbnormal"     orm:"pay_overtime_abnormal"      description:"请求支付超时费接口是否异常"`
	IsAdminComplete         int         `json:"isAdminComplete"         orm:"is_admin_complete"          description:"是否是后台强制完成"`
	AdminCompleteOperatorId uint        `json:"adminCompleteOperatorId" orm:"admin_complete_operator_id" description:"强制完成处理操作人ID"`
	CreatedAt               *gtime.Time `json:"createdAt"               orm:"created_at"                 description:""`
	UpdatedAt               *gtime.Time `json:"updatedAt"               orm:"updated_at"                 description:""`
	IsFx                    string      `json:"isFx"                    orm:"is_fx"                      description:"是否是分销订单"`
}
