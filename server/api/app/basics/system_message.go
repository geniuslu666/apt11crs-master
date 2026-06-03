package basics

import (
	"APT/internal/model/input/input_basics"

	"github.com/gogf/gf/v2/frame/g"
)

type AppLatestMessageReq struct {
	g.Meta `path:"/message/latestMessage" method:"post" tags:"APP_BASICS" summary:"[系统消息]最新消息"`
}

type AppLatestMessageRes struct {
	*input_basics.MessageAppLatestModel
}

type AppMessageListReq struct {
	g.Meta `path:"/message/appList" method:"post" tags:"APP_BASICS" summary:"[系统消息]列表"`
	input_basics.MessageAppListInp
}

type AppMessageListRes struct {
	List  []*input_basics.MessageAppListModel `json:"list"   dc:"数据列表"`
	Count int                                 `json:"count"   dc:"数据总数"`
}

type AppMessageReadReq struct {
	g.Meta `path:"/message/read" method:"post" tags:"APP_BASICS" summary:"[系统消息]已读"`
	input_basics.MessageAppReadInp
}

type AppMessageReadRes struct {
}
