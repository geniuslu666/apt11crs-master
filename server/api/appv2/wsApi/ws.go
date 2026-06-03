package wsApi

import (
	"github.com/gogf/gf/v2/frame/g"
)

type WsReq struct {
	g.Meta `path:"/ws" method:"get" tags:"APP_WS" summary:"链接_websocket"`
}

type WsRes struct {
}

type SendWebsocketMessageReq struct {
	g.Meta   `path:"/send_message" method:"get" tags:"APP_WS" summary:"发送消息给客户端"`
	ClientId string
	Message  string
}

type SendWebsocketMessageRes struct {
}
