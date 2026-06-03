// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package appv2

import (
	"context"

	"APT/api/appv2/hotel"
	"APT/api/appv2/wsApi"
)

type IAppv2Hotel interface {
	PreOrderCreate(ctx context.Context, req *hotel.PreOrderCreateReq) (res *hotel.PreOrderCreateRes, err error)
	PreMoreOrderCreate(ctx context.Context, req *hotel.PreMoreOrderCreateReq) (res *hotel.PreMoreOrderCreateRes, err error)
	PreOrderCreateDetail(ctx context.Context, req *hotel.PreOrderCreateDetailReq) (res *hotel.PreOrderCreateDetailRes, err error)
	OrderReceipt(ctx context.Context, req *hotel.OrderReceiptReq) (res *hotel.OrderReceiptRes, err error)
	OrderReceiptView(ctx context.Context, req *hotel.OrderReceiptViewReq) (res *hotel.OrderReceiptViewRes, err error)
	RoomTypeList(ctx context.Context, req *hotel.RoomTypeListReq) (res *hotel.RoomTypeListRes, err error)
}

type IAppv2WsApi interface {
	ChatSendMessage(ctx context.Context, req *wsApi.ChatSendMessageReq) (res *wsApi.ChatSendMessageRes, err error)
	ChatGetMessage(ctx context.Context, req *wsApi.ChatGetMessageReq) (res *wsApi.ChatGetMessageRes, err error)
	TranslateMessage(ctx context.Context, req *wsApi.TranslateMessageReq) (res *wsApi.TranslateMessageRes, err error)
	ChatGetMessageUnread(ctx context.Context, req *wsApi.ChatGetMessageUnreadReq) (res *wsApi.ChatGetMessageUnreadRes, err error)
	Ws(ctx context.Context, req *wsApi.WsReq) (res *wsApi.WsRes, err error)
	SendWebsocketMessage(ctx context.Context, req *wsApi.SendWebsocketMessageReq) (res *wsApi.SendWebsocketMessageRes, err error)
}
