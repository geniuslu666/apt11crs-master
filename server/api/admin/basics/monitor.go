package basics

import (
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type MonitorUserOfflineReq struct {
	g.Meta `path:"/monitor/userOffline" method:"post" tags:"ADMIN" summary:"在线用户_下线用户"`
	Id     string `json:"id" v:"required#SID不能为空" description:"SID"`
}

type MonitorUserOfflineRes struct{}

type MonitorUserOnlineListReq struct {
	g.Meta `path:"/monitor/userOnlineList" method:"get" tags:"ADMIN" summary:"在线用户_获取在线用户列表"`
	input_form.PageReq
	UserId    int64         `json:"userId"      description:"用户ID"`
	Username  string        `json:"username"    description:"用户名"`
	IP        string        `json:"ip"        description:"登录IP"`
	FirstTime []*gtime.Time `json:"firstTime"   description:"登录时间"`
}

type MonitorUserOnlineListRes struct {
	List []*MonitorUserOnlineModel `json:"list"   description:"数据列表"`
	input_form.PageRes
}

type MonitorUserOnlineModel struct {
	ID            string `json:"id"`
	IP            string `json:"ip"`
	Os            string `json:"os"`
	Browser       string `json:"browser"`
	FirstTime     int64  `json:"firstTime"`
	HeartbeatTime uint64 `json:"heartbeatTime"`
	App           string `json:"app"`
	UserId        int64  `json:"userId"`
	Username      string `json:"username"`
	Avatar        string `json:"avatar"`
}
