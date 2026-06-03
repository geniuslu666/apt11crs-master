package basics

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type NoticeListReq struct {
	g.Meta `path:"/notice/list" method:"get" tags:"ADMIN" summary:"公告_获取公告列表"`
	input_basics.NoticeListInp
}

type NoticeListRes struct {
	List []*input_basics.NoticeListModel `json:"list"   dc:"数据列表"`
	input_form.PageRes
}

type NoticeViewReq struct {
	g.Meta `path:"/notice/view" method:"get" tags:"ADMIN" summary:"公告_获取指定信息"`
	input_basics.NoticeViewInp
}

type NoticeViewRes struct {
	*input_basics.NoticeViewModel
}

type NoticeEditReq struct {
	g.Meta `path:"/notice/edit" method:"post" tags:"ADMIN" summary:"公告_修改/新增公告"`
	input_basics.NoticeEditInp
}

type NoticeEditRes struct{}

type NoticeDeleteReq struct {
	g.Meta `path:"/notice/delete" method:"post" tags:"ADMIN" summary:"公告_删除公告"`
	input_basics.NoticeDeleteInp
}

type NoticeDeleteRes struct{}

type NoticeMaxSortReq struct {
	g.Meta `path:"/notice/maxSort" method:"get" tags:"ADMIN" summary:"公告最大排序"`
	input_basics.NoticeMaxSortInp
}

type NoticeMaxSortRes struct {
	*input_basics.NoticeMaxSortModel
}

type NoticeStatusReq struct {
	g.Meta `path:"/notice/status" method:"post" tags:"ADMIN" summary:"更新公告状态"`
	input_basics.NoticeStatusInp
}

type NoticeStatusRes struct{}

type NoticeEditNotifyReq struct {
	g.Meta `path:"/notice/editNotify" method:"post" tags:"ADMIN" summary:"修改/新增通知"`
	input_basics.NoticeEditInp
}

type NoticeEditNotifyRes struct{}

type NoticeEditNoticeReq struct {
	g.Meta `path:"/notice/editNotice" method:"post" tags:"ADMIN" summary:"修改/新增公告"`
	input_basics.NoticeEditInp
}

type NoticeEditNoticeRes struct{}

type NoticeEditLetterReq struct {
	g.Meta `path:"/notice/editLetter" method:"post" tags:"ADMIN" summary:"修改/新增私信"`
	input_basics.NoticeEditInp
}

type NoticeEditLetterRes struct{}

type NoticePullMessagesReq struct {
	g.Meta `path:"/notice/pullMessages" method:"get" tags:"ADMIN" summary:"拉取我的消息"`
	input_basics.PullMessagesInp
}

type NoticePullMessagesRes struct {
	*input_basics.PullMessagesModel
}

type NoticeReadAllReq struct {
	g.Meta `path:"/notice/readAll" method:"post" tags:"ADMIN" summary:"全部已读"`
	input_basics.NoticeReadAllInp
}

type NoticeReadAllRes struct {
}

type NoticeUpReadReq struct {
	g.Meta `path:"/notice/upRead" method:"post" tags:"ADMIN" summary:"更新已读"`
	input_basics.NoticeUpReadInp
}
type NoticeUpReadRes struct{}

type NoticeMessageListReq struct {
	g.Meta `path:"/notice/messageList" method:"get" tags:"ADMIN" summary:"我的消息列表"`
	input_basics.NoticeMessageListInp
}

type NoticeMessageListRes struct {
	List []*input_basics.NoticeMessageListModel `json:"list"   dc:"数据列表"`
	input_form.PageRes
}
