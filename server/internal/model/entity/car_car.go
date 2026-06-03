// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CarCar is the golang structure for table car_car.
type CarCar struct {
	Id                  int         `json:"id"                  orm:"id"                    description:""`
	CarName             string      `json:"carName"             orm:"car_name"              description:"车辆名称"`
	TypeId              int         `json:"typeId"              orm:"type_id"               description:"车型ID"`
	Brand               string      `json:"brand"               orm:"brand"                 description:"车辆品牌型号"`
	LicenseNo           string      `json:"licenseNo"           orm:"license_no"            description:"车牌号码"`
	LicenseColor        string      `json:"licenseColor"        orm:"license_color"         description:"牌照颜色"`
	SeatNum             int         `json:"seatNum"             orm:"seat_num"              description:"座位数(废弃)"`
	PassengerNum        int         `json:"passengerNum"        orm:"passenger_num"         description:"建议乘员人数(废弃)"`
	MaxPackageNum       int         `json:"maxPackageNum"       orm:"max_package_num"       description:"最大容纳行李件数(废弃)"`
	ChildrenSeatSupport string      `json:"childrenSeatSupport" orm:"children_seat_support" description:"是否支持儿童座椅(废弃)"`
	ChildrenSeatSpace   float64     `json:"childrenSeatSpace"   orm:"children_seat_space"   description:"儿童座椅占几个作为(废弃)"`
	WorkStatus          string      `json:"workStatus"          orm:"work_status"           description:"工作状态"`
	Status              uint        `json:"status"              orm:"status"                description:"状态1、启用 2、禁用"`
	Sort                int         `json:"sort"                orm:"sort"                  description:"排序(越大越靠前)"`
	QualityMaterials    string      `json:"qualityMaterials"    orm:"quality_materials"     description:"资质信息(多图)"`
	TotalOrderNum       int         `json:"totalOrderNum"       orm:"total_order_num"       description:"预约单总数量（包含退款）"`
	TotalOrderAmount    float64     `json:"totalOrderAmount"    orm:"total_order_amount"    description:"预约单总金额（包含退款）"`
	PayOrderNum         int         `json:"payOrderNum"         orm:"pay_order_num"         description:"预约单支付数量（不包含退款）"`
	PayOrderAmount      float64     `json:"payOrderAmount"      orm:"pay_order_amount"      description:"预约单支付金额（不包含退款）"`
	CreateAt            *gtime.Time `json:"createAt"            orm:"create_at"             description:"创建时间"`
	UpdateAt            *gtime.Time `json:"updateAt"            orm:"update_at"             description:"更新时间"`
	DeletedAt           *gtime.Time `json:"deletedAt"           orm:"deleted_at"            description:""`
}
