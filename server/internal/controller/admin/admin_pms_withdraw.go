package admin

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"

	"APT/api/admin/pms"
)

func (c *ControllerPms) WithdrawStaffList(ctx context.Context, req *pms.WithdrawStaffListReq) (res *pms.WithdrawStaffListRes, err error) {
	req.PmsWithdrawListInp.Type = "STAFF"
	list, totalCount, err := service.BasicsWithdraw().List(ctx, &req.PmsWithdrawListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_basics.PmsWithdrawListModel{}
	}

	res = new(pms.WithdrawStaffListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerPms) WithdrawChannelList(ctx context.Context, req *pms.WithdrawChannelListReq) (res *pms.WithdrawChannelListRes, err error) {
	req.PmsWithdrawListInp.Type = "CHANNEL"
	list, totalCount, err := service.BasicsWithdraw().List(ctx, &req.PmsWithdrawListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_basics.PmsWithdrawListModel{}
	}

	res = new(pms.WithdrawChannelListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerPms) WithdrawView(ctx context.Context, req *pms.WithdrawViewReq) (res *pms.WithdrawViewRes, err error) {
	data, err := service.BasicsWithdraw().View(ctx, &req.PmsWithdrawViewInp)
	if err != nil {
		return
	}

	res = new(pms.WithdrawViewRes)
	res.PmsWithdrawViewModel = data
	return
}
func (c *ControllerPms) WithdrawAgreeStaff(ctx context.Context, req *pms.WithdrawAgreeStaffReq) (res *pms.WithdrawAgreeStaffRes, err error) {
	err = service.BasicsWithdraw().Agree(ctx, &req.PmsWithdrawAgreeInp)
	return
}
func (c *ControllerPms) WithdrawDisagreeStaff(ctx context.Context, req *pms.WithdrawDisagreeStaffReq) (res *pms.WithdrawDisagreeStaffRes, err error) {
	err = service.BasicsWithdraw().Disagree(ctx, &req.PmsWithdrawDisagreeInp)
	return
}
func (c *ControllerPms) WithdrawAgreeChannel(ctx context.Context, req *pms.WithdrawAgreeChannelReq) (res *pms.WithdrawAgreeChannelRes, err error) {
	err = service.BasicsWithdraw().Agree(ctx, &req.PmsWithdrawAgreeInp)
	return
}
func (c *ControllerPms) WithdrawDisagreeChannel(ctx context.Context, req *pms.WithdrawDisagreeChannelReq) (res *pms.WithdrawDisagreeChannelRes, err error) {
	err = service.BasicsWithdraw().Disagree(ctx, &req.PmsWithdrawDisagreeInp)
	return
}
func (c *ControllerPms) WithdrawTransferStaff(ctx context.Context, req *pms.WithdrawTransferStaffReq) (res *pms.WithdrawTransferStaffRes, err error) {
	err = service.BasicsWithdraw().Transfer(ctx, &req.PmsWithdrawTransferInp)
	return
}
func (c *ControllerPms) WithdrawTransferChannel(ctx context.Context, req *pms.WithdrawTransferChannelReq) (res *pms.WithdrawTransferChannelRes, err error) {
	err = service.BasicsWithdraw().Transfer(ctx, &req.PmsWithdrawTransferInp)
	return
}
