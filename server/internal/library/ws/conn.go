package ws

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

var (
	WsConnMap map[int64]*WsConn
)

type WsConn struct {
	Ctx       context.Context
	Conn      *websocket.Conn
	InChan    chan []byte
	OutChan   chan []byte
	closeChan chan byte
	IsClose   bool
	mutex     sync.Mutex
	Key       string
	MemberId  int64
	MemberNo  string
	Role      string `json:"role" dc:"角色  MEMBER   会员   CS  客服"`
}

type BaseMessage struct {
	Code    int         `json:"code"`
	Event   string      `json:"event"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Login struct {
	ClientKey string `json:"client_key"`
}

type ChatMessage struct {
	MessageId    int64       `json:"messageId"`
	Source       string      `json:"source" dc:"Member、Driver、System、Technician、DriverLeader、TechnicianLeader、TechnicianISP"`
	SourceId     int         `json:"sourceId"`
	SourceName   string      `json:"sourceName"`
	SourcePhoto  string      `json:"sourcePhoto"`
	FromMemberId int64       `json:"fromMemberId"`
	ToMemberId   int64       `json:"toMemberId"`
	OrderSn      string      `json:"orderSn"`
	Message      string      `json:"message"`
	MessageType  string      `json:"messageType"`
	CreatedAt    *gtime.Time `json:"createdAt"`
}

type Verify struct {
}

// SendWebsocketMessageInp 发送socket消息
type SendWebsocketMessageInp struct {
	MemberId int         `json:"memberId" v:"required#memberId不能为空" description:"会员ID"`
	Code     int         `json:"code"     description:"响应码"`
	Event    string      `json:"event"    description:"事件"`
	Message  string      `json:"message"  description:"响应信息"`
	Data     interface{} `json:"data"`
}

// InitWsConn 初始化websocket链接
func InitWsConn(ctx context.Context, conn *websocket.Conn, MemberId int64, MemberNo string) *WsConn {
	return &WsConn{
		Ctx:       ctx,
		Conn:      conn,
		InChan:    make(chan []byte, 1000),
		OutChan:   make(chan []byte, 1000),
		closeChan: make(chan byte, 1),
		Key:       guid.S(),
		IsClose:   false,
		MemberId:  MemberId,
		MemberNo:  MemberNo,
	}
}

func GetWebSocketConn(memberId int64) (Ws *WsConn, err error) {
	_, ok := WsConnMap[memberId]
	if !ok {
		err = errors.New("websocket is not exist")
		return
	}
	Ws = WsConnMap[memberId]
	return
}
func (conn *WsConn) AddWsConnMap() {
	g.Dump("AddWsConnMap key:%s", conn.Key)
	if WsConnMap == nil {
		WsConnMap = make(map[int64]*WsConn)
	}
	online, err := FindWsIsOnline(conn.MemberId)
	if err != nil {
		return
	}
	if online {
		// 如果用户当前存在链接   则关闭这个链接
		WsConnMap[conn.MemberId].Conn.Close()
	}
	WsConnMap[conn.MemberId] = conn
	returnMessageObj := &BaseMessage{
		Code:    200,
		Event:   "CONNECT",
		Message: "链接成功",
		Data: Login{
			ClientKey: conn.Key,
		},
	}
	conn.OutChan <- []byte(gjson.New(returnMessageObj).String())
	return
}

// FindWsIsOnline 查询用户是否在线
func FindWsIsOnline(MemberId int64) (Online bool, err error) {
	for _, v := range WsConnMap {
		if v.MemberId == MemberId {
			Online = true
			return
		}
	}
	return
}

// FindOnlineList 查询用户在线列表
func FindOnlineList() (OnlineList []*WsConn) {
	for _, v := range WsConnMap {
		OnlineList = append(OnlineList, v)
	}
	return
}

func (conn *WsConn) Close() {
	conn.Conn.Close()
	conn.mutex.Lock()
	defer conn.mutex.Unlock()
	if !conn.IsClose {
		close(conn.closeChan)
		delete(WsConnMap, conn.MemberId)
		conn.IsClose = true
		g.Log().Debug(conn.Ctx, "websocket close")
	}
	g.Log().Debug(conn.Ctx, "websocket closed")
}

// ReadLoop 自动读取消息
func (conn *WsConn) ReadLoop() {
	var (
		data []byte
		err  error
	)
	for {
		if _, data, err = conn.Conn.ReadMessage(); err != nil {
			goto ERR
		}
		select {
		case conn.InChan <- data:
		case <-conn.closeChan:
			goto ERR
		}
	}
ERR:
	g.Log().Debug(conn.Ctx, "ReadLoop err:%v", err)
	conn.Close()
}

// WriteLoop 自动发送消息
func (conn *WsConn) WriteLoop() {
	var (
		data []byte
		err  error
	)
	for {
		select {
		case data = <-conn.OutChan:
		case <-conn.closeChan:
			goto ERR
		}
		g.Log().Debug(conn.Ctx, "WriteLoop data:%s", string(data))
		if err = conn.Conn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	g.Log().Debug(conn.Ctx, "WriteLoop err:%v", err)
	conn.Close()
}

func (conn *WsConn) HartBeatLoop() {
	var (
		err error
	)
	for {
		if err = conn.WriteMessage([]byte("ping")); err != nil {
			goto ERR
		}
		time.Sleep(time.Second * 2)
	}
ERR:
	g.Log().Debug(conn.Ctx, "HartBeatLoop err:%v", err)
	conn.Close()
}

func (conn *WsConn) ReadMessage() (data []byte, err error) {

	select {
	case data = <-conn.InChan:
	case <-conn.closeChan:
		err = errors.New("connection is close")
		g.Log().Debug(conn.Ctx, "ReadMessage err:%v", err)
	}
	return
}

func (conn *WsConn) WriteMessage(data []byte) (err error) {
	select {
	case conn.OutChan <- data:
	case <-conn.closeChan:
		err = errors.New("connection is close")
		g.Log().Debug(conn.Ctx, "WriteMessage err:%v", err)
	}
	return
}
