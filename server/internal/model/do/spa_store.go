// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SpaStore is the golang structure of table hg_spa_store for DAO operations like Where/Data.
type SpaStore struct {
	g.Meta        `orm:"table:hg_spa_store, do:true"`
	Id            interface{} //
	Name          interface{} // 名称
	HeadName      interface{} // 负责人
	PhoneArea     interface{} // 区号
	Phone         interface{} // 联系电话
	OpenTime      interface{} // 营业时间
	AreaPid       interface{} // 地区省级ID
	AreaId        interface{} // 地区市级ID
	DetailAddress interface{} // 详细地址
	GgLat         interface{} // 谷歌纬度
	GgLng         interface{} // 谷歌经度
	Lat           interface{} // 纬度
	Lng           interface{} // 经度
	CreateAt      *gtime.Time // 创建时间
	UpdateAt      *gtime.Time // 更新时间
	DeletedAt     *gtime.Time //
}
