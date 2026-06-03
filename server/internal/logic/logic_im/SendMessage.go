package logic_im

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/MobilePush"
	"APT/internal/library/contexts"
	"APT/internal/library/ws"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_ws"
	"APT/internal/service"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
)

type sWsService struct{}

func NewWsService() *sWsService {
	return &sWsService{}
}

func init() {
	service.RegisterWsService(NewWsService())
}

func (s *sWsService) ChatSendMessage(ctx context.Context, req *input_ws.ChatSendMessageIpt) (err error) {
	var (
		Member     *model.MemberIdentity
		SocketConn *ws.WsConn
		MemberIds  []*input_ws.ChatMemberIds
		Msg        *ws.BaseMessage
		SourceName string
	)
	// 识别身份
	Member = contexts.GetMemberUser(ctx)

	if g.IsEmpty(req.Source) {
		req.Source = "Member"
	}

	switch req.OrderType {
	case "car":
		if MemberIds, Msg, SourceName, err = s.CarSendMessage(ctx, req); err != nil {
			return
		}
		break
	case "spa":
		if MemberIds, Msg, SourceName, err = s.SpaSendMessage(ctx, req); err != nil {
			return
		}
	}
	g.Log().Debug(ctx, "开始发送消息")
	g.Log().Debug(ctx, gjson.New(MemberIds).String())
	if Msg == nil {
		return gerror.New("消息内容不能为空")
	}

	var (
		SystemMessageConfig *model.SystemMessageConfig
	)
	if SystemMessageConfig, err = service.BasicsConfig().GetSystemMessageConfig(ctx); err != nil {
		return
	}

	for _, v := range MemberIds {
		if v.MemberId == 0 {
			continue
		}
		MemberInfo := new(entity.PmsMember)
		if err = dao.PmsMember.Ctx(ctx).
			Where(dao.PmsMember.Columns().Id, v.MemberId).
			Scan(&MemberInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}

		var pushData g.MapStrStr
		if SystemMessageConfig.IsMqPushOpen == 2 {
			// 原来的参数
			pushData = g.MapStrStr{
				"orderid": req.OrderSn,
				"name":    MemberInfo.FullName,
				"phone":   fmt.Sprintf("%s%s", gstr.Trim(MemberInfo.PhoneArea, "-"), gstr.Trim(MemberInfo.Phone, "-")),
				"member":  v.Source,
				"status":  "1",
			}
			if req.OrderSn[:1] == "C" {
				pushData["type"] = "C"
			} else if req.OrderSn[:1] == "S" {
				pushData["type"] = "S"
			}
		} else {
			pushData = g.MapStrStr{
				"orderid": req.OrderSn,
				"name":    MemberInfo.FullName,
				"phone":   fmt.Sprintf("%s%s", gstr.Trim(MemberInfo.PhoneArea, "-"), gstr.Trim(MemberInfo.Phone, "-")),
				"member":  v.Source,
				"status":  "start",
			}
			if req.OrderSn[:1] == "C" {
				pushData["orderType"] = "car"
			} else if req.OrderSn[:1] == "S" {
				pushData["orderType"] = "spa"
			}
		}

		// 队列发消息
		var OperatorRole string
		switch v.Source {
		case "Member":
			OperatorRole = "MEMBER"
			break
		case "Driver":
			OperatorRole = "DRIVER"
			break
		case "DriverLeader":
			OperatorRole = "DRIVER_LEADER"
			break
		case "TechnicianISP":
			OperatorRole = "TECHNICIAN_ISP"
			break
		case "TechnicianLeader":
			OperatorRole = "TECHNICIAN_LEADER"
			break
		case "Technician":
			OperatorRole = "TECHNICIAN"
			break
		default:
			OperatorRole = "UNKNOWN"
		}
		systemMessageTitle := map[string]string{
			"zh":    "服务消息",
			"en":    "Service Message",
			"ja":    "サービスメッセージ",
			"ko":    "서비스 메시지",
			"zh_CN": "服务消息",
		}
		pushDataJson, _ := json.Marshal(pushData)
		appPushData := g.MapStrStr{
			"type":  "2",
			"param": string(pushDataJson),
		}

		if SocketConn, err = ws.GetWebSocketConn(v.MemberId); err != nil {
			g.Log().Errorf(ctx, "当前用户ID【%d】socket 链接获取失败", v.MemberId)
			g.Log().Debug(ctx, "发送给用户ID【%d】", v.MemberId)

			// 发推送要排除掉自己
			if Member.Id == gvar.New(v.MemberId).Int() {
				continue
			}

			if SystemMessageConfig.IsMqPushOpen == 2 {
				if err = MobilePush.SendPush(ctx, "service message", gjson.New(Msg.Data).Get("message").String(), []string{v.MemberNo}, g.MapStrStr{
					"path": "/service_detail",
					"data": gjson.New(pushData).String(),
				}); err != nil {
					return
				}
			}

			// 发送到消息队列

			service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
				SystemMessageTitle: systemMessageTitle,
				SystemMessageContent: map[string]string{
					"zh":    SourceName + "：" + gjson.New(Msg.Data).Get("message").String(),
					"en":    SourceName + ": " + gjson.New(Msg.Data).Get("message").String(),
					"ja":    SourceName + "：" + gjson.New(Msg.Data).Get("message").String(),
					"ko":    SourceName + "：" + gjson.New(Msg.Data).Get("message").String(),
					"zh_CN": SourceName + "：" + gjson.New(Msg.Data).Get("message").String(),
				},
				Scene:        pushData["orderType"],
				Type:         "im",
				MemberId:     int(v.MemberId),
				Language:     contexts.GetLanguage(ctx),
				AppPushData:  appPushData,
				AppLink:      "/service_detail",
				WxLink:       fmt.Sprintf("/pages/im/online?orderType=%s&orderSn=%s", req.OrderType, req.OrderSn),
				EnablePush:   true,
				EnableSms:    false,
				PushTitle:    systemMessageTitle[contexts.GetLanguage(ctx)],
				PushContent:  SourceName + "：" + gjson.New(Msg.Data).Get("message").String(),
				OperatorId:   int(v.MemberId),
				OperatorRole: OperatorRole,
				OrderSn:      req.OrderSn,
			})

			err = nil
			continue
		}
		if err = SocketConn.WriteMessage([]byte(gjson.New(Msg).String())); err != nil {
			g.Log().Errorf(ctx, "当前用户ID【%d】消息推送失败", v.MemberId)
			g.Log().Debug(ctx, "发送给用户ID【%d】", v.MemberId)
			// 发推送要排除掉自己
			if Member.Id == gvar.New(v.MemberId).Int() {
				continue
			}

			var (
				SystemMessageConfig *model.SystemMessageConfig
			)
			if SystemMessageConfig, err = service.BasicsConfig().GetSystemMessageConfig(ctx); err != nil {
				return
			}

			if SystemMessageConfig.IsMqPushOpen == 2 {
				if err = MobilePush.SendPush(ctx, "service message", gjson.New(Msg.Data).Get("message").String(), []string{v.MemberNo}, g.MapStrStr{
					"path": "/service_detail",
					"data": gjson.New(pushData).String(),
				}); err != nil {
					return
				}

				// if err = MobilePush.SendPush(ctx, "service message", gjson.New(Msg.Data).Get("message").String(), []string{v.MemberNo}, pushData); err != nil {
				// 	return
				// }
			}

			// 发送到消息队列
			service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
				SystemMessageTitle: systemMessageTitle,
				SystemMessageContent: map[string]string{
					"zh":    SourceName + "：" + gjson.New(Msg.Data).Get("message").String(),
					"en":    SourceName + ": " + gjson.New(Msg.Data).Get("message").String(),
					"ja":    SourceName + "：" + gjson.New(Msg.Data).Get("message").String(),
					"ko":    SourceName + "：" + gjson.New(Msg.Data).Get("message").String(),
					"zh_CN": SourceName + "：" + gjson.New(Msg.Data).Get("message").String(),
				},
				Scene:        pushData["orderType"],
				Type:         "im",
				MemberId:     int(v.MemberId),
				Language:     contexts.GetLanguage(ctx),
				AppPushData:  appPushData,
				AppLink:      "/service_detail",
				WxLink:       fmt.Sprintf("/pages/im/online?orderType=%s&orderSn=%s", req.OrderType, req.OrderSn),
				EnablePush:   true,
				EnableSms:    false,
				PushTitle:    systemMessageTitle[contexts.GetLanguage(ctx)],
				PushContent:  SourceName + "：" + gjson.New(Msg.Data).Get("message").String(),
				OperatorId:   int(v.SourceId),
				OperatorRole: OperatorRole,
				OrderSn:      req.OrderSn,
			})

			err = nil
			continue
		}

		chatMsg, ok := Msg.Data.(*ws.ChatMessage)
		if ok {
			// 更新发送会员最新已读消息
			_ = s.UpdateLatestMsgRead(ctx, &input_ws.UpdateLatestMsgReadInput{
				OrderSn:  req.OrderSn,
				MemberId: v.MemberId,
				ImId:     chatMsg.MessageId,
			})
		}

	}

	return
}

func (s *sWsService) CarSendMessage(ctx context.Context, req *input_ws.ChatSendMessageIpt) (MemberIds []*input_ws.ChatMemberIds, msg *ws.BaseMessage, SourceName string, err error) {
	var (
		CarOrder           *entity.CarOrder
		Member             *model.MemberIdentity
		DriverInfo         *entity.CarDriver
		MemberDriverLeader *entity.CarDriver
		DriverLeader       []*entity.CarDriver
		FromMemberId       int64
		ToMemberId         int64
		Source             string
		SourceId           int
		SourcePhoto        string
		MessageId          int64
		MemberInfo         *entity.PmsMember
	)
	// 识别身份
	Member = contexts.GetMemberUser(ctx)
	if g.IsEmpty(Member) {
		Member.Id = 0
	}
	// 查询当前订单的人员ID
	if err = dao.CarOrder.Ctx(ctx).
		Where(dao.CarOrder.Columns().OrderSn, req.OrderSn).
		Scan(&CarOrder); err != nil {
		return
	}
	if g.IsEmpty(CarOrder) {
		err = gerror.New("不存在该订单")
		return
	}
	// 查询司机表中司机对应的身份信息
	if err = dao.CarDriver.Ctx(ctx).
		Where(dao.CarDriver.Columns().Id, CarOrder.DriverId).
		Scan(&DriverInfo); err != nil {
		return
	}
	if g.IsEmpty(DriverInfo) {
		err = gerror.New("司机不存在")
		return
	}
	MemberIds = append(MemberIds, &input_ws.ChatMemberIds{
		MemberId: gvar.New(CarOrder.MemberId).Int64(),
		Source:   "Member",
		SourceId: gvar.New(CarOrder.MemberId).Int64(),
		Name:     Member.FullName,
	})
	MemberIds = append(MemberIds, &input_ws.ChatMemberIds{
		MemberId: gvar.New(DriverInfo.MemberId).Int64(),
		Source:   "Driver",
		SourceId: gvar.New(DriverInfo.Id).Int64(),
		Name:     DriverInfo.Nickname,
	})
	if err = dao.CarDriver.Ctx(ctx).
		Where(dao.CarDriver.Columns().IsLeader, 1).
		Where(dao.CarDriver.Columns().Status, 1).
		WhereNot(dao.CarDriver.Columns().Id, DriverInfo.Id).
		Scan(&DriverLeader); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	for _, v := range DriverLeader {
		MemberIds = append(MemberIds, &input_ws.ChatMemberIds{
			MemberId: gvar.New(v.MemberId).Int64(),
			Source:   "DriverLeader",
			SourceId: gvar.New(v.Id).Int64(),
			Name:     v.Nickname,
		})
	}
	if !g.IsEmpty(req.Source) {
		for k, v := range MemberIds {
			if err = dao.PmsMember.Ctx(ctx).
				Where(dao.PmsMember.Columns().Id, v.MemberId).
				Scan(&MemberInfo); err != nil {
				return
			}
			if !g.IsEmpty(MemberInfo) {
				MemberIds[k].MemberNo = gvar.New(MemberInfo.MemberNo).String()
			}
			MemberInfo = nil
		}
		for _, v := range MemberIds {
			if v.MemberId == gvar.New(Member.Id).Int64() {
				Source = v.Source
				break
			}
		}
	}

	if req.Source == "Member" {
		// 发送方是乘客
		Source = "Member"
		SourceId = int(CarOrder.MemberId)
		SourceName = Member.FullName
		SourcePhoto = Member.Avatar
		FromMemberId = gvar.New(CarOrder.MemberId).Int64()
		ToMemberId = 0
	} else if req.Source == "Driver" {
		// 发送方是司机或车队长
		Source = req.Source
		SourceId = DriverInfo.Id
		SourceName = DriverInfo.Nickname
		SourcePhoto = DriverInfo.Photo
		FromMemberId = gvar.New(DriverInfo.MemberId).Int64()
		ToMemberId = 0
	} else if req.Source == "DriverLeader" {
		// 查询司机表中司机对应的身份信息
		if err = dao.CarDriver.Ctx(ctx).
			Where(dao.CarDriver.Columns().MemberId, Member.Id).
			Where(dao.CarDriver.Columns().IsLeader, 1).
			Where(dao.CarDriver.Columns().Status, 1).
			Scan(&MemberDriverLeader); err != nil {
			return
		}
		if g.IsEmpty(MemberDriverLeader) {
			err = gerror.New("司机不存在")
			return
		}
		// 发送方是司机或车队长
		Source = req.Source
		SourceId = MemberDriverLeader.Id
		SourceName = MemberDriverLeader.Nickname
		SourcePhoto = MemberDriverLeader.Photo
		FromMemberId = gvar.New(MemberDriverLeader.MemberId).Int64()
		ToMemberId = 0
	} else {
		// 发送发是系统
		Source = "System"
		FromMemberId = 0
		ToMemberId = 0
	}

	//if CarOrder.MemberId == gvar.New(Member.Id).Uint() {
	//	// 发送方是乘客
	//	Source = "Member"
	//	SourceId = int(CarOrder.MemberId)
	//	SourceName = Member.FullName
	//	SourcePhoto = Member.Avatar
	//	FromMemberId = gvar.New(CarOrder.MemberId).Int64()
	//	ToMemberId = 0
	//} else if DriverInfo.MemberId == Member.Id {
	//	if DriverInfo.IsLeader == 1 {
	//		// 发送方是车队长
	//		Source = "DriverLeader"
	//	} else {
	//		// 发送方是司机
	//		Source = "Driver"
	//	}
	//	SourceId = DriverInfo.Id
	//	SourceName = DriverInfo.Nickname
	//	SourcePhoto = DriverInfo.Photo
	//	FromMemberId = gvar.New(DriverInfo.MemberId).Int64()
	//	ToMemberId = 0
	//} else {
	//	// 发送发是系统
	//	Source = "System"
	//	FromMemberId = 0
	//	ToMemberId = 0
	//}

	// 插入聊天记录表
	if MessageId, err = dao.ImMessage.Ctx(ctx).Data(g.MapStrAny{
		dao.ImMessage.Columns().Source:       Source,
		dao.ImMessage.Columns().SourceId:     SourceId,
		dao.ImMessage.Columns().SourceName:   SourceName,
		dao.ImMessage.Columns().SourcePhoto:  SourcePhoto,
		dao.ImMessage.Columns().FromMemberId: FromMemberId,
		dao.ImMessage.Columns().ToMemberId:   ToMemberId,
		dao.ImMessage.Columns().OrderSn:      req.OrderSn,
	}).InsertAndGetId(); err != nil {
		return
	}
	// 构造消息格式内容
	msg = &ws.BaseMessage{
		Code:    200,
		Event:   "chatMessage",
		Message: "im",
		Data: &ws.ChatMessage{
			MessageId:    MessageId,
			Source:       Source,
			SourceId:     SourceId,
			SourceName:   SourceName,
			SourcePhoto:  SourcePhoto,
			FromMemberId: FromMemberId,
			ToMemberId:   ToMemberId,
			OrderSn:      req.OrderSn,
			Message:      req.Message,
			MessageType:  req.Type,
			CreatedAt:    gtime.Now(),
		},
	}
	if _, err = dao.ImMessage.Ctx(ctx).WherePri(MessageId).Update(g.Map{
		dao.ImMessage.Columns().Content: gjson.New(msg),
	}); err != nil {
		return
	}
	return
}
func (s *sWsService) SpaSendMessage(ctx context.Context, req *input_ws.ChatSendMessageIpt) (MemberIds []*input_ws.ChatMemberIds, msg *ws.BaseMessage, SourceName string, err error) {
	var (
		SpaOrder               *entity.SpaOrder
		Member                 *model.MemberIdentity
		SpaTechnician          *entity.SpaTechnician
		MemberTechnicianLeader *entity.SpaTechnician
		TechnicianLeader       []*entity.SpaTechnician
		SpaIsp                 *entity.SpaIsp
		FromMemberId           int64
		ToMemberId             int64
		Source                 string
		SourceId               int
		SourcePhoto            string
		MessageId              int64
		MemberInfo             *entity.PmsMember
	)
	// 识别身份
	Member = contexts.GetMemberUser(ctx)
	if g.IsEmpty(Member) {
		Member.Id = 0
	}
	// 查询当前订单的人员ID
	if err = dao.SpaOrder.Ctx(ctx).
		Where(dao.SpaOrder.Columns().OrderSn, req.OrderSn).
		Scan(&SpaOrder); err != nil {
		return
	}
	if g.IsEmpty(SpaOrder) {
		err = gerror.New("不存在该订单")
		return
	}
	// 查询技师表中技师对应的身份信息
	if err = dao.SpaTechnician.Ctx(ctx).
		Where(dao.SpaTechnician.Columns().Id, SpaOrder.TechnicianIds).
		Scan(&SpaTechnician); err != nil {
		return
	}
	if g.IsEmpty(SpaTechnician) {
		err = gerror.New("技师不存在")
		return
	}

	MemberIds = append(MemberIds, &input_ws.ChatMemberIds{
		MemberId: gvar.New(SpaOrder.MemberId).Int64(),
		Source:   "Member",
		SourceId: gvar.New(SpaOrder.MemberId).Int64(),
		Name:     Member.FullName,
	})
	MemberIds = append(MemberIds, &input_ws.ChatMemberIds{
		MemberId: gvar.New(SpaTechnician.MemberId).Int64(),
		Source:   "Technician",
		SourceId: gvar.New(SpaTechnician.Id).Int64(),
		Name:     SpaTechnician.Nickname,
	})
	// 查询店长
	if err = dao.SpaTechnician.Ctx(ctx).
		Where(dao.SpaTechnician.Columns().IsLeader, 1).
		Where(dao.SpaTechnician.Columns().Status, 1).
		WhereNot(dao.SpaTechnician.Columns().Id, SpaTechnician.Id).
		Scan(&TechnicianLeader); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	// 查询技师表中技师对应的身份信息
	if err = dao.SpaIsp.Ctx(ctx).
		Where(dao.SpaIsp.Columns().Id, SpaOrder.IspId).
		Where(dao.SpaIsp.Columns().Status, 1).
		WhereNot(dao.SpaIsp.Columns().MemberId, SpaTechnician.MemberId).
		Scan(&SpaIsp); err != nil {
		return
	}
	if !g.IsEmpty(SpaIsp) {
		MemberIds = append(MemberIds, &input_ws.ChatMemberIds{
			MemberId: gvar.New(SpaIsp.MemberId).Int64(),
			Source:   "TechnicianISP",
			SourceId: gvar.New(SpaIsp.Id).Int64(),
			Name:     SpaIsp.Name,
		})
	}

	for _, v := range TechnicianLeader {
		MemberIds = append(MemberIds, &input_ws.ChatMemberIds{
			MemberId: gvar.New(v.MemberId).Int64(),
			Source:   "TechnicianLeader",
			SourceId: gvar.New(v.Id).Int64(),
			Name:     v.Nickname,
		})
	}
	if !g.IsEmpty(req.Source) {

		for k, v := range MemberIds {
			if err = dao.PmsMember.Ctx(ctx).
				Where(dao.PmsMember.Columns().Id, v.MemberId).
				Scan(&MemberInfo); err != nil {
				return
			}
			if !g.IsEmpty(MemberInfo) {
				MemberIds[k].MemberNo = gvar.New(MemberInfo.MemberNo).String()
			}
			MemberInfo = nil
		}

		for _, v := range MemberIds {
			if v.MemberId == gvar.New(Member.Id).Int64() {
				Source = v.Source
				break
			}
		}
	}
	if req.Source == "Member" {
		// 发送方是用户
		Source = "Member"
		SourceId = int(SpaOrder.MemberId)
		SourceName = Member.FullName
		SourcePhoto = Member.Avatar
		FromMemberId = gvar.New(SpaOrder.MemberId).Int64()
		ToMemberId = 0
	} else if req.Source == "Technician" {
		// 发送方是普通技师或店长
		Source = req.Source
		SourceId = SpaTechnician.Id
		SourceName = SpaTechnician.Nickname
		SourcePhoto = SpaTechnician.Photo
		FromMemberId = gvar.New(SpaTechnician.MemberId).Int64()
		ToMemberId = 0
	} else if req.Source == "TechnicianLeader" {
		// 查询技师表中技师对应的身份信息
		if err = dao.SpaTechnician.Ctx(ctx).
			Where(dao.SpaTechnician.Columns().MemberId, Member.Id).
			Where(dao.SpaTechnician.Columns().IsLeader, 1).
			Where(dao.SpaTechnician.Columns().Status, 1).
			Scan(&MemberTechnicianLeader); err != nil {
			return
		}
		if g.IsEmpty(MemberTechnicianLeader) {
			err = gerror.New("技师不存在")
			return
		}

		// 发送方是普通技师或店长
		Source = req.Source
		SourceId = MemberTechnicianLeader.Id
		SourceName = MemberTechnicianLeader.Nickname
		SourcePhoto = MemberTechnicianLeader.Photo
		FromMemberId = gvar.New(MemberTechnicianLeader.MemberId).Int64()
		ToMemberId = 0
	} else if req.Source == "TechnicianISP" {
		// 发送方是服务商
		Source = "TechnicianISP"
		SourceId = SpaIsp.Id
		SourceName = SpaIsp.Name
		FromMemberId = gvar.New(SpaIsp.MemberId).Int64()
		ToMemberId = 0
	} else {
		// 发送发是系统
		Source = "System"
		FromMemberId = 0
		ToMemberId = 0
	}

	//if SpaOrder.MemberId == gvar.New(Member.Id).Uint() {
	//	// 发送方是乘客
	//	Source = "Member"
	//	SourceId = int(SpaOrder.MemberId)
	//	SourceName = Member.FullName
	//	SourcePhoto = Member.Avatar
	//	FromMemberId = gvar.New(SpaOrder.MemberId).Int64()
	//	ToMemberId = 0
	//} else if SpaTechnician.MemberId == Member.Id {
	//	if SpaTechnician.IsLeader == 1 {
	//		// 发送方是店长
	//		Source = "TechnicianLeader"
	//	} else {
	//		// 发送方是技师
	//		Source = "Technician"
	//	}
	//	SourceId = SpaTechnician.Id
	//	SourceName = SpaTechnician.Nickname
	//	SourcePhoto = SpaTechnician.Photo
	//	FromMemberId = gvar.New(SpaTechnician.MemberId).Int64()
	//	ToMemberId = 0
	//} else if !g.IsEmpty(SpaIsp) && SpaIsp.MemberId == Member.Id {
	//	// 发送方是服务商
	//	Source = "TechnicianISP"
	//	SourceId = SpaIsp.Id
	//	SourceName = SpaIsp.Name
	//	FromMemberId = gvar.New(SpaIsp.MemberId).Int64()
	//	ToMemberId = 0
	//} else {
	//	// 发送发是系统
	//	Source = "System"
	//	FromMemberId = 0
	//	ToMemberId = 0
	//}

	// 插入聊天记录表
	if MessageId, err = dao.ImMessage.Ctx(ctx).Data(g.MapStrAny{
		dao.ImMessage.Columns().Source:       Source,
		dao.ImMessage.Columns().SourceId:     SourceId,
		dao.ImMessage.Columns().SourceName:   SourceName,
		dao.ImMessage.Columns().SourcePhoto:  SourcePhoto,
		dao.ImMessage.Columns().FromMemberId: FromMemberId,
		dao.ImMessage.Columns().ToMemberId:   ToMemberId,
		dao.ImMessage.Columns().OrderSn:      req.OrderSn,
	}).InsertAndGetId(); err != nil {
		return
	}
	// 构造消息格式内容
	msg = &ws.BaseMessage{
		Code:    200,
		Event:   "chatMessage",
		Message: "im",
		Data: &ws.ChatMessage{
			MessageId:    MessageId,
			Source:       Source,
			SourceId:     SourceId,
			SourceName:   SourceName,
			SourcePhoto:  SourcePhoto,
			FromMemberId: FromMemberId,
			ToMemberId:   ToMemberId,
			OrderSn:      req.OrderSn,
			Message:      req.Message,
			MessageType:  req.Type,
			CreatedAt:    gtime.Now(),
		},
	}
	if _, err = dao.ImMessage.Ctx(ctx).WherePri(MessageId).Update(g.Map{
		dao.ImMessage.Columns().Content: gjson.New(msg),
	}); err != nil {
		return
	}
	return
}

// ChatMessageUnreadCount 会员订单未读消息数
func (s *sWsService) ChatMessageUnreadCount(ctx context.Context, in *input_ws.OrderUnreadListInput) (list []*input_ws.OrderUnreadListModel, err error) {
	var (
		MemberInfo = contexts.GetMemberUser(ctx)
	)
	err = g.Model("hg_im_message m").
		LeftJoin(" hg_im_message_read r", "m.order_sn = r.order_sn AND r.member_id = "+gvar.New(MemberInfo.Id).String()).
		Fields("m.order_sn, COUNT(m.id) AS unread_count").
		WhereIn("m.order_sn", strings.Split(in.OrderSn, ",")).
		Where("(r.im_id IS NULL OR m.id > r.im_id)").
		Group("m.order_sn").
		Scan(&list)

	if err != nil {
		err = gerror.Wrap(err, "获取订单未读消息数失败，请稍后重试！")
		return
	}
	return
}

// UpdateLatestMsgRead 更新已读消息
func (s *sWsService) UpdateLatestMsgRead(ctx context.Context, in *input_ws.UpdateLatestMsgReadInput) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		mod := dao.ImMessageRead.Ctx(ctx).Safe()
		// 查询会员订单已读消息
		var models *entity.ImMessageRead
		if err = dao.ImMessageRead.Ctx(ctx).Safe().
			Where(dao.ImMessageRead.Columns().MemberId, in.MemberId).
			Where(dao.ImMessageRead.Columns().OrderSn, in.OrderSn).
			Scan(&models); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return
		}

		if !g.IsEmpty(models) {
			if _, err = mod.WherePri(models.Id).Data(g.Map{
				dao.ImMessageRead.Columns().ImId: in.ImId,
			}).Update(); err != nil {
				err = gerror.Wrap(err, "更新已读消息失败，请稍后重试！")
				return
			}
		} else {
			if _, err = mod.Data(g.Map{
				dao.ImMessageRead.Columns().MemberId: in.MemberId,
				dao.ImMessageRead.Columns().OrderSn:  in.OrderSn,
				dao.ImMessageRead.Columns().ImId:     in.ImId,
			}).OmitEmptyData().Insert(); err != nil {
				err = gerror.Wrap(err, "新增已读消息失败，请稍后重试！")
			}
		}
		return
	})
}
