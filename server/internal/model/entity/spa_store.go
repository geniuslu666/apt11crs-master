// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SpaStore is the golang structure for table spa_store.
type SpaStore struct {
	Id            int         `json:"id"            orm:"id"             description:""`
	Name          string      `json:"name"          orm:"name"           description:"名称"`
	HeadName      string      `json:"headName"      orm:"head_name"      description:"负责人"`
	PhoneArea     string      `json:"phoneArea"     orm:"phone_area"     description:"区号"`
	Phone         string      `json:"phone"         orm:"phone"          description:"联系电话"`
	OpenTime      string      `json:"openTime"      orm:"open_time"      description:"营业时间"`
	AreaPid       int         `json:"areaPid"       orm:"area_pid"       description:"地区省级ID"`
	AreaId        int         `json:"areaId"        orm:"area_id"        description:"地区市级ID"`
	DetailAddress string      `json:"detailAddress" orm:"detail_address" description:"详细地址"`
	GgLat         string      `json:"ggLat"         orm:"gg_lat"         description:"谷歌纬度"`
	GgLng         string      `json:"ggLng"         orm:"gg_lng"         description:"谷歌经度"`
	Lat           string      `json:"lat"           orm:"lat"            description:"纬度"`
	Lng           string      `json:"lng"           orm:"lng"            description:"经度"`
	CreateAt      *gtime.Time `json:"createAt"      orm:"create_at"      description:"创建时间"`
	UpdateAt      *gtime.Time `json:"updateAt"      orm:"update_at"      description:"更新时间"`
	DeletedAt     *gtime.Time `json:"deletedAt"     orm:"deleted_at"     description:""`
}
