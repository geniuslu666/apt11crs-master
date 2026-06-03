// Package sysin

package input_app_member

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/utility/validate"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gmeta"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsMemberUpdateFields 修改会员信息字段过滤
type PmsMemberUpdateFields struct {
	MemberNo     string      `json:"memberNo"     dc:"会员号"`
	FirstName    string      `json:"firstName"    dc:"名"`
	LastName     string      `json:"lastName"     dc:"姓"`
	FullName     string      `json:"fullName"     dc:"全名"`
	Phone        string      `json:"phone"        dc:"手机号"`
	Mail         string      `json:"mail"         dc:"邮箱"`
	Password     string      `json:"password"     dc:"密码"`
	Birthday     *gtime.Time `json:"birthday"     dc:"生日"`
	Balance      float64     `json:"balance"      dc:"积分"`
	Referrer     int         `json:"referrer"     dc:"推荐人"`
	Source       string      `json:"source"       dc:"注册来源  IOS,Andriod,h5"`
	LastLoginIp  string      `json:"lastLoginIp"  dc:"上次登录IP"`
	LastLogin    *gtime.Time `json:"lastLogin"    dc:"上次登录时间"`
	RegisterIp   string      `json:"registerIp"   dc:"注册IP"`
	RegisterTime *gtime.Time `json:"registerTime" dc:"注册时间"`
	Address      string      `json:"address"      dc:"地址"`
	Nationality  string      `json:"nationality"  dc:"国籍"`
	CreateAt     string      `json:"createAt"     dc:"创建时间"`
	UpdateAt     string      `json:"updateAt"     dc:"更新时间"`
}

// PmsMemberInsertFields 新增会员信息字段过滤
type PmsMemberInsertFields struct {
	MemberNo     string      `json:"memberNo"     dc:"会员号"`
	FirstName    string      `json:"firstName"    dc:"名"`
	LastName     string      `json:"lastName"     dc:"姓"`
	FullName     string      `json:"fullName"     dc:"全名"`
	Phone        string      `json:"phone"        dc:"手机号"`
	Mail         string      `json:"mail"         dc:"邮箱"`
	Password     string      `json:"password"     dc:"密码"`
	Birthday     *gtime.Time `json:"birthday"     dc:"生日"`
	Balance      float64     `json:"balance"      dc:"积分"`
	Referrer     int         `json:"referrer"     dc:"推荐人"`
	Source       string      `json:"source"       dc:"注册来源  IOS,Andriod,h5"`
	LastLoginIp  string      `json:"lastLoginIp"  dc:"上次登录IP"`
	LastLogin    *gtime.Time `json:"lastLogin"    dc:"上次登录时间"`
	RegisterIp   string      `json:"registerIp"   dc:"注册IP"`
	RegisterTime *gtime.Time `json:"registerTime" dc:"注册时间"`
	Address      string      `json:"address"      dc:"地址"`
	Nationality  string      `json:"nationality"  dc:"国籍"`
	CreateAt     string      `json:"createAt"     dc:"创建时间"`
	UpdateAt     string      `json:"updateAt"     dc:"更新时间"`
}

// PmsMemberEditInp 修改/新增会员信息
type PmsMemberEditInp struct {
	entity.PmsMember
}

func (in *PmsMemberEditInp) Filter(ctx context.Context) (err error) {
	// 验证名
	if err := g.Validator().Rules("required").Data(in.FirstName).Messages("名不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证姓
	if err := g.Validator().Rules("required").Data(in.LastName).Messages("姓不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if !g.IsEmpty(in.Password) {
		// 验证密码
		if err := g.Validator().Rules("regex:^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,18}$").Data(in.Password).Messages("密码必须包含6-18为字母和数字").Run(ctx); err != nil {
			return err.Current()
		}
	}

	return
}

type PmsMemberEditModel struct{}

// PmsMemberBaseEditInp 修改/新增会员信息
type PmsMemberBaseEditInp struct {
	Type string `json:"type"     dc:"修改类型"`
	entity.PmsMember
}

func (in *PmsMemberBaseEditInp) Filter(ctx context.Context) (err error) {
	if in.Type == "firstName" {
		// 验证名
		if err := g.Validator().Rules("required").Data(in.FirstName).Messages("名不能为空").Run(ctx); err != nil {
			return err.Current()
		}
	}

	if in.Type == "lastName" {
		// 验证姓
		if err := g.Validator().Rules("required").Data(in.LastName).Messages("姓不能为空").Run(ctx); err != nil {
			return err.Current()
		}
	}
	if !g.IsEmpty(in.Password) {
		// 验证密码
		if err := g.Validator().Rules("regex:^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,18}$").Data(in.Password).Messages("密码必须包含6-18为字母和数字").Run(ctx); err != nil {
			return err.Current()
		}
	}

	return
}

type PmsMemberBaseEditModel struct{}

// PmsMemberBalanceEditInp PmsMemberBaseEditInp 修改/新增会员信息
type PmsMemberBalanceEditInp struct {
	Id      int     `json:"id"                   description:"主键"`
	Balance float64 `json:"balance"            description:"积分"`
	Value   float64 `json:"value"                   description:"调整值"`
	Des     string  `json:"des"                   description:"备注"`
	Reason  string  `json:"reason"                   description:"原因"`
}

func (in *PmsMemberBalanceEditInp) Filter(ctx context.Context) (err error) {
	if (in.Balance + in.Value) < 0 {
		err = gerror.New("调整值与当前积分数相加不能小于0")
		return
	}

	return
}

type PmsMemberBalanceEditModel struct{}

// PmsMemberExpEditInp 修改/新增会员信息
type PmsMemberExpEditInp struct {
	Id    int     `json:"id"                   description:"主键"`
	Exp   int     `json:"exp"        description:"成长值"`
	Value float64 `json:"value"                   description:"调整值"`
	Des   string  `json:"des"                   description:"备注"`
}

func (in *PmsMemberExpEditInp) Filter(ctx context.Context) (err error) {
	if (float64(in.Exp) + in.Value) < 0 {
		err = gerror.New("调整值与当前成长值数相加不能小于0")
		return
	}

	return
}

type PmsMemberExpEditModel struct{}

// PmsMemberDeleteInp 删除会员信息
type PmsMemberDeleteInp struct {
	Id interface{} `json:"id" v:"required#主键不能为空" dc:"主键"`
}

func (in *PmsMemberDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsMemberDeleteModel struct{}

// PmsMemberViewInp 获取指定会员信息信息
type PmsMemberViewInp struct {
	Id int `json:"id" v:"required#主键不能为空" dc:"主键"`
}

func (in *PmsMemberViewInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsMemberViewModel struct {
	entity.PmsMember
	MemberLevel *struct {
		gmeta.Meta `orm:"table:hg_pms_member_level"`
		Id         int    `json:"id"    orm:"id"      dc:"id"`
		LevelName  string `json:"levelName"                 orm:"levelName"                    description:"会员等级名称"`
	} `json:"memberLevel" orm:"with:id=level" dc:"会员等级信息"`
	ReferrerDetail *struct {
		gmeta.Meta `orm:"table:hg_pms_member"`
		Id         int    `json:"id"    orm:"id"      dc:"id"`
		MemberNo   string `json:"memberNo"                 orm:"memberNo"                    description:"会员号"`
		RebateMode string `json:"rebateMode"   orm:"rebateMode"   description:"MEMBER 会员     STAFF    员工 CHANNEL  渠道"`
	} `json:"referrerDetail" orm:"with:id=referrer" dc:"会员推荐人信息"`
	LastReferrerDetail *struct {
		gmeta.Meta `orm:"table:hg_pms_member"`
		Id         int    `json:"id"    orm:"id"      dc:"id"`
		MemberNo   string `json:"memberNo"                 orm:"memberNo"                    description:"会员号"`
		RebateMode string `json:"rebateMode"   orm:"rebateMode"   description:"MEMBER 会员     STAFF    员工 CHANNEL  渠道"`
	} `json:"LastReferrerDetail" orm:"with:id=last_referrer" dc:"会员最后推荐人信息"`
	MemberGroup *struct {
		gmeta.Meta  `orm:"table:hg_pms_member_group"`
		Id          int    `json:"id"    orm:"id"      dc:"分组id"`
		MemberGroup string `json:"memberGroup"         orm:"memberGroup"     dc:"会员分组名称"`
	} `json:"memberGroup" orm:"with:id=group_id" dc:"会员分组信息"`
}

// PmsMemberListInp 获取会员信息列表
type PmsMemberListInp struct {
	input_form.PageReq
	Id         int           `json:"id"          dc:"主键"`
	MemberNo   string        `json:"memberNo"    dc:"会员号"`
	Phone      string        `json:"phone"       dc:"手机号"`
	Mail       string        `json:"mail"        dc:"邮箱"`
	Level      int           `json:"level"        dc:"等级"`
	FullName   string        `json:"fullName"    dc:"全称"`
	CreatedAt  []*gtime.Time `json:"createdAt" dc:"created_at"`
	LastLogin  []*gtime.Time `json:"lastLogin" dc:"last_login"`
	MemberIds  string        `json:"memberIds"   dc:"会员Ids"`
	Sort       string        `json:"sort"      dc:"排序"`
	IsJunior   int           `json:"isJunior"    dc:"是否下级 1-查询下级 0-查询全部"`
	RebateMode string        `json:"rebateMode"     dc:"MEMBER 会员     STAFF    员工 CHANNEL  渠道"`
	Status     int           `json:"status"          dc:"状态1、启用 2、禁用"`
}

func (in *PmsMemberListInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsMemberListModel struct {
	entity.PmsMember
	MemberLevel *struct {
		gmeta.Meta `orm:"table:hg_pms_member_level"`
		Id         int    `json:"id"    orm:"id"      dc:"id"`
		LevelName  string `json:"levelName"                 orm:"levelName"                    description:"会员等级名称"`
	} `json:"memberLevel" orm:"with:id=level" dc:"会员等级信息"`
	ReferrerDetail *struct {
		gmeta.Meta `orm:"table:hg_pms_member"`
		Id         int    `json:"id"    orm:"id"      dc:"id"`
		MemberNo   string `json:"memberNo"                 orm:"memberNo"                    description:"会员号"`
		RebateMode string `json:"rebateMode"   orm:"rebateMode"   description:"MEMBER 会员     STAFF    员工 CHANNEL  渠道"`
	} `json:"referrerDetail" orm:"with:id=referrer" dc:"会员推荐人信息"`
	LastReferrerDetail *struct {
		gmeta.Meta `orm:"table:hg_pms_member"`
		Id         int    `json:"id"    orm:"id"      dc:"id"`
		MemberNo   string `json:"memberNo"                 orm:"memberNo"                    description:"会员号"`
		RebateMode string `json:"rebateMode"   orm:"rebateMode"   description:"MEMBER 会员     STAFF    员工 CHANNEL  渠道"`
	} `json:"LastReferrerDetail" orm:"with:id=last_referrer" dc:"会员最后推荐人信息"`
	MemberGroup *struct {
		gmeta.Meta  `orm:"table:hg_pms_member_group"`
		Id          int    `json:"id"    orm:"id"      dc:"分组id"`
		MemberGroup string `json:"memberGroup"         orm:"memberGroup"     dc:"会员分组名称"`
	} `json:"memberGroup" orm:"with:id=group_id" dc:"会员分组信息"`
	MemberCancelArr []*struct {
		gmeta.Meta  `orm:"table:hg_pms_member_cancel"`
		Id          int    `json:"id"    orm:"id"      dc:"id"`
		MemberId    int    `json:"memberId"    orm:"member_id"      dc:"用户ID"`
		AuditStatus string `json:"auditStatus"         orm:"audit_status"     dc:"审核状态"`
	} `json:"memberCancelArr" orm:"with:member_id=id, where:audit_status!='3'" dc:"会员注销信息"`
}

// PmsMemberSelectInp 获取会员信息列表
type PmsMemberSelectInp struct {
	Keyword string `json:"keyword"    dc:"关键词"`
}

// PmsMemberExportModel 导出会员信息
type PmsMemberExportModel struct {
	Id           int         `json:"id"           dc:"主键"`
	MemberNo     string      `json:"memberNo"     dc:"会员号"`
	FullName     string      `json:"fullName"     dc:"姓名"`
	ReferInfo    string      `json:"level"        dc:"推荐信息"`
	Phone        string      `json:"phone"        dc:"手机号"`
	Mail         string      `json:"mail"         dc:"邮箱"`
	RegisterTime *gtime.Time `json:"registerTime"     dc:"注册时间"`
}

// PmsMemberStatInp 获取会员概况
type PmsMemberStatInp struct {
	Type string `json:"type"    dc:"类型"`
}

type PmsMemberStatModel struct {
	MemberBaseStat struct {
		MemberTotal            int     `json:"memberTotal"           dc:"会员总数"`
		TodayMemberAdd         int     `json:"todayMemberAdd"           dc:"今日增加会员"`
		YesterdayMemberAdd     int     `json:"yesterdayMemberAdd"           dc:"昨日增加会员"`
		MemberBalanceTotal     float64 `json:"memberBalanceTotal"      dc:"会员积分池"`
		TodayConsumedBalance   float64 `json:"todayConsumedBalance"      dc:"今日消耗积分"`
		TodayMemberOrderCount  int     `json:"todayMemberOrderCount"      dc:"今日下单会员"`
		TodayHotelOrderNum     int     `json:"todayHotelOrderNum"         dc:"今日民宿下单数"`
		YesterdayHotelOrderNum int     `json:"yesterdayHotelOrderNum"         dc:"昨日民宿下单数"`
		MemberOrder            int     `json:"memberOrder"           dc:"下单会员数"`
		MemberOrderPercent     float64 `json:"memberOrderPercent"           dc:"下单会员数比例"`
		MemberNotOrder         int     `json:"memberNotOrder"           dc:"未下单会员数"`
		MemberNotOrderPercent  float64 `json:"memberNotOrderPercent"           dc:"未下单会员数比例"`
	} `json:"memberBaseStat" dc:"会员基础统计数据"`
	MemberOrderPie struct {
		MemberOrder           int     `json:"memberOrder"           dc:"下单会员数"`
		MemberOrderPercent    float64 `json:"memberOrderPercent"           dc:"下单会员数比例"`
		MemberNotOrder        int     `json:"memberNotOrder"           dc:"未下单会员数"`
		MemberNotOrderPercent float64 `json:"memberNotOrderPercent"           dc:"未下单会员数比例"`
	} `json:"memberOrderPie" dc:"会员下单饼状图"`
	MemberCreateTrend []*struct {
		Num  int    `json:"num"           description:"数量"`
		Date string `json:"date"    description:"日期"`
	} `json:"memberCreateTrend" dc:"新增会员趋势"`
	MemberOrderTrend []*struct {
		Money float64 `json:"money"      dc:"金额"`
		Date  string  `json:"date"    description:"日期"`
	} `json:"memberOrderTrend" dc:"会员消费趋势"`
	MemberSource []*struct {
		Source  string  `json:"source"       dc:"注册来源  IOS,Andriod,h5"`
		Num     int     `json:"num"           description:"数量"`
		Percent float64 `json:"percent"      dc:"百分比"`
	} `json:"memberSource" dc:"会员来源"`
	MemberLevel []*struct {
		LevelName string  `json:"levelName"       dc:"等级名称"`
		Num       int     `json:"num"           description:"数量"`
		Percent   float64 `json:"percent"      dc:"百分比"`
	} `json:"memberLevel" dc:"会员等级比率"`
}

// PmsMemberStatusInp 更新会员状态
type PmsMemberStatusInp struct {
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
	Status int `json:"status" dc:"状态"`
}

func (in *PmsMemberStatusInp) Filter(ctx context.Context) (err error) {
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

type PmsMemberStatusModel struct{}

// PmsMemberCancelInp 会员注销
type PmsMemberCancelInp struct {
	Id           int    `json:"id" v:"required#id不能为空" dc:"id"`
	CancelReason string `json:"cancelReason" dc:"注销原因"`
}

type PmsMemberCancelModel struct{}
