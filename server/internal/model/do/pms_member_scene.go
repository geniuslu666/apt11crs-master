// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsMemberScene is the golang structure of table hg_pms_member_scene for DAO operations like Where/Data.
type PmsMemberScene struct {
	g.Meta              `orm:"table:hg_pms_member_scene, do:true"`
	Id                  interface{} //
	SceneName           interface{} // 场景名称
	IsGetOpen           interface{} // 是否允许积分获取
	IsPayOpen           interface{} // 是否允许积分抵扣
	LimitMoney          interface{} // 限制金额
	GetRate             interface{} // 获得积分的抵扣比例
	PayRate             interface{} // 能够使用的总金额比例积分
	IsOpenReward        interface{} // 1、启用 2、禁用
	RewardType          interface{} // 奖励类型（balance：积分 | coupon：优惠券 | thcoupon：礼品券）
	RewardCouponTypeIds interface{} // 奖励优惠券ID，逗号分隔
	RewardThCouponIds   interface{} // 奖励礼品券ID，逗号分隔
	CreatedAt           *gtime.Time //
	UpdatedAt           *gtime.Time //
	DeletedAt           *gtime.Time //
}
