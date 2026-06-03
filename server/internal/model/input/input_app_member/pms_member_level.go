// Package sysin

package input_app_member

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsMemberLevelUpdateFields 修改会员等级字段过滤
type PmsMemberLevelUpdateFields struct {
	LevelName      string  `json:"levelName" dc:"会员等级名称"`
	Exp            int     `json:"exp"       dc:"会员经验"`
	Desc           string  `json:"desc"            dc:"等级说明"`
	WordColor      string  `json:"wordColor"       dc:"等级字体颜色"`
	LevelBadge     string  `json:"levelBadge"      dc:"等级徽章（单图）"`
	LevelCard      string  `json:"levelCard"       dc:"等级卡片（单图）"`
	LevelBigPic    string  `json:"levelBigPic"  dc:"等级大图（单图）"`
	HotelGetRate   float64 `json:"hotelGetRate" dc:"酒店场景获取积分倍率"`
	FoodGetRate    float64 `json:"foodGetRate"  dc:"餐饮场景获取积分倍率"`
	SpaGetRate     float64 `json:"spaGetRate"  dc:"按摩场景获取积分倍率"`
	CarGetRate     float64 `json:"carGetRate"  dc:"接送机场景获取积分倍率"`
	CabinetGetRate float64 `json:"cabinetGetRate"  dc:"储物柜场景获取积分倍率"`
}

// PmsMemberLevelInsertFields 新增会员等级字段过滤
type PmsMemberLevelInsertFields struct {
	LevelName      string  `json:"levelName" dc:"会员等级名称"`
	Exp            int     `json:"exp"       dc:"会员经验"`
	Desc           string  `json:"desc"            dc:"等级说明"`
	WordColor      string  `json:"wordColor"       dc:"等级字体颜色"`
	LevelBadge     string  `json:"levelBadge"      dc:"等级徽章（单图）"`
	LevelCard      string  `json:"levelCard"       dc:"等级卡片（单图）"`
	LevelBigPic    string  `json:"levelBigPic"  dc:"等级大图（单图）"`
	HotelGetRate   float64 `json:"hotelGetRate" dc:"酒店场景获取积分倍率"`
	FoodGetRate    float64 `json:"foodGetRate"  dc:"餐饮场景获取积分倍率"`
	SpaGetRate     float64 `json:"spaGetRate"  dc:"按摩场景获取积分倍率"`
	CarGetRate     float64 `json:"carGetRate"  dc:"接送机场景获取积分倍率"`
	CabinetGetRate float64 `json:"cabinetGetRate"  dc:"储物柜场景获取积分倍率"`
}

// PmsMemberLevelEditInp 修改/新增会员等级
type PmsMemberLevelEditInp struct {
	entity.PmsMemberLevel
}

func (in *PmsMemberLevelEditInp) Filter(ctx context.Context) (err error) {

	// 验证会员等级名称
	if err := g.Validator().Rules("required").Data(in.LevelName).Messages("会员等级名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证会员经验
	if err := g.Validator().Rules("required").Data(in.Exp).Messages("会员经验不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type PmsMemberLevelEditModel struct{}

// PmsMemberLevelDeleteInp 删除会员等级
type PmsMemberLevelDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsMemberLevelDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsMemberLevelDeleteModel struct{}

// PmsMemberLevelViewInp 获取指定会员等级信息
type PmsMemberLevelViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsMemberLevelViewInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsMemberLevelViewModel struct {
	entity.PmsMemberLevel
}

// PmsMemberLevelListInp 获取会员等级列表
type PmsMemberLevelListInp struct {
	input_form.PageReq
	LevelName string `json:"levelName" dc:"会员等级名称"`
}

func (in *PmsMemberLevelListInp) Filter(ctx context.Context) (err error) {
	return
}

// PmsMemberLevelAllInp 获取会员等级列表
type PmsMemberLevelAllInp struct {
	LevelName string `json:"levelName" dc:"会员等级名称"`
}

func (in *PmsMemberLevelAllInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsMemberLevelListModel struct {
	Id           int         `json:"id"        dc:"id"`
	LevelName    string      `json:"levelName" dc:"会员等级名称"`
	Exp          int         `json:"exp"       dc:"会员经验"`
	Desc         string      `json:"desc"            dc:"等级说明"`
	WordColor    string      `json:"wordColor"       dc:"等级字体颜色"`
	LevelBadge   string      `json:"levelBadge"      dc:"等级徽章（单图）"`
	LevelCard    string      `json:"levelCard"       dc:"等级卡片（单图）"`
	LevelBigPic  string      `json:"levelBigPic"  dc:"等级大图（单图）"`
	HotelGetRate float64     `json:"hotelGetRate" dc:"酒店场景获取积分倍率"`
	FoodGetRate  float64     `json:"foodGetRate"  dc:"餐饮场景获取积分倍率"`
	SpaGetRate   float64     `json:"spaGetRate"  dc:"按摩场景获取积分倍率"`
	CarGetRate   float64     `json:"carGetRate"  dc:"接送机/包车场景获取积分倍率"`
	CreatedAt    *gtime.Time `json:"createdAt" dc:"created_at"`
	MemberCount  int         `json:"memberCount" dc:"持有人数"`
}

type PmsMemberLevelAllModel struct {
	Id           int         `json:"id"        dc:"id"`
	LevelName    string      `json:"levelName" dc:"会员等级名称"`
	Exp          int         `json:"exp"       dc:"会员经验"`
	Desc         string      `json:"desc"            dc:"等级说明"`
	WordColor    string      `json:"wordColor"       dc:"等级字体颜色"`
	LevelBadge   string      `json:"levelBadge"      dc:"等级徽章（单图）"`
	LevelCard    string      `json:"levelCard"       dc:"等级卡片（单图）"`
	HotelGetRate float64     `json:"hotelGetRate" dc:"酒店场景获取积分倍率"`
	FoodGetRate  float64     `json:"foodGetRate"  dc:"餐饮场景获取积分倍率"`
	SpaGetRate   float64     `json:"spaGetRate"  dc:"按摩场景获取积分倍率"`
	CarGetRate   float64     `json:"carGetRate"  dc:"接送机/包车场景获取积分倍率"`
	CreatedAt    *gtime.Time `json:"createdAt" dc:"created_at"`
	MemberDetail []*struct {
		g.Meta `orm:"table:hg_pms_member"`
		Id     int `json:"id"               orm:"id"                 description:"主键"`
		Level  int `json:"level"        dc:"等级"`
	} `json:"memberDetail" orm:"with:level=id" dc:"会员信息"`
}
