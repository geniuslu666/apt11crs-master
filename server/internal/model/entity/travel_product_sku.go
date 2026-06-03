// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TravelProductSku is the golang structure for table travel_product_sku.
type TravelProductSku struct {
	Id               uint64      `json:"id"               orm:"id"                 description:""`
	ProductId        uint        `json:"productId"        orm:"product_id"         description:"产品ID"`
	Name             string      `json:"name"             orm:"name"               description:"车型名称（默认语言；多语言存 hg_pms_language）"`
	Price            float64     `json:"price"            orm:"price"              description:"售价（元）"`
	DailyCapacity    uint        `json:"dailyCapacity"    orm:"daily_capacity"     description:"每日最大接待人数"`
	Status           int         `json:"status"           orm:"status"             description:"状态（1启用 2禁用）"`
	Sort             int         `json:"sort"             orm:"sort"               description:"排序（越大越靠前）"`
	SalesNum         uint        `json:"salesNum"         orm:"sales_num"          description:"已售"`
	ContactMobile    string      `json:"contactMobile"    orm:"contact_mobile"     description:"联系电话"`
	MeetingPlace     string      `json:"meetingPlace"     orm:"meeting_place"      description:"集合地点"`
	MeetingTime      string      `json:"meetingTime"      orm:"meeting_time"       description:"集合时间（格式：HH:MM）"`
	GgLat            string      `json:"ggLat"            orm:"gg_lat"             description:"谷歌纬度"`
	GgLng            string      `json:"ggLng"            orm:"gg_lng"             description:"谷歌经度"`
	DeletedAt        *gtime.Time `json:"deletedAt"        orm:"deleted_at"         description:"软删除时间（NULL=正常）"`
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"         description:"创建时间"`
	UpdatedAt        *gtime.Time `json:"updatedAt"        orm:"updated_at"         description:"更新时间"`
	AllowCancel      int         `json:"allowCancel"      orm:"allow_cancel"       description:"是否允许取消"`
	FreeCancelHours  int         `json:"freeCancelHours"  orm:"free_cancel_hours"  description:"几小时前免费"`
	CancelFeePercent int         `json:"cancelFeePercent" orm:"cancel_fee_percent" description:"取消费率%"`
}
