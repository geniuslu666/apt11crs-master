// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/input/input_terminal"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	ITerminalTerminal interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_terminal.VerifyListInp) (list []*input_terminal.VerifyListModel, totalCount int, err error)
		View(ctx context.Context, in *input_terminal.VerifyLogViewInp) (res *input_terminal.VerifyLogViewModel, err error)
		CodeView(ctx context.Context, in *input_terminal.CodeViewInp) (res *input_terminal.CodeViewModel, err error)
		CodeVerify(ctx context.Context, in *input_terminal.CodeVerifyInp) (res *input_terminal.CodeVerifyModel, err error)
		FoodOrderVerify(ctx context.Context, in *input_terminal.FoodOrderVerifyInp) (err error)
		MemberCouponVerify(ctx context.Context, in *input_terminal.MemberCouponVerifyInp) (err error)
	}
)

var (
	localTerminalTerminal ITerminalTerminal
)

func TerminalTerminal() ITerminalTerminal {
	if localTerminalTerminal == nil {
		panic("implement not found for interface ITerminalTerminal, forgot register?")
	}
	return localTerminalTerminal
}

func RegisterTerminalTerminal(i ITerminalTerminal) {
	localTerminalTerminal = i
}
