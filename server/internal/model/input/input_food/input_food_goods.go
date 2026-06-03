// Package sysin

package input_food

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_language"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodGoodsUpdateFields 修改套餐管理字段过滤
type FoodGoodsUpdateFields struct {
	RestaurantId        int64   `json:"restaurantId" dc:"餐厅ID"`
	ToretaCourseId      string  `json:"toretaCourseId" dc:"Toreta课程ID"`
	GoodsName           string  `json:"goodsName"    dc:"套餐名称"`
	Introduction        string  `json:"introduction" dc:"促销语"`
	LabelIds            string  `json:"labelIds"     dc:"标签（多选）"`
	Images              string  `json:"images"       dc:"图集"`
	Price               float64 `json:"price"        dc:"套餐售价"`
	MarketPrice         float64 `json:"marketPrice"  dc:"套餐原价"`
	MaxOrderNum         int     `json:"maxOrderNum"  dc:"单次最大可预定数"`
	TimeDuration        string  `json:"timeDuration" dc:"用餐停留时间"`
	GoodsContent        string  `json:"goodsContent" dc:"套餐详情"`
	Notice              string  `json:"notice"       dc:"注意事项"`
	GoodsState          int     `json:"goodsState"   dc:"状态（1.正常2下架）"`
	IsThCouponExclusive int     `json:"isThCouponExclusive" dc:"是否是礼品券兑换专属：0-否，1-是"`
	ThCouponId          int     `json:"thCouponId"          dc:"绑定的礼品券ID，用于礼品券兑换专属商品"`
	IsNoPay             int     `json:"isNoPay" dc:"是否是无需支付：0-否，1-是"`
	MaxTimeOrderOpen    uint    `json:"maxTimeOrderOpen"    dc:"是否开启时间限制 1-开始  2-关闭"`
	OrderTimeForm       string  `json:"orderTimeForm"       dc:"限制时段"`
}

// FoodGoodsInsertFields 新增套餐管理字段过滤
type FoodGoodsInsertFields struct {
	RestaurantId        int64   `json:"restaurantId" dc:"餐厅ID"`
	ToretaCourseId      string  `json:"toretaCourseId" dc:"Toreta课程ID"`
	GoodsName           string  `json:"goodsName"    dc:"套餐名称"`
	Introduction        string  `json:"introduction" dc:"促销语"`
	LabelIds            string  `json:"labelIds"     dc:"标签（多选）"`
	Images              string  `json:"images"       dc:"图集"`
	Price               float64 `json:"price"        dc:"套餐售价"`
	MarketPrice         float64 `json:"marketPrice"  dc:"套餐原价"`
	MaxOrderNum         int     `json:"maxOrderNum"  dc:"单次最大可预定数"`
	TimeDuration        string  `json:"timeDuration" dc:"用餐停留时间"`
	GoodsContent        string  `json:"goodsContent" dc:"套餐详情"`
	Notice              string  `json:"notice"       dc:"注意事项"`
	GoodsState          int     `json:"goodsState"   dc:"状态（1.正常2下架）"`
	IsThCouponExclusive int     `json:"isThCouponExclusive" dc:"是否是礼品券兑换专属：0-否，1-是"`
	ThCouponId          int     `json:"thCouponId"          dc:"绑定的礼品券ID，用于礼品券兑换专属商品"`
	IsNoPay             int     `json:"isNoPay" dc:"是否是无需支付：0-否，1-是"`
	MaxTimeOrderOpen    uint    `json:"maxTimeOrderOpen"    dc:"是否开启时间限制 1-开始  2-关闭"`
	OrderTimeForm       string  `json:"orderTimeForm"       dc:"限制时段"`
}

// FoodGoodsEditInp 修改/新增套餐管理
type FoodGoodsEditInp struct {
	entity.FoodGoods
	NameLanguage        input_language.LanguageModel `json:"nameLanguage"          dc:"多语言套餐名称"`
	DescriptionLanguage input_language.LanguageModel `json:"descriptionLanguage"          dc:"多语言套餐详情"`
}

func (in *FoodGoodsEditInp) Filter(ctx context.Context) (err error) {
	// 验证餐厅ID
	if err := g.Validator().Rules("required").Data(in.RestaurantId).Messages("餐厅ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证套餐名称
	if err := g.Validator().Rules("required").Data(in.GoodsName).Messages("套餐名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证促销语
	if err := g.Validator().Rules("required").Data(in.Introduction).Messages("促销语不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证标签（多选）
	if err := g.Validator().Rules("required").Data(in.LabelIds).Messages("标签（多选）不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证套餐售价
	if err := g.Validator().Rules("regex:(^[0-9]{1,10}$)|(^[0-9]{1,10}[\\.]{1}[0-9]{1,2}$)").Data(in.Price).Messages("套餐售价最多允许输入10位整数及2位小数").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证套餐原价
	if err := g.Validator().Rules("required").Data(in.MarketPrice).Messages("套餐原价不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证单次最大可预定数
	if err := g.Validator().Rules("required").Data(in.MaxOrderNum).Messages("单次最大可预定数不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证用餐停留时间
	if err := g.Validator().Rules("required").Data(in.TimeDuration).Messages("用餐停留时间不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证套餐详情
	if err := g.Validator().Rules("required").Data(in.GoodsContent).Messages("套餐详情不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证注意事项
	if err := g.Validator().Rules("required").Data(in.Notice).Messages("注意事项不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证状态（1.正常2下架）
	if err := g.Validator().Rules("required").Data(in.GoodsState).Messages("状态（1.正常2下架）不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type FoodGoodsEditModel struct{}

// FoodGoodsDeleteInp 删除套餐管理
type FoodGoodsDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *FoodGoodsDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodGoodsDeleteModel struct{}

// FoodGoodsViewInp 获取指定套餐管理信息
type FoodGoodsViewInp struct {
	Id         int64 `json:"id" v:"required#id不能为空" dc:"id"`
	IsLanguage bool  `json:"isLanguage" dc:"是否获取多语言数据"`
}

func (in *FoodGoodsViewInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodGoodsViewModel struct {
	entity.FoodGoods
	NameLanguage        []*input_hotel.LanguageType `json:"nameLanguage"         dc:"房型名称"   orm:"with:uuid=goods_name"`
	DescriptionLanguage []*input_hotel.LanguageType `json:"descriptionLanguage"         dc:"房型名称"   orm:"with:uuid=goods_content"`
}

// FoodGoodsListInp 获取套餐管理列表
type FoodGoodsListInp struct {
	input_form.PageReq
	GoodsName    string    `json:"goodsName" dc:"套餐名称"`
	RestaurantId uint64    `json:"restaurantId" dc:"餐厅ID"`
	GoodsState   int       `json:"goodsState"             dc:"状态"`
	PriceRange   []float64 `json:"priceRange"        dc:"价格区间"`
}

func (in *FoodGoodsListInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodGoodsListModel struct {
	Id           int64   `json:"id"           dc:"id"`
	RestaurantId int64   `json:"restaurantId" dc:"餐厅ID"`
	GoodsName    string  `json:"goodsName"    dc:"套餐名称"`
	Introduction string  `json:"introduction" dc:"促销语"`
	LabelIds     string  `json:"labelIds"     dc:"标签（多选）"`
	Images       string  `json:"images"       dc:"图集"`
	Price        float64 `json:"price"        dc:"套餐售价"`
	MarketPrice  float64 `json:"marketPrice"  dc:"套餐原价"`
	MaxOrderNum  int     `json:"maxOrderNum"  dc:"单次最大可预定数"`
	TimeDuration string  `json:"timeDuration" dc:"用餐停留时间"`
	GoodsContent string  `json:"goodsContent" dc:"套餐详情"`
	Notice       string  `json:"notice"       dc:"注意事项"`
	GoodsState   int     `json:"goodsState"   dc:"状态（1.正常2下架）"`
}

// FoodGoodsExportModel 导出套餐管理
type FoodGoodsExportModel struct {
	Id           int64       `json:"id"           dc:"id"`
	RestaurantId int64       `json:"restaurantId" dc:"餐厅ID"`
	GoodsName    string      `json:"goodsName"    dc:"套餐名称"`
	Introduction string      `json:"introduction" dc:"促销语"`
	LabelIds     string      `json:"labelIds"     dc:"标签（多选）"`
	Price        float64     `json:"price"        dc:"套餐售价"`
	MarketPrice  float64     `json:"marketPrice"  dc:"套餐原价"`
	MaxOrderNum  int         `json:"maxOrderNum"  dc:"单次最大可预定数"`
	TimeDuration string      `json:"timeDuration" dc:"用餐停留时间"`
	Notice       string      `json:"notice"       dc:"注意事项"`
	GoodsState   int         `json:"goodsState"   dc:"状态（1.正常2下架）"`
	CreateAt     *gtime.Time `json:"createAt"     dc:"创建时间"`
	UpdateAt     *gtime.Time `json:"updateAt"     dc:"更新时间"`
}

// FoodGoodsStatusInp 更新餐厅标签状态
type FoodGoodsStatusInp struct {
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
	Status int `json:"status" dc:"状态"`
}

func (in *FoodGoodsStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("id不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}
	return
}

type FoodGoodsStatusModel struct{}
