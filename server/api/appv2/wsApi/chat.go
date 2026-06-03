package wsApi

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_ws"
	"github.com/gogf/gf/v2/frame/g"
)

type ChatSendMessageReq struct {
	g.Meta `path:"/chatSendMessage" method:"post" tags:"APP_WS_V2" summary:"聊天发送消息"`
	*input_ws.ChatSendMessageIpt
}

type ChatSendMessageRes struct {
}

type ChatGetMessageReq struct {
	g.Meta   `path:"/chatGetMessage" method:"post" tags:"APP_WS_V2" summary:"订单聊天获取消息"`
	OrderSn  string `json:"orderSn" v:"required#please_enter_the_order_number" dc:"订单号"`
	LastId   int64  `json:"lastId" v:"required#please_enter_the_last_msg_id" dc:"最后一条消息ID"`
	PageSize int    `p:"pageSize" v:"required#page_length_unknown" dc:"页长"`
}

type ChatGetMessageRes struct {
	LastId      int64               `json:"lastId"`
	MessageList []*entity.ImMessage `json:"messageList"`
}

type TranslateMessageReq struct {
	g.Meta `path:"/translateMessage" method:"post" tags:"APP_WS_V2" summary:"翻译消息"`
	Text   string `json:"text" v:"required#please_enter_the_text_to_be_translated" dc:"要翻译的文本"`
}

type TranslateMessageRes struct {
	Text string `json:"text"`
}

type ChatGetMessageUnreadReq struct {
	g.Meta `path:"/chatGetMessageUnread" method:"post" tags:"APP_WS_V2" summary:"获取订单未读消息数"`
	input_ws.OrderUnreadListInput
}

type ChatGetMessageUnreadRes struct {
	OrderUnreadList []*input_ws.OrderUnreadListModel `json:"orderUnreadList"`
}
