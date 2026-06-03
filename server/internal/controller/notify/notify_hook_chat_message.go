package notify

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/MobilePush"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"APT/internal/websocket"
	"APT/utility/simple"
	"context"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/notify/hook"
)

func (c *ControllerHook) VisitorMessageCallUrl(ctx context.Context, req *hook.VisitorMessageCallUrlReq) (res *hook.VisitorMessageCallUrlRes, err error) {
	var (
		in            *input_basics.NoticeEditInp
		MemberInfo    *entity.AdminMember
		PmsMemberInfo *entity.PmsMember
	)
	g.Log().Path("logs/HOOK/KEFU_HOOK").Debugf(ctx, "req: %#v", req)
	res = new(hook.VisitorMessageCallUrlRes)
	if req.Type == "in" {
		if err = dao.AdminMember.Ctx(ctx).Where(g.MapStrAny{
			dao.AdminMember.Columns().KefuAccount: req.To,
		}).Scan(&MemberInfo); err != nil {
			g.Log().Path("logs/HOOK/KEFU_HOOK").Error(ctx, err)
			err = gerror.New("用户信息异常")
			return
		}
		if g.IsEmpty(MemberInfo) {
			err = gerror.New("用户信息异常")
			return
		}
		in = new(input_basics.NoticeEditInp)
		if req.From == "order" {
			in.Title = "订单新消息通知"
			in.Type = consts.NoticeTypeOrderHandle
		} else {
			in.Title = "客服新消息通知"
			in.Type = consts.NoticeTypeLetter
		}
		in.Content = req.Content
		in.Receiver = []int64{MemberInfo.Id}
		// 新增
		in.CreatedBy = MemberInfo.Id
		in.CreatedAt = gtime.Now()
		// 推送通知
		response := &websocket.WResponse{
			Event: "notice",
			Data:  in,
		}
		simple.SafeGo(ctx, func(ctx context.Context) {
			if in.Type == consts.NoticeTypeLetter {
				for _, receiverId := range in.Receiver {
					websocket.SendToUser(receiverId, response)
				}
			} else {
				websocket.SendToAll(response)
			}
		})
	} else if req.Type == "out" {
		// 向客户推送通知
		memberIds := strings.Split(req.To, "|")
		if len(memberIds) > 1 && !g.IsEmpty(memberIds[1]) {
			memberId := memberIds[1]
			// 查询用户信息
			if err = dao.PmsMember.Ctx(ctx).Where(g.MapStrAny{
				dao.PmsMember.Columns().Id: memberId,
			}).Scan(&PmsMemberInfo); err != nil {
				g.Log().Path("logs/HOOK/KEFU_HOOK").Error(ctx, err)
				return
			}

			systemMessageTitle := map[string]string{
				"zh":    "客服消息",
				"en":    "Customer Service Message",
				"ja":    "カスタマーサービスメッセージ",
				"ko":    "고객센터 메시지",
				"zh_CN": "客服消息",
			}
			systemMessageContent := map[string]string{
				"zh":    "你有一条新客服消息",
				"en":    "You have a new customer service message.",
				"ja":    "新しいカスタマーサービスメッセージがあります。",
				"ko":    "새로운 고객센터 메시지가 있습니다.",
				"zh_CN": "您有一條新的客服消息。",
			}
			appPushData := g.MapStrStr{
				"type": "0",
			}
			var (
				SystemMessageConfig *model.SystemMessageConfig
			)
			if SystemMessageConfig, err = service.BasicsConfig().GetSystemMessageConfig(ctx); err != nil {
				return
			}

			// 查询用户的手机号区号 来判断用户语言
			phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(PmsMemberInfo.Id).Value()
			var memberLanguage string
			if phoneArea.String() == "+86" {
				memberLanguage = "zh"
			} else if phoneArea.String() == "+81" {
				memberLanguage = "ja"
			} else if phoneArea.String() == "+82" {
				memberLanguage = "ko"
			} else if phoneArea.String() == "+886" || phoneArea.String() == "+852" || phoneArea.String() == "+853" {
				memberLanguage = "zh_CN"
			} else {
				memberLanguage = "en"
			}

			if !g.IsEmpty(PmsMemberInfo) {
				if SystemMessageConfig.IsMqPushOpen == 2 {
					if err = MobilePush.SendPush(ctx, "customer service message", "You have a customer service message", g.SliceStr{PmsMemberInfo.MemberNo}, g.MapStrStr{
						"path": "/service",
					}); err != nil {
						g.Log().Path("logs/HOOK/KEFU_HOOK").Error(ctx, err)
						return
					}
				}

				// 发送到消息队列
				service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
					SystemMessageTitle:   systemMessageTitle,
					SystemMessageContent: systemMessageContent,
					Scene:                "system",
					Type:                 "im",
					MemberId:             PmsMemberInfo.Id,
					Language:             memberLanguage,
					AppPushData:          appPushData,
					AppLink:              "/service",
					WxLink:               "/pages/cs/cs",
					EnablePush:           true,
					EnableSms:            false,
					PushTitle:            systemMessageTitle[memberLanguage],
					PushContent:          systemMessageContent[memberLanguage],
					OperatorId:           0,
					OperatorRole:         "SYSTEM",
				})
			}
		}
	}
	g.Log().Path("logs/HOOK/KEFU_HOOK").Info(ctx, "——————————————————————————————————————————————————————————————————————————————————————————")
	g.Log().Path("logs/HOOK/KEFU_HOOK").Info(ctx, req)
	return
}
