// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IWarningWorkWx interface {
		SendWarningWorkWx(ctx context.Context, Content string) (err error)
	}
)

var (
	localWarningWorkWx IWarningWorkWx
)

func WarningWorkWx() IWarningWorkWx {
	if localWarningWorkWx == nil {
		panic("implement not found for interface IWarningWorkWx, forgot register?")
	}
	return localWarningWorkWx
}

func RegisterWarningWorkWx(i IWarningWorkWx) {
	localWarningWorkWx = i
}
