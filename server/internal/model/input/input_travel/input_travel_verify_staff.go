package input_travel

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
)

// TravelVerifyStaffScopeItem 核销员权限范围项
// skuId=0 表示该产品下全部SKU
// skuId>0 表示精确到SKU
type TravelVerifyStaffScopeItem struct {
	ProductId uint64 `json:"productId" v:"required#请选择产品" dc:"产品ID"`
	SkuId     uint64 `json:"skuId" dc:"SKU ID（0=该产品全部SKU）"`
}

// TravelVerifyStaffUpdateFields 修改核销人员字段过滤
type TravelVerifyStaffUpdateFields struct {
	Name         string `json:"name" dc:"姓名"`
	Mobile       string `json:"mobile" dc:"电话"`
	Username     string `json:"username" dc:"登录账号"`
	PasswordHash string `json:"passwordHash" dc:"密码哈希"`
	Salt         string `json:"salt" dc:"密码盐"`
	Status       int    `json:"status" dc:"状态（1启用2禁用）"`
}

// TravelVerifyStaffInsertFields 新增核销人员字段过滤
type TravelVerifyStaffInsertFields struct {
	Name         string `json:"name" dc:"姓名"`
	Mobile       string `json:"mobile" dc:"电话"`
	Username     string `json:"username" dc:"登录账号"`
	PasswordHash string `json:"passwordHash" dc:"密码哈希"`
	Salt         string `json:"salt" dc:"密码盐"`
	Status       int    `json:"status" dc:"状态（1启用2禁用）"`
}

// TravelVerifyStaffListInp 获取核销人员列表
type TravelVerifyStaffListInp struct {
	input_form.PageReq
	Name   string `json:"name" dc:"姓名关键词"`
	Status int    `json:"status" dc:"状态（0全部1启用2禁用）"`
}

// TravelVerifyStaffListModel 核销人员列表模型
type TravelVerifyStaffListModel struct {
	Id        int64  `json:"id" dc:"核销人员ID"`
	Name      string `json:"name" dc:"姓名"`
	Mobile    string `json:"mobile" dc:"电话"`
	Username  string `json:"username" dc:"登录账号"`
	Status    int    `json:"status" dc:"状态（1启用2禁用）"`
	CreatedAt string `json:"createdAt" dc:"创建时间"`
}

// TravelVerifyStaffViewInp 获取核销人员详情
type TravelVerifyStaffViewInp struct {
	Id int64 `json:"id" v:"required#请选择要查看的核销人员" dc:"核销人员ID"`
}

// TravelVerifyStaffViewModel 核销人员详情模型
type TravelVerifyStaffViewModel struct {
	entity.TravelVerifyStaff
	IsAllScope bool                          `json:"isAllScope" dc:"是否全部可核销"`
	ScopeItems []*TravelVerifyStaffScopeItem `json:"scopeItems" dc:"授权范围列表"`
}

// TravelVerifyStaffEditInp 新增/编辑核销人员
type TravelVerifyStaffEditInp struct {
	Id           int64                         `json:"id" dc:"核销人员ID（为0时新增）"`
	Name         string                        `json:"name" v:"required#请输入姓名" dc:"姓名"`
	Mobile       string                        `json:"mobile" v:"required#请输入电话" dc:"电话"`
	Username     string                        `json:"username" v:"required#请输入登录账号" dc:"登录账号"`
	Password     string                        `json:"password" dc:"密码（编辑时不填则不修改）"`
	PasswordHash string                        `json:"passwordHash" dc:"密码哈希"`
	Salt         string                        `json:"salt" dc:"密码盐"`
	Status       int                           `json:"status" v:"in:1,2#状态值错误" dc:"状态（1启用2禁用）"`
	IsAllScope   bool                          `json:"isAllScope" dc:"是否全部可核销"`
	ScopeItems   []*TravelVerifyStaffScopeItem `json:"scopeItems" dc:"授权范围列表"`
}

// TravelVerifyStaffScopeSkuOptionModel 核销权限范围SKU选项
// 仅用于前端下拉展示
type TravelVerifyStaffScopeSkuOptionModel struct {
	Id        uint64 `json:"id" dc:"SKU ID"`
	ProductId uint64 `json:"productId" dc:"产品ID"`
	Name      string `json:"name" dc:"SKU名称"`
}

// TravelVerifyStaffScopeOptionModel 核销权限范围选项
// 仅用于前端下拉展示
type TravelVerifyStaffScopeOptionModel struct {
	Id      uint64                                  `json:"id" dc:"产品ID"`
	Title   string                                  `json:"title" dc:"产品名称"`
	SkuList []*TravelVerifyStaffScopeSkuOptionModel `json:"skuList" dc:"SKU选项"`
}

// TravelVerifyStaffDeleteInp 删除核销人员
type TravelVerifyStaffDeleteInp struct {
	Id []int64 `json:"id" v:"required#请选择要删除的核销人员" dc:"核销人员ID列表"`
}

// TravelVerifyStaffStatusInp 更新核销人员状态
type TravelVerifyStaffStatusInp struct {
	Id     int64 `json:"id" v:"required#请选择要操作的核销人员" dc:"核销人员ID"`
	Status int   `json:"status" v:"in:1,2#状态值错误" dc:"状态（1启用2禁用）"`
}
