// Package hook
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package hook

import (
	"APT/internal/dao"
	"APT/internal/library/MobilePush"

	"APT/internal/model/do"
	"context"
	"encoding/json"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemMessageHookParams 系统消息钩子参数
type SystemMessageHookParams struct {
	Scene      string                 `json:"scene" v:"required|in:system,hotel,food,car,spa" dc:"场景：system-系统，hotel-酒店，food-餐饮，car-接送机，spa-按摩"`
	Type       string                 `json:"type" v:"required|in:order,im" dc:"类型：order-订单流转，im-聊天"`
	MemberId   uint64                 `json:"memberId" v:"required|min:1" dc:"会员ID"`
	Title      map[string]string      `json:"title" v:"required" dc:"消息标题，多语言格式：{\"zh\":\"中文标题\",\"en\":\"English Title\"}"`
	Content    map[string]string      `json:"content" dc:"消息内容，多语言格式：{\"zh\":\"中文内容\",\"en\":\"English Content\"}"`
	Image      string                 `json:"image" dc:"消息图片URL"`
	AppLink    string                 `json:"appLink" dc:"APP跳转链接"`
	WxLink     string                 `json:"wxLink" dc:"微信跳转链接"`
	UrlParam   map[string]interface{} `json:"urlParam" dc:"跳转链接参数，JSON格式存储"`
	OperatorId uint64                 `json:"operatorId" dc:"操作员ID"`
	// 推送和短信配置
	EnablePush  bool              `json:"enablePush" dc:"是否启用手机推送"`
	EnableSms   bool              `json:"enableSms" dc:"是否启用短信发送"`
	PushTitle   string            `json:"pushTitle" dc:"推送标题（如果为空则使用消息标题的中文）"`
	PushContent string            `json:"pushContent" dc:"推送内容（如果为空则使用消息内容的中文）"`
	SmsTemplate string            `json:"smsTemplate" dc:"短信模板代码"`
	SmsParams   map[string]string `json:"smsParams" dc:"短信模板参数"`
}

// SystemMessageHookResult 系统消息钩子返回结果
type SystemMessageHookResult struct {
	MessageId   uint64 `json:"messageId" dc:"消息ID"`
	PushSuccess *bool  `json:"pushSuccess" dc:"推送是否成功"`
	SmsSuccess  *bool  `json:"smsSuccess" dc:"短信是否成功"`
	PushError   string `json:"pushError,omitempty" dc:"推送错误信息"`
	SmsError    string `json:"smsError,omitempty" dc:"短信错误信息"`
}

// ProcessSystemMessage 处理系统消息的钩子函数
// 功能：1. 消息入库 2. 手机推送 3. 短信发送
func ProcessSystemMessage(ctx context.Context, params *SystemMessageHookParams) (*SystemMessageHookResult, error) {
	var (
		result = &SystemMessageHookResult{}
		err    error
	)

	// 参数验证
	if err = g.Validator().Data(params).Run(ctx); err != nil {
		return nil, gerror.Wrap(err, "参数验证失败")
	}

	// 1. 消息入库
	messageId, err := insertSystemMessage(ctx, params)
	if err != nil {
		return nil, gerror.Wrap(err, "消息入库失败")
	}
	result.MessageId = uint64(messageId)

	// 2. 异步处理推送和短信（避免阻塞主流程）
	go func() {
		asyncCtx := context.Background()

		// 处理手机推送
		if params.EnablePush {
			pushSuccess, pushErr := handleMobilePush(asyncCtx, params)
			updatePushStatus(asyncCtx, messageId, pushSuccess)
			if pushErr != nil {
				result.PushError = pushErr.Error()
				g.Log().Error(asyncCtx, "手机推送失败", "messageId", messageId, "error", pushErr)
			}
			result.PushSuccess = &pushSuccess
		}

		// 处理短信发送
		if params.EnableSms {
			smsSuccess, smsErr := handleSmsSend(asyncCtx, params)
			updateSmsStatus(asyncCtx, messageId, smsSuccess)
			if smsErr != nil {
				result.SmsError = smsErr.Error()
				g.Log().Error(asyncCtx, "短信发送失败", "messageId", messageId, "error", smsErr)
			}
			result.SmsSuccess = &smsSuccess
		}
	}()

	return result, nil
}

// insertSystemMessage 插入系统消息到数据库
func insertSystemMessage(ctx context.Context, params *SystemMessageHookParams) (int64, error) {
	var (
		titleJson    *gjson.Json
		contentJson  *gjson.Json
		urlParamJson *gjson.Json
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

	// 构建插入数据
	data := do.SystemMessage{
		Scene:      params.Scene,
		Type:       params.Type,
		MemberId:   params.MemberId,
		Title:      titleJson,
		Content:    contentJson,
		Image:      params.Image,
		AppLink:    params.AppLink,
		WxLink:     params.WxLink,
		UrlParam:   urlParamJson,
		OperatorId: params.OperatorId,
	}

	// 使用DAO插入数据库
	result, err := dao.SystemMessage.Ctx(ctx).Data(data).Insert()
	if err != nil {
		return 0, gerror.Wrap(err, "数据库插入失败")
	}

	messageId, err := result.LastInsertId()
	if err != nil {
		return 0, gerror.Wrap(err, "获取插入ID失败")
	}

	return messageId, nil
}

// handleMobilePush 处理手机推送
func handleMobilePush(ctx context.Context, params *SystemMessageHookParams) (bool, error) {
	// 查询用户的推送ID（这里需要根据member_id查询用户的推送标识）
	var memberNos []string

	// 查询用户推送ID的逻辑（需要根据实际表结构调整）
	err := dao.PmsMember.Ctx(ctx).
		Fields("member_no").
		Where("id", params.MemberId).
		Where("member_no IS NOT NULL AND member_no != ''").
		Scan(&memberNos)

	if err != nil {
		return false, gerror.Wrap(err, "查询用户推送ID失败")
	}

	if len(memberNos) == 0 {
		g.Log().Warning(ctx, "用户没有推送ID", "memberId", params.MemberId)
		return false, gerror.New("用户没有推送ID")
	}

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
		if paramJson, err := json.Marshal(params.UrlParam); err == nil {
			extras["data"] = string(paramJson)
		}
	}

	// 发送推送
	err = MobilePush.SendPush(ctx, pushTitle, pushContent, memberNos, extras)
	if err != nil {
		return false, gerror.Wrap(err, "推送发送失败")
	}

	return true, nil
}

// handleSmsSend 处理短信发送
func handleSmsSend(ctx context.Context, params *SystemMessageHookParams) (bool, error) {
	// 获取用户手机号
	// var mobile string
	// err := dao.PmsMember.Ctx(ctx).
	// 	Fields("mobile").
	// 	Where("id", params.MemberId).
	// 	Where("mobile IS NOT NULL AND mobile != ''").
	// 	Scan(&mobile)

	// if err != nil {
	// 	return false, gerror.Wrap(err, "查询用户手机号失败")
	// }

	// if mobile == "" {
	// 	g.Log().Warning(ctx, "用户没有手机号", "memberId", params.MemberId)
	// 	return false, gerror.New("用户没有手机号")
	// }

	// // 发送短信
	// smsDriver := sms.New()
	// sendMsgInp := &input_basics.SendMsgInp{
	// 	Mobile:   mobile,
	// 	Template: params.SmsTemplate,
	// }

	// err = smsDriver.SendMsg(ctx, sendMsgInp)
	// if err != nil {
	// 	return false, gerror.Wrap(err, "短信发送失败")
	// }

	return true, nil
}

// updatePushStatus 更新推送状态
func updatePushStatus(ctx context.Context, messageId int64, success bool) {
	var status interface{} = nil
	if success {
		status = 1
	} else {
		status = 0
	}

	_, err := dao.SystemMessage.Ctx(ctx).
		Where("id", messageId).
		Update(g.Map{"mobile_push_success": status})

	if err != nil {
		g.Log().Error(ctx, "更新推送状态失败", "messageId", messageId, "error", err)
	}
}

// updateSmsStatus 更新短信状态
func updateSmsStatus(ctx context.Context, messageId int64, success bool) {
	var status interface{} = nil
	if success {
		status = 1
	} else {
		status = 0
	}

	_, err := dao.SystemMessage.Ctx(ctx).
		Where("id", messageId).
		Update(g.Map{"mobile_sms_success": status})

	if err != nil {
		g.Log().Error(ctx, "更新短信状态失败", "messageId", messageId, "error", err)
	}
}
