package th

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_th"
	"github.com/gogf/gf/v2/frame/g"
)

type MchStoreListReq struct {
	g.Meta `path:"/thMchStore/list" method:"get" tags:"ADMIN_TH" summary:"获取商户门店列表"`
	input_th.ThMchStoreListInp
}

type MchStoreListRes struct {
	input_form.PageRes
	List []*input_th.ThMchStoreListModel `json:"list"   dc:"数据列表"`
}

type MchStoreViewReq struct {
	g.Meta `path:"/thMchStore/view" method:"get" tags:"ADMIN_TH" summary:"获取商户门店指定信息"`
	input_th.ThMchStoreViewInp
}

type MchStoreViewRes struct {
	*input_th.ThMchStoreViewModel
}

type MchStoreEditReq struct {
	g.Meta `path:"/thMchStore/edit" method:"post" tags:"ADMIN_TH" summary:"修改/新增商户门店"`
	input_th.ThMchStoreEditInp
}

type MchStoreEditRes struct{}

type MchStoreDeleteReq struct {
	g.Meta `path:"/thMchStore/delete" method:"post" tags:"ADMIN_TH" summary:"删除商户门店"`
	input_th.ThMchStoreDeleteInp
}

type MchStoreDeleteRes struct{}

type MchStoreStatusReq struct {
	g.Meta `path:"/thMchStore/status" method:"post" tags:"ADMIN_TH" summary:"更新商户门店状态"`
	input_th.ThMchStoreStatusInp
}

type MchStoreStatusRes struct{}

type MchStoreSwitchReq struct {
	g.Meta `path:"/thMchStore/switch" method:"post" tags:"ADMIN_TH" summary:"商户门店_更新状态"`
	input_th.ThMchStoreSwitchInp
}

type MchStoreSwitchRes struct{}

type MchStoreAllListReq struct {
	g.Meta `path:"/thMchStore/all" method:"get" tags:"ADMIN_TH" summary:"获取指定下商户门店列表"`
	input_th.ThMchStoreAllInp
}

type MchStoreAllListRes struct {
	List []*input_th.ThMchStoreAllModel `json:"list"   dc:"数据列表"`
}
