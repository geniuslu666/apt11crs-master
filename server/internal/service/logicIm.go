// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"APT/internal/library/ws"
	"APT/internal/model/input/input_ws"
	"context"
)

type (
	IWsService interface {
		ChatSendMessage(ctx context.Context, req *input_ws.ChatSendMessageIpt) (err error)
		CarSendMessage(ctx context.Context, req *input_ws.ChatSendMessageIpt) (MemberIds []*input_ws.ChatMemberIds, msg *ws.BaseMessage, SourceName string, err error)
		SpaSendMessage(ctx context.Context, req *input_ws.ChatSendMessageIpt) (MemberIds []*input_ws.ChatMemberIds, msg *ws.BaseMessage, SourceName string, err error)
		// ChatMessageUnreadCount 会员订单未读消息数
		ChatMessageUnreadCount(ctx context.Context, in *input_ws.OrderUnreadListInput) (list []*input_ws.OrderUnreadListModel, err error)
		// UpdateLatestMsgRead 更新已读消息
		UpdateLatestMsgRead(ctx context.Context, in *input_ws.UpdateLatestMsgReadInput) (err error)
	}
)

var (
	localWsService IWsService
)

func WsService() IWsService {
	if localWsService == nil {
		panic("implement not found for interface IWsService, forgot register?")
	}
	return localWsService
}

func RegisterWsService(i IWsService) {
	localWsService = i
}
