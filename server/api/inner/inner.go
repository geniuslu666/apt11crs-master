// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package inner

import (
	"context"

	"APT/api/inner/order"
	"APT/api/inner/system_message"
)

type IInnerOrder interface {
	OrderList(ctx context.Context, req *order.OrderListReq) (res *order.OrderListRes, err error)
}

type IInnerSystem_message interface {
	GetSystemMessageList(ctx context.Context, req *system_message.GetSystemMessageListReq) (res *system_message.GetSystemMessageListRes, err error)
	MarkMessageRead(ctx context.Context, req *system_message.MarkMessageReadReq) (res *system_message.MarkMessageReadRes, err error)
	GetUnreadCount(ctx context.Context, req *system_message.GetUnreadCountReq) (res *system_message.GetUnreadCountRes, err error)
}
