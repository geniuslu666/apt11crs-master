// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SpaIspBalanceChange is the golang structure of table hg_spa_isp_balance_change for DAO operations like Where/Data.
type SpaIspBalanceChange struct {
	g.Meta            `orm:"table:hg_spa_isp_balance_change, do:true"`
	Id                interface{} // 主键
	IspId             interface{} // 服务商ID
	Type              interface{} // 金额变动方式   SETTLEMENT 结算   VERIFY核账    RETURN_ORDER退单    SYS 系统调整
	ChangePrice       interface{} // 变更金额
	SettlementOrderId interface{} // 结算单ID
	WithdrawOrderId   interface{} // 提现单ID
	Des               interface{} // 描述
	CreatedAt         *gtime.Time // 创建时间
	UpdatedAt         *gtime.Time // 更新时间
}
