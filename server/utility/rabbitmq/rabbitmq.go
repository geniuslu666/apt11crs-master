package rabbitmq

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gmlock"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	Conn *amqp.Connection
)

type MQConnection struct {
	Conn         *amqp.Connection // RabbitMQ 连接
	Chan         *amqp.Channel    // RabbitMQ 通道
	ExchangeName string           // 交换机名
	QueueName    string           // 队列名
	RouteKey     string           // 路由名
}

func init() {
	var (
		err       error
		MQLink    g.MapStrStr
		MQLinkVar *gvar.Var
	)
	if MQLinkVar, err = g.Cfg().Get(context.TODO(), "rabbitmq"); err != nil {
		panic(err)
	}
	MQLink = MQLinkVar.MapStrStr()
	MQLinkUrl := fmt.Sprintf("amqp://%s:%s@%s:%s%s", MQLink["user"], MQLink["pwd"], MQLink["link"], MQLink["port"], MQLink["vhost"])
	if Conn, err = amqp.Dial(MQLinkUrl); err != nil {
		return
	}
}

func (ChanApi *MQConnection) ReGetConn() (err error) {
	var (
		MQLink    g.MapStrStr
		MQLinkVar *gvar.Var
	)
	if !ChanApi.Conn.IsClosed() {
		return
	}
	if MQLinkVar, err = g.Cfg().Get(context.TODO(), "rabbitmq"); err != nil {
		panic(err)
	}
	MQLink = MQLinkVar.MapStrStr()
	MQLinkUrl := fmt.Sprintf("amqp://%s:%s@%s:%s%s", MQLink["user"], MQLink["pwd"], MQLink["link"], MQLink["port"], MQLink["vhost"])
	gmlock.Lock("rabbitmq_conn")
	defer gmlock.Unlock("rabbitmq_conn")
	if ChanApi.Conn, err = amqp.Dial(MQLinkUrl); err != nil {
		return
	}
	Conn = ChanApi.Conn
	return
}

func (ChanApi *MQConnection) Channel() (err error) {
	var (
		Chan *amqp.Channel
	)
	if g.IsEmpty(ChanApi.Conn) || ChanApi.Conn.IsClosed() {
		if err = ChanApi.ReGetConn(); err != nil {
			return err
		}
	}
	if Chan, err = ChanApi.Conn.Channel(); err != nil {
		return err
	}
	ChanApi.Chan = Chan
	return
}

// Exchange
/*
描述
	创建一个exchange交换机
参数
	kind 			【string】	交换机类型
	durable			【bool】		是否持久化
	autoDeleted     【bool】		是否自动删除
返回
	err 			【error】	错误信息
*/
func (ChanApi *MQConnection) Exchange(kind string, durable bool, autoDeleted bool) error {
	err := ChanApi.Chan.ExchangeDeclare(
		ChanApi.ExchangeName, //交换机名
		kind,                 //类型
		durable,              // is durable
		autoDeleted,          // is auto-deleted
		false,                // is internal
		false,                // is noWait
		amqp.Table{
			"x-delayed-type": "direct",
		}, //arguments
	)
	return err
}

// Queue
/*
描述
	创建一个queue通道
参数
	durable			【bool】		是否持久化
	autoDeleted     【bool】		是否自动删除
返回
	err 			【error】	错误信息
*/
func (ChanApi *MQConnection) Queue(durable bool, autoDeleted bool, args amqp.Table) (err error) {
	if _, err = ChanApi.Chan.QueueDeclare(
		ChanApi.QueueName, // name
		durable,           // durable
		autoDeleted,       // delete when unused
		false,             // exclusive
		false,             // no-wait
		args,              // arguments
	); err != nil {
		return
	}
	return err
}

// Bind
/*
描述
	交换机绑定队列
参数
	durable			【bool】		是否持久化
	autoDeleted     【bool】		是否自动删除
返回
	err 			【error】	错误信息
*/
func (ChanApi *MQConnection) Bind() error {
	err := ChanApi.Chan.QueueBind(ChanApi.QueueName, ChanApi.RouteKey, ChanApi.ExchangeName, false, nil)
	return err
}

// Publish
/*
描述
	生产一条消息
参数
	message		【byte】		是否自动删除
返回
	err			【error】	错误信息
*/
func (ChanApi *MQConnection) Publish(message []byte, header amqp.Table, CorrelationId string) error {
	err := ChanApi.Chan.Publish(
		ChanApi.ExchangeName, // exchange
		ChanApi.RouteKey,     // routing key
		false,                // mandatory
		false,                // immediate
		amqp.Publishing{
			Headers:         header,
			ContentType:     "text/html",
			ContentEncoding: "",
			CorrelationId:   CorrelationId,
			Body:            message,
		},
	)
	return err
}

// DiyPublish
/*
描述
	自定义生产一条消息
参数
	message		【byte】		是否自动删除
返回
	err			【error】	错误信息
*/
func (ChanApi *MQConnection) DiyPublish(Publish amqp.Publishing) error {
	err := ChanApi.Chan.Publish(
		ChanApi.ExchangeName, // exchange
		ChanApi.RouteKey,     // routing key
		false,                // mandatory
		false,                // immediate
		Publish,
	)
	return err
}

// Consume
/*
描述
	创建消费通道来监听消费消息
参数
	consumerName		【string】		消费通道名 用于区分不同的消费者
	autoAck				【bool】			是否自动回复    true 会在收到消息后直接消费掉这条消息  false 收到消息过后需要手动消费这条消息 不然的话将回到队列中去
返回
	err			【error】	错误信息
*/
func (ChanApi *MQConnection) Consume(consumerName string, autoAck bool) (<-chan amqp.Delivery, error) {
	messageChannel, err := ChanApi.Chan.Consume(
		ChanApi.QueueName,
		consumerName,
		autoAck, // auto-ack 是否自动回复，【自动回复表示该条队列消息已经被处理过了，消费掉了】
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // args
	)
	return messageChannel, err
}
