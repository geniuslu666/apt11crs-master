// Package input_car
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.0.0
package input_car

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

// CarAddressUpdateFields 修改接送机地点管理字段过滤
type CarAddressUpdateFields struct {
	TypeId        int    `json:"typeId"        dc:"地点类型id"`
	Name          string `json:"name"          dc:"地点名称（后台）"`
	SubName       string `json:"subName"       dc:"地点名称（app-多语）"`
	AirportCode   string `json:"airportCode"   dc:"机场代码"`
	TerminalName  string `json:"terminalName"  dc:"航站楼"`
	DetailAddress string `json:"detailAddress" dc:"详细地址-多语"`
	GgLat         string `json:"ggLat"         dc:"谷歌纬度"`
	GgLng         string `json:"ggLng"         dc:"谷歌经度"`
	Lat           string `json:"lat"           dc:"纬度"`
	Lng           string `json:"lng"           dc:"经度"`
	Status        int    `json:"status"        dc:"状态1、启用 2、禁用"`
	PropertyId    int    `json:"propertyId"    dc:"物业id"`
}

// CarAddressInsertFields 新增接送机地点管理字段过滤
type CarAddressInsertFields struct {
	TypeId        int    `json:"typeId"        dc:"地点类型id"`
	Name          string `json:"name"          dc:"地点名称（后台）"`
	SubName       string `json:"subName"       dc:"地点名称（app-多语）"`
	AirportCode   string `json:"airportCode"   dc:"机场代码"`
	TerminalName  string `json:"terminalName"  dc:"航站楼"`
	DetailAddress string `json:"detailAddress" dc:"详细地址-多语"`
	GgLat         string `json:"ggLat"         dc:"谷歌纬度"`
	GgLng         string `json:"ggLng"         dc:"谷歌经度"`
	Lat           string `json:"lat"           dc:"纬度"`
	Lng           string `json:"lng"           dc:"经度"`
	Status        int    `json:"status"        dc:"状态1、启用 2、禁用"`
	PropertyId    int    `json:"propertyId"    dc:"物业id"`
}

// CarAddressEditInp 修改/新增接送机地点管理
type CarAddressEditInp struct {
	entity.CarAddress
	NameLanguage    input_language.LanguageModel `json:"nameLanguage"          dc:"多语言地点名称"`
	AddressLanguage input_language.LanguageModel `json:"addressLanguage"          dc:"多语言详细地址"`
}

func (in *CarAddressEditInp) Filter(ctx context.Context) (err error) {
	// 验证地点类型id
	if err := g.Validator().Rules("required").Data(in.TypeId).Messages("地点类型id不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证地点名称（后台）
	if err := g.Validator().Rules("required").Data(in.Name).Messages("地点名称（后台）不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证谷歌纬度
	if err := g.Validator().Rules("required").Data(in.GgLat).Messages("谷歌纬度不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证谷歌经度
	if err := g.Validator().Rules("required").Data(in.GgLng).Messages("谷歌经度不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证纬度
	if err := g.Validator().Rules("required").Data(in.Lat).Messages("纬度不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证经度
	if err := g.Validator().Rules("required").Data(in.Lng).Messages("经度不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证状态1、启用 2、禁用
	if err := g.Validator().Rules("required").Data(in.Status).Messages("状态1、启用 2、禁用不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:1,2").Data(in.Status).Messages("状态1、启用 2、禁用值不正确").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type CarAddressEditModel struct{}

// CarAddressDeleteInp 删除接送机地点管理
type CarAddressDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *CarAddressDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type CarAddressDeleteModel struct{}

// CarAddressViewInp 获取指定接送机地点管理信息
type CarAddressViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *CarAddressViewInp) Filter(ctx context.Context) (err error) {
	return
}

type CarAddressViewModel struct {
	entity.CarAddress
	NameLanguage      []*input_hotel.LanguageType `json:"nameLanguage"         dc:"地点名称"   orm:"with:uuid=sub_name"`
	AddressLanguage   []*input_hotel.LanguageType `json:"addressLanguage"      dc:"详细地址"   orm:"with:uuid=detail_address"`
	AddressTypeDetail *struct {
		gmeta.Meta `orm:"table:hg_car_address_type"`
		*entity.CarAddressType
	} `json:"addressTypeDetail" orm:"with:id=type_id"  dc:"地点类型"`
}

// CarAddressListInp 获取接送机地点管理列表
type CarAddressListInp struct {
	input_form.PageReq
	Id          int           `json:"id"          dc:"id"`
	Name        string        `json:"name"        dc:"地点名称（后台）"`
	AirportCode string        `json:"airportCode" dc:"机场代码"`
	Status      int           `json:"status"      dc:"状态1、启用 2、禁用"`
	CreateAt    []*gtime.Time `json:"createAt"    dc:"创建时间"`
}

func (in *CarAddressListInp) Filter(ctx context.Context) (err error) {
	return
}

type CarAddressListModel struct {
	Id                int         `json:"id"            dc:"id"`
	TypeId            int         `json:"typeId"        dc:"地点类型id"`
	Name              string      `json:"name"          dc:"地点名称（后台）"`
	SubName           string      `json:"subName"       dc:"地点名称（app-多语）"`
	AirportCode       string      `json:"airportCode"   dc:"机场代码"`
	DetailAddress     string      `json:"detailAddress" dc:"详细地址-多语"`
	GgLat             string      `json:"ggLat"         dc:"谷歌纬度"`
	GgLng             string      `json:"ggLng"         dc:"谷歌经度"`
	Lat               string      `json:"lat"           dc:"纬度"`
	Lng               string      `json:"lng"           dc:"经度"`
	Status            int         `json:"status"        dc:"状态1、启用 2、禁用"`
	CreateAt          *gtime.Time `json:"createAt"      dc:"创建时间"`
	UpdateAt          *gtime.Time `json:"updateAt"      dc:"更新时间"`
	AddressTypeDetail *struct {
		gmeta.Meta `orm:"table:hg_car_address_type"`
		*entity.CarAddressType
	} `json:"addressTypeDetail" orm:"with:id=type_id"  dc:"地点类型"`
}

// CarAddressExportModel 导出接送机地点管理
type CarAddressExportModel struct {
	Id            int         `json:"id"            dc:"id"`
	TypeId        int         `json:"typeId"        dc:"地点类型id"`
	Name          string      `json:"name"          dc:"地点名称（后台）"`
	SubName       string      `json:"subName"       dc:"地点名称（app-多语）"`
	AirportCode   string      `json:"airportCode"   dc:"机场代码"`
	DetailAddress string      `json:"detailAddress" dc:"详细地址-多语"`
	GgLat         string      `json:"ggLat"         dc:"谷歌纬度"`
	GgLng         string      `json:"ggLng"         dc:"谷歌经度"`
	Lat           string      `json:"lat"           dc:"纬度"`
	Lng           string      `json:"lng"           dc:"经度"`
	Status        int         `json:"status"        dc:"状态1、启用 2、禁用"`
	CreateAt      *gtime.Time `json:"createAt"      dc:"创建时间"`
	UpdateAt      *gtime.Time `json:"updateAt"      dc:"更新时间"`
}

// CarAddressStatusInp 更新接送机地点管理状态
type CarAddressStatusInp struct {
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
	Status int `json:"status" dc:"状态"`
}

func (in *CarAddressStatusInp) Filter(ctx context.Context) (err error) {
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

type CarAddressStatusModel struct{}
