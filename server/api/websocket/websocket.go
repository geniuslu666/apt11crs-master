// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package websocket

import (
	"context"

	"APT/api/websocket/basics"
)

type IWebsocketBasics interface {
	SendToTag(ctx context.Context, req *basics.SendToTagReq) (res *basics.SendToTagRes, err error)
}
