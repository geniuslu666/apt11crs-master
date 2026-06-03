// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsAvailabilities is the golang structure for table pms_availabilities.
type PmsAvailabilities struct {
	Id        int         `json:"id"        orm:"id"         description:"主键"`
	Puid      string      `json:"puid"      orm:"puid"       description:"物业UID"`
	Tuid      string      `json:"tuid"      orm:"tuid"       description:"房型UID"`
	Date      string      `json:"date"      orm:"date"       description:"日期"`
	Allotment int         `json:"allotment" orm:"allotment"  description:"库存"`
	Currency  string      `json:"currency"  orm:"currency"   description:"货币"`
	Price     float64     `json:"price"     orm:"price"      description:"实时价格"`
	BasePrice float64     `json:"basePrice" orm:"base_price" description:"基础价格"`
	Close     int         `json:"close"     orm:"close"      description:"1、开启 2、关闭"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`
}
