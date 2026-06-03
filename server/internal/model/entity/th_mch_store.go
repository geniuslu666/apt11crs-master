// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ThMchStore is the golang structure for table th_mch_store.
type ThMchStore struct {
	Id                 int64       `json:"id"                 orm:"id"                   description:""`
	MchId              uint64      `json:"mchId"              orm:"mch_id"               description:"商户ID"`
	StoreName          string      `json:"storeName"          orm:"store_name"           description:"门店名称（多语）"`
	Images             string      `json:"images"             orm:"images"               description:"图集"`
	PhoneArea          string      `json:"phoneArea"          orm:"phone_area"           description:"区号"`
	Phone              string      `json:"phone"              orm:"phone"                description:"电话"`
	DetailAddress      string      `json:"detailAddress"      orm:"detail_address"       description:"详细地址"`
	GgLat              string      `json:"ggLat"              orm:"gg_lat"               description:"谷歌纬度"`
	GgLng              string      `json:"ggLng"              orm:"gg_lng"               description:"谷歌经度"`
	VerifyNum          int         `json:"verifyNum"          orm:"verify_num"           description:"核销数量"`
	Account            string      `json:"account"            orm:"account"              description:"账号"`
	PasswordHash       string      `json:"passwordHash"       orm:"password_hash"        description:"密码"`
	Salt               string      `json:"salt"               orm:"salt"                 description:"密码盐"`
	PasswordResetToken string      `json:"passwordResetToken" orm:"password_reset_token" description:"密码重置令牌"`
	Status             uint        `json:"status"             orm:"status"               description:"1、启用 2、禁用"`
	CreateAt           *gtime.Time `json:"createAt"           orm:"create_at"            description:"创建时间"`
	UpdateAt           *gtime.Time `json:"updateAt"           orm:"update_at"            description:"更新时间"`
	DeletedAt          *gtime.Time `json:"deletedAt"          orm:"deleted_at"           description:""`
}
