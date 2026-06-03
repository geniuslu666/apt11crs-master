// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SpaService is the golang structure for table spa_service.
type SpaService struct {
	Id               int64       `json:"id"               orm:"id"                 description:""`
	IspId            int         `json:"ispId"            orm:"isp_id"             description:"服务商ID"`
	Name             string      `json:"name"             orm:"name"               description:"服务名称(多语言)"`
	SubName          string      `json:"subName"          orm:"sub_name"           description:"副标题(多语言)"`
	LabelIds         string      `json:"labelIds"         orm:"label_ids"          description:"标签（多选）"`
	Images           string      `json:"images"           orm:"images"             description:"图集"`
	Channel          int         `json:"channel"          orm:"channel"            description:"服务渠道，1-到店和上门 2-仅上门 3-仅到店"`
	PropertyIds      string      `json:"propertyIds"      orm:"property_ids"       description:"适用物业(以逗号分割)"`
	ServiceState     int         `json:"serviceState"     orm:"service_state"      description:"状态（1-立即上架 2-放入仓库）"`
	Sort             int         `json:"sort"             orm:"sort"               description:"排序(越大越靠前)"`
	Content          string      `json:"content"          orm:"content"            description:"服务详情(多语言)"`
	TotalOrderNum    int         `json:"totalOrderNum"    orm:"total_order_num"    description:"预约单总数量（包含退款）"`
	TotalOrderAmount float64     `json:"totalOrderAmount" orm:"total_order_amount" description:"预约单总金额（包含退款）"`
	PayOrderNum      int         `json:"payOrderNum"      orm:"pay_order_num"      description:"预约单支付数量（不包含退款）"`
	PayOrderAmount   float64     `json:"payOrderAmount"   orm:"pay_order_amount"   description:"预约单支付金额（不包含退款）"`
	CreateAt         *gtime.Time `json:"createAt"         orm:"create_at"          description:"创建时间"`
	UpdateAt         *gtime.Time `json:"updateAt"         orm:"update_at"          description:"更新时间"`
	DeletedAt        *gtime.Time `json:"deletedAt"        orm:"deleted_at"         description:""`
}
