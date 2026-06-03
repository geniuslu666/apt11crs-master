// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysTerminal is the golang structure for table sys_terminal.
type SysTerminal struct {
	Id           int         `json:"id"           orm:"id"            description:""`
	TerminalName string      `json:"terminalName" orm:"terminal_name" description:"终端名称"`
	TerminalType string      `json:"terminalType" orm:"terminal_type" description:"终端类型"`
	BrandModel   string      `json:"brandModel"   orm:"brand_model"   description:"品牌型号"`
	Sn           string      `json:"sn"           orm:"sn"            description:"终端编号"`
	OnlineStatus int         `json:"onlineStatus" orm:"online_status" description:"在线状态 1-在线 2离线"`
	StoreId      int         `json:"storeId"      orm:"store_id"      description:"绑定门店"`
	RestaurantId int         `json:"restaurantId" orm:"restaurant_id" description:"绑定餐厅"`
	PrintTimes   uint        `json:"printTimes"   orm:"print_times"   description:"打印联数(次数)"`
	CreateAt     *gtime.Time `json:"createAt"     orm:"create_at"     description:"创建时间"`
	UpdateAt     *gtime.Time `json:"updateAt"     orm:"update_at"     description:"更新时间"`
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"    description:""`
}
