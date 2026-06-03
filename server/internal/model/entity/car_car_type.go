// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CarCarType is the golang structure for table car_car_type.
type CarCarType struct {
	Id                  int         `json:"id"                  orm:"id"                    description:""`
	Name                string      `json:"name"                orm:"name"                  description:"名称"`
	Desc                string      `json:"desc"                orm:"desc"                  description:"简介"`
	Image               string      `json:"image"               orm:"image"                 description:"图片"`
	SeatNum             int         `json:"seatNum"             orm:"seat_num"              description:"座位数"`
	PassengerNum        int         `json:"passengerNum"        orm:"passenger_num"         description:"建议乘员人数"`
	MaxPackageNum       int         `json:"maxPackageNum"       orm:"max_package_num"       description:"最大容纳行李件数"`
	ChildrenSeatSupport string      `json:"childrenSeatSupport" orm:"children_seat_support" description:"是否支持儿童座椅"`
	ChildrenSeatSpace   float64     `json:"childrenSeatSpace"   orm:"children_seat_space"   description:"儿童座椅占几个作为"`
	Content             string      `json:"content"             orm:"content"               description:"内容"`
	Status              uint        `json:"status"              orm:"status"                description:"状态1、启用 2、禁用"`
	Sort                int         `json:"sort"                orm:"sort"                  description:"排序(越大越靠前)"`
	CreateAt            *gtime.Time `json:"createAt"            orm:"create_at"             description:"创建时间"`
	UpdateAt            *gtime.Time `json:"updateAt"            orm:"update_at"             description:"更新时间"`
	DeletedAt           *gtime.Time `json:"deletedAt"           orm:"deleted_at"            description:"删除时间"`
}
