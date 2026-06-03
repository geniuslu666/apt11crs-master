package appv2

import (
	"APT/api/appv2/wsApi"
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/translate"
	"APT/internal/model/input/input_ws"
	"APT/internal/service"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerWsApi) ChatSendMessage(ctx context.Context, req *wsApi.ChatSendMessageReq) (res *wsApi.ChatSendMessageRes, err error) {
	res = new(wsApi.ChatSendMessageRes)
	if err = service.WsService().ChatSendMessage(ctx, req.ChatSendMessageIpt); err != nil {
		return
	}
	return
}
func (c *ControllerWsApi) ChatGetMessage(ctx context.Context, req *wsApi.ChatGetMessageReq) (res *wsApi.ChatGetMessageRes, err error) {
	res = new(wsApi.ChatGetMessageRes)
	Map := g.Map{
		dao.ImMessage.Columns().OrderSn: req.OrderSn,
	}
	if req.LastId > 0 {
		Map[dao.ImMessage.Columns().Id+"<"] = req.LastId
	}

	if err = dao.ImMessage.Ctx(ctx).
		Where(Map).
		Limit(req.PageSize).
		OrderDesc(dao.ImMessage.Columns().Id).
		Scan(&res.MessageList); err != nil {
		return
	}
	if len(res.MessageList) > 0 {
		res.LastId = res.MessageList[len(res.MessageList)-1].Id
	}

	if g.IsEmpty(req.LastId) && !g.IsEmpty(res.MessageList) {
		// 更新会员最新已读消息
		var (
			MemberInfo = contexts.GetMemberUser(ctx)
		)
		_ = service.WsService().UpdateLatestMsgRead(ctx, &input_ws.UpdateLatestMsgReadInput{
			OrderSn:  req.OrderSn,
			MemberId: int64(MemberInfo.Id),
			ImId:     res.MessageList[0].Id,
		})
	}

	return
}

func (c *ControllerWsApi) TranslateMessage(ctx context.Context, req *wsApi.TranslateMessageReq) (res *wsApi.TranslateMessageRes, err error) {
	var (
		language          = contexts.GetLanguage(ctx)
		DetectLanguageRes *translate.DetectObj
	)
	res = new(wsApi.TranslateMessageRes)
	// 校验当前语言
	if DetectLanguageRes, err = translate.DetectLanguage(ctx, req.Text); err != nil {
		fmt.Println("Error:", err)
		return
	}
	switch DetectLanguageRes.Data.Detections[0][0].Language {
	case "zh-CN":
		if language == "zh" {
			res.Text = req.Text
			return
		}
		break
	case "zh-TW":
		if language == "zh" {
			res.Text = req.Text
			return
		}
		break
	case "en":
		if language == "en" {
			res.Text = req.Text
			return
		}
		break
	case "ja":
		if language == "ja" {
			res.Text = req.Text
			return
		}
		break
	case "ko":
		if language == "ko" {
			res.Text = req.Text
			return
		}
		break
	}

	if language == "zh" {
		language = "zh-CN"
	} else if language == "zh_CN" {
		language = "zh-TW"
	}
	// 翻译
	if res.Text, err = translate.TranslateText(ctx, req.Text, language); err != nil {
		return
	}
	return
}
func (c *ControllerWsApi) ChatGetMessageUnread(ctx context.Context, req *wsApi.ChatGetMessageUnreadReq) (res *wsApi.ChatGetMessageUnreadRes, err error) {

	res = new(wsApi.ChatGetMessageUnreadRes)

	if res.OrderUnreadList, err = service.WsService().ChatMessageUnreadCount(ctx, &req.OrderUnreadListInput); err != nil {
		return
	}

	return
}
