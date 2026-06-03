package appv2

import (
	"APT/internal/library/contexts"
	"APT/internal/library/ws"
	"APT/internal/model"
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gorilla/websocket"
	"log"
	"net/http"

	"APT/api/appv2/wsApi"
)

func (c *ControllerWsApi) Ws(ctx context.Context, req *wsApi.WsReq) (res *wsApi.WsRes, err error) {
	var (
		r        = ghttp.RequestFromCtx(ctx)
		upgrader = websocket.Upgrader{
			// socket配置跨域
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
		wsConn     *websocket.Conn
		SocketConn *ws.WsConn
		MemberInfo *model.MemberIdentity
	)
	MemberInfo = contexts.GetMemberUser(ctx)
	if g.IsEmpty(MemberInfo) {
		// 未登录
		err = gerror.New(gi18n.T(ctx, "not_login"))
		return
	}
	//建立socket链接
	if wsConn, err = upgrader.Upgrade(r.Response.Writer, r.Request, nil); err != nil {
		log.Println(err)
		return
	}
	defer wsConn.Close()
	SocketConn = ws.InitWsConn(ctx, wsConn, gvar.New(MemberInfo.Id).Int64(), MemberInfo.MemberNo)
	//memberId := 39
	//fmt.Println("memberId", memberId)
	// 是否存在
	defer SocketConn.Close()
	SocketConn.AddWsConnMap()
	go SocketConn.ReadLoop()
	//go SocketConn.HartBeatLoop()
	SocketConn.WriteLoop()
	return
}
func (c *ControllerWsApi) SendWebsocketMessage(ctx context.Context, req *wsApi.SendWebsocketMessageReq) (res *wsApi.SendWebsocketMessageRes, err error) {

	//var (
	//	Conn *ws.WsConn
	//)
	//if Conn, err = ws.GetWebSocketConn(req.ClientId); err != nil {
	//	return
	//}
	//if err = Conn.WriteMessage([]byte(req.Message)); err != nil {
	//	return
	//}
	return
}
