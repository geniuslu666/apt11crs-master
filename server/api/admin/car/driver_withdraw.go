package car

import (
	"APT/internal/model/input/input_car"
	"APT/internal/model/input/input_form"

	"github.com/gogf/gf/v2/frame/g"
)

// DriverWithdrawListReq 获取司机提现记录列表
type DriverWithdrawListReq struct {
	g.Meta `path:"/carDriverWithdraw/list" method:"get" tags:"ADMIN_CAR" summary:"获取司机提现记录列表"`
	input_car.DriverWithdrawListInp
}

type DriverWithdrawListRes struct {
	input_form.PageRes
	List []*input_car.DriverWithdrawListModel `json:"list"   dc:"数据列表"`
}

// DriverWithdrawViewReq 获取司机提现信息详情
type DriverWithdrawViewReq struct {
	g.Meta `path:"/carDriverWithdraw/view" method:"get" tags:"ADMIN_CAR" summary:"获取司机提现信息详情"`
	input_car.DriverWithdrawViewInp
}

type DriverWithdrawViewRes struct {
	*input_car.DriverWithdrawViewModel
}

type DriverWithdrawAgreeReq struct {
	g.Meta `path:"/carDriverWithdraw/agree" method:"post" tags:"ADMIN_CAR" summary:"司机提现申请_同意提现"`
	input_car.DriverWithdrawAgreeInp
}

type DriverWithdrawAgreeRes struct {
}

type DriverWithdrawDisagreeReq struct {
	g.Meta `path:"/carDriverWithdraw/disagree" method:"post" tags:"ADMIN_CAR" summary:"司机提现申请_拒绝提现"`
	input_car.DriverWithdrawDisagreeInp
}

type DriverWithdrawDisagreeRes struct{}

type DriverWithdrawTransferReq struct {
	g.Meta `path:"/carDriverWithdraw/transfer" method:"post" tags:"ADMIN_CAR" summary:"司机提现申请_提现转账"`
	input_car.DriverWithdrawTransferInp
}

type DriverWithdrawTransferRes struct {
}
