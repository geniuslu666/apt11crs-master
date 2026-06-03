package queue

import (
	"APT/internal/consts"
	"APT/internal/library/systemMessage"
	_ "APT/internal/logic"
	"APT/internal/model"
	"APT/internal/service"
	"APT/utility/rabbitmq"
	"context"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
	amqp "github.com/rabbitmq/amqp091-go"
)

// SystemMessageQueueData 系统消息队列数据结构
type SystemMessageQueueData struct {
	TaskId       string                 `json:"taskId"`       // 任务ID
	Scene        string                 `json:"scene"`        // 场景
	Type         string                 `json:"type"`         // 类型
	MemberId     uint64                 `json:"memberId"`     // 会员ID
	Title        map[string]string      `json:"title"`        // 消息标题
	Content      map[string]string      `json:"content"`      // 消息内容
	Image        string                 `json:"image"`        // 消息图片URL
	AppLink      string                 `json:"appLink"`      // APP跳转链接
	WxLink       string                 `json:"wxLink"`       // 微信跳转链接
	UrlParam     map[string]interface{} `json:"urlParam"`     // 跳转链接参数
	OperatorId   uint64                 `json:"operatorId"`   // 操作员ID
	OperatorRole string                 `json:"operatorRole"` // 操作员角色
	OrderSn      string                 `json:"orderSn"`      // 订单号
	EnablePush   bool                   `json:"enablePush"`   // 是否启用手机推送
	EnableSms    bool                   `json:"enableSms"`    // 是否启用短信发送
	PushTitle    string                 `json:"pushTitle"`    // 推送标题
	PushContent  string                 `json:"pushContent"`  // 推送内容
	SmsTemplate  string                 `json:"smsTemplate"`  // 短信模板代码
	SmsParams    map[string]string      `json:"smsParams"`    // 短信模板参数
	ShowIndex    bool                   `json:"showIndex"`    // 是否显示在首页
	RetryCount   int                    `json:"retryCount"`   // 重试次数
	CreatedAt    gtime.Time             `json:"createdAt"`    // 创建时间
}

var (
	SystemMessageLogger = g.Log().Path("logs/MQ/" + consts.RabbitMQQueueNameSystemMessage)
)

func SystemMessage() {
	var (
		ctx          = gctx.New()
		MQMsg        <-chan amqp.Delivery
		exchangeName = consts.RabbitMQExchangeName
		QueueName    = consts.RabbitMQQueueNameSystemMessage
		MQConnStruct *rabbitmq.MQConnection
		err          error
	)
	g.DB().SetLogger(SystemMessageLogger)
	MQConnStruct = &rabbitmq.MQConnection{
		Conn:         rabbitmq.Conn,
		Chan:         nil,
		ExchangeName: exchangeName,
		QueueName:    QueueName,
		RouteKey:     fmt.Sprintf("%s.%s", exchangeName, QueueName),
	}
	if err = MQConnStruct.Channel(); err != nil {
		err = gerror.New("创建队列失败")
		goto ERR
	}
	if err = MQConnStruct.Exchange("topic", true, false); err != nil {
		err = gerror.New("创建交换机失败")
		goto ERR
	}
	if err = MQConnStruct.Queue(true, false, nil); err != nil {
		err = gerror.New("创建通道失败")
		goto ERR
	}
	if err = MQConnStruct.Bind(); err != nil {
		err = gerror.New("交换机绑定队列失败")
		goto ERR
	}
	SystemMessageLogger.Info(ctx, "System Message Queue START SUCCESSFUL")
	if MQMsg, err = MQConnStruct.Consume(guid.S(), false); err != nil {
		SystemMessageLogger.Error(ctx, err.Error())
		return
	}
	for msg := range MQMsg {
		ctx = gctx.New()
		SystemMessageLogger.Info(ctx, "--[start]----------------------------------------------")
		SystemMessageLogger.Info(ctx, string(msg.Body))
		var MsgBody *SystemMessageQueueData

		if err = json.Unmarshal(msg.Body, &MsgBody); err != nil {
			SystemMessageLogger.Error(ctx, "解析消息体失败", "error", err, "body", string(msg.Body))
			_ = msg.Nack(false, false) // 拒绝消息，不重新入队
			continue
		}

		if g.IsEmpty(MsgBody.TaskId) {
			// 生成任务ID
			MsgBody.TaskId = guid.S()
		}

		if err = HandleSystemMessageMQ(ctx, MsgBody); err != nil {
			SystemMessageLogger.Error(ctx, "处理系统消息失败", "error", err, "taskId", MsgBody.TaskId)

			// 检查重试次数
			MsgBody.RetryCount++
			if MsgBody.RetryCount < 3 { // 最多重试3次
				SystemMessageLogger.Info(ctx, "消息处理失败，准备重试", "taskId", MsgBody.TaskId, "retryCount", MsgBody.RetryCount)

				// 延迟重新发布到队列
				go func(retryData *SystemMessageQueueData) {
					time.Sleep(time.Duration(retryData.RetryCount) * time.Minute) // 递增延迟
					if err := republishMessage(ctx, retryData); err != nil {
						SystemMessageLogger.Error(ctx, "重试消息发布失败", "error", err, "taskId", retryData.TaskId)
					}
				}(MsgBody)
			} else {
				SystemMessageLogger.Error(ctx, "消息处理失败，超过最大重试次数", "taskId", MsgBody.TaskId, "retryCount", MsgBody.RetryCount)
			}

			_ = msg.Nack(false, false) // 拒绝消息
			continue
		}

		SystemMessageLogger.Info(ctx, "系统消息处理成功", "taskId", MsgBody.TaskId, "memberId", MsgBody.MemberId)
		_ = msg.Ack(false)
	}
ERR:
	SystemMessageLogger.Error(ctx, err)
	panic(err)
}

func HandleSystemMessageMQ(ctx context.Context, data *SystemMessageQueueData) (err error) {
	SystemMessageLogger.Info(ctx, "开始处理系统消息", "taskId", data.TaskId, "memberId", data.MemberId, "scene", data.Scene)
	defer func() {
		if r := recover(); r != nil {
			SystemMessageLogger.Error(ctx, r)
		}
	}()
	// 转换为系统消息参数
	// EnablePush字段根据sys_config表中的isMqPushOpen来判断，如果isMqPushOpen为2，则false，否则就是传入的data.EnablePush
	var (
		SystemMessageConfig *model.SystemMessageConfig
	)
	if SystemMessageConfig, err = service.BasicsConfig().GetSystemMessageConfig(ctx); err != nil {
		return
	}

	if SystemMessageConfig.IsMqPushOpen == 2 {
		return nil
	}

	params := &systemMessage.SystemMessageParams{
		Scene:        data.Scene,
		Type:         data.Type,
		MemberId:     data.MemberId,
		Title:        data.Title,
		Content:      data.Content,
		Image:        data.Image,
		AppLink:      data.AppLink,
		WxLink:       data.WxLink,
		UrlParam:     data.UrlParam,
		OperatorId:   data.OperatorId,
		OperatorRole: data.OperatorRole,
		OrderSn:      data.OrderSn,
		EnablePush:   data.EnablePush,
		EnableSms:    data.EnableSms,
		PushTitle:    data.PushTitle,
		PushContent:  data.PushContent,
		SmsTemplate:  data.SmsTemplate,
		SmsParams:    data.SmsParams,
		ShowIndex:    data.ShowIndex,
	}

	// 直接调用系统消息处理
	result, err := systemMessage.ProcessSystemMessage(ctx, params)
	if err != nil {
		return gerror.Wrap(err, "系统消息处理失败")
	}

	SystemMessageLogger.Info(ctx, "系统消息处理成功",
		"taskId", data.TaskId,
		"messageId", result.MessageId,
		"pushSuccess", result.PushSuccess,
		"smsSuccess", result.SmsSuccess,
	)
	return nil
}

// republishMessage 重新发布消息到队列
func republishMessage(ctx context.Context, data *SystemMessageQueueData) error {
	// 序列化消息数据
	messageBody, err := json.Marshal(data)
	if err != nil {
		return gerror.Wrap(err, "消息序列化失败")
	}

	// 使用项目的RabbitMQ连接模式
	mqConn := &rabbitmq.MQConnection{
		Conn:         rabbitmq.Conn,
		Chan:         nil,
		ExchangeName: consts.RabbitMQExchangeName,
		QueueName:    consts.RabbitMQQueueNameSystemMessage,
		RouteKey:     fmt.Sprintf("%s.%s", consts.RabbitMQExchangeName, consts.RabbitMQQueueNameSystemMessage),
	}

	// 创建通道
	if err := mqConn.Channel(); err != nil {
		return gerror.Wrap(err, "创建队列通道失败")
	}

	// 创建交换机
	if err := mqConn.Exchange("topic", true, false); err != nil {
		return gerror.Wrap(err, "创建交换机失败")
	}

	// 创建队列
	if err := mqConn.Queue(true, false, nil); err != nil {
		return gerror.Wrap(err, "创建队列失败")
	}

	// 绑定队列
	if err := mqConn.Bind(); err != nil {
		return gerror.Wrap(err, "绑定队列失败")
	}

	// 发布消息
	err = mqConn.Chan.Publish(
		consts.RabbitMQExchangeName,
		fmt.Sprintf("%s.%s", consts.RabbitMQExchangeName, consts.RabbitMQQueueNameSystemMessage),
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         messageBody,
			DeliveryMode: amqp.Persistent,
			Priority:     0,
			Timestamp:    time.Now(),
			MessageId:    data.TaskId,
		},
	)
	if err != nil {
		return gerror.Wrap(err, "发布消息失败")
	}

	SystemMessageLogger.Info(ctx, "重试消息已重新发布到队列", "taskId", data.TaskId, "retryCount", data.RetryCount)
	return nil
}

// PublishSystemMessageToQueue 发布系统消息到队列的便捷方法
func PublishSystemMessageToQueue(ctx context.Context, data *SystemMessageQueueData) error {
	// 生成任务ID
	if g.IsEmpty(data.TaskId) {
		data.TaskId = guid.S()
	}

	// 设置创建时间
	if data.CreatedAt.IsZero() {
		data.CreatedAt = *gtime.Now()
	}

	// 序列化消息数据
	messageBody, err := json.Marshal(data)
	if err != nil {
		SystemMessageLogger.Error(ctx, "系统消息队列数据序列化失败", "error", err, "data", data)
		return gerror.Wrap(err, "消息序列化失败")
	}

	// 使用项目的RabbitMQ连接模式
	mqConn := &rabbitmq.MQConnection{
		Conn:         rabbitmq.Conn,
		Chan:         nil,
		ExchangeName: consts.RabbitMQExchangeName,
		QueueName:    consts.RabbitMQQueueNameSystemMessage,
		RouteKey:     fmt.Sprintf("%s.%s", consts.RabbitMQExchangeName, consts.RabbitMQQueueNameSystemMessage),
	}

	// 创建通道
	if err := mqConn.Channel(); err != nil {
		SystemMessageLogger.Error(ctx, "创建队列通道失败", "error", err)
		return gerror.Wrap(err, "创建队列通道失败")
	}

	// 创建交换机
	if err := mqConn.Exchange("topic", true, false); err != nil {
		SystemMessageLogger.Error(ctx, "创建交换机失败", "error", err)
		return gerror.Wrap(err, "创建交换机失败")
	}

	// 创建队列
	if err := mqConn.Queue(true, false, nil); err != nil {
		SystemMessageLogger.Error(ctx, "创建队列失败", "error", err)
		return gerror.Wrap(err, "创建队列失败")
	}

	// 绑定队列
	if err := mqConn.Bind(); err != nil {
		SystemMessageLogger.Error(ctx, "绑定队列失败", "error", err)
		return gerror.Wrap(err, "绑定队列失败")
	}

	// 发布消息
	err = mqConn.Chan.Publish(
		consts.RabbitMQExchangeName,
		fmt.Sprintf("%s.%s", consts.RabbitMQExchangeName, consts.RabbitMQQueueNameSystemMessage),
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         messageBody,
			DeliveryMode: amqp.Persistent,
			Priority:     0,
			Timestamp:    time.Now(),
			MessageId:    data.TaskId,
		},
	)
	if err != nil {
		SystemMessageLogger.Error(ctx, "发布消息失败", "error", err, "taskId", data.TaskId)
		return gerror.Wrap(err, "发布消息失败")
	}

	SystemMessageLogger.Info(ctx, "系统消息已发布到队列", "taskId", data.TaskId, "memberId", data.MemberId, "scene", data.Scene)
	return nil
}
