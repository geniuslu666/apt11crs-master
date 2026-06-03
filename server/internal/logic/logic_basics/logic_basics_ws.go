package logic_basics

import (
	"APT/internal/library/ws"
	"APT/internal/service"
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
)

type sBasicsWs struct{}

func NewBasicsWs() *sBasicsWs {
	return &sBasicsWs{}
}

func init() {
	service.RegisterBasicsWs(NewBasicsWs())
}

func (s *sBasicsWs) SendMemberWebsocketMessage(ctx context.Context, in *ws.SendWebsocketMessageInp) (err error) {
	var (
		Conn *ws.WsConn
	)

	if Conn, err = ws.GetWebSocketConn(gvar.New(in.MemberId).Int64()); err != nil {
		return
	}

	returnMessageObj := &ws.BaseMessage{
		Code:    in.Code,
		Event:   in.Event,
		Message: in.Message,
	}

	if err = Conn.WriteMessage([]byte(gjson.New(returnMessageObj).String())); err != nil {
		return
	}
	return
}
