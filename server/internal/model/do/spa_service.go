// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SpaService is the golang structure of table hg_spa_service for DAO operations like Where/Data.
type SpaService struct {
	g.Meta           `orm:"table:hg_spa_service, do:true"`
	Id               interface{} //
	IspId            interface{} // 服务商ID
	Name             interface{} // 服务名称(多语言)
	SubName          interface{} // 副标题(多语言)
	LabelIds         interface{} // 标签（多选）
	Images           interface{} // 图集
	Channel          interface{} // 服务渠道，1-到店和上门 2-仅上门 3-仅到店
	PropertyIds      interface{} // 适用物业(以逗号分割)
	ServiceState     interface{} // 状态（1-立即上架 2-放入仓库）
	Sort             interface{} // 排序(越大越靠前)
	Content          interface{} // 服务详情(多语言)
	TotalOrderNum    interface{} // 预约单总数量（包含退款）
	TotalOrderAmount interface{} // 预约单总金额（包含退款）
	PayOrderNum      interface{} // 预约单支付数量（不包含退款）
	PayOrderAmount   interface{} // 预约单支付金额（不包含退款）
	CreateAt         *gtime.Time // 创建时间
	UpdateAt         *gtime.Time // 更新时间
	DeletedAt        *gtime.Time //
}
