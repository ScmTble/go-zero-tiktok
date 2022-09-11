package listen

import (
	"fmt"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/queue"
	"log"
	"tiktok/like/mq/internal/svc"
)

type (
	ConsumeHandler func(svcCtx *svc.ServiceContext, message string) error

	RabbitListener struct {
		svcCtx  *svc.ServiceContext
		conn    *amqp.Connection
		channel *amqp.Channel
		forever chan bool
		handler ConsumeHandler
	}
)

func (q RabbitListener) Start() {
	q.ExchangeDeclare()
	q.DeclareQueueAndBind()

	for _, que := range q.svcCtx.Config.Mq.ListenerQueues {
		msg, err := q.channel.Consume(
			que.Name,
			"",
			que.ConsumerConf.AutoAck,
			que.ConsumerConf.Exclusive,
			que.ConsumerConf.NoLocal,
			que.ConsumerConf.NoWait,
			nil,
		)
		if err != nil {
			log.Fatalf("failed to listener, error: %v", err)
		}

		go func() {
			for d := range msg {
				if err := q.handler(q.svcCtx, string(d.Body)); err != nil {
					logx.Errorf("Error on consuming: %s, error: %v", string(d.Body), err)
				}
			}
		}()
	}
	fmt.Println(q.svcCtx.Config.Name + " listening...")
	<-q.forever
}

func (q RabbitListener) Stop() {
	q.channel.Close()
	q.conn.Close()
	close(q.forever)
}

// MustNewListener 创建RabbitMQ监听者
func MustNewListener(svcCtx *svc.ServiceContext, handler ConsumeHandler) queue.MessageQueue {
	listener := RabbitListener{
		svcCtx:  svcCtx,
		handler: handler,
		forever: make(chan bool),
	}
	conn, err := amqp.Dial(svcCtx.Config.Mq.GetUrl())
	if err != nil {
		log.Fatalf("failed to connect rabbitmq, error: %v", err)
	}

	listener.conn = conn
	channel, err := listener.conn.Channel()
	if err != nil {
		log.Fatalf("failed to open a channel: %v", err)
	}

	listener.channel = channel
	return listener
}

// ExchangeDeclare 声明交换机
func (q RabbitListener) ExchangeDeclare() error {
	exchange := q.svcCtx.Config.Mq.Exchange
	return q.channel.ExchangeDeclare(
		exchange.Name,
		exchange.Type,
		exchange.Durable,
		exchange.AutoDelete,
		exchange.Internal,
		exchange.NoWait,
		nil,
	)
}

func (q RabbitListener) DeclareQueueAndBind() error {
	if len(q.svcCtx.Config.Mq.ListenerQueues) == 0 {
		return nil
	}
	var err error
	for _, conf := range q.svcCtx.Config.Mq.ListenerQueues {
		_, err = q.channel.QueueDeclare(
			conf.Name,
			conf.Durable,
			conf.AutoDelete,
			conf.Exclusive,
			conf.NoWait,
			nil,
		)
		err = q.bind(conf.Name, conf.BindConf.RouterKey, conf.BindConf.Exchange, conf.BindConf.NotWait)
	}
	return err
}

func (q RabbitListener) bind(queueName string, RouterKey string, exchange string, notWait bool) error {
	return q.channel.QueueBind(
		queueName,
		RouterKey,
		exchange,
		notWait,
		nil,
	)
}
