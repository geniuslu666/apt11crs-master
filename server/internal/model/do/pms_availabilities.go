// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsAvailabilities is the golang structure of table hg_pms_availabilities for DAO operations like Where/Data.
type PmsAvailabilities struct {
	g.Meta    `orm:"table:hg_pms_availabilities, do:true"`
	Id        interface{} // 主键
	Puid      interface{} // 物业UID
	Tuid      interface{} // 房型UID
	Date      interface{} // 日期
	Allotment interface{} // 库存
	Currency  interface{} // 货币
	Price     interface{} // 实时价格
	BasePrice interface{} // 基础价格
	Close     interface{} // 1、开启 2、关闭
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
