package basics

import (
	"APT/internal/model/input/input_basics"
	"github.com/gogf/gf/v2/frame/g"
)

// SendToTagReq 发送标签消息
type SendToTagReq struct {
	g.Meta `path:"/send/toTag" method:"post" tags:"WebSocket" summary:"发送标签消息"`
	input_basics.SendToTagInp
}

type SendToTagRes struct {
}
