// Package sysin

package input_hotel

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_language"
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/util/gmeta"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsRoomTypeUpdateFields 修改房型字段过滤
type PmsRoomTypeUpdateFields struct {
	Uid                    string  `json:"uid"          dc:"房型ID"`
	Puid                   string  `json:"puid"         dc:"物业ID"`
	Cover                  string  `json:"cover"        dc:"封面"`
	CoverList              string  `json:"coverList"    dc:"照片墙"`
	Name                   string  `json:"name"         dc:"房型名称"`
	CheckinAt              string  `json:"checkinAt"    dc:"入住时间"`
	CheckoutAt             string  `json:"checkoutAt"   dc:"退房时间"`
	Occupancy              int     `json:"occupancy"    dc:"占用"`
	Size                   string  `json:"size"         dc:"房间大小"`
	Bedrooms               string  `json:"bedrooms"     dc:"卧室"`
	RoomNum                string  `json:"roomNum"      dc:"房间数量"`
	RatePlanId             string  `json:"ratePlanId"   dc:"费率ID"`
	AdditionalGuestAmounts float64 `json:"additionalGuestAmounts" dc:"额外客人金额"`
	OccupantsForBaseRate   int     `json:"occupantsForBaseRate"   dc:"无需增加额外客人金额人数"`
}

// PmsRoomTypeInsertFields 新增房型字段过滤
type PmsRoomTypeInsertFields struct {
	Uid                    string  `json:"uid"          dc:"房型ID"`
	Puid                   string  `json:"puid"         dc:"物业ID"`
	Cover                  string  `json:"cover"        dc:"封面"`
	CoverList              string  `json:"coverList"    dc:"照片墙"`
	Name                   string  `json:"name"         dc:"房型名称"`
	BasePrice              float64 `json:"basePrice"     dc:"最低价格"`
	CheckinAt              string  `json:"checkinAt"    dc:"入住时间"`
	CheckoutAt             string  `json:"checkoutAt"   dc:"退房时间"`
	BookingStyle           string  `json:"bookingStyle" dc:"预订方式"`
	RoomStyle              string  `json:"roomStyle"    dc:"房间风格"`
	Occupancy              int     `json:"occupancy"    dc:"占用"`
	Size                   string  `json:"size"         dc:"房间大小"`
	Bedrooms               string  `json:"bedrooms"     dc:"卧室"`
	Bathrooms              string  `json:"bathrooms"    dc:"浴室"`
	CleaningFee            float64 `json:"cleaningFee"  dc:"清理费"`
	AdditionalGuestAmounts float64 `json:"additionalGuestAmounts" dc:"额外客人金额"`
	OccupantsForBaseRate   int     `json:"occupantsForBaseRate"   dc:"无需增加额外客人金额人数"`
}

// PmsRoomTypeEditInp 修改/新增房型
type PmsRoomTypeEditInp struct {
	entity.PmsRoomType
	BedType []*struct {
		BedTypeName string `json:"bedTypeName"     dc:"床型名称"`
		BedWidth    string `json:"bedWidth"     dc:"床宽"`
		BedNum      int    `json:"bedNum"     dc:"数量"`
	}
	NameLanguage input_language.LanguageModel `json:"nameLanguage"          dc:"多语言房型名称"`
}

func (in *PmsRoomTypeEditInp) Filter(ctx context.Context) (err error) {
	// 验证房型ID
	if err := g.Validator().Rules("required").Data(in.Uid).Messages("房型ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证物业ID
	if err := g.Validator().Rules("required").Data(in.Puid).Messages("物业ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type PmsRoomTypeEditModel struct{}

// PmsRoomTypeDeleteInp 删除房型
type PmsRoomTypeDeleteInp struct {
	Id interface{} `json:"id" v:"required#主键ID不能为空" dc:"主键ID"`
}

func (in *PmsRoomTypeDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsRoomTypeDeleteModel struct{}

// PmsRoomTypeViewInp 获取指定房型信息
type PmsRoomTypeViewInp struct {
	Id  int    `json:"id" v:"required-without:uid#主键ID不能为空" dc:"主键ID"`
	Uid string `json:"uid" v:"required-without:id#请输入房型UID" dc:"房型UID"`
}

func (in *PmsRoomTypeViewInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsRoomTypeViewModel struct {
	entity.PmsRoomType
	NameLanguage []*LanguageType `json:"nameLanguage"         dc:"房型名称"   orm:"with:uuid=name"`
	BedTypeDes   string          `json:"bedType" dc:"床型信息"`
	BedTypeStr   string          `json:"bedTypeStr" dc:"床型信息(显示床型名称)"`
	BedType      []*struct {
		gmeta.Meta  `orm:"table:hg_pms_bed_type"`
		Id          int64   `json:"id"                  orm:"id"                     description:"床型ID"`
		BedTypeName string  `json:"bedTypeName"         orm:"bed_type_name"          description:"床型名称"`
		BedWidth    float64 `json:"bedWidth"            orm:"bed_width"              description:"床宽"`
		BedNum      float64 `json:"bedNum"              orm:"bed_num"                description:"数量"`
	} `json:"bedTypes,omitempty" orm:"with:room_type_id=id" dc:"床型信息"`
	RoomRatePlanInfos []*struct {
		gmeta.Meta `orm:"table:hg_pms_room_rate_plan"`
		*entity.PmsRoomRatePlan
	} `json:"roomRatePlanInfos,omitempty" orm:"with:tuid=uid" dc:"房型价格计划"`
}

// PmsRoomTypeListInp 获取房型列表
type PmsRoomTypeListInp struct {
	input_form.PageReq
	Puid            string        `json:"puid"            dc:"物业ID"`
	Uid             string        `json:"uid"             dc:"房型ID"`
	Id              int           `json:"id"              dc:"主键ID"`
	Name            string        `json:"name"            dc:"房型名称"`
	CreateAt        []*gtime.Time `json:"createAt"        dc:"创建时间"`
	PmsPropertyName string        `json:"pmsPropertyName" dc:"物业名称"`
}

func (in *PmsRoomTypeListInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsRoomTypeListModel struct {
	Id           int             `json:"id"              dc:"主键ID"`
	Uid          string          `json:"uid"             dc:"第三方系统的ID"`
	Puid         string          `json:"puid"             dc:"puid"`
	Cover        string          `json:"cover"           dc:"封面"`
	CoverList    *gjson.Json     `json:"coverList"       dc:"照片墙"`
	Name         string          `json:"name"            dc:"房型名称"`
	NameLanguage []*LanguageType `json:"nameLanguage"        dc:"多语言房型名称"   orm:"with:uuid=name"`
	BasePrice    float64         `json:"basePrice"        dc:"最低价格"`
	CheckinAt    string          `json:"checkinAt"       dc:"入住时间"`
	CheckoutAt   string          `json:"checkoutAt"      dc:"退房时间"`
	BookingStyle string          `json:"bookingStyle"    dc:"预订方式"`
	RoomStyle    string          `json:"roomStyle"       dc:"房间风格"`
	Occupancy    int             `json:"occupancy"       dc:"占用"`
	Size         string          `json:"size"            dc:"房间大小"`
	Bedrooms     string          `json:"bedrooms"        dc:"卧室"`
	Bathrooms    string          `json:"bathrooms"       dc:"浴室"`
	RoomNum      int             `json:"roomNum"         dc:"房间数量"`
	CleaningFee  float64         `json:"cleaningFee"     dc:"清理费"`
	IsShow       int             `json:"isShow"          dc:"是否显示 1、显示 0 隐藏"`
	CreateAt     *gtime.Time     `json:"createAt"        dc:"创建时间"`
	UpdateAt     *gtime.Time     `json:"updateAt"        dc:"更新时间"`
	Property     *struct {
		gmeta.Meta `orm:"table:hg_pms_property"`
		*entity.PmsProperty
	} `json:"property" orm:"with:uid=puid" dc:"物业信息"`
}

// PmsRoomTypeIsShowInp 更新房型状态
type PmsRoomTypeIsShowInp struct {
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
	IsShow int `json:"isShow" dc:"是否显示 1、显示 0 隐藏"`
}
