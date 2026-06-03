package input_car

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// CarOrderLogListInp 获取餐厅套餐预订单列表
type CarOrderLogListInp struct {
	input_form.PageReq
	OrderSn     string        `json:"orderSn" dc:"订单编号"`
	OrderStatus string        `json:"orderStatus" dc:"订单状态"`
	ActionWay   string        `json:"actionWay" dc:"操作名"`
	CreatedAt   []*gtime.Time `json:"createdAt"    dc:"创建时间"`
}

func (in *CarOrderLogListInp) Filter(ctx context.Context) (err error) {
	return
}

type CarOrderLogListModel struct {
	entity.CarOrderLog
	AdminMemberUsername string `json:"adminMemberUsername"      dc:"操作员"`
	OperateName         string `json:"operateName"      dc:"操作人姓名"`
	OperatePhone        string `json:"operatePhone"      dc:"操作人电话"`
	OrderDetail         *struct {
		gmeta.Meta `orm:"table:hg_car_order"`
		*entity.CarOrder
	} `json:"orderDetail" orm:"with:id=order_id"  dc:"订单详情"`
}
