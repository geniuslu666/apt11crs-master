package systemMessage

import (
	"APT/internal/dao"
	"APT/internal/library/MobilePush"
	"APT/internal/model/do"
	"APT/internal/model/entity"
	"context"
	"encoding/json"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// SystemMessageParams 系统消息参数
type SystemMessageParams struct {
	Scene        string                 `json:"scene" v:"required|in:system,hotel,food,car,spa#场景不能为空|场景值无效"`
	Type         string                 `json:"type" v:"required|in:order,im#类型不能为空|类型值无效"`
	MemberId     uint64                 `json:"memberId" v:"required|min:1#会员ID不能为空|会员ID无效"`
	Title        map[string]string      `json:"title" v:"required#消息标题不能为空"`
	Content      map[string]string      `json:"content"`
	Image        string                 `json:"image"`
	AppLink      string                 `json:"appLink"`
	WxLink       string                 `json:"wxLink"`
	UrlParam     map[string]interface{} `json:"urlParam"`
	OperatorId   uint64                 `json:"operatorId"`
	OperatorRole string                 `json:"operatorRole"`
	OrderSn      string                 `json:"orderSn"`
	EnablePush   bool                   `json:"enablePush"`
	EnableSms    bool                   `json:"enableSms"`
	PushTitle    string                 `json:"pushTitle"`
	PushContent  string                 `json:"pushContent"`
	SmsTemplate  string                 `json:"smsTemplate"`
	SmsParams    map[string]string      `json:"smsParams"`
	ShowIndex    bool                   `json:"showIndex"`
}

// SystemMessageResult 系统消息处理结果
type SystemMessageResult struct {
	MessageId   uint64 `json:"messageId"`
	PushSuccess bool   `json:"pushSuccess"`
	SmsSuccess  bool   `json:"smsSuccess"`
}

var (
	SystemMessageLogger = g.Log().Path("logs/systemMessage")
)

// ProcessSystemMessage 处理系统消息（入库+推送+短信）
func ProcessSystemMessage(ctx context.Context, params *SystemMessageParams) (*SystemMessageResult, error) {
	// 参数验证
	if err := g.Validator().Data(params).Run(ctx); err != nil {
		SystemMessageLogger.Error(ctx, "系统消息参数验证失败", "error", err, "params", params)
		return nil, gerror.Wrap(err, "参数验证失败")
	}

	SystemMessageLogger.Info(ctx, "开始处理系统消息", "scene", params.Scene, "type", params.Type, "memberId", params.MemberId, "enablePush", params.EnablePush, "enableSms", params.EnableSms, "showIndex", params.ShowIndex)

	// 1. 消息入库
	messageId, err := saveMessageToDB(ctx, params)
	if err != nil {
		SystemMessageLogger.Error(ctx, "消息入库失败", "error", err, "params", params)
		return nil, gerror.Wrap(err, "消息入库失败")
	}

	result := &SystemMessageResult{
		MessageId:   messageId,
		PushSuccess: false,
		SmsSuccess:  false,
	}

	// 2. 异步处理推送和短信
	go func() {
		asyncCtx := gctx.New()

		// 手机推送
		if params.EnablePush {
			pushSuccess := handleMobilePush(asyncCtx, params, messageId)
			result.PushSuccess = pushSuccess

			// 更新推送状态
			updatePushStatus(asyncCtx, messageId, pushSuccess)
		}

		// 短信发送
		if params.EnableSms {
			smsSuccess := handleSmsMessage(asyncCtx, params, messageId)
			result.SmsSuccess = smsSuccess

			// 更新短信状态
			updateSmsStatus(asyncCtx, messageId, smsSuccess)
		}

		SystemMessageLogger.Info(asyncCtx, "异步处理完成",
			"messageId", messageId,
			"pushSuccess", result.PushSuccess,
			"smsSuccess", result.SmsSuccess,
		)
	}()

	SystemMessageLogger.Info(ctx, "系统消息处理完成", "messageId", messageId, "memberId", params.MemberId)
	return result, nil
}

// saveMessageToDB 保存消息到数据库
func saveMessageToDB(ctx context.Context, params *SystemMessageParams) (uint64, error) {
	var (
		titleJson    *gjson.Json
		contentJson  *gjson.Json
		urlParamJson *gjson.Json
		showIndexInt int
		err          error
	)

	// 转换多语言标题为JSON
	titleJson = gjson.New(params.Title)

	// 转换多语言内容为JSON（如果存在）
	if len(params.Content) > 0 {
		contentJson = gjson.New(params.Content)
	}

	// 转换URL参数为JSON（如果存在）
	if len(params.UrlParam) > 0 {
		urlParamJson = gjson.New(params.UrlParam)
	}

	if params.ShowIndex {
		showIndexInt = 1
	} else {
		showIndexInt = 0
	}

	if g.IsEmpty(params.OperatorRole) {
		params.OperatorRole = "SYSTEM"
	}

	// 构建数据对象
	data := &do.SystemMessage{
		Scene:        params.Scene,
		Type:         params.Type,
		MemberId:     params.MemberId,
		Title:        titleJson,
		Content:      contentJson,
		Image:        params.Image,
		AppLink:      params.AppLink,
		WxLink:       params.WxLink,
		UrlParam:     urlParamJson,
		ShowIndex:    showIndexInt,
		OperatorId:   params.OperatorId,
		OperatorRole: params.OperatorRole,
		OrderSn:      params.OrderSn,
		CreatedAt:    gtime.Now(),
		UpdatedAt:    gtime.Now(),
	}

	// 插入数据库
	result, err := dao.SystemMessage.Ctx(ctx).Data(data).Insert()
	if err != nil {
		return 0, gerror.Wrap(err, "数据库插入失败")
	}

	messageId, err := result.LastInsertId()
	if err != nil {
		return 0, gerror.Wrap(err, "获取插入ID失败")
	}

	SystemMessageLogger.Info(ctx, "消息入库成功", "messageId", messageId, "scene", params.Scene, "memberId", params.MemberId)
	return uint64(messageId), nil
}

// handleMobilePush 处理手机推送
func handleMobilePush(ctx context.Context, params *SystemMessageParams, messageId uint64) bool {
	var (
		memberInfo *entity.PmsMember
	)

	if !params.EnablePush {
		return false
	}

	// 查询用户的推送ID（这里需要根据member_id查询用户的推送标识）
	// 查询用户推送ID的逻辑（需要根据实际表结构调整）
	err := dao.PmsMember.Ctx(ctx).
		Where("id", params.MemberId).
		Where("member_no IS NOT NULL AND member_no != ''").
		Scan(&memberInfo)

	if err != nil {
		SystemMessageLogger.Error(ctx, "查询用户失败", "messageId", messageId)
		return false
	}

	if g.IsEmpty(memberInfo.MemberNo) {
		SystemMessageLogger.Warning(ctx, "用户没有推送ID", "memberId", params.MemberId)
		return false
	}

	SystemMessageLogger.Info(ctx, "开始发送手机推送", "messageId", messageId, "memberId", params.MemberId)

	// 准备推送内容
	pushTitle := params.PushTitle
	if pushTitle == "" {
		// 使用消息标题的中文版本
		if zhTitle, exists := params.Title["zh"]; exists {
			pushTitle = zhTitle
		} else {
			pushTitle = "系统消息"
		}
	}

	pushContent := params.PushContent
	if pushContent == "" {
		// 使用消息内容的中文版本
		if len(params.Content) > 0 {
			if zhContent, exists := params.Content["zh"]; exists {
				pushContent = zhContent
			}
		}
		if pushContent == "" {
			pushContent = pushTitle
		}
	}

	// 准备推送额外参数
	extras := make(map[string]string)
	if params.AppLink != "" {
		extras["path"] = params.AppLink
	}
	if len(params.UrlParam) > 0 {
		// 将UrlParam中的参数合并到extras中
		for key, value := range params.UrlParam {
			extras[key] = gconv.String(value)
		}
		// if paramJson, err := json.Marshal(params.UrlParam); err == nil {
		// 	extras["data"] = string(paramJson)
		// }
	}

	// 发送推送
	extrasJson, _ := json.Marshal(extras)
	SystemMessageLogger.Info(ctx, "发送手机推送参数", "pushTitle", pushTitle, "pushContent", pushContent, "memberNos", memberInfo.MemberNo, "extras", string(extrasJson))
	err = MobilePush.SendPush(ctx, pushTitle, pushContent, g.SliceStr{memberInfo.MemberNo}, extras)
	if err != nil {
		SystemMessageLogger.Error(ctx, "手机推送发送失败", "error", err, "messageId", messageId, "memberId", params.MemberId)
		return false
	}

	return true
}

// handleSmsMessage 处理短信发送
func handleSmsMessage(ctx context.Context, params *SystemMessageParams, messageId uint64) bool {
	if !params.EnableSms {
		return false
	}

	SystemMessageLogger.Info(ctx, "开始发送短信", "messageId", messageId, "memberId", params.MemberId)

	// TODO: 集成实际的短信服务
	// 这里模拟短信逻辑
	time.Sleep(200 * time.Millisecond) // 模拟网络延迟

	// 模拟短信成功/失败
	success := true // 实际应该根据短信服务返回结果判断

	if success {
		SystemMessageLogger.Info(ctx, "短信发送成功", "messageId", messageId, "memberId", params.MemberId)
	} else {
		SystemMessageLogger.Error(ctx, "短信发送失败", "messageId", messageId, "memberId", params.MemberId)
	}

	return success
}

// updatePushStatus 更新推送状态
func updatePushStatus(ctx context.Context, messageId uint64, success bool) {
	status := 0
	if success {
		status = 1
	}

	_, err := dao.SystemMessage.Ctx(ctx).
		Where(dao.SystemMessage.Columns().Id, messageId).
		Update(g.Map{
			dao.SystemMessage.Columns().MobilePushSuccess: status,
			dao.SystemMessage.Columns().UpdatedAt:         gtime.Now(),
		})

	if err != nil {
		SystemMessageLogger.Error(ctx, "更新推送状态失败", "error", err, "messageId", messageId)
	} else {
		SystemMessageLogger.Info(ctx, "推送状态更新成功", "messageId", messageId, "success", success)
	}
}

// updateSmsStatus 更新短信状态
func updateSmsStatus(ctx context.Context, messageId uint64, success bool) {
	status := 0
	if success {
		status = 1
	}

	_, err := dao.SystemMessage.Ctx(ctx).
		Where(dao.SystemMessage.Columns().Id, messageId).
		Update(g.Map{
			dao.SystemMessage.Columns().MobileSmsSuccess: status,
			dao.SystemMessage.Columns().UpdatedAt:        gtime.Now(),
		})

	if err != nil {
		SystemMessageLogger.Error(ctx, "更新短信状态失败", "error", err, "messageId", messageId)
	} else {
		SystemMessageLogger.Info(ctx, "短信状态更新成功", "messageId", messageId, "success", success)
	}
}

// GetMessagesByMember 获取会员的系统消息列表
func GetMessagesByMember(ctx context.Context, memberId uint64, page, pageSize int) ([]*entity.SystemMessage, int, error) {
	// 查询总数
	count, err := dao.SystemMessage.Ctx(ctx).
		Where(dao.SystemMessage.Columns().MemberId, memberId).
		Count()
	if err != nil {
		return nil, 0, gerror.Wrap(err, "查询消息总数失败")
	}

	// 查询消息列表
	var messages []*entity.SystemMessage
	err = dao.SystemMessage.Ctx(ctx).
		Where(dao.SystemMessage.Columns().MemberId, memberId).
		OrderDesc(dao.SystemMessage.Columns().CreatedAt).
		Page(page, pageSize).
		Scan(&messages)

	if err != nil {
		return nil, 0, gerror.Wrap(err, "查询消息列表失败")
	}

	return messages, count, nil
}

// MarkMessageAsRead 标记消息为已读
func MarkMessageAsRead(ctx context.Context, messageId, memberId uint64) error {
	// 检查消息是否存在且属于该用户
	var message *entity.SystemMessage
	err := dao.SystemMessage.Ctx(ctx).
		Where(dao.SystemMessage.Columns().Id, messageId).
		Where(dao.SystemMessage.Columns().MemberId, memberId).
		Scan(&message)

	if err != nil {
		return gerror.Wrap(err, "查询消息失败")
	}

	if message == nil {
		return gerror.New("消息不存在或无权限")
	}

	// 检查是否已读
	readCount, err := dao.SystemMessageRead.Ctx(ctx).
		Where(dao.SystemMessageRead.Columns().MessageId, messageId).
		Where(dao.SystemMessageRead.Columns().MemberId, memberId).
		Count()

	if err != nil {
		return gerror.Wrap(err, "查询已读状态失败")
	}

	if readCount > 0 {
		return nil // 已经标记为已读
	}

	// 插入已读记录
	_, err = dao.SystemMessageRead.Ctx(ctx).Data(&do.SystemMessageRead{
		MessageId: messageId,
		MemberId:  memberId,
		CreatedAt: gtime.Now(),
	}).Insert()

	if err != nil {
		return gerror.Wrap(err, "标记已读失败")
	}

	SystemMessageLogger.Info(ctx, "消息标记为已读", "messageId", messageId, "memberId", memberId)
	return nil
}

// GetUnreadCount 获取未读消息数量
func GetUnreadCount(ctx context.Context, memberId uint64) (int, error) {
	count, err := dao.SystemMessage.Ctx(ctx).
		LeftJoin("hg_system_message_read smr", "hg_system_message.id = smr.message_id AND smr.member_id = ?", gconv.String(memberId)).
		Where("hg_system_message.member_id", memberId).
		Where("smr.id IS NULL"). // 未读的消息
		Count()

	if err != nil {
		return 0, gerror.Wrap(err, "查询未读消息数量失败")
	}

	return count, nil
}
