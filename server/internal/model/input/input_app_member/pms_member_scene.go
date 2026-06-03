// Package sysin

package input_app_member

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsMemberSceneUpdateFields 修改会员等级场景字段过滤
type PmsMemberSceneUpdateFields struct {
	SceneName           string  `json:"sceneName"  dc:"场景名称"`
	IsGetOpen           string  `json:"isGetOpen"  dc:"是否允许积分获取"`
	IsPayOpen           string  `json:"isPayOpen"  dc:"是否允许积分抵扣"`
	LimitMoney          float64 `json:"limitMoney" dc:"限制金额"`
	GetRate             float64 `json:"getRate"    dc:"获得积分的抵扣比例"`
	PayRate             float64 `json:"payRate"    dc:"能够使用的总金额比例积分"`
	IsOpenReward        uint    `json:"isOpenReward"        dc:"1、启用 2、禁用"`
	RewardType          string  `json:"rewardType"          dc:"奖励类型（balance：积分 | coupon：优惠券 | thcoupon：礼品券）"`
	RewardCouponTypeIds string  `json:"rewardCouponTypeIds" dc:"奖励优惠券ID，逗号分隔"`
	RewardThCouponIds   string  `json:"rewardThCouponIds"   dc:"奖励礼品券ID，逗号分隔"`
}

// PmsMemberSceneInsertFields 新增会员等级场景字段过滤
type PmsMemberSceneInsertFields struct {
	SceneName           string  `json:"sceneName"  dc:"场景名称"`
	IsGetOpen           string  `json:"isGetOpen"  dc:"是否允许积分获取"`
	IsPayOpen           string  `json:"isPayOpen"  dc:"是否允许积分抵扣"`
	LimitMoney          float64 `json:"limitMoney" dc:"限制金额"`
	GetRate             float64 `json:"getRate"    dc:"获得积分的抵扣比例"`
	PayRate             float64 `json:"payRate"    dc:"能够使用的总金额比例积分"`
	IsOpenReward        uint    `json:"isOpenReward"        dc:"1、启用 2、禁用"`
	RewardType          string  `json:"rewardType"          dc:"奖励类型（balance：积分 | coupon：优惠券 | thcoupon：礼品券）"`
	RewardCouponTypeIds string  `json:"rewardCouponTypeIds" dc:"奖励优惠券ID，逗号分隔"`
	RewardThCouponIds   string  `json:"rewardThCouponIds"   dc:"奖励礼品券ID，逗号分隔"`
}

// PmsMemberSceneEditInp 修改/新增会员等级场景
type PmsMemberSceneEditInp struct {
	entity.PmsMemberScene
}

func (in *PmsMemberSceneEditInp) Filter(ctx context.Context) (err error) {
	// 验证限制金额
	if err := g.Validator().Rules("regex:(^[0-9]{1,10}$)|(^[0-9]{1,10}[\\.]{1}[0-9]{1,2}$)").Data(in.LimitMoney).Messages("限制金额最多允许输入10位整数及2位小数").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证获得积分的抵扣比例
	if err := g.Validator().Rules("min:0|max:100").Data(in.GetRate).Messages("获得积分的抵扣比例必须0-100之间").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证能够使用的总金额比例积分
	if err := g.Validator().Rules("min:0|max:100").Data(in.PayRate).Messages("能够使用的总金额比例积分必须0-100之间").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type PmsMemberSceneEditModel struct{}

// PmsMemberSceneDeleteInp 删除会员等级场景
type PmsMemberSceneDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsMemberSceneDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsMemberSceneDeleteModel struct{}

// PmsMemberSceneViewInp 获取指定会员等级场景信息
type PmsMemberSceneViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsMemberSceneViewInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsMemberSceneViewModel struct {
	entity.PmsMemberScene
}

// PmsMemberSceneListInp 获取会员等级场景列表
type PmsMemberSceneListInp struct {
	input_form.PageReq
	SceneName string        `json:"sceneName" dc:"场景名称"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *PmsMemberSceneListInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsMemberSceneListModel struct {
	Id         int         `json:"id"         dc:"id"`
	SceneName  string      `json:"sceneName"  dc:"场景名称"`
	IsGetOpen  string      `json:"isGetOpen"  dc:"是否允许积分获取"`
	IsPayOpen  string      `json:"isPayOpen"  dc:"是否允许积分抵扣"`
	LimitMoney float64     `json:"limitMoney" dc:"限制金额"`
	GetRate    float64     `json:"getRate"    dc:"获得积分的抵扣比例"`
	PayRate    float64     `json:"payRate"    dc:"能够使用的总金额比例积分"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  dc:"修改时间"`
}

// PmsMemberSceneSwitchInp 更新会员等级场景开关状态
type PmsMemberSceneSwitchInp struct {
	input_form.SwitchReq
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsMemberSceneSwitchInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsMemberSceneSwitchModel struct{}
