// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodRestaurant is the golang structure for table food_restaurant.
type FoodRestaurant struct {
	Id                      int64       `json:"id"                      orm:"id"                         description:""`
	OrderMode               string      `json:"orderMode"               orm:"order_mode"                 description:"预定模式"`
	ToretaId                string      `json:"toretaId"                orm:"toreta_id"                  description:"Toreta的餐厅ID"`
	Name                    string      `json:"name"                    orm:"name"                       description:"名称"`
	CuisineIds              string      `json:"cuisineIds"              orm:"cuisine_ids"                description:"菜系（多选）"`
	LabelIds                string      `json:"labelIds"                orm:"label_ids"                  description:"标签（多选）"`
	CooperateTypeId         int         `json:"cooperateTypeId"         orm:"cooperate_type_id"          description:"合作类型ID"`
	Logo                    string      `json:"logo"                    orm:"logo"                       description:"LOGO"`
	Images                  string      `json:"images"                  orm:"images"                     description:"图集"`
	Content                 string      `json:"content"                 orm:"content"                    description:"简介"`
	Phone                   string      `json:"phone"                   orm:"phone"                      description:"联系电话"`
	OpenTime                string      `json:"openTime"                orm:"open_time"                  description:"营业时间"`
	AreaPid                 int         `json:"areaPid"                 orm:"area_pid"                   description:"地区省级ID"`
	AreaId                  int         `json:"areaId"                  orm:"area_id"                    description:"地区市级ID"`
	DetailAddress           string      `json:"detailAddress"           orm:"detail_address"             description:"详细地址"`
	GgLat                   string      `json:"ggLat"                   orm:"gg_lat"                     description:"谷歌纬度"`
	GgLng                   string      `json:"ggLng"                   orm:"gg_lng"                     description:"谷歌经度"`
	Lat                     string      `json:"lat"                     orm:"lat"                        description:"纬度"`
	Lng                     string      `json:"lng"                     orm:"lng"                        description:"经度"`
	OpenStatus              string      `json:"openStatus"              orm:"open_status"                description:"营业状态"`
	OrderTimeType           int         `json:"orderTimeType"           orm:"order_time_type"            description:"预约日期  1每天 2自定义"`
	OrderTimeWeek           string      `json:"orderTimeWeek"           orm:"order_time_week"            description:"预约日期自定义周数"`
	RestTimeWeek            string      `json:"restTimeWeek"            orm:"rest_time_week"             description:"定休日周数"`
	RestTimeDate            string      `json:"restTimeDate"            orm:"rest_time_date"             description:"定休日固定日期"`
	OrderConfirmType        int         `json:"orderConfirmType"        orm:"order_confirm_type"         description:"预约确定模式 1手动确认 2自动确认(废除)"`
	OrderConfirmDayType     int         `json:"orderConfirmDayType"     orm:"order_confirm_day_type"     description:"1 手动确认在前   2自动确认在前"`
	OrderConfirmBeforeDays  int         `json:"orderConfirmBeforeDays"  orm:"order_confirm_before_days"  description:"多少天内手动确认"`
	OrderConfirmTime        int         `json:"orderConfirmTime"        orm:"order_confirm_time"         description:"预约确认自定义分钟数"`
	OpenTimeType            int         `json:"openTimeType"            orm:"open_time_type"             description:"1 按时段  2 按时间点"`
	AmOpenTime              string      `json:"amOpenTime"              orm:"am_open_time"               description:"早市开始时间"`
	AmCloseTime             string      `json:"amCloseTime"             orm:"am_close_time"              description:"早市结束时间"`
	PmOpenTime              string      `json:"pmOpenTime"              orm:"pm_open_time"               description:"晚市开始时间"`
	PmCloseTime             string      `json:"pmCloseTime"             orm:"pm_close_time"              description:"晚市结束时间"`
	TimePoints              string      `json:"timePoints"              orm:"time_points"                description:"时间点"`
	TimeDuration            string      `json:"timeDuration"            orm:"time_duration"              description:"时间间隔"`
	OrderMaxType            int         `json:"orderMaxType"            orm:"order_max_type"             description:"最大容纳数  1按时段  2按时间点"`
	TimeDurationMax         int         `json:"timeDurationMax"         orm:"time_duration_max"          description:"时段最大容纳预定数量  0不限制"`
	DayTimeLimit            string      `json:"dayTimeLimit"            orm:"day_time_limit"             description:"当天预约处理截止时间"`
	AdvanceOrderDay         int         `json:"advanceOrderDay"         orm:"advance_order_day"          description:"至少提前预约天数  0当日可约"`
	MaxOrderDay             int         `json:"maxOrderDay"             orm:"max_order_day"              description:"最长预约天数"`
	CancelPolicyOpen        int         `json:"cancelPolicyOpen"        orm:"cancel_policy_open"         description:"是否允许取消  1允许  2不允许"`
	BeforeConfirmCancelRate float64     `json:"beforeConfirmCancelRate" orm:"before_confirm_cancel_rate" description:"订单确认前取消费用为订单的%，0则免费"`
	AfterConfirmCancelDay   int         `json:"afterConfirmCancelDay"   orm:"after_confirm_cancel_day"   description:"订单确认后距离到店多少天"`
	AfterConfirmCancelRate1 float64     `json:"afterConfirmCancelRate1" orm:"after_confirm_cancel_rate1" description:"订单确认后距离到店天数前取消费用为订单的%，0则免费"`
	AfterConfirmCancelRate2 float64     `json:"afterConfirmCancelRate2" orm:"after_confirm_cancel_rate2" description:"订单确认后距离到店天数后取消费用为订单的%，0则免费"`
	MaxSeat                 int         `json:"maxSeat"                 orm:"max_seat"                   description:"最大席位数"`
	CanSmoking              int         `json:"canSmoking"              orm:"can_smoking"                description:"是否允许抽烟  1允许  2不允许"`
	SettlementType          int         `json:"settlementType"          orm:"settlement_type"            description:"门店抽成类型 1跟随系统  2自定义"`
	SettlementId            int         `json:"settlementId"            orm:"settlement_id"              description:"结算模式ID"`
	SettlementRate          float64     `json:"settlementRate"          orm:"settlement_rate"            description:"门店结算比例"`
	Desc                    string      `json:"desc"                    orm:"desc"                       description:"开业/停业原因"`
	TotalOrderNum           int         `json:"totalOrderNum"           orm:"total_order_num"            description:"预约单总数量（包含退款）"`
	TotalOrderAmount        float64     `json:"totalOrderAmount"        orm:"total_order_amount"         description:"预约单总金额（包含退款）"`
	PayOrderNum             int         `json:"payOrderNum"             orm:"pay_order_num"              description:"预约单支付数量（不包含退款）"`
	PayOrderAmount          float64     `json:"payOrderAmount"          orm:"pay_order_amount"           description:"预约单支付金额（不包含退款）"`
	WaitConfirmOrderNum     int         `json:"waitConfirmOrderNum"     orm:"wait_confirm_order_num"     description:"预约单待确认数量（已支付待确认）"`
	SettlementOrderNum      int         `json:"settlementOrderNum"      orm:"settlement_order_num"       description:"预约单已结算数量"`
	SettlementOrderAmount   float64     `json:"settlementOrderAmount"   orm:"settlement_order_amount"    description:"已结算预约单订单金额"`
	TotalSettlementAmount   float64     `json:"totalSettlementAmount"   orm:"total_settlement_amount"    description:"已结算金额"`
	VerifyCode              string      `json:"verifyCode"              orm:"verify_code"                description:"核销码"`
	CanOrder                int         `json:"canOrder"                orm:"can_order"                  description:"是否开放预定  1开放  2关闭"`
	Sort                    int         `json:"sort"                    orm:"sort"                       description:"排序(越大越靠前)"`
	DepositRate             float64     `json:"depositRate"             orm:"deposit_rate"               description:"定金比例"`
	Account                 string      `json:"account"                 orm:"account"                    description:"账号"`
	PasswordHash            string      `json:"passwordHash"            orm:"password_hash"              description:"密码"`
	Salt                    string      `json:"salt"                    orm:"salt"                       description:"密码盐"`
	PasswordResetToken      string      `json:"passwordResetToken"      orm:"password_reset_token"       description:"密码重置令牌"`
	ToretaCancelEnable      int         `json:"toretaCancelEnable"      orm:"toreta_cancel_enable"       description:"Toreta是否允许取消 1 允许  2 不允许"`
	ToretaCancelLimitDay    int         `json:"toretaCancelLimitDay"    orm:"toreta_cancel_limit_day"    description:"Toreta允许取消几天前"`
	ToretaCancelLimitTime   string      `json:"toretaCancelLimitTime"   orm:"toreta_cancel_limit_time"   description:"Toreta允许取消时间前"`
	MaxDatetimeOrderOpen    uint        `json:"maxDatetimeOrderOpen"    orm:"max_datetime_order_open"    description:"日期限制 1 开启  2关闭"`
	MaxDatetimeOrderDate    string      `json:"maxDatetimeOrderDate"    orm:"max_datetime_order_date"    description:"限制的日期 逗号分隔"`
	CreateAt                *gtime.Time `json:"createAt"                orm:"create_at"                  description:"创建时间"`
	UpdateAt                *gtime.Time `json:"updateAt"                orm:"update_at"                  description:"更新时间"`
	DeletedAt               *gtime.Time `json:"deletedAt"               orm:"deleted_at"                 description:""`
}
