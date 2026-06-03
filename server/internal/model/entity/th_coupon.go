// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ThCoupon is the golang structure for table th_coupon.
type ThCoupon struct {
	Id                       int         `json:"id"                       orm:"id"                         description:"ID"`
	CouponName               string      `json:"couponName"               orm:"coupon_name"                description:"提货券名称"`
	CouponSubName            string      `json:"couponSubName"            orm:"coupon_sub_name"            description:"提货券副标题"`
	IdentityName             string      `json:"identityName"             orm:"identity_name"              description:"券识别名称"`
	CouponNoPrefix           string      `json:"couponNoPrefix"           orm:"coupon_no_prefix"           description:"编号前缀"`
	Logo                     string      `json:"logo"                     orm:"logo"                       description:"LOGO"`
	Status                   int         `json:"status"                   orm:"status"                     description:"发放状态（1立即启用  2暂不启用）"`
	FixedTerm                int         `json:"fixedTerm"                orm:"fixed_term"                 description:"激活后几天内有效"`
	Desc                     string      `json:"desc"                     orm:"desc"                       description:"券说明"`
	Count                    int         `json:"count"                    orm:"count"                      description:"发放数量"`
	UseStatus                int         `json:"useStatus"                orm:"use_status"                 description:"使用状态（1开始使用  2停止使用）"`
	UseMode                  string      `json:"useMode"                  orm:"use_mode"                   description:"使用模式"`
	CategoryId               uint        `json:"categoryId"               orm:"category_id"                description:"分类ID"`
	UsedCount                int         `json:"usedCount"                orm:"used_count"                 description:"已使用数量"`
	Sort                     int         `json:"sort"                     orm:"sort"                       description:"排序"`
	NeedReservation          int         `json:"needReservation"          orm:"need_reservation"           description:"是否需要预约：0-不需要，1-需要"`
	ReservationRestaurantIds string      `json:"reservationRestaurantIds" orm:"reservation_restaurant_ids" description:"需要预约的餐厅IDs，多个用逗号分隔，如：1,2,3"`
	CreateAt                 *gtime.Time `json:"createAt"                 orm:"create_at"                  description:"创建时间"`
	UpdateAt                 *gtime.Time `json:"updateAt"                 orm:"update_at"                  description:"修改时间"`
	DeletedAt                *gtime.Time `json:"deletedAt"                orm:"deleted_at"                 description:"删除时间"`
}
