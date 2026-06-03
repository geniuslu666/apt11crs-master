package input_hotel

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_language"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type FindOneHotelPricePlanInp struct {
	Id int `json:"id"              dc:""`
}

type FindOneHotelPricePlanRes struct {
	Id                   int             `json:"id"              dc:""`
	PlanName             string          `json:"planName"        dc:"价格plan名称"`
	PlanShowName         string          `json:"planShowName"     dc:"价格plan展示名称"`
	PlanShowNameLanguage []*LanguageType `json:"planShowNameLanguage"         dc:"价格plan展示名称"   orm:"with:uuid=plan_show_name"`
	PropertyId           string          `json:"propertyId"      dc:"物业ID"`
	PropertyDetail       *struct {
		g.Meta `orm:"table:hg_pms_property"`
		*entity.PmsProperty
	} `json:"propertyDetail" orm:"with:uid=propertyId"`
	RoomTypeId        string                   `json:"roomTypeId"      dc:"房型ID"`
	RoomTypeDetail    *entity.PmsRoomType      `json:"roomTypeDetail" dc:"房型"`
	BookingDays       int                      `json:"bookingDays"     dc:"预订天数"`
	MemberGroupId     string                   `json:"memberGroupId"   dc:"用户组"`
	MemberGroupDetail []*entity.PmsMemberGroup `json:"memberGroupDetail"`
	MemberLevelId     string                   `json:"memberLevelId"   dc:"会员等级"`
	MemberLevelDetail []*entity.PmsMemberLevel `json:"memberLevelDetail" dc:"会员等级详情"`
	IsCancel          string                   `json:"isCancel"        dc:"是否可取消   Y  是  N  否"`
	IsOpenPriceMode   string                   `json:"isOpenPriceMode" dc:"是否开启价格模式"`
	PriceMode         string                   `json:"priceMode"       dc:"模式    +  贵   - 便宜"`
	PriceStandard     string                   `json:"priceStandard"   dc:"基准  PERCENT 倍率  AMOUNT  金额"`
	PlanValue         int                      `json:"planValue"       dc:"价格基准值"`
	PlanTips          string                   `json:"planTips"        dc:"价格计划提示"`
	PricePlanStatus   string                   `json:"pricePlanStatus" dc:"Y 开启  N 关闭"`
	CreatedAt         *gtime.Time              `json:"createdAt"       dc:"创建时间"`
}
type SearchHotelPricePlanInp struct {
	PlanName        string `json:"planName"`
	PropertyId      string `json:"propertyId"`
	RoomTypeId      string `json:"roomTypeId"`
	IsCancel        string `json:"isCancel"`
	PricePlanStatus string `json:"pricePlanStatus"`
	Type            int    `json:"type"  dc:"类型：0-正常价格计划 1-回收站"`
	input_form.PageReq
}

type SearchHotelPricePlanRes struct {
	Id             int    `json:"id"              dc:""`
	PlanName       string `json:"planName"        dc:"价格plan名称"`
	PlanShowName   string `json:"planShowName"     dc:"价格plan展示名称"`
	PropertyId     string `json:"propertyId"      dc:"物业ID"`
	PropertyDetail *struct {
		g.Meta `orm:"table:hg_pms_property"`
		*entity.PmsProperty
	} `json:"propertyDetail" orm:"with:uid=propertyId"`
	RoomTypeId     string `json:"roomTypeId"      dc:"房型ID"`
	RoomTypeDetail *struct {
		g.Meta `orm:"table:hg_pms_room_type"`
		*entity.PmsRoomType
	} `json:"roomTypeDetail" orm:"with:uid=room_type_id"`
	BookingDays       int                      `json:"bookingDays"     dc:"预订天数"`
	MemberGroupId     string                   `json:"memberGroupId"   dc:"用户组"`
	MemberGroupDetail []*entity.PmsMemberGroup `json:"memberGroupDetail" dc:"用户组详情"`
	MemberLevelId     string                   `json:"memberLevelId"   dc:"会员等级"`
	MemberLevelDetail []*entity.PmsMemberLevel `json:"memberLevelDetail" dc:"会员等级详情"`
	IsCancel          string                   `json:"isCancel"        dc:"是否可取消   Y  是  N  否"`
	IsOpenPriceMode   string                   `json:"isOpenPriceMode" dc:"是否开启价格模式"`
	PriceMode         string                   `json:"priceMode"       dc:"模式    +  贵   - 便宜"`
	PriceStandard     string                   `json:"priceStandard"   dc:"基准  PERCENT 倍率  AMOUNT  金额"`
	PlanValue         int                      `json:"planValue"       dc:"价格基准值"`
	PlanTips          string                   `json:"planTips"        dc:"价格计划提示"`
	PricePlanStatus   string                   `json:"pricePlanStatus" dc:"Y 开启  N 关闭"`
	CreatedAt         *gtime.Time              `json:"createdAt"       dc:"创建时间"`
	UpdatedAt         *gtime.Time              `json:"updatedAt"       dc:"更新时间"`
	DeletedAt         *gtime.Time              `json:"deletedAt"       dc:"删除时间"`
}

type EditHotelPricePlanInp struct {
	Id                   int                          `json:"id"                       v:"" dc:""`
	PlanName             string                       `json:"planName"                 v:"required#价格计划名称必填" dc:"价格plan名称"`
	PlanShowName         string                       `json:"planShowName"             v:"" dc:"价格plan展示名称"`
	PlanShowNameLanguage input_language.LanguageModel `json:"planShowNameLanguage"     v:"required#多语言价格计划名称必填" dc:"多语言价格plan展示名称"`
	PropertyId           string                       `json:"propertyId"               v:"required#物业必填" dc:"物业ID"`
	RoomTypeId           string                       `json:"roomTypeId"               v:"required#房型必填" dc:"房型ID"`
	BookingDays          int                          `json:"bookingDays"              v:"" dc:"预订天数"`
	MemberGroupId        string                       `json:"memberGroupId"            v:"" dc:"用户组"`
	MemberLevelId        string                       `json:"memberLevelId"            v:"" dc:"会员等级"`
	IsCancel             string                       `json:"isCancel"                 v:"required#取消政策必填" dc:"是否可取消   Y  是  N  否"`
	IsOpenPriceMode      string                       `json:"isOpenPriceMode"          v:"" dc:"是否开启价格模式"`
	PriceMode            string                       `json:"priceMode"                v:"required#价格模式必填" dc:"模式    +  贵   - 便宜"`
	PriceStandard        string                       `json:"priceStandard"            v:"required#价格基准必填" dc:"基准  PERCENT 倍率  AMOUNT  金额"`
	PlanValue            int                          `json:"planValue"                v:"required#价格基准值必填" dc:"价格基准值"`
	PlanTips             string                       `json:"planTips"                 v:"" dc:"价格计划提示"`
	PricePlanStatus      string                       `json:"pricePlanStatus"          v:"" dc:"Y 开启  N 关闭"`
}

type DeleteHotelPricePlanInp struct {
	Id int `json:"id"              dc:""`
}

// PricePlanStatusInp 更新价格计划状态
type PricePlanStatusInp struct {
	Id     int    `json:"id" v:"required#id不能为空" dc:"id"`
	Status string `json:"status" dc:"状态"`
}

func (in *PricePlanStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("id不能为空")
		return
	}
	return
}

type PricePlanStatusModel struct{}
