package input_car

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/utility/validate"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gmeta"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// CarCarUpdateFields 修改车辆字段过滤
type CarCarUpdateFields struct {
	CarName          string `json:"carName" dc:"车辆名称"`
	TypeId           int    `json:"typeId"                dc:"车型ID"`
	Brand            string `json:"brand" dc:"车辆品牌型号"`
	LicenseNo        string `json:"licenseNo"           dc:"车牌号码"`
	LicenseColor     string `json:"licenseColor"       dc:"牌照颜色"`
	QualityMaterials string `json:"qualityMaterials" dc:"资质信息(多图)"`
	Status           int    `json:"status"    dc:"状态"`
	Sort             int    `json:"sort"      dc:"排序"`
}

// CarCarInsertFields 新增车辆字段过滤
type CarCarInsertFields struct {
	CarName          string `json:"carName" dc:"车辆名称"`
	TypeId           int    `json:"typeId"                dc:"车型ID"`
	Brand            string `json:"brand" dc:"车辆品牌型号"`
	LicenseNo        string `json:"licenseNo"           dc:"车牌号码"`
	LicenseColor     string `json:"licenseColor"       dc:"牌照颜色"`
	QualityMaterials string `json:"qualityMaterials" dc:"资质信息(多图)"`
	Status           int    `json:"status"    dc:"状态"`
	Sort             int    `json:"sort"      dc:"排序"`
}

// CarCarEditInp 修改/新增车辆
type CarCarEditInp struct {
	entity.CarCar
}

func (in *CarCarEditInp) Filter(ctx context.Context) (err error) {
	return
}

type CarCarEditModel struct{}

// CarCarDeleteInp 删除车辆
type CarCarDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *CarCarDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type CarCarDeleteModel struct{}

// CarCarViewInp 获取指定车辆信息
type CarCarViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *CarCarViewInp) Filter(ctx context.Context) (err error) {
	return
}

type CarCarViewModel struct {
	entity.CarCar
}

// CarCarListInp 获取车辆列表
type CarCarListInp struct {
	input_form.PageReq
}

func (in *CarCarListInp) Filter(ctx context.Context) (err error) {
	return
}

type CarCarListModel struct {
	Id               int         `json:"id"        dc:"id"`
	CarName          string      `json:"carName" dc:"车辆名称"`
	TypeId           int         `json:"typeId"                dc:"车型ID"`
	Brand            string      `json:"brand" dc:"车辆品牌型号"`
	LicenseNo        string      `json:"licenseNo"           dc:"车牌号码"`
	LicenseColor     string      `json:"licenseColor"       dc:"牌照颜色"`
	QualityMaterials string      `json:"qualityMaterials" dc:"资质信息(多图)"`
	Status           int         `json:"status"    dc:"状态"`
	WorkStatus       string      `json:"workStatus"    dc:"工作状态"`
	Sort             int         `json:"sort"      dc:"排序"`
	CreateAt         *gtime.Time `json:"createAt"  dc:"创建时间"`
	UpdateAt         *gtime.Time `json:"updateAt"  dc:"更新时间"`
	CarTypeDetail    *struct {
		gmeta.Meta `orm:"table:hg_car_car_type"`
		*entity.CarCarType
	} `json:"carTypeDetail" orm:"with:id=type_id"  dc:"车辆车型"`
}

// CarCarMaxSortInp 获取车辆最大排序
type CarCarMaxSortInp struct{}

func (in *CarCarMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type CarCarMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// CarCarStatusInp 更新车辆状态
type CarCarStatusInp struct {
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
	Status int `json:"status" dc:"状态"`
}

func (in *CarCarStatusInp) Filter(ctx context.Context) (err error) {
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

type CarCarStatusModel struct{}

// CarCarWorkStatusInp 更新工作状态
type CarCarWorkStatusInp struct {
	Id         int    `json:"id" v:"required#id不能为空" dc:"id"`
	WorkStatus string `json:"workStatus" dc:"状态"`
}

func (in *CarCarWorkStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("id不能为空")
		return
	}

	if g.IsEmpty(in.WorkStatus) {
		err = gerror.New("状态不能为空")
		return
	}

	return
}

type CarCarWorkStatusModel struct{}
