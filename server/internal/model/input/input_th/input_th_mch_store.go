// Package sysin

package input_th

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_language"
	"APT/utility/validate"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gmeta"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// ThMchStoreUpdateFields 修改商户门店字段过滤
type ThMchStoreUpdateFields struct {
	MchId         uint64 `json:"mchId"         dc:"商户ID"`
	StoreName     string `json:"storeName"     dc:"门店名称（多语）"`
	Images        string `json:"images"        dc:"图集"`
	PhoneArea     string `json:"phoneArea"     dc:"区号"`
	Phone         string `json:"phone"         dc:"电话"`
	DetailAddress string `json:"detailAddress"                 dc:"详细地址"`
	GgLat         string `json:"ggLat"                   dc:"谷歌纬度"`
	GgLng         string `json:"ggLng"                   dc:"谷歌经度"`
	TerminalType  string `json:"terminalType"  dc:"终端类型"`
	Account       string `json:"account"       dc:"账号"`
	Password      string `json:"password"     dc:"密码"`
	PasswordHash  string `json:"passwordHash"            dc:"密码"`
	Salt          string `json:"salt"                    dc:"密码盐"`
}

// ThMchStoreInsertFields 新增商户门店字段过滤
type ThMchStoreInsertFields struct {
	MchId         uint64 `json:"mchId"         dc:"商户ID"`
	StoreName     string `json:"storeName"     dc:"门店名称（多语）"`
	Images        string `json:"images"        dc:"图集"`
	PhoneArea     string `json:"phoneArea"     dc:"区号"`
	Phone         string `json:"phone"         dc:"电话"`
	DetailAddress string `json:"detailAddress"                 dc:"详细地址"`
	GgLat         string `json:"ggLat"                   dc:"谷歌纬度"`
	GgLng         string `json:"ggLng"                   dc:"谷歌经度"`
	TerminalType  string `json:"terminalType"  dc:"终端类型"`
	Account       string `json:"account"       dc:"账号"`
	Password      string `json:"password"     dc:"密码"`
	PasswordHash  string `json:"passwordHash"            dc:"密码"`
	Salt          string `json:"salt"                    dc:"密码盐"`
}

// ThMchStoreEditInp 修改/新增商户门店
type ThMchStoreEditInp struct {
	entity.ThMchStore
	Password     string `json:"password"     dc:"密码"`
	TerminalList []*struct {
		Id         int `json:"id"     dc:"终端ID"`
		PrintTimes int `json:"printTimes"     dc:"打印联数"`
	} `json:"terminalList"          dc:"终端列表"`
	NameLanguage input_language.LanguageModel `json:"nameLanguage"          dc:"多语言门店名称"`
}

func (in *ThMchStoreEditInp) Filter(ctx context.Context) (err error) {

	return
}

type ThMchStoreEditModel struct{}

// ThMchStoreDeleteInp 删除商户门店
type ThMchStoreDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *ThMchStoreDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type ThMchStoreDeleteModel struct{}

// ThMchStoreViewInp 获取指定商户门店信息
type ThMchStoreViewInp struct {
	Id         int  `json:"id" v:"required#id不能为空" dc:"id"`
	IsLanguage bool `json:"isLanguage" dc:"是否获取多语言数据"`
}

func (in *ThMchStoreViewInp) Filter(ctx context.Context) (err error) {
	return
}

type ThMchStoreViewModel struct {
	entity.ThMchStore
	NameLanguage []*input_hotel.LanguageType `json:"nameLanguage"         dc:"门店名称"   orm:"with:uuid=store_name"`
	ThMchName    string                      `json:"thMchName" dc:"商户名称"`
	TerminalList []*struct {
		gmeta.Meta `orm:"table:hg_th_store_terminal"`
		*entity.ThStoreTerminal
		TerminalInfo *struct {
			gmeta.Meta   `orm:"table:hg_sys_terminal"`
			Id           int    `json:"id"           dc:""`
			TerminalName string `json:"terminalName" dc:"终端名称"`
			TerminalType string `json:"terminalType" dc:"终端类型"`
			BrandModel   string `json:"brandModel"   dc:"品牌型号"`
			Sn           string `json:"sn"           dc:"终端编号"`
			PrintTimes   uint   `json:"printTimes" dc:"打印联数(次数)"`
		} `json:"terminalInfo" orm:"with:id=terminal_id" dc:"终端信息"`
	} `json:"terminalList" orm:"with:store_id=id" dc:"终端列表"`
}

// ThMchStoreListInp 获取商户门店列表
type ThMchStoreListInp struct {
	input_form.PageReq
	MchId     uint64 `json:"mchId"         dc:"商户ID"`
	StoreName string `json:"storeName"     dc:"门店名称（多语）"`
}

func (in *ThMchStoreListInp) Filter(ctx context.Context) (err error) {
	return
}

type ThMchStoreListModel struct {
	Id            int         `json:"id"       dc:"id"`
	ThMchName     string      `json:"thMchName" dc:"商户名称"`
	StoreName     string      `json:"storeName"     dc:"门店名称（多语）"`
	PhoneArea     string      `json:"phoneArea"     dc:"区号"`
	Phone         string      `json:"phone"         dc:"电话"`
	DetailAddress string      `json:"detailAddress"                 dc:"详细地址"`
	Status        int         `json:"status"   dc:"状态1、启用 2、禁用"`
	VerifyNum     int         `json:"verifyNum"     dc:"核销数量"`
	CreatedAt     *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// ThMchStoreStatusInp 更新商户门店状态
type ThMchStoreStatusInp struct {
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
	Status int `json:"status" dc:"状态"`
}

func (in *ThMchStoreStatusInp) Filter(ctx context.Context) (err error) {
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

type ThMchStoreStatusModel struct{}

// ThMchStoreSwitchInp 更新状态
type ThMchStoreSwitchInp struct {
	Id     int  `json:"id" v:"required#id不能为空" dc:"id"`
	Status uint `json:"status"    dc:"1、启用 2、禁用"`
}

func (in *ThMchStoreSwitchInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("id不能为空")
		return
	}

	if g.IsEmpty(in.Status) {
		err = gerror.New("状态不能为空")
		return
	}

	return
}

type ThMchStoreSwitchModel struct{}

// ThMchStoreAllInp 礼品券关联门店
type ThMchStoreAllInp struct {
	CouponId int `json:"couponId" v:"required#礼品券ID不能为空" dc:"礼品券ID"`
}

type ThMchStoreAllModel struct {
	Id        int    `json:"id"       dc:"id"`
	StoreName string `json:"storeName"     dc:"门店名称（多语）"`
	Status    int    `json:"status"   dc:"状态1、启用 2、禁用"`
}
