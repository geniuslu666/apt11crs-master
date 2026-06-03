package logic_basics

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	"APT/internal/library/systemMessage"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"APT/utility/rabbitmq"
	"context"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/util/gconv"

	"APT/internal/library/contexts"
	"encoding/json"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
)

type sBasicsSystemMessage struct{}

func NewBasicsSystemMessage() *sBasicsSystemMessage {
	return &sBasicsSystemMessage{}
}

func init() {
	service.RegisterBasicsSystemMessage(NewBasicsSystemMessage())
}

func (s *sBasicsSystemMessage) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.SystemMessage.Ctx(ctx), option...)
}

func (s *sBasicsSystemMessage) SendMessage(ctx context.Context, in *input_basics.SendMessageInp) (err error) {
	queueData := &systemMessage.SystemMessageParams{
		Scene:        in.Scene,
		Type:         in.Type,
		MemberId:     uint64(in.MemberId),
		Title:        in.SystemMessageTitle,
		Content:      in.SystemMessageContent,
		AppLink:      in.AppLink,
		UrlParam:     gconv.Map(in.AppPushData),
		WxLink:       in.WxLink,
		EnablePush:   in.EnablePush,
		EnableSms:    in.EnableSms,
		PushTitle:    in.PushTitle,
		PushContent:  in.PushContent,
		ShowIndex:    in.ShowIndex,
		OperatorId:   uint64(in.OperatorId),
		OperatorRole: in.OperatorRole,
		OrderSn:      in.OrderSn,
	}

	// 序列化为JSON
	var dataBytes []byte
	dataBytes, err = gjson.New(queueData).MarshalJSON()
	if err != nil {
		return
	}

	// 发送到消息队列
	_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
		ExchangeName: consts.RabbitMQExchangeName,
		QueueName:    consts.RabbitMQQueueNameSystemMessage,
		DataByte:     dataBytes,
		Header:       nil,
	})
	return
}

func (s *sBasicsSystemMessage) AppList(ctx context.Context, in *input_basics.MessageAppListInp) (list []*input_basics.MessageAppListModel, totalCount int, err error) {
	var (
		SystemMessage   []*entity.SystemMessage
		TitleJson       *input_basics.LanguageJson
		ContentJson     *input_basics.LanguageJson
		AppUrlParamJson *input_basics.AppUrlParamModel
	)
	mod := s.Model(ctx)

	if !g.IsEmpty(in.MemberId) {
		mod = mod.Where(dao.SystemMessage.Columns().MemberId, in.MemberId)
	}

	if !g.IsEmpty(in.Scene) {
		mod = mod.Where(dao.SystemMessage.Columns().Scene, in.Scene)
	}

	if !g.IsEmpty(in.Type) {
		mod = mod.Where(dao.SystemMessage.Columns().Type, in.Type)
	}

	if in.ShowIndex {
		mod = mod.Where(dao.SystemMessage.Columns().ShowIndex, 1)
	}

	if in.UnRead {
		// 查询未读消息：LEFT JOIN hg_system_message_read 表，条件是已读表中没有对应记录
		mod = mod.Fields("hg_system_message.*"). // 明确指定查询主表的所有字段
								LeftJoin("hg_system_message_read smr", "hg_system_message.id = smr.message_id AND smr.member_id = "+gconv.String(in.MemberId)).
								Where("smr.id IS NULL") // 未读的消息（已读表中没有记录）
	}

	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.OrderDesc(dao.SystemMessage.Columns().CreatedAt)

	if in.Pagination {
		if err = mod.ScanAndCount(&SystemMessage, &totalCount, false); err != nil {
			err = gerror.Wrap(err, gi18n.T(ctx, "get_message_list_failed"))
			return
		}
	} else {
		if err = mod.Scan(&SystemMessage); err != nil {
			err = gerror.Wrap(err, gi18n.T(ctx, "get_message_list_failed"))
			return
		}
	}

	var Language = contexts.GetLanguage(ctx)
	if !g.IsEmpty(SystemMessage) {
		// 处理聊天消息发送者信息字段
		// 收集所有 im 类型消息的操作员ID，用于批量查询发送者信息
		// var operatorIds []uint64
		// operatorIdMap := make(map[uint64][]*entity.SystemMessage) // operatorId -> messages

		// for _, v := range SystemMessage {
		// 	if v.Type == "im" && v.OperatorId > 0 {
		// 		if _, exists := operatorIdMap[v.OperatorId]; !exists {
		// 			operatorIds = append(operatorIds, v.OperatorId)
		// 			operatorIdMap[v.OperatorId] = make([]*entity.SystemMessage, 0)
		// 		}
		// 		operatorIdMap[v.OperatorId] = append(operatorIdMap[v.OperatorId], v)
		// 	}
		// }

		// 批量查询发送者信息（这里需要根据 OperatorRole 查询不同的表）
		// 假设需要查询员工表或会员表，具体实现可能需要调整
		// var senderInfoMap = make(map[uint64]struct {
		// 	Avatar   string
		// 	Nickname string
		// })

		// if len(operatorIds) > 0 {
		// 	// 根据 operator_role 分组操作员ID
		// 	memberIds := make([]uint64, 0)
		// 	driverIds := make([]uint64, 0)
		// 	technicianIds := make([]uint64, 0)
		// 	ispIds := make([]uint64, 0)

		// 	// 分组收集不同角色的操作员ID
		// 	for _, v := range SystemMessage {
		// 		if v.Type == "im" && v.OperatorId > 0 {
		// 			switch v.OperatorRole {
		// 			case "MEMBER":
		// 				memberIds = append(memberIds, v.OperatorId)
		// 			case "DRIVER", "DRIVER_LEADER":
		// 				driverIds = append(driverIds, v.OperatorId)
		// 			case "TECHNICIAN", "TECHNICIAN_LEADER":
		// 				technicianIds = append(technicianIds, v.OperatorId)
		// 			case "TECHNICIAN_ISP":
		// 				ispIds = append(ispIds, v.OperatorId)
		// 			}
		// 		}
		// 	}

		// 	// 查询会员信息
		// 	if len(memberIds) > 0 {
		// 		var members []*entity.PmsMember
		// 		if err = dao.PmsMember.Ctx(ctx).
		// 			Fields("id", "avatar", "full_name").
		// 			WhereIn("id", memberIds).
		// 			Scan(&members); err == nil {
		// 			for _, member := range members {
		// 				senderInfoMap[uint64(member.Id)] = struct {
		// 					Avatar   string
		// 					Nickname string
		// 				}{
		// 					Avatar:   member.Avatar,   // 会员表有 avatar 字段
		// 					Nickname: member.FullName, // 会员表使用 full_name 字段
		// 				}
		// 			}
		// 		}
		// 	}

		// 	// 查询司机信息
		// 	if len(driverIds) > 0 {
		// 		var drivers []*entity.CarDriver
		// 		if err = dao.CarDriver.Ctx(ctx).
		// 			Fields("id", "photo", "nickname").
		// 			WhereIn("id", driverIds).
		// 			Scan(&drivers); err == nil {
		// 			for _, driver := range drivers {
		// 				senderInfoMap[uint64(driver.Id)] = struct {
		// 					Avatar   string
		// 					Nickname string
		// 				}{
		// 					Avatar:   driver.Photo, // 司机表使用 photo 字段
		// 					Nickname: driver.Nickname,
		// 				}
		// 			}
		// 		}
		// 	}

		// 	// 查询技师信息
		// 	if len(technicianIds) > 0 {
		// 		var technicians []*entity.SpaTechnician
		// 		if err = dao.SpaTechnician.Ctx(ctx).
		// 			Fields("id", "photo", "nickname").
		// 			WhereIn("id", technicianIds).
		// 			Scan(&technicians); err == nil {
		// 			for _, technician := range technicians {
		// 				senderInfoMap[uint64(technician.Id)] = struct {
		// 					Avatar   string
		// 					Nickname string
		// 				}{
		// 					Avatar:   technician.Photo, // 技师表使用 photo 字段
		// 					Nickname: technician.Nickname,
		// 				}
		// 			}
		// 		}
		// 	}

		// 	// 查询ISP技师信息
		// 	if len(ispIds) > 0 {
		// 		var isps []*entity.SpaIsp
		// 		if err = dao.SpaIsp.Ctx(ctx).
		// 			Fields("id", "name").
		// 			WhereIn("id", ispIds).
		// 			Scan(&isps); err == nil {
		// 			for _, isp := range isps {
		// 				senderInfoMap[uint64(isp.Id)] = struct {
		// 					Avatar   string
		// 					Nickname string
		// 				}{
		// 					Avatar:   "",       // ISP表没有头像字段
		// 					Nickname: isp.Name, // ISP表使用 name 字段
		// 				}
		// 			}
		// 		}
		// 	}
		// }

		// 批量查询已读状态，避免N+1查询问题
		messageIds := make([]uint64, 0, len(SystemMessage))
		for _, v := range SystemMessage {
			messageIds = append(messageIds, v.Id)
		}

		// 一次性查询所有消息的已读状态
		var readRecords []*entity.SystemMessageRead
		if err = dao.SystemMessageRead.Ctx(ctx).
			Fields(dao.SystemMessageRead.Columns().MessageId).
			Where(dao.SystemMessageRead.Columns().MemberId, in.MemberId).
			WhereIn(dao.SystemMessageRead.Columns().MessageId, messageIds).
			Scan(&readRecords); err != nil {
			err = gerror.Wrap(err, gi18n.T(ctx, "get_message_list_failed"))
			return
		}

		// 构建已读消息ID的map，用于快速查找
		readMap := make(map[uint64]bool, len(readRecords))
		for _, record := range readRecords {
			readMap[record.MessageId] = true
		}

		list = make([]*input_basics.MessageAppListModel, 0, len(SystemMessage))
		for _, v := range SystemMessage {
			item := &input_basics.MessageAppListModel{
				Id:        int(v.Id),
				Scene:     v.Scene,
				Type:      v.Type,
				WxLink:    v.WxLink,
				AppLink:   v.AppLink,
				IsRead:    readMap[v.Id], // 从map中快速查找是否已读
				OrderSn:   v.OrderSn,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			}

			// 处理聊天消息发送者信息
			// if v.Type == "im" && v.OperatorId > 0 {
			// 	if senderInfo, exists := senderInfoMap[v.OperatorId]; exists {
			// 		item.ImInfo = &struct {
			// 			Avatar   string `json:"avatar" dc:"头像"`
			// 			Nickname string `json:"nickname" dc:"昵称"`
			// 		}{
			// 			Avatar:   senderInfo.Avatar,
			// 			Nickname: senderInfo.Nickname,
			// 		}
			// 	}
			// }
			if !g.IsEmpty(v.Title) && !g.IsEmpty(v.Content) {
				TitleJson = new(input_basics.LanguageJson)
				ContentJson = new(input_basics.LanguageJson)
				if err = json.Unmarshal([]byte(v.Title.String()), &TitleJson); err != nil {
					err = gerror.New(gi18n.T(ctx, "failed_to_parse_path"))
					return
				}
				if err = json.Unmarshal([]byte(v.Content.String()), &ContentJson); err != nil {
					err = gerror.New(gi18n.T(ctx, "failed_to_parse_path"))
					return
				}

				switch Language {
				case "zh":
					item.Title = TitleJson.Zh
					item.Content = ContentJson.Zh
				case "zh_CN":
					item.Title = TitleJson.ZhCn
					item.Content = ContentJson.ZhCn
				case "en":
					item.Title = TitleJson.En
					item.Content = ContentJson.En
				case "ja":
					item.Title = TitleJson.Ja
					item.Content = ContentJson.Ja
				case "ko":
					item.Title = TitleJson.Ko
					item.Content = ContentJson.Ko
				}
			}

			if g.IsEmpty(v.AppLink) && g.IsEmpty(v.WxLink) {
				item.AppLinkType = 4
			} else {
				if err = json.Unmarshal([]byte(v.UrlParam.String()), &AppUrlParamJson); err != nil {
					err = gerror.New(gi18n.T(ctx, "failed_to_parse_path"))
					return
				}
				item.AppLinkType = gconv.Int(AppUrlParamJson.Type)
				var pushUrlParam g.MapStrStr
				switch item.AppLinkType {
				case 1:
					pushUrlParam = g.MapStrStr{
						"string": AppUrlParamJson.String,
					}
				case 2:
					pushUrlParam = g.MapStrStr{
						"param": AppUrlParamJson.Param,
					}
				}

				if !g.IsEmpty(pushUrlParam) {
					pushUrlParamJson, _ := json.Marshal(pushUrlParam)
					item.UrlParam = string(pushUrlParamJson)
				} else {
					item.UrlParam = ""
				}
			}

			list = append(list, item)
		}
	}
	return
}

func (s *sBasicsSystemMessage) Read(ctx context.Context, in *input_basics.MessageAppReadInp) (err error) {
	var (
		lastInsertId int64
		messageInfo  *entity.SystemMessage
		messageRead  *entity.SystemMessageRead
	)
	// 判断消息是不是这个用户的
	if err = s.Model(ctx).Where(dao.SystemMessage.Columns().Id, in.Id).Scan(&messageInfo); err != nil {
		return
	}
	if messageInfo.MemberId != uint64(in.MemberId) {
		// 操作失败，请稍后重试！
		err = gerror.Wrap(err, gi18n.T(ctx, "operation_failed"))
		return
	}

	// 判断消息是否已经读取
	if err = dao.SystemMessageRead.Ctx(ctx).Where(dao.SystemMessageRead.Columns().MemberId, in.MemberId).Where(dao.SystemMessageRead.Columns().MessageId, in.Id).Scan(&messageRead); err != nil {
		return
	}
	if !g.IsEmpty(messageRead) {
		// 批量标记相关消息为已读
		var relatedMessages []*entity.SystemMessage

		// 情况1：非system场景，相同scene、type、orderSn的消息
		if messageInfo.Scene != "system" && !g.IsEmpty(messageInfo.OrderSn) {
			if err = dao.SystemMessage.Ctx(ctx).
				Fields("id").
				Where(dao.SystemMessage.Columns().Scene, messageInfo.Scene).
				Where(dao.SystemMessage.Columns().Type, messageInfo.Type).
				Where(dao.SystemMessage.Columns().OrderSn, messageInfo.OrderSn).
				Where(dao.SystemMessage.Columns().MemberId, in.MemberId).
				WhereNotIn("id", dao.SystemMessageRead.Ctx(ctx).
					Fields("message_id").
					Where("member_id", in.MemberId)).
				Scan(&relatedMessages); err != nil {
				return
			}
		}

		// 情况2：system场景，type为im，wx_link为/pages/cs/cs的消息
		if messageInfo.Scene == "system" && messageInfo.Type == "im" && messageInfo.WxLink == "/pages/cs/cs" {
			if err = dao.SystemMessage.Ctx(ctx).
				Fields("id").
				Where(dao.SystemMessage.Columns().Scene, "system").
				Where(dao.SystemMessage.Columns().Type, "im").
				Where(dao.SystemMessage.Columns().WxLink, "/pages/cs/cs").
				Where(dao.SystemMessage.Columns().MemberId, in.MemberId).
				WhereNotIn("id", dao.SystemMessageRead.Ctx(ctx).
					Fields("message_id").
					Where("member_id", in.MemberId)).
				Scan(&relatedMessages); err != nil {
				return
			}
		}

		// 批量插入已读记录
		if len(relatedMessages) > 0 {
			var readRecords []entity.SystemMessageRead
			for _, msg := range relatedMessages {
				readRecords = append(readRecords, entity.SystemMessageRead{
					MemberId:  uint64(in.MemberId),
					MessageId: msg.Id,
				})
			}

			if len(readRecords) > 0 {
				dao.SystemMessageRead.Ctx(ctx).Data(readRecords).Insert()
			}
		}
		return
	}

	if lastInsertId, err = dao.SystemMessageRead.Ctx(ctx).Data(&entity.SystemMessageRead{
		MemberId:  uint64(in.MemberId),
		MessageId: uint64(in.Id),
	}).InsertAndGetId(); err != nil {
		// 操作失败，请稍后重试！
		err = gerror.Wrap(err, gi18n.T(ctx, "operation_failed"))
		return
	}
	if lastInsertId < 1 {
		// 操作失败，请稍后重试！
		err = gerror.Wrap(err, gi18n.T(ctx, "operation_failed"))
		return
	}

	// 批量标记相关消息为已读
	var relatedMessages []*entity.SystemMessage

	// 情况1：非system场景，相同scene、type、orderSn的消息
	if messageInfo.Scene != "system" && !g.IsEmpty(messageInfo.OrderSn) {
		if err = dao.SystemMessage.Ctx(ctx).
			Fields("id").
			Where(dao.SystemMessage.Columns().Scene, messageInfo.Scene).
			Where(dao.SystemMessage.Columns().Type, messageInfo.Type).
			Where(dao.SystemMessage.Columns().OrderSn, messageInfo.OrderSn).
			Where(dao.SystemMessage.Columns().MemberId, in.MemberId).
			WhereNotIn("id", dao.SystemMessageRead.Ctx(ctx).
				Fields("message_id").
				Where("member_id", in.MemberId)).
			Scan(&relatedMessages); err != nil {
			return
		}
	}

	// 情况2：system场景，type为im，wx_link为/pages/cs/cs的消息
	if messageInfo.Scene == "system" && messageInfo.Type == "im" && messageInfo.WxLink == "/pages/cs/cs" {
		if err = dao.SystemMessage.Ctx(ctx).
			Fields("id").
			Where(dao.SystemMessage.Columns().Scene, "system").
			Where(dao.SystemMessage.Columns().Type, "im").
			Where(dao.SystemMessage.Columns().WxLink, "/pages/cs/cs").
			Where(dao.SystemMessage.Columns().MemberId, in.MemberId).
			WhereNotIn("id", dao.SystemMessageRead.Ctx(ctx).
				Fields("message_id").
				Where("member_id", in.MemberId)).
			Scan(&relatedMessages); err != nil {
			return
		}
	}

	// 批量插入已读记录
	if len(relatedMessages) > 0 {
		var readRecords []entity.SystemMessageRead
		for _, msg := range relatedMessages {
			readRecords = append(readRecords, entity.SystemMessageRead{
				MemberId:  uint64(in.MemberId),
				MessageId: msg.Id,
			})
		}

		if len(readRecords) > 0 {
			dao.SystemMessageRead.Ctx(ctx).Data(readRecords).Insert()
		}
	}
	return
}

func (s *sBasicsSystemMessage) AppLatest(ctx context.Context, memberId int) (res *input_basics.MessageAppLatestModel, err error) {
	var (
		TitleJson       *input_basics.LanguageJson
		ContentJson     *input_basics.LanguageJson
		AppUrlParamJson *input_basics.AppUrlParamModel
		Language        = contexts.GetLanguage(ctx)
	)

	res = new(input_basics.MessageAppLatestModel)

	// 定义需要查询的场景
	scenes := []string{"system", "hotel", "food", "spa", "car"}

	for _, scene := range scenes {
		var message *entity.SystemMessage

		// 查询每个场景的最新消息
		if err = s.Model(ctx).
			Where(dao.SystemMessage.Columns().MemberId, memberId).
			Where(dao.SystemMessage.Columns().Scene, scene).
			OrderDesc(dao.SystemMessage.Columns().CreatedAt).
			Limit(1).
			Scan(&message); err != nil {
			continue // 如果查询失败，跳过这个场景
		}

		if message == nil {
			continue // 如果没有消息，跳过这个场景
		}

		// 构建消息项
		item := &input_basics.MessageAppLatestItem{
			Id:        int(message.Id),
			WxLink:    message.WxLink,
			AppLink:   message.AppLink,
			CreatedAt: message.CreatedAt,
			UpdatedAt: message.UpdatedAt,
		}

		// 检查是否已读
		var readRecord *entity.SystemMessageRead
		if err = dao.SystemMessageRead.Ctx(ctx).
			Where(dao.SystemMessageRead.Columns().MemberId, memberId).
			Where(dao.SystemMessageRead.Columns().MessageId, message.Id).
			Scan(&readRecord); err == nil {
			item.IsRead = readRecord != nil
		}

		// 统计同场景下的未读消息数量
		var unreadCount int
		if count, err := dao.SystemMessage.Ctx(ctx).
			Where(dao.SystemMessage.Columns().MemberId, memberId).
			Where(dao.SystemMessage.Columns().Scene, scene).
			WhereNotIn("id", dao.SystemMessageRead.Ctx(ctx).
				Fields("message_id").
				Where("member_id", memberId)).
			Count(); err == nil {
			unreadCount = count
		}
		item.UnReadNum = unreadCount

		// 处理多语言标题和内容
		if !g.IsEmpty(message.Title) && !g.IsEmpty(message.Content) {
			TitleJson = new(input_basics.LanguageJson)
			ContentJson = new(input_basics.LanguageJson)

			if err = json.Unmarshal([]byte(message.Title.String()), &TitleJson); err == nil {
				switch Language {
				case "zh":
					item.Title = TitleJson.Zh
				case "zh_CN":
					item.Title = TitleJson.ZhCn
				case "en":
					item.Title = TitleJson.En
				case "ja":
					item.Title = TitleJson.Ja
				case "ko":
					item.Title = TitleJson.Ko
				}
			}

			if err = json.Unmarshal([]byte(message.Content.String()), &ContentJson); err == nil {
				switch Language {
				case "zh":
					item.Content = ContentJson.Zh
				case "zh_CN":
					item.Content = ContentJson.ZhCn
				case "en":
					item.Content = ContentJson.En
				case "ja":
					item.Content = ContentJson.Ja
				case "ko":
					item.Content = ContentJson.Ko
				}
			}
		}

		// 处理链接参数
		if g.IsEmpty(message.AppLink) && g.IsEmpty(message.WxLink) {
			item.AppLinkType = 4
		} else {
			if err = json.Unmarshal([]byte(message.UrlParam.String()), &AppUrlParamJson); err == nil {
				item.AppLinkType = gconv.Int(AppUrlParamJson.Type)
				var pushUrlParam g.MapStrStr
				switch item.AppLinkType {
				case 1:
					pushUrlParam = g.MapStrStr{
						"string": AppUrlParamJson.String,
					}
				case 2:
					pushUrlParam = g.MapStrStr{
						"param": AppUrlParamJson.Param,
					}
				}

				if !g.IsEmpty(pushUrlParam) {
					pushUrlParamJson, _ := json.Marshal(pushUrlParam)
					item.UrlParam = string(pushUrlParamJson)
				} else {
					item.UrlParam = ""
				}
			}
		}

		// 根据场景分配到对应的字段
		switch scene {
		case "system":
			res.SystemMessage = item
		case "hotel":
			res.HotelMessage = item
		case "food":
			res.FoodMessage = item
		case "spa":
			res.SpaMessage = item
		case "car":
			res.CarMessage = item
		}
	}

	err = nil // 重置错误，避免最后一次查询的错误影响整体结果
	return
}
