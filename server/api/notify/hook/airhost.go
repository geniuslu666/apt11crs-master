package hook

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AirhostHookReq struct {
	g.Meta     `path:"/airhost_hook" method:"post" tags:"NOTIFY_HOOK" summary:"airhost消息推送" group:"SYSTEM"`
	Id         string `json:"id"`
	ObjectType string `json:"object_type"`
	Event      string `json:"event"`
	Paginator  struct {
		NextPageUrl interface{} `json:"next_page_url"`
	} `json:"paginator"`
	Body interface{} `json:"body"`
}

type AirhostHookRes struct {
	//Context interface{} `json:"context"`
}
