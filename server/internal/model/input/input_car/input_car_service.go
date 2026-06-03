package input_car

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/utility/validate"
	"context"
	"github.com/gogf/gf/v2/util/gmeta"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// CarServiceUpdateFields 修改服务字段过滤
type CarServiceUpdateFields struct {
	ServiceType     string  `json:"serviceType" dc:"服务类型"`
	ServiceName     string  `json:"serviceName" dc:"名称"`
	CarTypeId       int     `json:"carTypeId"    dc:"车型ID"`
	StartIds        string  `json:"startIds"     dc:"出发地ID  多选"`
	EndIds          string  `json:"endIds"     dc:"目的地ID 多选"`
	Distance        int     `json:"distance"    dc:"路线长度 KM"`
	UseTime         int     `json:"useTime"    dc:"线路时长  分钟"`
	Price           float64 `json:"price"        dc:"基础费用"`
	LinePrice       float64 `json:"linePrice"        dc:"划线价"`
	FreeWaitTime    int     `json:"freeWaitTime"    dc:"免费等待时长  分钟"`
	MaxWaitTime     int     `json:"maxWaitTime"    dc:"最大等待时长  分钟"`
	TimeoutPreTime  int     `json:"timeoutPreTime"    dc:"超时每xx分钟"`
	TimeoutPrePrice float64 `json:"timeoutPrePrice"        dc:"超时价格"`
	Status          int     `json:"status"    dc:"状态"`
	Sort            int     `json:"sort"      dc:"排序"`
}

// CarServiceInsertFields 新增服务字段过滤
type CarServiceInsertFields struct {
	ServiceType     string  `json:"serviceType" dc:"服务类型"`
	ServiceName     string  `json:"serviceName" dc:"名称"`
	CarTypeId       int     `json:"carTypeId"    dc:"车型ID"`
	StartIds        string  `json:"startIds"     dc:"出发地ID  多选"`
	EndIds          string  `json:"endIds"     dc:"目的地ID 多选"`
	Distance        int     `json:"distance"    dc:"路线长度 KM"`
	UseTime         int     `json:"useTime"    dc:"线路时长  分钟"`
	Price           float64 `json:"price"        dc:"基础费用"`
	LinePrice       float64 `json:"linePrice"        dc:"划线价"`
	FreeWaitTime    int     `json:"freeWaitTime"    dc:"免费等待时长  分钟"`
	MaxWaitTime     int     `json:"maxWaitTime"    dc:"最大等待时长  分钟"`
	TimeoutPreTime  int     `json:"timeoutPreTime"    dc:"超时每xx分钟"`
	TimeoutPrePrice float64 `json:"timeoutPrePrice"        dc:"超时价格"`
	Status          int     `json:"status"    dc:"状态"`
	Sort            int     `json:"sort"      dc:"排序"`
}

// CarServiceEditInp 修改/新增服务
type CarServiceEditInp struct {
	entity.CarService
}

func (in *CarServiceEditInp) Filter(ctx context.Context) (err error) {
	return
}

type CarServiceEditModel struct{}

// CarServiceDeleteInp 删除服务
type CarServiceDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *CarServiceDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type CarServiceDeleteModel struct{}

// CarServiceViewInp 获取指定服务信息
type CarServiceViewInp struct {
	Id         int  `json:"id" v:"required#id不能为空" dc:"id"`
	IsLanguage bool `json:"isLanguage" dc:"是否获取多语言数据"`
}

func (in *CarServiceViewInp) Filter(ctx context.Context) (err error) {
	return
}

type CarServiceViewModel struct {
	entity.CarService
}

// CarServiceListInp 获取服务列表
type CarServiceListInp struct {
	input_form.PageReq
	ServiceType string `json:"serviceType" dc:"服务类型"`
	Sort        string `json:"sort"      dc:"排序"`
}

func (in *CarServiceListInp) Filter(ctx context.Context) (err error) {
	return
}

type CarServiceListModel struct {
	Id              int         `json:"id"        dc:"id"`
	ServiceType     string      `json:"serviceType" dc:"服务类型"`
	ServiceName     string      `json:"serviceName" dc:"名称"`
	CarTypeId       int         `json:"carTypeId"    dc:"车型ID"`
	StartIds        string      `json:"startIds"     dc:"出发地ID  多选"`
	EndIds          string      `json:"endIds"     dc:"目的地ID 多选"`
	Distance        int         `json:"distance"    dc:"路线长度 KM"`
	UseTime         int         `json:"useTime"    dc:"线路时长  分钟"`
	Price           float64     `json:"price"        dc:"基础费用"`
	FreeWaitTime    int         `json:"freeWaitTime"    dc:"免费等待时长  分钟"`
	MaxWaitTime     int         `json:"maxWaitTime"    dc:"最大等待时长  分钟"`
	TimeoutPreTime  int         `json:"timeoutPreTime"    dc:"超时每xx分钟"`
	TimeoutPrePrice float64     `json:"timeoutPrePrice"        dc:"超时价格"`
	Status          int         `json:"status"    dc:"状态"`
	Sort            int         `json:"sort"      dc:"排序"`
	CreateAt        *gtime.Time `json:"createAt"  dc:"创建时间"`
	UpdateAt        *gtime.Time `json:"updateAt"  dc:"更新时间"`
	CarTypeDetail   *struct {
		gmeta.Meta `orm:"table:hg_car_car_type"`
		*entity.CarCarType
	} `json:"carTypeDetail" orm:"with:id=car_type_id"  dc:"车辆车型"`
	StartServiceAddress []*struct {
		gmeta.Meta `orm:"table:hg_car_service_address"`
		*entity.CarServiceAddress
		AddressDetail *struct {
			gmeta.Meta `orm:"table:hg_car_address"`
			*entity.CarAddress
		} `json:"addressDetail" orm:"with:id=address_id"  dc:"地点详情"`
	} `json:"startServiceAddress" orm:"with:service_id=id, where:type=1"  dc:"出发机场"`
	EndServiceAddress []*struct {
		gmeta.Meta `orm:"table:hg_car_service_address"`
		*entity.CarServiceAddress
		AddressDetail *struct {
			gmeta.Meta `orm:"table:hg_car_address"`
			*entity.CarAddress
		} `json:"addressDetail" orm:"with:id=address_id"  dc:"地点详情"`
	} `json:"endServiceAddress" orm:"with:service_id=id, where:type=2"  dc:"目的地"`
}

// CarServiceMaxSortInp 获取服务最大排序
type CarServiceMaxSortInp struct{}

func (in *CarServiceMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type CarServiceMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// CarServiceStatusInp 更新服务状态
type CarServiceStatusInp struct {
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
	Status int `json:"status" dc:"状态"`
}

func (in *CarServiceStatusInp) Filter(ctx context.Context) (err error) {
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

type CarServiceStatusModel struct{}

// CarServiceSortInp 编辑服务排序
type CarServiceSortInp struct {
	Id   interface{} `json:"id" v:"required#主键不能为空" dc:"主键"`
	Sort int         `json:"sort"        dc:"排序"`
}

func (in *CarServiceSortInp) Filter(ctx context.Context) (err error) {
	return
}

type CarServiceSortModel struct{}
