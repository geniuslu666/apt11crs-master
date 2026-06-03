// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CarService is the golang structure for table car_service.
type CarService struct {
	Id               int         `json:"id"               orm:"id"                 description:""`
	ServiceType      string      `json:"serviceType"      orm:"service_type"       description:"服务类型"`
	ServiceName      string      `json:"serviceName"      orm:"service_name"       description:"路线名称"`
	CarTypeId        int         `json:"carTypeId"        orm:"car_type_id"        description:"车型ID"`
	StartIds         string      `json:"startIds"         orm:"start_ids"          description:"出发地ID  多选"`
	EndIds           string      `json:"endIds"           orm:"end_ids"            description:"目的地ID 多选"`
	Distance         int         `json:"distance"         orm:"distance"           description:"路线长度 KM"`
	UseTime          int         `json:"useTime"          orm:"use_time"           description:"线路时长  分钟"`
	Price            float64     `json:"price"            orm:"price"              description:"基础费用"`
	LinePrice        float64     `json:"linePrice"        orm:"line_price"         description:"划线价"`
	FreeWaitTime     int         `json:"freeWaitTime"     orm:"free_wait_time"     description:"免费等待时长  分钟"`
	MaxWaitTime      int         `json:"maxWaitTime"      orm:"max_wait_time"      description:"最大等待时长  分钟"`
	TimeoutPreTime   int         `json:"timeoutPreTime"   orm:"timeout_pre_time"   description:"超时每xx分钟"`
	TimeoutPrePrice  float64     `json:"timeoutPrePrice"  orm:"timeout_pre_price"  description:"超时价格"`
	Status           uint        `json:"status"           orm:"status"             description:"状态1、启用 2、禁用"`
	Sort             int         `json:"sort"             orm:"sort"               description:"排序(越大越靠前)"`
	TotalOrderNum    int         `json:"totalOrderNum"    orm:"total_order_num"    description:"预约单总数量（包含退款）"`
	TotalOrderAmount float64     `json:"totalOrderAmount" orm:"total_order_amount" description:"预约单总金额（包含退款）"`
	PayOrderNum      int         `json:"payOrderNum"      orm:"pay_order_num"      description:"预约单支付数量（不包含退款）"`
	PayOrderAmount   float64     `json:"payOrderAmount"   orm:"pay_order_amount"   description:"预约单支付金额（不包含退款）"`
	CreateAt         *gtime.Time `json:"createAt"         orm:"create_at"          description:"创建时间"`
	UpdateAt         *gtime.Time `json:"updateAt"         orm:"update_at"          description:"更新时间"`
	DeletedAt        *gtime.Time `json:"deletedAt"        orm:"deleted_at"         description:"删除时间"`
}
