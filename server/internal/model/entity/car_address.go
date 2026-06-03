// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CarAddress is the golang structure for table car_address.
type CarAddress struct {
	Id            int         `json:"id"            orm:"id"             description:""`
	TypeId        uint        `json:"typeId"        orm:"type_id"        description:"地点类型id"`
	Name          string      `json:"name"          orm:"name"           description:"地点名称（后台）"`
	SubName       string      `json:"subName"       orm:"sub_name"       description:"地点名称（app-多语）"`
	AirportCode   string      `json:"airportCode"   orm:"airport_code"   description:"机场代码"`
	TerminalName  string      `json:"terminalName"  orm:"terminal_name"  description:"机场航站楼名称"`
	DetailAddress string      `json:"detailAddress" orm:"detail_address" description:"详细地址-多语"`
	GgLat         string      `json:"ggLat"         orm:"gg_lat"         description:"谷歌纬度"`
	GgLng         string      `json:"ggLng"         orm:"gg_lng"         description:"谷歌经度"`
	Lat           string      `json:"lat"           orm:"lat"            description:"纬度"`
	Lng           string      `json:"lng"           orm:"lng"            description:"经度"`
	Status        uint        `json:"status"        orm:"status"         description:"状态1、启用 2、禁用"`
	PropertyId    int         `json:"propertyId"    orm:"property_id"    description:"物业id"`
	CreateAt      *gtime.Time `json:"createAt"      orm:"create_at"      description:"创建时间"`
	UpdateAt      *gtime.Time `json:"updateAt"      orm:"update_at"      description:"更新时间"`
	DeletedAt     *gtime.Time `json:"deletedAt"     orm:"deleted_at"     description:"删除时间"`
}
