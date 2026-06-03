// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsRoomGuest is the golang structure for table pms_room_guest.
type PmsRoomGuest struct {
	Uid            string      `json:"uid"            orm:"uid"              description:"外部ID"`
	Id             int         `json:"id"             orm:"id"               description:"主键ID"`
	FirstName      string      `json:"firstName"      orm:"first_name"       description:"名"`
	LastName       string      `json:"lastName"       orm:"last_name"        description:"姓"`
	FirstNameKana  string      `json:"firstNameKana"  orm:"first_name_kana"  description:"名的假名"`
	LastNameKana   string      `json:"lastNameKana"   orm:"last_name_kana"   description:"姓的假名"`
	FullName       string      `json:"fullName"       orm:"full_name"        description:"全名"`
	Language       string      `json:"language"       orm:"language"         description:"语言"`
	Email          string      `json:"email"          orm:"email"            description:"电子邮件"`
	Phone          string      `json:"phone"          orm:"phone"            description:"电话"`
	Nationality    string      `json:"nationality"    orm:"nationality"      description:"国籍"`
	Address        string      `json:"address"        orm:"address"          description:"地址"`
	IsMainGuest    int         `json:"isMainGuest"    orm:"is_main_guest"    description:"是否为主要客人"`
	Gender         string      `json:"gender"         orm:"gender"           description:"性别"`
	Dob            *gtime.Time `json:"dob"            orm:"dob"              description:"出生日期"`
	VisaNo         string      `json:"visaNo"         orm:"visa_no"          description:"护照号或身份证号"`
	Occupation     string      `json:"occupation"     orm:"occupation"       description:"职业"`
	PostalCode     string      `json:"postalCode"     orm:"postal_code"      description:"邮政编码"`
	LastPortEmbark string      `json:"lastPortEmbark" orm:"last_port_embark" description:"最后登船港口"`
	NextPortEmbark string      `json:"nextPortEmbark" orm:"next_port_embark" description:"下一个登船港口"`
	Photo          string      `json:"photo"          orm:"photo"            description:"护照或身份证照片"`
	Selfie         string      `json:"selfie"         orm:"selfie"           description:"客人自拍照片"`
	Signature      string      `json:"signature"      orm:"signature"        description:"客人签名"`
	IdentPhoto     string      `json:"identPhoto"     orm:"ident_photo"      description:"客人到达酒店后拍摄的照片"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"       description:""`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"       description:""`
}
