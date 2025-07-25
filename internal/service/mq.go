package service

import (
	"DeviceManagement/internal/dao"
	"DeviceManagement/internal/model"
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	mqsd "github.com/yyboo586/MQSDK"
)

type IMQ interface {
	Subscribe(ctx context.Context, topic string, channel string, handler mqsd.MessageHandler) error
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
	TopicDeviceOnline  = "core.device.online"
	TopicDeviceOffline = "core.device.offline"
	TopicDeviceAlarm   = "core.device.alarm"
)

func RegisterMQService() {
	localMQ = NewMQ()
	localMQ.Subscribe(context.Background(), TopicDeviceOnline, "DeviceManagement", localMQ.handleDeviceOnline)
	localMQ.Subscribe(context.Background(), TopicDeviceOffline, "DeviceManagement", localMQ.handleDeviceOffline)
	localMQ.Subscribe(context.Background(), TopicDeviceAlarm, "DeviceManagement", localMQ.handleDeviceAlarm)
}

type MQ struct {
	consumer mqsd.Consumer
}

func NewMQ() *MQ {
	nsqConfig := &mqsd.NSQConfig{
		Type:     "nsq",
		NSQDAddr: "124.220.236.38:4150",
		// 暂时不使用nsqlookupd，直接连接nsqd
		// NSQLookup: []string{"127.0.0.1:4160"},
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

func (m *MQ) Subscribe(ctx context.Context, topic string, channel string, handler mqsd.MessageHandler) (err error) {
	err = m.consumer.Subscribe(ctx, topic, channel, handler)
	if err != nil {
		g.Log().Error(ctx, "subscribe failed", err)
	}
	return
}

func (m *MQ) handleDeviceOnline(msg *mqsd.Message) (err error) {
	ctx := context.Background()
	g.Log().Info(ctx, "device online", *msg)
	body, ok := msg.Body.(map[string]interface{})
	if !ok {
		g.Log().Error(ctx, "device online message body is not map[string]interface{}")
		return nil
	}

	orgID, ok := body["org_id"].(string)
	if !ok {
		g.Log().Error(ctx, "device online message body org_id is not string")
		return nil
	}

	deviceID, ok := body["device_id"].(float64)
	if !ok {
		g.Log().Error(ctx, "device online message body device_id is not float64")
		return nil
	}

	deviceName, ok := body["device_name"].(string)
	if !ok {
		g.Log().Error(ctx, "device online message body device_name is not string")
		return nil
	}

	deviceKey, ok := body["device_key"].(string)
	if !ok {
		g.Log().Error(ctx, "device online message body device_key is not string")
		return nil
	}

	insertDeviceLogData := map[string]interface{}{
		dao.DeviceLog.Columns().OrgID:      orgID,
		dao.DeviceLog.Columns().DeviceID:   int64(deviceID),
		dao.DeviceLog.Columns().DeviceName: deviceName,
		dao.DeviceLog.Columns().DeviceKey:  deviceKey,
		dao.DeviceLog.Columns().Type:       model.DeviceLogTypeOnline,
		dao.DeviceLog.Columns().Timestamp:  msg.Timestamp,
		dao.DeviceLog.Columns().CreatedAt:  time.Now().Unix(),
	}
	updateDeviceData := map[string]interface{}{
		dao.Device.Columns().OnlineStatus:   model.DeviceOnlineStatusOnline,
		dao.Device.Columns().LastOnlineTime: time.Now(),
	}
	g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		_, err = dao.DeviceLog.Ctx(ctx).TX(tx).Data(insertDeviceLogData).Insert()
		if err != nil {
			g.Log().Error(ctx, "insert device log failed", err)
			return err
		}
		_, err = dao.Device.Ctx(ctx).TX(tx).Data(updateDeviceData).Where(dao.Device.Columns().ID, deviceID).Update()
		if err != nil {
			g.Log().Error(ctx, "update device failed", err)
			return err
		}
		return
	})
	return
}

func (m *MQ) handleDeviceOffline(msg *mqsd.Message) (err error) {
	ctx := context.Background()
	g.Log().Info(ctx, "device offline", *msg)
	body, ok := msg.Body.(map[string]interface{})
	if !ok {
		g.Log().Error(ctx, "device offline message body is not map[string]interface{}")
		return nil
	}

	orgID, ok := body["org_id"].(string)
	if !ok {
		g.Log().Error(ctx, "device offline message body org_id is not string")
		return nil
	}

	deviceID, ok := body["device_id"].(float64)
	if !ok {
		g.Log().Error(ctx, "device offline message body device_id is not float64")
		return nil
	}

	deviceName, ok := body["device_name"].(string)
	if !ok {
		g.Log().Error(ctx, "device offline message body device_name is not string")
		return nil
	}

	deviceKey, ok := body["device_key"].(string)
	if !ok {
		g.Log().Error(ctx, "device offline message body device_key is not string")
		return nil
	}

	insertDeviceLogData := map[string]interface{}{
		dao.DeviceLog.Columns().OrgID:      orgID,
		dao.DeviceLog.Columns().DeviceID:   int64(deviceID),
		dao.DeviceLog.Columns().DeviceName: deviceName,
		dao.DeviceLog.Columns().DeviceKey:  deviceKey,
		dao.DeviceLog.Columns().Type:       model.DeviceLogTypeOffline,
		dao.DeviceLog.Columns().Timestamp:  msg.Timestamp,
		dao.DeviceLog.Columns().CreatedAt:  time.Now().Unix(),
	}
	updateDeviceData := map[string]interface{}{
		dao.Device.Columns().OnlineStatus:    model.DeviceOnlineStatusOffline,
		dao.Device.Columns().LastOfflineTime: time.Now(),
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		_, err = dao.DeviceLog.Ctx(ctx).TX(tx).Data(insertDeviceLogData).Insert()
		if err != nil {
			g.Log().Error(ctx, "insert device log failed", err)
			return err
		}
		_, err = dao.Device.Ctx(ctx).TX(tx).Data(updateDeviceData).Where(dao.Device.Columns().ID, deviceID).Update()
		if err != nil {
			g.Log().Error(ctx, "update device failed", err)
			return err
		}
		return
	})
	return
}

func (m *MQ) handleDeviceAlarm(msg *mqsd.Message) (err error) {
	g.Log().Info(context.Background(), "device alarm", *msg)

	var body map[string]interface{}
	body, ok := msg.Body.(map[string]interface{})
	if !ok {
		g.Log().Error(context.Background(), "device alarm message body is not map[string]interface{}")
		return errors.New("device alarm message body is not map[string]interface{}")
	}

	orgID, ok := body["org_id"].(string)
	if !ok {
		g.Log().Error(context.Background(), "device alarm message body org_id is not string")
		return errors.New("device alarm message body org_id is not string")
	}

	deviceID, ok := body["device_id"].(float64)
	if !ok {
		g.Log().Error(context.Background(), "device alarm message body device_id is not int64")
		return errors.New("device alarm message body device_id is not int64")
	}

	deviceName, ok := body["device_name"].(string)
	if !ok {
		g.Log().Error(context.Background(), "device alarm message body device_name is not string")
		return errors.New("device alarm message body device_name is not string")
	}

	deviceKey, ok := body["device_key"].(string)
	if !ok {
		g.Log().Error(context.Background(), "device alarm message body device_key is not string")
		return errors.New("device alarm message body device_key is not string")
	}

	content, ok := body["content"].(map[string]interface{})
	if !ok {
		g.Log().Error(context.Background(), "device alarm message body content is not map[string]interface{}")
		return errors.New("device alarm message body content is not map[string]interface{}")
	}

	contentJSON, err := json.Marshal(content)
	if err != nil {
		g.Log().Error(context.Background(), "marshal content failed", err)
		return err
	}

	insertData := map[string]interface{}{
		dao.DeviceLog.Columns().OrgID:      orgID,
		dao.DeviceLog.Columns().DeviceID:   int64(deviceID),
		dao.DeviceLog.Columns().DeviceName: deviceName,
		dao.DeviceLog.Columns().DeviceKey:  deviceKey,
		dao.DeviceLog.Columns().Type:       model.DeviceLogTypeAlarm,
		dao.DeviceLog.Columns().Content:    string(contentJSON),
		dao.DeviceLog.Columns().Timestamp:  msg.Timestamp,
		dao.DeviceLog.Columns().CreatedAt:  time.Now().Unix(),
	}
	_, err = dao.DeviceLog.Ctx(context.Background()).Data(insertData).Insert()
	if err != nil {
		g.Log().Error(context.Background(), "insert device log failed", err)
	}

	message, ok := content["message"].(string)
	if !ok {
		g.Log().Error(context.Background(), "device alarm message body content message is not string")
		return errors.New("device alarm message body content message is not string")
	}

	Device().InvokeAlarm(context.Background(), orgID, message)

	return nil
}
