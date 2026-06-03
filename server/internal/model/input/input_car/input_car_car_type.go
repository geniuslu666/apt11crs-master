package input_car

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_language"
	"APT/utility/validate"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// CarCarTypeUpdateFields 修改车辆车型字段过滤
type CarCarTypeUpdateFields struct {
	Name                string  `json:"Name" dc:"名称"`
	Desc                string  `json:"desc" dc:"简介"`
	Image               string  `json:"image" dc:"图片"`
	SeatNum             int     `json:"seatNum"             dc:"座位数"`
	PassengerNum        int     `json:"passengerNum"        dc:"建议乘员人数"`
	MaxPackageNum       int     `json:"maxPackageNum"       dc:"最大容纳行李件数"`
	ChildrenSeatSupport string  `json:"childrenSeatSupport" dc:"是否支持儿童座椅"`
	ChildrenSeatSpace   float64 `json:"childrenSeatSpace"   dc:"儿童座椅占几个作为"`
	Content             string  `json:"content"             dc:"内容"`
	Status              int     `json:"status"    dc:"状态"`
	Sort                int     `json:"sort"      dc:"排序"`
}

// CarCarTypeInsertFields 新增车辆车型字段过滤
type CarCarTypeInsertFields struct {
	Name                string  `json:"Name" dc:"名称"`
	Desc                string  `json:"desc" dc:"简介"`
	Image               string  `json:"image" dc:"图片"`
	SeatNum             int     `json:"seatNum"             dc:"座位数"`
	PassengerNum        int     `json:"passengerNum"        dc:"建议乘员人数"`
	MaxPackageNum       int     `json:"maxPackageNum"       dc:"最大容纳行李件数"`
	ChildrenSeatSupport string  `json:"childrenSeatSupport" dc:"是否支持儿童座椅"`
	ChildrenSeatSpace   float64 `json:"childrenSeatSpace"   dc:"儿童座椅占几个作为"`
	Content             string  `json:"content"             dc:"内容"`
	Status              int     `json:"status"    dc:"状态"`
	Sort                int     `json:"sort"      dc:"排序"`
}

// CarCarTypeEditInp 修改/新增车辆车型
type CarCarTypeEditInp struct {
	entity.CarCarType
	NameLanguage input_language.LanguageModel `json:"nameLanguage"          dc:"多语言名称"`
	DescLanguage input_language.LanguageModel `json:"descLanguage"          dc:"多语言简介"`
}

func (in *CarCarTypeEditInp) Filter(ctx context.Context) (err error) {
	return
}

type CarCarTypeEditModel struct{}

// CarCarTypeDeleteInp 删除车辆车型
type CarCarTypeDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *CarCarTypeDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type CarCarTypeDeleteModel struct{}

// CarCarTypeViewInp 获取指定车辆车型信息
type CarCarTypeViewInp struct {
	Id         int  `json:"id" v:"required#id不能为空" dc:"id"`
	IsLanguage bool `json:"isLanguage" dc:"是否获取多语言数据"`
}

func (in *CarCarTypeViewInp) Filter(ctx context.Context) (err error) {
	return
}

type CarCarTypeViewModel struct {
	entity.CarCarType
	NameLanguage []*input_hotel.LanguageType `json:"nameLanguage"         dc:"名称"   orm:"with:uuid=name"`
	DescLanguage []*input_hotel.LanguageType `json:"descLanguage"         dc:"简介"   orm:"with:uuid=desc"`
}

// CarCarTypeListInp 获取车辆车型列表
type CarCarTypeListInp struct {
	input_form.PageReq
}

func (in *CarCarTypeListInp) Filter(ctx context.Context) (err error) {
	return
}

type CarCarTypeListModel struct {
	Id                  int         `json:"id"        dc:"id"`
	Name                string      `json:"name" dc:"名称"`
	Desc                string      `json:"desc" dc:"简介"`
	Image               string      `json:"image" dc:"图片"`
	SeatNum             int         `json:"seatNum"             dc:"座位数"`
	PassengerNum        int         `json:"passengerNum"        dc:"建议乘员人数"`
	MaxPackageNum       int         `json:"maxPackageNum"       dc:"最大容纳行李件数"`
	ChildrenSeatSupport string      `json:"childrenSeatSupport" dc:"是否支持儿童座椅"`
	ChildrenSeatSpace   float64     `json:"childrenSeatSpace"   dc:"儿童座椅占几个作为"`
	Status              int         `json:"status"    dc:"状态"`
	Sort                int         `json:"sort"      dc:"排序"`
	CreateAt            *gtime.Time `json:"createAt"  dc:"创建时间"`
	UpdateAt            *gtime.Time `json:"updateAt"  dc:"更新时间"`
}

// CarCarTypeMaxSortInp 获取车辆车型最大排序
type CarCarTypeMaxSortInp struct{}

func (in *CarCarTypeMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type CarCarTypeMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// CarCarTypeStatusInp 更新车辆车型状态
type CarCarTypeStatusInp struct {
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
	Status int `json:"status" dc:"状态"`
}

func (in *CarCarTypeStatusInp) Filter(ctx context.Context) (err error) {
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

type CarCarTypeStatusModel struct{}
