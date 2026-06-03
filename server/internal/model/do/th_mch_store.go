// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ThMchStore is the golang structure of table hg_th_mch_store for DAO operations like Where/Data.
type ThMchStore struct {
	g.Meta             `orm:"table:hg_th_mch_store, do:true"`
	Id                 interface{} //
	MchId              interface{} // 商户ID
	StoreName          interface{} // 门店名称（多语）
	Images             interface{} // 图集
	PhoneArea          interface{} // 区号
	Phone              interface{} // 电话
	DetailAddress      interface{} // 详细地址
	GgLat              interface{} // 谷歌纬度
	GgLng              interface{} // 谷歌经度
	VerifyNum          interface{} // 核销数量
	Account            interface{} // 账号
	PasswordHash       interface{} // 密码
	Salt               interface{} // 密码盐
	PasswordResetToken interface{} // 密码重置令牌
	Status             interface{} // 1、启用 2、禁用
	CreateAt           *gtime.Time // 创建时间
	UpdateAt           *gtime.Time // 更新时间
	DeletedAt          *gtime.Time //
}
