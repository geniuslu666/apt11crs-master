package app

import (
	"APT/api/app/terminal"
	"APT/internal/library/contexts"
	"APT/internal/model/input/input_terminal"
	"APT/internal/service"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
)

func (c *ControllerTerminal) VerifyLog(ctx context.Context, req *terminal.VerifyLogReq) (res *terminal.VerifyLogRes, err error) {
	var (
		MemberInfo = contexts.GetTerminalUser(ctx)
	)
	fmt.Println(gjson.New(MemberInfo))
	res = new(terminal.VerifyLogRes)
	if res.List, res.Count, err = service.TerminalTerminal().List(ctx, &input_terminal.VerifyListInp{
		PageReq:      req.PageReq,
		StoreId:      int(MemberInfo.StoreId),
		RestaurantId: int(MemberInfo.RestaurantId),
	}); err != nil {
		return
	}
	return
}

func (c *ControllerTerminal) VerifyLogView(ctx context.Context, req *terminal.VerifyLogViewReq) (res *terminal.VerifyLogViewRes, err error) {
	res = new(terminal.VerifyLogViewRes)
	if res.VerifyLogViewModel, err = service.TerminalTerminal().View(ctx, &req.VerifyLogViewInp); err != nil {
		return
	}
	return
}
func (c *ControllerTerminal) CodeView(ctx context.Context, req *terminal.CodeViewReq) (res *terminal.CodeViewRes, err error) {
	res = new(terminal.CodeViewRes)
	if res.CodeViewModel, err = service.TerminalTerminal().CodeView(ctx, &req.CodeViewInp); err != nil {
		return
	}
	return
}
func (c *ControllerTerminal) CodeVerify(ctx context.Context, req *terminal.CodeVerifyReq) (res *terminal.CodeVerifyRes, err error) {
	res = new(terminal.CodeVerifyRes)
	if res.CodeVerifyModel, err = service.TerminalTerminal().CodeVerify(ctx, &req.CodeVerifyInp); err != nil {
		return
	}
	return
}
