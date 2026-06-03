// Package sysin

package input_food

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/utility/validate"
	"context"
	"github.com/gogf/gf/v2/util/gmeta"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodRestaurantNoticeUpdateFields 修改通知管理字段过滤
type FoodRestaurantNoticeUpdateFields struct {
	RestaurantId    int    `json:"restaurantId"    dc:"餐厅ID"`
	Title           string `json:"title"           dc:"通知标题"`
	Content         string `json:"content"         dc:"公告内容"`
	Sort            int    `json:"sort"            dc:"排序(越大越靠前)"`
	Status          int    `json:"status"          dc:"状态1、启用 2、禁用"`
	NeedUserConfirm int    `json:"needUserConfirm" dc:"是否需要用户确认 1、需要  2、不需要"`
}

// FoodRestaurantNoticeInsertFields 新增通知管理字段过滤
type FoodRestaurantNoticeInsertFields struct {
	RestaurantId    int    `json:"restaurantId"    dc:"餐厅ID"`
	Title           string `json:"title"           dc:"通知标题"`
	Content         string `json:"content"         dc:"公告内容"`
	Sort            int    `json:"sort"            dc:"排序(越大越靠前)"`
	Status          int    `json:"status"          dc:"状态1、启用 2、禁用"`
	NeedUserConfirm int    `json:"needUserConfirm" dc:"是否需要用户确认 1、需要  2、不需要"`
}

// FoodRestaurantNoticeEditInp 修改/新增通知管理
type FoodRestaurantNoticeEditInp struct {
	entity.FoodRestaurantNotice
}

func (in *FoodRestaurantNoticeEditInp) Filter(ctx context.Context) (err error) {
	// 验证餐厅ID
	if err := g.Validator().Rules("required").Data(in.RestaurantId).Messages("餐厅ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证通知标题
	if err := g.Validator().Rules("required").Data(in.Title).Messages("通知标题不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证公告内容
	if err := g.Validator().Rules("required").Data(in.Content).Messages("公告内容不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证排序(越大越靠前)
	if err := g.Validator().Rules("required").Data(in.Sort).Messages("排序(越大越靠前)不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type FoodRestaurantNoticeEditModel struct{}

// FoodRestaurantNoticeDeleteInp 删除通知管理
type FoodRestaurantNoticeDeleteInp struct {
	Id interface{} `json:"id" v:"required# 不能为空" dc:" "`
}

func (in *FoodRestaurantNoticeDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodRestaurantNoticeDeleteModel struct{}

// FoodRestaurantNoticeViewInp 获取指定通知管理信息
type FoodRestaurantNoticeViewInp struct {
	Id int `json:"id" v:"required# 不能为空" dc:" "`
}

func (in *FoodRestaurantNoticeViewInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodRestaurantNoticeViewModel struct {
	entity.FoodRestaurantNotice
	RestaurantDetail *struct {
		gmeta.Meta `orm:"table:hg_food_restaurant"`
		Id         int    `json:"id"                 dc:"ID"`
		Name       string `json:"name"              dc:"餐厅名称 多语言"`
	} `json:"restaurantDetail" orm:"with:id=restaurant_id"  dc:"餐厅信息"`
}

// FoodRestaurantNoticeListInp 获取通知管理列表
type FoodRestaurantNoticeListInp struct {
	input_form.PageReq
	RestaurantId int           `json:"restaurantId" dc:"餐厅ID"`
	Title        string        `json:"title"        dc:"通知标题"`
	Status       int           `json:"status"       dc:"状态1、启用 2、禁用"`
	CreatedAt    []*gtime.Time `json:"createdAt"                 dc:"created_at"`
}

func (in *FoodRestaurantNoticeListInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodRestaurantNoticeListModel struct {
	Id               int         `json:"id"                 dc:"ID"`
	RestaurantId     int         `json:"restaurantId"       dc:"餐厅ID"`
	Title            string      `json:"title"              dc:"通知标题"`
	Sort             int         `json:"sort"               dc:"排序(越大越靠前)"`
	Status           int         `json:"status"             dc:"状态1、启用 2、禁用"`
	NeedUserConfirm  int         `json:"needUserConfirm"    dc:"是否需要用户确认 1、需要  2、不需要"`
	CreateAt         *gtime.Time `json:"createAt"           dc:"创建时间"`
	UpdateAt         *gtime.Time `json:"updateAt"           dc:"更新时间"`
	RestaurantDetail *struct {
		gmeta.Meta `orm:"table:hg_food_restaurant"`
		Id         int    `json:"id"                 dc:"ID"`
		Name       string `json:"name"              dc:"餐厅名称 多语言"`
	} `json:"restaurantDetail" orm:"with:id=restaurant_id"  dc:"餐厅信息"`
}

// FoodRestaurantNoticeExportModel 导出通知管理
type FoodRestaurantNoticeExportModel struct {
	Id              int         `json:"id"              dc:""`
	RestaurantId    int         `json:"restaurantId"    dc:"餐厅ID"`
	Title           string      `json:"title"           dc:"通知标题"`
	Sort            int         `json:"sort"            dc:"排序(越大越靠前)"`
	Status          int         `json:"status"          dc:"状态1、启用 2、禁用"`
	NeedUserConfirm int         `json:"needUserConfirm" dc:"是否需要用户确认 1、需要  2、不需要"`
	CreatedAt       *gtime.Time `json:"createdAt"        dc:"创建时间"`
	UpdatedAt       *gtime.Time `json:"updatedAt"        dc:"更新时间"`
}

// FoodRestaurantNoticeMaxSortInp 获取通知管理最大排序
type FoodRestaurantNoticeMaxSortInp struct{}

func (in *FoodRestaurantNoticeMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodRestaurantNoticeMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// FoodRestaurantNoticeStatusInp 更新通知管理状态
type FoodRestaurantNoticeStatusInp struct {
	Id     int `json:"id" v:"required# 不能为空" dc:" "`
	Status int `json:"status" dc:"状态"`
}

func (in *FoodRestaurantNoticeStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New(" 不能为空")
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

type FoodRestaurantNoticeStatusModel struct{}
