package input_basics

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// TerminalUpdateFields 修改终端字段过滤
type TerminalUpdateFields struct {
	TerminalName string `json:"printerName"  dc:"终端名称"`
	TerminalType string `json:"terminalType" dc:"终端类型"`
	BrandModel   string `json:"brandModel"   dc:"品牌型号"`
	Sn           string `json:"sn"           dc:"终端编号"`
}

// TerminalInsertFields 新增终端字段过滤
type TerminalInsertFields struct {
	TerminalName string `json:"printerName"  dc:"终端名称"`
	TerminalType string `json:"terminalType" dc:"终端类型"`
	BrandModel   string `json:"brandModel"   dc:"品牌型号"`
	Sn           string `json:"sn"           dc:"终端编号"`
}

// TerminalEditInp 修改/新增终端
type TerminalEditInp struct {
	entity.SysTerminal
}

func (in *TerminalEditInp) Filter(ctx context.Context) (err error) {
	if in.TerminalName == "" {
		err = gerror.New("请输入终端名称")
		return
	}

	if in.Sn == "" {
		err = gerror.New("请输入终端号")
		return
	}
	return
}

type TerminalEditModel struct{}

// TerminalDeleteInp 删除终端
type TerminalDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *TerminalDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type TerminalDeleteModel struct{}

// TerminalViewInp 获取指定终端信息
type TerminalViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *TerminalViewInp) Filter(ctx context.Context) (err error) {
	return
}

type TerminalViewModel struct {
	*entity.SysTerminal
	BrandInfo *struct {
		gmeta.Meta   `orm:"table:hg_sys_terminal_model"`
		Id           int    `json:"id"           orm:"id"            description:""`
		BrandModel   string `json:"brandModel"   orm:"brand_model"   description:"品牌型号"`
		ClientId     string `json:"clientId"     orm:"client_id"     description:"开发者ID"`
		ClientSecret string `json:"clientSecret" orm:"client_secret" description:"开发者秘钥"`
	} `json:"brandInfo" orm:"with:brand_model=brand_model" dc:"品牌型号"`
	StoreInfo *struct {
		gmeta.Meta `orm:"table:hg_th_mch_store"`
		MchId      int    `json:"mchId"       orm:"mch_id"       description:"商户ID"`
		Id         int    `json:"id"          orm:"id"           description:""`
		StoreName  string `json:"storeName"   orm:"store_name"   description:"门店名称（多语言）"`
	} `json:"storeInfo" orm:"with:id=store_id" dc:"绑定门店信息"`
}

// TerminalListInp 获取终端列表
type TerminalListInp struct {
	input_form.PageReq
	TerminalIds                string        `json:"terminalIds"    dc:"终端IDs"`
	TerminalName               string        `json:"terminalName"    dc:"终端名称"`
	Sn                         string        `json:"sn"             dc:"终端号"`
	StoreId                    int           `json:"storeId"             dc:"绑定门店ID"`
	TerminalStoreId            int           `json:"terminalStoreId"             dc:"当前门店ID"`
	RestaurantId               int           `json:"restaurantId"             dc:"绑定餐厅ID"`
	TerminalRestaurantId       int           `json:"terminalRestaurantId"             dc:"当前餐厅ID"`
	CreatedAt                  []*gtime.Time `json:"createdAt"      dc:"created_at"`
	TerminalType               string        `json:"terminalType" dc:"终端类型"`
	FromStore                  int           `json:"fromStore" dc:"是否来自门店编辑 1-是 0-否"`
	NeedPrintTimesRestaurantId int           `json:"needPrintTimesRestaurantId" dc:"需要返回打印联数的餐厅ID，返回对应关联表的打印联数"`
	NeedPrintTimesStoreId      int           `json:"needPrintTimesStoreId" dc:"需要返回打印联数的门店ID, 返回对应关联表的打印联数"`
}

func (in *TerminalListInp) Filter(ctx context.Context) (err error) {
	return
}

type TerminalListModel struct {
	Id           int         `json:"id"            dc:"id"`
	TerminalName string      `json:"terminalName"  dc:"终端名称"`
	TerminalType string      `json:"terminalType"  dc:"终端类型"`
	BrandModel   string      `json:"brandModel"    dc:"品牌型号"`
	Sn           string      `json:"sn"            dc:"终端号"`
	OnlineStatus int         `json:"onlineStatus"  dc:"在线状态"`
	StoreId      int         `json:"storeId"  dc:"绑定门店"`
	RestaurantId int         `json:"restaurantId"  dc:"绑定餐厅"`
	PrintTimes   uint        `json:"printTimes"   dc:"打印联数(次数)"`
	CreateAt     *gtime.Time `json:"createAt"      dc:"创建时间"`
	UpdateAt     *gtime.Time `json:"updateAt"      dc:"更新时间"`
	StoreInfo    *struct {
		gmeta.Meta `orm:"table:hg_th_mch_store"`
		Id         int    `json:"id"          orm:"id"           description:""`
		StoreName  string `json:"storeName"   orm:"store_name"   description:"门店名称（多语言）"`
	} `json:"storeInfo" orm:"with:id=store_id" dc:"绑定门店信息"`
	RestaurantInfo *struct {
		gmeta.Meta `orm:"table:hg_food_restaurant"`
		Id         int    `json:"id"          orm:"id"           description:""`
		Name       string `json:"name"   orm:"name"   description:"餐厅名称"`
	} `json:"restaurantInfo" orm:"with:id=restaurant_id" dc:"绑定餐厅信息"`
}

// BrandListInp 获取品牌型号列表
type BrandListInp struct {
	input_form.PageReq
	BrandModel string        `json:"brandModel"    dc:"品牌型号"`
	CreatedAt  []*gtime.Time `json:"createdAt"      dc:"created_at"`
}

func (in *BrandListInp) Filter(ctx context.Context) (err error) {
	return
}

type BrandListModel struct {
	Id           int         `json:"id"            dc:"id"`
	BrandModel   string      `json:"brandModel"    dc:"品牌型号"`
	ClientId     string      `json:"clientId"      dc:"开发者ID"`
	ClientSecret string      `json:"clientSecret"  dc:"开发者秘钥"`
	CreateAt     *gtime.Time `json:"createAt"      dc:"创建时间"`
	UpdateAt     *gtime.Time `json:"updateAt"      dc:"更新时间"`
}

// BrandViewInp 获取指定终端信息
type BrandViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *BrandViewInp) Filter(ctx context.Context) (err error) {
	return
}

type BrandViewModel struct {
	entity.SysTerminalModel
}

// BrandUpdateFields 修改品牌型号字段过滤
type BrandUpdateFields struct {
	BrandModel   string `json:"brandModel"   dc:"品牌型号"`
	ClientId     string `json:"clientId"     dc:"开发者ID"`
	ClientSecret string `json:"clientSecret" dc:"开发者秘钥"`
}

// BrandEditInp 修改/新增品牌型号
type BrandEditInp struct {
	entity.SysTerminalModel
}

func (in *BrandEditInp) Filter(ctx context.Context) (err error) {

	if in.ClientId == "" {
		err = gerror.New("请输入开发者ID")
		return
	}
	if in.ClientSecret == "" {
		err = gerror.New("请输入开发者秘钥")
		return
	}

	return
}

type BrandEditModel struct{}

// PrinterTestInp 测试打印
type PrinterTestInp struct {
	TerminalId int `json:"terminalId"  dc:"终端ID"`
}

// PrinterInp 打印
type PrinterInp struct {
	Sn           string `json:"sn"            dc:"终端编号"`
	PrintContent string `json:"printContent"  dc:"打印内容"`
	PrintTimes   int    `json:"printTimes"  dc:"打印联数"`
}
