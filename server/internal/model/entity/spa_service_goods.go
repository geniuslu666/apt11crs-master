// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SpaServiceGoods is the golang structure for table spa_service_goods.
type SpaServiceGoods struct {
	Id               int64       `json:"id"               orm:"id"                 description:""`
	ServiceId        uint        `json:"serviceId"        orm:"service_id"         description:"服务ID"`
	GoodsName        string      `json:"goodsName"        orm:"goods_name"         description:"服务套餐名称(多语言)"`
	Image            string      `json:"image"            orm:"image"              description:"图片"`
	DetailImage      string      `json:"detailImage"      orm:"detail_image"       description:"详情图"`
	Price            float64     `json:"price"            orm:"price"              description:"套餐售价"`
	Duration         int         `json:"duration"         orm:"duration"           description:"时长(分钟)"`
	TotalOrderNum    int         `json:"totalOrderNum"    orm:"total_order_num"    description:"预约单总数量（包含退款）"`
	TotalOrderAmount float64     `json:"totalOrderAmount" orm:"total_order_amount" description:"预约单总金额（包含退款）"`
	PayOrderNum      int         `json:"payOrderNum"      orm:"pay_order_num"      description:"预约单支付数量（不包含退款）"`
	PayOrderAmount   float64     `json:"payOrderAmount"   orm:"pay_order_amount"   description:"预约单支付金额（不包含退款）"`
	Status           uint        `json:"status"           orm:"status"             description:"状态1、启用 2、禁用"`
	CreateAt         *gtime.Time `json:"createAt"         orm:"create_at"          description:"创建时间"`
	UpdateAt         *gtime.Time `json:"updateAt"         orm:"update_at"          description:"更新时间"`
	DeletedAt        *gtime.Time `json:"deletedAt"        orm:"deleted_at"         description:""`
}
