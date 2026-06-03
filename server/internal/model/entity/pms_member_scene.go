// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsMemberScene is the golang structure for table pms_member_scene.
type PmsMemberScene struct {
	Id                  int         `json:"id"                  orm:"id"                     description:""`
	SceneName           string      `json:"sceneName"           orm:"scene_name"             description:"场景名称"`
	IsGetOpen           string      `json:"isGetOpen"           orm:"is_get_open"            description:"是否允许积分获取"`
	IsPayOpen           string      `json:"isPayOpen"           orm:"is_pay_open"            description:"是否允许积分抵扣"`
	LimitMoney          float64     `json:"limitMoney"          orm:"limit_money"            description:"限制金额"`
	GetRate             float64     `json:"getRate"             orm:"get_rate"               description:"获得积分的抵扣比例"`
	PayRate             float64     `json:"payRate"             orm:"pay_rate"               description:"能够使用的总金额比例积分"`
	IsOpenReward        uint        `json:"isOpenReward"        orm:"is_open_reward"         description:"1、启用 2、禁用"`
	RewardType          string      `json:"rewardType"          orm:"reward_type"            description:"奖励类型（balance：积分 | coupon：优惠券 | thcoupon：礼品券）"`
	RewardCouponTypeIds string      `json:"rewardCouponTypeIds" orm:"reward_coupon_type_ids" description:"奖励优惠券ID，逗号分隔"`
	RewardThCouponIds   string      `json:"rewardThCouponIds"   orm:"reward_th_coupon_ids"   description:"奖励礼品券ID，逗号分隔"`
	CreatedAt           *gtime.Time `json:"createdAt"           orm:"created_at"             description:""`
	UpdatedAt           *gtime.Time `json:"updatedAt"           orm:"updated_at"             description:""`
	DeletedAt           *gtime.Time `json:"deletedAt"           orm:"deleted_at"             description:""`
}
