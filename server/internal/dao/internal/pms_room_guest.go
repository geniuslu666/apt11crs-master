// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsRoomGuestDao is the data access object for the table hg_pms_room_guest.
type PmsRoomGuestDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of the current DAO.
	columns PmsRoomGuestColumns // columns contains all the column names of Table for convenient usage.
}

// PmsRoomGuestColumns defines and stores column names for the table hg_pms_room_guest.
type PmsRoomGuestColumns struct {
	Uid            string // 外部ID
	Id             string // 主键ID
	FirstName      string // 名
	LastName       string // 姓
	FirstNameKana  string // 名的假名
	LastNameKana   string // 姓的假名
	FullName       string // 全名
	Language       string // 语言
	Email          string // 电子邮件
	Phone          string // 电话
	Nationality    string // 国籍
	Address        string // 地址
	IsMainGuest    string // 是否为主要客人
	Gender         string // 性别
	Dob            string // 出生日期
	VisaNo         string // 护照号或身份证号
	Occupation     string // 职业
	PostalCode     string // 邮政编码
	LastPortEmbark string // 最后登船港口
	NextPortEmbark string // 下一个登船港口
	Photo          string // 护照或身份证照片
	Selfie         string // 客人自拍照片
	Signature      string // 客人签名
	IdentPhoto     string // 客人到达酒店后拍摄的照片
	CreatedAt      string //
	UpdatedAt      string //
}

// pmsRoomGuestColumns holds the columns for the table hg_pms_room_guest.
var pmsRoomGuestColumns = PmsRoomGuestColumns{
	Uid:            "uid",
	Id:             "id",
	FirstName:      "first_name",
	LastName:       "last_name",
	FirstNameKana:  "first_name_kana",
	LastNameKana:   "last_name_kana",
	FullName:       "full_name",
	Language:       "language",
	Email:          "email",
	Phone:          "phone",
	Nationality:    "nationality",
	Address:        "address",
	IsMainGuest:    "is_main_guest",
	Gender:         "gender",
	Dob:            "dob",
	VisaNo:         "visa_no",
	Occupation:     "occupation",
	PostalCode:     "postal_code",
	LastPortEmbark: "last_port_embark",
	NextPortEmbark: "next_port_embark",
	Photo:          "photo",
	Selfie:         "selfie",
	Signature:      "signature",
	IdentPhoto:     "ident_photo",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewPmsRoomGuestDao creates and returns a new DAO object for table data access.
func NewPmsRoomGuestDao() *PmsRoomGuestDao {
	return &PmsRoomGuestDao{
		group:   "default",
		table:   "hg_pms_room_guest",
		columns: pmsRoomGuestColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsRoomGuestDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsRoomGuestDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsRoomGuestDao) Columns() PmsRoomGuestColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsRoomGuestDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsRoomGuestDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsRoomGuestDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
