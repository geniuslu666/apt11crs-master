package input_spa

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

// SpaTechnicianUpdateFields 修改技师管理字段过滤
type SpaTechnicianUpdateFields struct {
	IspId            int     `json:"ispId"       dc:"服务商ID"`
	Name             string  `json:"name"                    dc:"真实姓名"`
	Nickname         string  `json:"nickname"              dc:"昵称"`
	Sex              int     `json:"sex"                dc:"性别"`
	Phone            string  `json:"phone"                   dc:"手机号"`
	PhoneArea        string  `json:"phoneArea"                dc:"区号"`
	Photo            string  `json:"photo"                 dc:"照片"`
	Age              int     `json:"age"                 dc:"年龄"`
	WorkYears        int     `json:"workYears"                  dc:"从业年数"`
	Status           int     `json:"status"                  dc:"状态"`
	WechatNo         string  `json:"wechatNo"           dc:"微信号"`
	QualityMaterials string  `json:"qualityMaterials" dc:"资质信息(多图)"`
	SettlementId     int     `json:"settlementId"            dc:"结算模式ID"`
	SettlementType   int     `json:"settlementType"          dc:"分成抽成类型 1跟随系统  2自定义"`
	SettlementRate   float64 `json:"settlementRate"          dc:"分成结算比例"`
	OrderTimeType    int     `json:"orderTimeType"    dc:"预约日期  1每天 2自定义"`
	OrderTimeWeek    string  `json:"orderTimeWeek"    dc:"预约日期自定义周数"`
	RestTimeWeek     string  `json:"restTimeWeek"     dc:"定休日周数"`
}

// SpaTechnicianBasicUpdateFields 修改技师基础信息字段过滤
type SpaTechnicianBasicUpdateFields struct {
	IspId     int    `json:"ispId"       dc:"服务商ID"`
	Name      string `json:"name"                    dc:"真实姓名"`
	Nickname  string `json:"nickname"              dc:"昵称"`
	Sex       int    `json:"sex"                dc:"性别"`
	Phone     string `json:"phone"                   dc:"手机号"`
	PhoneArea string `json:"phoneArea"                dc:"区号"`
	Photo     string `json:"photo"                 dc:"照片"`
	Age       int    `json:"age"                 dc:"年龄"`
	WorkYears int    `json:"workYears"                  dc:"从业年数"`
	Status    int    `json:"status"                  dc:"状态"`
}

// SpaTechnicianQualificationUpdateFields 修改技师资质信息信息字段过滤
type SpaTechnicianQualificationUpdateFields struct {
	WechatNo         string `json:"wechatNo"           dc:"微信号"`
	QualityMaterials string `json:"qualityMaterials" dc:"资质信息(多图)"`
}

// SpaTechnicianSettleUpdateFields 修改技师结算信息字段过滤
type SpaTechnicianSettleUpdateFields struct {
	SettlementId   int     `json:"settlementId"            dc:"结算模式ID"`
	SettlementType int     `json:"settlementType"          dc:"分成抽成类型 1跟随系统  2自定义"`
	SettlementRate float64 `json:"settlementRate"          dc:"分成结算比例"`
}

// SpaTechnicianOrderUpdateFields 修改技师预约信息字段过滤
type SpaTechnicianOrderUpdateFields struct {
	OrderTimeType int    `json:"orderTimeType"    dc:"预约日期  1每天 2自定义"`
	OrderTimeWeek string `json:"orderTimeWeek"    dc:"预约日期自定义周数"`
	RestTimeWeek  string `json:"restTimeWeek"     dc:"定休日周数"`
}

// SpaTechnicianInsertFields 新增技师管理字段过滤
type SpaTechnicianInsertFields struct {
	IspId          int     `json:"ispId"       dc:"服务商ID"`
	Name           string  `json:"name"                    dc:"真实姓名"`
	Nickname       string  `json:"nickname"              dc:"昵称"`
	Sex            int     `json:"sex"                dc:"性别"`
	Phone          string  `json:"phone"                   dc:"手机号"`
	PhoneArea      string  `json:"phoneArea"                dc:"区号"`
	Photo          string  `json:"photo"                 dc:"照片"`
	Age            int     `json:"age"                 dc:"年龄"`
	WorkYears      int     `json:"workYears"                  dc:"从业年数"`
	Status         int     `json:"status"                  dc:"状态"`
	SettlementId   int     `json:"settlementId"            dc:"结算模式ID"`
	SettlementType int     `json:"settlementType"          dc:"分成抽成类型 1跟随系统  2自定义"`
	SettlementRate float64 `json:"settlementRate"          dc:"分成结算比例"`
}

// SpaTechnicianEditInp 修改/新增技师管理
type SpaTechnicianEditInp struct {
	entity.SpaTechnician
	Type string `json:"type"     dc:"修改类型：basic、qualification、settle、order"`
}

func (in *SpaTechnicianEditInp) Filter(ctx context.Context) (err error) {

	return
}

type SpaTechnicianEditModel struct{}

// SpaTechnicianDeleteInp 删除技师管理
type SpaTechnicianDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *SpaTechnicianDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaTechnicianDeleteModel struct{}

// SpaTechnicianViewInp 获取指定技师管理信息
type SpaTechnicianViewInp struct {
	Id         int64 `json:"id" v:"required#id不能为空" dc:"id"`
	IsLanguage bool  `json:"isLanguage" dc:"是否获取多语言数据"`
}

func (in *SpaTechnicianViewInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaTechnicianViewModel struct {
	entity.SpaTechnician
}

// SpaTechnicianListInp 获取技师管理列表
type SpaTechnicianListInp struct {
	input_form.PageReq
	Keywords   string        `json:"keywords" dc:"真实姓名/手机号"`
	Status     int           `json:"status"             dc:"状态"`
	WorkStatus string        `json:"workStatus"             dc:"工作状态"`
	CreateAt   []*gtime.Time `json:"createAt" dc:"加入时间"`
}

func (in *SpaTechnicianListInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaTechnicianListModel struct {
	Id           int64       `json:"id"                    dc:"id"`
	IspId        int         `json:"ispId"                 dc:"服务商ID"`
	IsLeader     int         `json:"isLeader"    dc:"是否是店长 1是  2否"`
	Photo        string      `json:"photo"                    dc:"照片"`
	Nickname     string      `json:"nickname"              dc:"昵称"`
	Name         string      `json:"name"                dc:"真实姓名"`
	Phone        string      `json:"phone"         dc:"手机号"`
	PhoneArea    string      `json:"phoneArea"                    dc:"区号"`
	Status       int         `json:"status"    dc:"状态"`
	WorkStatus   string      `json:"workStatus"       dc:"工作状态"`
	MemberId     int         `json:"memberId"         dc:"会员ID"`
	CreateAt     *gtime.Time `json:"createAt"                dc:"创建时间"`
	UpdateAt     *gtime.Time `json:"updateAt"                dc:"更新时间"`
	MemberDetail *struct {
		gmeta.Meta `orm:"table:hg_pms_member"`
		Id         int    `json:"id"    orm:"id"      dc:"id"`
		FullName   string `json:"fullName"     orm:"full_name"     dc:"全名"`
		Avatar     string `json:"avatar"       orm:"avatar"        dc:"头像"`
		Phone      string `json:"phone"        orm:"phone"         dc:"手机号"`
		PhoneArea  string `json:"phoneArea"    orm:"phone_area"    dc:"手机区号"`
		MemberNo   string `json:"memberNo"    orm:"member_no"    dc:"会员号"`
	} `json:"memberDetail" orm:"with:id=member_id" dc:"会员信息"`
	SpaIspDetail *struct {
		gmeta.Meta `orm:"table:hg_spa_isp"`
		Id         int    `json:"id"      orm:"id"    dc:""`
		Name       string `json:"name"    orm:"name"  dc:"服务商名称"`
	} `json:"spaIspDetail" orm:"with:id=isp_id" dc:"服务商信息"`
}

type SpaTechnicianAllListModel struct {
	Id        int64  `json:"id"                      dc:"id"`
	Photo     string `json:"photo"                    dc:"照片"`
	Nickname  string `json:"nickname"              dc:"昵称"`
	Name      string `json:"name"                dc:"真实姓名"`
	Phone     string `json:"phone"         dc:"手机号"`
	PhoneArea string `json:"phoneArea"                    dc:"区号"`
}

// SpaTechnicianStatusInp 更新技师状态
type SpaTechnicianStatusInp struct {
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
	Status int `json:"status" dc:"状态"`
}

func (in *SpaTechnicianStatusInp) Filter(ctx context.Context) (err error) {
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

type SpaTechnicianStatusModel struct{}

// SpaTechnicianWorkStatusInp 更新工作状态
type SpaTechnicianWorkStatusInp struct {
	Id         int    `json:"id" v:"required#id不能为空" dc:"id"`
	WorkStatus string `json:"workStatus" dc:"状态"`
}

func (in *SpaTechnicianWorkStatusInp) Filter(ctx context.Context) (err error) {
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

type SpaTechnicianWorkStatusModel struct{}

// SpaTechnicianBindInp 绑定用户
type SpaTechnicianBindInp struct {
	Id       int `json:"id" dc:"ID"`
	MemberId int `json:"memberId" dc:"会员ID"`
	IsLeader int `json:"isLeader" dc:"是否是店长 1是  2否"`
}

func (in *SpaTechnicianBindInp) Filter(ctx context.Context) (err error) {
	// 验证会员
	if in.MemberId == 0 {
		err = gerror.New("请选择会员")
		return
	}

	return
}

type SpaTechnicianBindModel struct{}

// SpaTechnicianUnbindInp 解绑渠道管理
type SpaTechnicianUnbindInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *SpaTechnicianUnbindInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaTechnicianUnbindModel struct{}
