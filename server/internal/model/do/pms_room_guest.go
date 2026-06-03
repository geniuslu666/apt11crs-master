// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsRoomGuest is the golang structure of table hg_pms_room_guest for DAO operations like Where/Data.
type PmsRoomGuest struct {
	g.Meta         `orm:"table:hg_pms_room_guest, do:true"`
	Uid            interface{} // 外部ID
	Id             interface{} // 主键ID
	FirstName      interface{} // 名
	LastName       interface{} // 姓
	FirstNameKana  interface{} // 名的假名
	LastNameKana   interface{} // 姓的假名
	FullName       interface{} // 全名
	Language       interface{} // 语言
	Email          interface{} // 电子邮件
	Phone          interface{} // 电话
	Nationality    interface{} // 国籍
	Address        interface{} // 地址
	IsMainGuest    interface{} // 是否为主要客人
	Gender         interface{} // 性别
	Dob            *gtime.Time // 出生日期
	VisaNo         interface{} // 护照号或身份证号
	Occupation     interface{} // 职业
	PostalCode     interface{} // 邮政编码
	LastPortEmbark interface{} // 最后登船港口
	NextPortEmbark interface{} // 下一个登船港口
	Photo          interface{} // 护照或身份证照片
	Selfie         interface{} // 客人自拍照片
	Signature      interface{} // 客人签名
	IdentPhoto     interface{} // 客人到达酒店后拍摄的照片
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
}
