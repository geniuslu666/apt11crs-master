package input_food

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// FoodOrderLogListInp 获取餐厅套餐预订单列表
type FoodOrderLogListInp struct {
	input_form.PageReq
	OrderSn   string        `json:"orderSn" dc:"订单编号"`
	ActionWay string        `json:"actionWay" dc:"操作名"`
	CreatedAt []*gtime.Time `json:"createdAt"    dc:"创建时间"`
}

func (in *FoodOrderLogListInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodOrderLogListModel struct {
	entity.FoodOrderLog
	AdminMemberUsername string `json:"adminMemberUsername"      dc:"操作员"`
	OperateName         string `json:"operateName"      dc:"操作人姓名"`
	OperatePhone        string `json:"operatePhone"      dc:"操作人电话"`
	OrderDetail         *struct {
		gmeta.Meta `orm:"table:hg_food_order"`
		*entity.FoodOrder
	} `json:"orderDetail" orm:"with:id=order_id"  dc:"订单详情"`
}
