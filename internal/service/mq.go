package service

import (
	"DeviceManagement/internal/dao"
	"DeviceManagement/internal/model"
	"context"
	"encoding/json"
	"log"

	"github.com/gogf/gf/v2/frame/g"
	mqsd "github.com/yyboo586/MQSDK"
)

type IMQ interface {
	Subscribe(ctx context.Context, topic string, handler mqsd.MessageHandler) error
}

var (
	localMQ *MQ
)

func MQService() IMQ {
	if localMQ == nil {
		panic("implement not found for interface IMQ, forgot register?")
	}
	return localMQ
}

const (
	DeviceOnlineTopic  = "core.device.online"
	DeviceOfflineTopic = "core.device.offline"
	DeviceAlarmTopic   = "core.device.alarm"
)

func RegisterMQService() {
	localMQ = NewMQ()
	localMQ.Subscribe(context.Background(), DeviceOnlineTopic, localMQ.handleDeviceOnline)
	localMQ.Subscribe(context.Background(), DeviceOfflineTopic, localMQ.handleDeviceOffline)
	localMQ.Subscribe(context.Background(), DeviceAlarmTopic, localMQ.handleDeviceAlarm)
}

type MQ struct {
	consumer mqsd.Consumer
}

func NewMQ() *MQ {
	nsqConfig := &mqsd.NSQConfig{
		Type:      "nsq",
		NSQDAddr:  "127.0.0.1:4151",
		NSQLookup: []string{"127.0.0.1:4161"},
	}
	factory := mqsd.NewFactory()
	consumer, err := factory.NewConsumer(nsqConfig)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	return &MQ{
		consumer: consumer,
	}
}

func (m *MQ) Subscribe(ctx context.Context, topic string, handler mqsd.MessageHandler) (err error) {
	err = m.consumer.Subscribe(ctx, topic, handler)
	if err != nil {
		g.Log().Error(ctx, "subscribe failed", err)
	}
	return
}

func (m *MQ) handleDeviceOnline(msg *mqsd.Message) (err error) {
	g.Log().Info(context.Background(), "device online", *msg)
	var body model.DeviceLog
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		g.Log().Error(context.Background(), "unmarshal device online message failed", err)
		return err
	}

	insertData := map[string]interface{}{
		dao.DeviceLog.Columns().OrgID:     body.OrgID,
		dao.DeviceLog.Columns().DeviceID:  body.DeviceID,
		dao.DeviceLog.Columns().DeviceKey: body.DeviceKey,
		dao.DeviceLog.Columns().Type:      model.DeviceLogTypeOnline,
		dao.DeviceLog.Columns().CreatedAt: msg.Timestamp,
	}
	_, err = dao.DeviceLog.Ctx(context.Background()).Data(insertData).Insert()
	if err != nil {
		g.Log().Error(context.Background(), "insert device log failed", err)
	}

	return nil
}

func (m *MQ) handleDeviceOffline(msg *mqsd.Message) (err error) {
	g.Log().Info(context.Background(), "device offline", *msg)
	var body model.DeviceLog
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		g.Log().Error(context.Background(), "unmarshal device offline message failed", err)
		return err
	}

	insertData := map[string]interface{}{
		dao.DeviceLog.Columns().OrgID:     body.OrgID,
		dao.DeviceLog.Columns().DeviceID:  body.DeviceID,
		dao.DeviceLog.Columns().DeviceKey: body.DeviceKey,
		dao.DeviceLog.Columns().Type:      model.DeviceLogTypeOffline,
		dao.DeviceLog.Columns().CreatedAt: msg.Timestamp,
	}
	_, err = dao.DeviceLog.Ctx(context.Background()).Data(insertData).Insert()
	if err != nil {
		g.Log().Error(context.Background(), "insert device log failed", err)
	}
	return nil
}

func (m *MQ) handleDeviceAlarm(msg *mqsd.Message) (err error) {
	g.Log().Info(context.Background(), "device alarm", *msg)
	var body model.DeviceLog
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		g.Log().Error(context.Background(), "unmarshal device alarm message failed", err)
		return err
	}

	content, err := json.Marshal(body.Content)
	if err != nil {
		g.Log().Error(context.Background(), "marshal content failed", err)
		return err
	}

	insertData := map[string]interface{}{
		dao.DeviceLog.Columns().OrgID:     body.OrgID,
		dao.DeviceLog.Columns().DeviceID:  body.DeviceID,
		dao.DeviceLog.Columns().DeviceKey: body.DeviceKey,
		dao.DeviceLog.Columns().Type:      model.DeviceLogTypeAlarm,
		dao.DeviceLog.Columns().Content:   string(content),
		dao.DeviceLog.Columns().CreatedAt: msg.Timestamp,
	}
	_, err = dao.DeviceLog.Ctx(context.Background()).Data(insertData).Insert()
	if err != nil {
		g.Log().Error(context.Background(), "insert device log failed", err)
	}

	if body.Content != nil {
		Device().InvokeAlarm(context.Background(), body.OrgID, body.Content.Message)
	}

	return nil
}
