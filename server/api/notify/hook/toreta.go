package hook

import "github.com/gogf/gf/v2/frame/g"

type ToretaReq struct {
	g.Meta         `path:"/toreta/callback" method:"post" tags:"NOTIFY_HOOK" summary:"Toreta通知"`
	RestaurantKey  string `json:"restaurant_key" dc:"餐厅key"`
	ResourceKey    string `json:"resource_key" dc:"资源key"`
	ResourceAction string `json:"resource_action" dc:"操作内容 create, update, destroy 三种类型"`
	ResourceType   string `json:"resource_type" dc:"资源类型 目前仅支持Reservation"`
	UpdatedAt      int    `json:"updated_at" dc:"更新时间"`
}

type ToretaRes struct{}
