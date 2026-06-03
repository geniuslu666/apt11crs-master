// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package system_message

import (
	"github.com/gogf/gf/v2/frame/g"
)


// GetSystemMessageListReq 获取系统消息列表请求
type GetSystemMessageListReq struct {
	g.Meta   `path:"/system-message/list" method:"get" summary:"获取系统消息列表" tags:"系统消息"`
	MemberId uint64 `json:"memberId" v:"required|min:1" dc:"会员ID"`
	Scene    string `json:"scene" v:"in:system,hotel,food,car,spa" dc:"场景筛选"`
	Type     string `json:"type" v:"in:order,im" dc:"类型筛选"`
	Page     int    `json:"page" v:"min:1" d:"1" dc:"页码"`
	PageSize int    `json:"pageSize" v:"min:1|max:100" d:"20" dc:"每页数量"`
	Language string `json:"language" v:"in:zh,en" d:"zh" dc:"语言"`
}

// SystemMessageItem 系统消息项
type SystemMessageItem struct {
	Id                uint64 `json:"id" dc:"消息ID"`
	Scene             string `json:"scene" dc:"场景"`
	Type              string `json:"type" dc:"类型"`
	Title             string `json:"title" dc:"消息标题"`
	Content           string `json:"content" dc:"消息内容"`
	Image             string `json:"image" dc:"消息图片URL"`
	AppLink           string `json:"appLink" dc:"APP跳转链接"`
	WxLink            string `json:"wxLink" dc:"微信跳转链接"`
	UrlParam          string `json:"urlParam" dc:"跳转链接参数"`
	MobilePushSuccess *int   `json:"mobilePushSuccess" dc:"手机消息推送是否成功"`
	MobileSmsSuccess  *int   `json:"mobileSmsSuccess" dc:"手机短信是否发送成功"`
	IsRead            bool   `json:"isRead" dc:"是否已读"`
	ReadAt            string `json:"readAt" dc:"已读时间"`
	CreatedAt         string `json:"createdAt" dc:"创建时间"`
}

// GetSystemMessageListRes 获取系统消息列表响应
type GetSystemMessageListRes struct {
	List       []SystemMessageItem `json:"list" dc:"消息列表"`
	Total      int                 `json:"total" dc:"总数"`
	Page       int                 `json:"page" dc:"当前页码"`
	PageSize   int                 `json:"pageSize" dc:"每页数量"`
	UnreadCount int                `json:"unreadCount" dc:"未读消息数量"`
}

// MarkMessageReadReq 标记消息已读请求
type MarkMessageReadReq struct {
	g.Meta     `path:"/system-message/read" method:"post" summary:"标记消息已读" tags:"系统消息"`
	MemberId   uint64   `json:"memberId" v:"required|min:1" dc:"会员ID"`
	MessageIds []uint64 `json:"messageIds" v:"required|min:1" dc:"消息ID列表"`
}

// MarkMessageReadRes 标记消息已读响应
type MarkMessageReadRes struct {
	SuccessCount int      `json:"successCount" dc:"成功标记的消息数量"`
	FailCount    int      `json:"failCount" dc:"失败的消息数量"`
	Errors       []string `json:"errors,omitempty" dc:"错误信息列表"`
}

// GetUnreadCountReq 获取未读消息数量请求
type GetUnreadCountReq struct {
	g.Meta   `path:"/system-message/unread-count" method:"get" summary:"获取未读消息数量" tags:"系统消息"`
	MemberId uint64 `json:"memberId" v:"required|min:1" dc:"会员ID"`
	Scene    string `json:"scene" v:"in:system,hotel,food,car,spa" dc:"场景筛选"`
}

// GetUnreadCountRes 获取未读消息数量响应
type GetUnreadCountRes struct {
	UnreadCount int `json:"unreadCount" dc:"未读消息数量"`
}
