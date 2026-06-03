package input_spa

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SpaStoreUpdateFields 修改门店管理字段过滤
type SpaStoreUpdateFields struct {
	Name          string `json:"name"          dc:"名称"`
	HeadName      string `json:"headName"      dc:"负责人"`
	PhoneArea     string `json:"phoneArea"     dc:"区号"`
	Phone         string `json:"phone"         dc:"联系电话"`
	OpenTime      string `json:"openTime"      dc:"营业时间"`
	AreaPid       int    `json:"areaPid"       dc:"地区省级ID"`
	AreaId        int    `json:"areaId"        dc:"地区市级ID"`
	DetailAddress string `json:"detailAddress" dc:"详细地址"`
	GgLat         string `json:"ggLat"         dc:"谷歌纬度"`
	GgLng         string `json:"ggLng"         dc:"谷歌经度"`
	Lat           string `json:"lat"           dc:"纬度"`
	Lng           string `json:"lng"           dc:"经度"`
}

// SpaStoreInsertFields 新增门店管理字段过滤
type SpaStoreInsertFields struct {
	Name          string `json:"name"          dc:"名称"`
	HeadName      string `json:"headName"      dc:"负责人"`
	PhoneArea     string `json:"phoneArea"     dc:"区号"`
	Phone         string `json:"phone"         dc:"联系电话"`
	OpenTime      string `json:"openTime"      dc:"营业时间"`
	AreaPid       int    `json:"areaPid"       dc:"地区省级ID"`
	AreaId        int    `json:"areaId"        dc:"地区市级ID"`
	DetailAddress string `json:"detailAddress" dc:"详细地址"`
	GgLat         string `json:"ggLat"         dc:"谷歌纬度"`
	GgLng         string `json:"ggLng"         dc:"谷歌经度"`
	Lat           string `json:"lat"           dc:"纬度"`
	Lng           string `json:"lng"           dc:"经度"`
}

// SpaStoreEditInp 修改/新增门店管理
type SpaStoreEditInp struct {
	entity.SpaStore
}

func (in *SpaStoreEditInp) Filter(ctx context.Context) (err error) {
	// 验证名称
	if err := g.Validator().Rules("required").Data(in.Name).Messages("名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证负责人
	if err := g.Validator().Rules("required").Data(in.HeadName).Messages("负责人不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证区号
	if err := g.Validator().Rules("required").Data(in.PhoneArea).Messages("区号不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证联系电话
	if err := g.Validator().Rules("required").Data(in.Phone).Messages("联系电话不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证营业时间
	if err := g.Validator().Rules("required").Data(in.OpenTime).Messages("营业时间不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证详细地址
	if err := g.Validator().Rules("required").Data(in.DetailAddress).Messages("详细地址不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type SpaStoreEditModel struct{}

// SpaStoreDeleteInp 删除门店管理
type SpaStoreDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *SpaStoreDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaStoreDeleteModel struct{}

// SpaStoreViewInp 获取指定门店管理信息
type SpaStoreViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *SpaStoreViewInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaStoreViewModel struct {
	entity.SpaStore
}

// SpaStoreListInp 获取门店管理列表
type SpaStoreListInp struct {
	input_form.PageReq
	Id int `json:"id" dc:"id"`
}

func (in *SpaStoreListInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaStoreListModel struct {
	Id            int         `json:"id"            dc:"id"`
	Name          string      `json:"name"          dc:"名称"`
	HeadName      string      `json:"headName"      dc:"负责人"`
	PhoneArea     string      `json:"phoneArea"     dc:"区号"`
	Phone         string      `json:"phone"         dc:"联系电话"`
	OpenTime      string      `json:"openTime"      dc:"营业时间"`
	AreaPid       int         `json:"areaPid"       dc:"地区省级ID"`
	AreaId        int         `json:"areaId"        dc:"地区市级ID"`
	DetailAddress string      `json:"detailAddress" dc:"详细地址"`
	GgLat         string      `json:"ggLat"         dc:"谷歌纬度"`
	GgLng         string      `json:"ggLng"         dc:"谷歌经度"`
	Lat           string      `json:"lat"           dc:"纬度"`
	Lng           string      `json:"lng"           dc:"经度"`
	CreateAt      *gtime.Time `json:"createAt"      dc:"创建时间"`
	UpdateAt      *gtime.Time `json:"updateAt"      dc:"更新时间"`
}
