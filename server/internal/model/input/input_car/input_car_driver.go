package input_car

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/utility/validate"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// CarDriverUpdateFields 修改司机管理字段过滤
type CarDriverUpdateFields struct {
	Name             string  `json:"name"                    dc:"真实姓名"`
	Nickname         string  `json:"nickname"              dc:"昵称"`
	CooperateTypeId  int     `json:"cooperateTypeId"                dc:"合作类型ID"`
	Sex              int     `json:"sex"                dc:"性别"`
	Phone            string  `json:"phone"                   dc:"手机号"`
	PhoneArea        string  `json:"phoneArea"                dc:"区号"`
	Photo            string  `json:"photo"                 dc:"照片"`
	Age              int     `json:"age"                 dc:"年龄"`
	WorkYears        int     `json:"workYears"                  dc:"从业年数"`
	Language         string  `json:"language"           dc:"语言能力"`
	Status           int     `json:"status"                  dc:"状态"`
	QualityMaterials string  `json:"qualityMaterials" dc:"资质信息(多图)"`
	SettlementId     int     `json:"settlementId"            dc:"结算模式ID"`
	SettlementType   int     `json:"settlementType"          dc:"分成抽成类型 1跟随系统  2自定义"`
	SettlementRate   float64 `json:"settlementRate"          dc:"分成结算比例"`
}

// CarDriverBasicUpdateFields 修改司机基础信息字段过滤
type CarDriverBasicUpdateFields struct {
	Name            string `json:"name"                    dc:"真实姓名"`
	Nickname        string `json:"nickname"              dc:"昵称"`
	Sex             int    `json:"sex"                dc:"性别"`
	Phone           string `json:"phone"                   dc:"手机号"`
	PhoneArea       string `json:"phoneArea"                dc:"区号"`
	Photo           string `json:"photo"                 dc:"照片"`
	Age             int    `json:"age"                 dc:"年龄"`
	WorkYears       int    `json:"workYears"                  dc:"从业年数"`
	Language        string `json:"language"           dc:"语言能力"`
	CooperateTypeId int    `json:"cooperateTypeId"                dc:"合作类型ID"`
	Status          int    `json:"status"                  dc:"状态"`
}

// CarDriverQualificationUpdateFields 修改司机资质信息信息字段过滤
type CarDriverQualificationUpdateFields struct {
	QualityMaterials string `json:"qualityMaterials" dc:"资质信息(多图)"`
}

// CarDriverSettleUpdateFields 修改司机结算信息字段过滤
type CarDriverSettleUpdateFields struct {
	SettlementId   int     `json:"settlementId"            dc:"结算模式ID"`
	SettlementType int     `json:"settlementType"          dc:"分成抽成类型 1跟随系统  2自定义"`
	SettlementRate float64 `json:"settlementRate"          dc:"分成结算比例"`
}

// CarDriverInsertFields 新增司机管理字段过滤
type CarDriverInsertFields struct {
	Name             string  `json:"name"                    dc:"真实姓名"`
	Nickname         string  `json:"nickname"              dc:"昵称"`
	CooperateTypeId  int     `json:"cooperateTypeId"                dc:"合作类型ID"`
	Sex              int     `json:"sex"                dc:"性别"`
	Phone            string  `json:"phone"                   dc:"手机号"`
	PhoneArea        string  `json:"phoneArea"                dc:"区号"`
	Photo            string  `json:"photo"                 dc:"照片"`
	Age              int     `json:"age"                 dc:"年龄"`
	WorkYears        int     `json:"workYears"                  dc:"从业年数"`
	Language         string  `json:"language"           dc:"语言能力"`
	Status           int     `json:"status"                  dc:"状态"`
	QualityMaterials string  `json:"qualityMaterials" dc:"资质信息(多图)"`
	SettlementId     int     `json:"settlementId"            dc:"结算模式ID"`
	SettlementType   int     `json:"settlementType"          dc:"分成抽成类型 1跟随系统  2自定义"`
	SettlementRate   float64 `json:"settlementRate"          dc:"分成结算比例"`
}

// CarDriverEditInp 修改/新增司机管理
type CarDriverEditInp struct {
	entity.CarDriver
	Type string `json:"type"     dc:"修改类型：basic、qualification、settle"`
}

func (in *CarDriverEditInp) Filter(ctx context.Context) (err error) {

	return
}

type CarDriverEditModel struct{}

// CarDriverDeleteInp 删除司机管理
type CarDriverDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *CarDriverDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type CarDriverDeleteModel struct{}

// CarDriverViewInp 获取指定司机管理信息
type CarDriverViewInp struct {
	Id         int64 `json:"id" v:"required#id不能为空" dc:"id"`
	IsLanguage bool  `json:"isLanguage" dc:"是否获取多语言数据"`
}

func (in *CarDriverViewInp) Filter(ctx context.Context) (err error) {
	return
}

type CarDriverViewModel struct {
	entity.CarDriver
}

// CarDriverListInp 获取司机管理列表
type CarDriverListInp struct {
	input_form.PageReq
	Keywords        string        `json:"keywords" dc:"真实姓名/昵称/手机号"`
	Status          int           `json:"status"             dc:"状态"`
	CooperateTypeId int           `json:"cooperateTypeId"  dc:"合作类型ID"`
	WorkStatus      string        `json:"workStatus"             dc:"工作状态"`
	OrderType       string        `json:"orderType"             dc:"排序"`
	CreateAt        []*gtime.Time `json:"createdAt" dc:"加入时间"`
}

func (in *CarDriverListInp) Filter(ctx context.Context) (err error) {
	return
}

type CarDriverListModel struct {
	Id                  int64       `json:"id"                      dc:"id"`
	Photo               string      `json:"photo"                    dc:"照片"`
	Nickname            string      `json:"nickname"              dc:"昵称"`
	Name                string      `json:"name"                dc:"真实姓名"`
	Phone               string      `json:"phone"         dc:"手机号"`
	PhoneArea           string      `json:"phoneArea"                    dc:"区号"`
	CooperateTypeId     int         `json:"cooperateTypeId"         dc:"合作类型ID"`
	Status              int         `json:"status"    dc:"状态"`
	IsLeader            int         `json:"isLeader"    dc:"是否是车队长 1是  2否"`
	CarId               int         `json:"carId"    dc:"车辆ID"`
	WorkStatus          string      `json:"workStatus"       dc:"工作状态"`
	WorkStatusEnum      int         `json:"workStatusEnum"       dc:"工作状态"`
	MemberId            int         `json:"memberId"         dc:"会员ID"`
	CreateAt            *gtime.Time `json:"createAt"                dc:"创建时间"`
	UpdateAt            *gtime.Time `json:"updateAt"                dc:"更新时间"`
	CooperateTypeDetail *struct {
		gmeta.Meta `orm:"table:hg_car_cooperate_type"`
		*entity.CarCooperateType
	} `json:"cooperateTypeDetail" orm:"with:id=cooperate_type_id"  dc:"合作类型"`
	MemberDetail *struct {
		gmeta.Meta `orm:"table:hg_pms_member"`
		Id         int    `json:"id"    orm:"id"      dc:"id"`
		FullName   string `json:"fullName"     orm:"full_name"     dc:"全名"`
		Avatar     string `json:"avatar"       orm:"avatar"        dc:"头像"`
		Phone      string `json:"phone"        orm:"phone"         dc:"手机号"`
		PhoneArea  string `json:"phoneArea"    orm:"phone_area"    dc:"手机区号"`
		MemberNo   string `json:"memberNo"    orm:"member_no"    dc:"会员号"`
	} `json:"memberDetail" orm:"with:id=member_id" dc:"会员信息"`
}

type CarDriverAllListModel struct {
	Id        int64  `json:"id"                      dc:"id"`
	Photo     string `json:"photo"                    dc:"照片"`
	Nickname  string `json:"nickname"              dc:"昵称"`
	Name      string `json:"name"                dc:"真实姓名"`
	Phone     string `json:"phone"         dc:"手机号"`
	PhoneArea string `json:"phoneArea"                    dc:"区号"`
}

// CarDriverStatusInp 更新司机状态
type CarDriverStatusInp struct {
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
	Status int `json:"status" dc:"状态"`
}

func (in *CarDriverStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("id不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}

	return
}

type CarDriverStatusModel struct{}

// CarDriverWorkStatusInp 更新工作状态
type CarDriverWorkStatusInp struct {
	Id         int    `json:"id" v:"required#id不能为空" dc:"id"`
	WorkStatus string `json:"workStatus" dc:"状态"`
}

func (in *CarDriverWorkStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("id不能为空")
		return
	}

	if g.IsEmpty(in.WorkStatus) {
		err = gerror.New("状态不能为空")
		return
	}

	return
}

type CarDriverWorkStatusModel struct{}

// CarDriverBindInp 绑定用户
type CarDriverBindInp struct {
	Id       int `json:"id" dc:"ID"`
	MemberId int `json:"memberId" dc:"会员ID"`
	IsLeader int `json:"isLeader" dc:"是否是车队长 1是  2否"`
}

func (in *CarDriverBindInp) Filter(ctx context.Context) (err error) {
	// 验证会员
	if in.MemberId == 0 {
		err = gerror.New("请选择会员")
		return
	}

	return
}

type CarDriverBindModel struct{}

// CarDriverUnbindInp 解绑渠道管理
type CarDriverUnbindInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *CarDriverUnbindInp) Filter(ctx context.Context) (err error) {
	return
}

type CarDriverUnbindModel struct{}

// CarDriverBindCarInp 绑定司机
type CarDriverBindCarInp struct {
	Id    int `json:"id" dc:"ID" v:"required#司机ID不能为空"`
	CarId int `json:"carId" dc:"司机ID"`
}

func (in *CarDriverBindCarInp) Filter(ctx context.Context) (err error) {
	return
}

type CarDriverBindCarModel struct{}
