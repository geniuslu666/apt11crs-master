package admin

import (
	"APT/api/admin/basics"
	"APT/internal/consts"
	"APT/internal/model/input/input_form"
	"APT/internal/websocket"
	"APT/utility/simple"
	"APT/utility/useragent"
	"context"
	"github.com/gogf/gf/v2/text/gstr"
	"sort"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerBasics) MonitorUserOffline(ctx context.Context, req *basics.MonitorUserOfflineReq) (res *basics.MonitorUserOfflineRes, err error) {
	client := websocket.Manager().GetClient(req.Id)
	if client == nil {
		err = gerror.New("客户端已离线")
		return
	}

	simple.SafeGo(ctx, func(ctx context.Context) {
		websocket.SendSuccess(client, "kick")
		websocket.Close(client)
	})
	return
}

func (c *ControllerBasics) MonitorUserOnlineList(ctx context.Context, req *basics.MonitorUserOnlineListReq) (res *basics.MonitorUserOnlineListRes, err error) {
	var (
		clients []*basics.MonitorUserOnlineModel
		i       int
	)

	if websocket.Manager().GetClientsLen() == 0 {
		return
	}

	for conn := range websocket.Manager().GetClients() {
		if conn.SendClose || conn.User == nil {
			continue
		}

		if req.UserId > 0 && req.UserId != conn.User.Id {
			continue
		}

		if req.Username != "" && req.Username != conn.User.Username {
			continue
		}

		if req.IP != "" && !gstr.Contains(conn.IP, req.IP) {
			continue
		}

		if len(req.FirstTime) == 2 && (conn.User.LoginAt.Before(req.FirstTime[0]) || conn.User.LoginAt.After(req.FirstTime[1])) {
			continue
		}

		clients = append(clients, &basics.MonitorUserOnlineModel{
			ID:            conn.ID,
			IP:            conn.IP,
			Os:            useragent.GetOs(conn.UserAgent),
			Browser:       useragent.GetBrowser(conn.UserAgent),
			FirstTime:     conn.User.LoginAt.Unix(),
			HeartbeatTime: conn.HeartbeatTime,
			App:           conn.User.App,
			UserId:        conn.User.Id,
			Username:      conn.User.Username,
			Avatar:        conn.User.Avatar,
		})
	}

	res = new(basics.MonitorUserOnlineListRes)
	res.PageRes.Pack(req, len(clients))

	sort.Slice(clients, func(i, j int) bool {
		if clients[i].FirstTime == clients[j].FirstTime {
			return clients[i].ID < clients[j].ID
		}
		return clients[i].FirstTime < clients[j].FirstTime
	})

	_, perPage, offset := input_form.CalPage(req.Page, req.PerPage)
	for k, v := range clients {
		if k >= offset && i <= perPage {
			if simple.IsDemo(ctx) {
				v.IP = consts.DemoTips
			}
			res.List = append(res.List, v)
			i++
		}
	}
	return
}
