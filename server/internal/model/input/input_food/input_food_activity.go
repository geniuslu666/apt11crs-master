// Package sysin

package input_food

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_language"
	"APT/utility/validate"
	"context"
	"github.com/gogf/gf/v2/util/gmeta"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodActivityUpdateFields 修改订餐活动表字段过滤
type FoodActivityUpdateFields struct {
	Date    *gtime.Time `json:"date"    dc:"活动日期"`
	Name    string      `json:"name"    dc:"标题（多语言）"`
	SubName string      `json:"subName" dc:"副标题"`
	Pic     string      `json:"pic"     dc:"图片"`
	Status  int         `json:"status"  dc:"状态"`
}

// FoodActivityInsertFields 新增订餐活动表字段过滤
type FoodActivityInsertFields struct {
	Date    *gtime.Time `json:"date"    dc:"活动日期"`
	Name    string      `json:"name"    dc:"标题（多语言）"`
	SubName string      `json:"subName" dc:"副标题"`
	Pic     string      `json:"pic"     dc:"图片"`
	Status  int         `json:"status"  dc:"状态"`
}

// FoodActivityEditInp 修改/新增订餐活动表
type FoodActivityEditInp struct {
	entity.FoodActivity
	NameLanguage  input_language.LanguageModel `json:"nameLanguage"          dc:"多语言优惠券名称"`
	RestaurantIds string                       `json:"restaurantIds" dc:"餐厅IDs"`
}

func (in *FoodActivityEditInp) Filter(ctx context.Context) (err error) {
	// 验证副标题
	if err := g.Validator().Rules("required").Data(in.SubName).Messages("副标题不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证图片
	if err := g.Validator().Rules("required").Data(in.Pic).Messages("图片不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type FoodActivityEditModel struct{}

// FoodActivityDeleteInp 删除订餐活动表
type FoodActivityDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *FoodActivityDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodActivityDeleteModel struct{}

// FoodActivityViewInp 获取指定订餐活动表信息
type FoodActivityViewInp struct {
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *FoodActivityViewInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodActivityViewModel struct {
	entity.FoodActivity
	NameLanguage   []*input_hotel.LanguageType `json:"nameLanguage"         dc:"优惠券名称"   orm:"with:uuid=name"`
	RestaurantList []*struct {
		gmeta.Meta   `orm:"table:hg_food_activity_restaurant"`
		RestaurantId int64 `json:"restaurantId"  orm:"restaurant_id"     description:"餐厅ID"`
		ActivityId   int   `json:"activityId"        orm:"activity_id"           description:"活动ID"`
	} `json:"restaurantList" orm:"with:activity_id=id" dc:"餐厅列表"`
}

// FoodActivityListInp 获取订餐活动表列表
type FoodActivityListInp struct {
	input_form.PageReq
	Date       []*gtime.Time `json:"date"     dc:"活动日期"`
	Status     int           `json:"status"   dc:"状态"`
	CreateAt   []*gtime.Time `json:"createAt" dc:"创建时间"`
	IsLanguage bool          `json:"isLanguage" dc:"是否获取多语言数据"`
}

func (in *FoodActivityListInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodActivityListModel struct {
	Id             int64                       `json:"id"       dc:"id"`
	Date           *gtime.Time                 `json:"date"     dc:"活动日期"`
	Name           string                      `json:"name"     dc:"标题（多语言）"`
	SubName        string                      `json:"subName"  dc:"副标题"`
	Pic            string                      `json:"pic"      dc:"图片"`
	Status         int                         `json:"status"   dc:"状态"`
	CreateAt       *gtime.Time                 `json:"createAt" dc:"创建时间"`
	UpdateAt       *gtime.Time                 `json:"updateAt" dc:"更新时间"`
	NameLanguage   []*input_hotel.LanguageType `json:"nameLanguage"         dc:"优惠券名称"   orm:"with:uuid=name"`
	RestaurantList []*struct {
		gmeta.Meta       `orm:"table:hg_food_activity_restaurant"`
		RestaurantId     int64 `json:"restaurantId"  orm:"restaurant_id"     description:"餐厅ID"`
		ActivityId       int   `json:"activityId"        orm:"activity_id"           description:"活动ID"`
		RestaurantDetail *struct {
			gmeta.Meta `orm:"table:hg_food_restaurant"`
			Id         int    `json:"id"              dc:"ID"`
			Name       string `json:"name"            dc:"餐厅名称 多语言"`
			Images     string `json:"images"          dc:"图片"`
		} `json:"restaurantDetail" orm:"with:id=restaurant_id"`
	} `json:"restaurantList" orm:"with:activity_id=id" dc:"餐厅列表"`
}

// FoodActivityExportModel 导出订餐活动表
type FoodActivityExportModel struct {
	Id       int64       `json:"id"       dc:"id"`
	Date     *gtime.Time `json:"date"     dc:"活动日期"`
	Name     string      `json:"name"     dc:"标题（多语言）"`
	SubName  string      `json:"subName"  dc:"副标题"`
	Pic      string      `json:"pic"      dc:"图片"`
	Status   int         `json:"status"   dc:"状态"`
	CreateAt *gtime.Time `json:"createAt" dc:"创建时间"`
	UpdateAt *gtime.Time `json:"updateAt" dc:"更新时间"`
}

// FoodActivityStatusInp 更新订餐活动表状态
type FoodActivityStatusInp struct {
	Id     int64 `json:"id" v:"required#id不能为空" dc:"id"`
	Status int   `json:"status" dc:"状态"`
}

func (in *FoodActivityStatusInp) Filter(ctx context.Context) (err error) {
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

type FoodActivityStatusModel struct{}
