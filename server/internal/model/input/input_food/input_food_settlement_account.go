// Package sysin

package input_food

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/utility/validate"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodSettlementAccountUpdateFields 修改结算模式字段过滤
type FoodSettlementAccountUpdateFields struct {
	RestaurantId  int    `json:"restaurantId"    dc:"餐厅ID"`
	Type          string `json:"type"            dc:"类型"`
	BankName      string `json:"bankName"        dc:"开户行"`
	BankUser      string `json:"bankUser"        dc:"开户人姓名"`
	BankCard      string `json:"bankCard"     dc:"银行账号"`
	WechatAccount string `json:"wechatAccount"   dc:"微信名"`
	AlipayName    string `json:"alipayName"      dc:"支付宝真实姓名"`
	AlipayAccount string `json:"alipayAccount" dc:"支付宝账户"`
	Status        int    `json:"status"          dc:"状态 1、启用 2、禁用"`
}

// FoodSettlementAccountInsertFields 新增结算模式字段过滤
type FoodSettlementAccountInsertFields struct {
	RestaurantId  int    `json:"restaurantId"    dc:"餐厅ID"`
	Type          string `json:"type"            dc:"类型"`
	BankName      string `json:"bankName"        dc:"开户行"`
	BankUser      string `json:"bankUser"        dc:"开户人姓名"`
	BankCard      string `json:"bankCard"     dc:"银行账号"`
	WechatAccount string `json:"wechatAccount"   dc:"微信名"`
	AlipayName    string `json:"alipayName"      dc:"支付宝真实姓名"`
	AlipayAccount string `json:"alipayAccount" dc:"支付宝账户"`
	Status        int    `json:"status"          dc:"状态 1、启用 2、禁用"`
}

// FoodSettlementAccountEditInp 修改/新增结算模式
type FoodSettlementAccountEditInp struct {
	entity.FoodSettlementAccount
}

func (in *FoodSettlementAccountEditInp) Filter(ctx context.Context) (err error) {
	// 验证餐厅ID
	if err := g.Validator().Rules("required").Data(in.RestaurantId).Messages("餐厅ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证类型
	if err := g.Validator().Rules("required").Data(in.Type).Messages("类型不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	if in.Type == "BANKCARD" {
		if err := g.Validator().Rules("required").Data(in.BankName).Messages("开户行不能为空").Run(ctx); err != nil {
			return err.Current()
		}
		if err := g.Validator().Rules("required").Data(in.BankUser).Messages("开户人姓名不能为空").Run(ctx); err != nil {
			return err.Current()
		}
		if err := g.Validator().Rules("required").Data(in.BankCard).Messages("银行账号不能为空").Run(ctx); err != nil {
			return err.Current()
		}
	} else if in.Type == "WECHAT" {
		if err := g.Validator().Rules("required").Data(in.WechatAccount).Messages("微信名不能为空").Run(ctx); err != nil {
			return err.Current()
		}
	} else if in.Type == "ALIPAY" {
		if err := g.Validator().Rules("required").Data(in.AlipayName).Messages("支付宝真实姓名不能为空").Run(ctx); err != nil {
			return err.Current()
		}
		if err := g.Validator().Rules("required").Data(in.AlipayAccount).Messages("支付宝账户不能为空").Run(ctx); err != nil {
			return err.Current()
		}
	}

	return
}

type FoodSettlementAccountEditModel struct{}

// FoodSettlementAccountDeleteInp 删除结算模式
type FoodSettlementAccountDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *FoodSettlementAccountDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodSettlementAccountDeleteModel struct{}

// FoodSettlementAccountListInp 获取结算模式列表
type FoodSettlementAccountListInp struct {
	input_form.PageReq
	RestaurantId int `json:"restaurantId" v:"required#餐厅ID不能为空" dc:"餐厅ID"`
	Status       int `json:"status" dc:"状态"`
}

func (in *FoodSettlementAccountListInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodSettlementAccountListModel struct {
	Id            int64       `json:"id"              dc:"id"`
	RestaurantId  int         `json:"restaurantId"    dc:"餐厅ID"`
	Type          string      `json:"type"            dc:"类型"`
	BankName      string      `json:"bankName"        dc:"开户行"`
	BankUser      string      `json:"bankUser"        dc:"开户人姓名"`
	BankCard      string      `json:"bankCard"     dc:"银行账号"`
	WechatAccount string      `json:"wechatAccount"   dc:"微信名"`
	AlipayName    string      `json:"alipayName"      dc:"支付宝真实姓名"`
	AlipayAccount string      `json:"alipayAccount" dc:"支付宝账户"`
	Status        int         `json:"status"          dc:"状态 1、启用 2、禁用"`
	CreateAt      *gtime.Time `json:"createAt"        dc:"创建时间"`
	UpdateAt      *gtime.Time `json:"updateAt"        dc:"更新时间"`
}

// FoodSettlementAccountStatusInp 更新结算模式状态
type FoodSettlementAccountStatusInp struct {
	Id     int64 `json:"id" v:"required#id不能为空" dc:"id"`
	Status int   `json:"status" dc:"状态"`
}

func (in *FoodSettlementAccountStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("id不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}
	return
}

type FoodSettlementAccountStatusModel struct{}

// FoodSettlementAccountViewInp 获取指定餐厅-结算账户信息
type FoodSettlementAccountViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *FoodSettlementAccountViewInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodSettlementAccountViewModel struct {
	entity.FoodSettlementAccount
}
