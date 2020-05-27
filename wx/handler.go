package wx

import (
	"encoding/xml"
	"errors"
	"github.com/songouku/gwx/constant"
	"github.com/songouku/gwx/model"
)

type IMessageHandler interface {
	TxtHandler() interface{}
	ImageHandler() interface{}
	VoiceHandler() interface{}
	VideoHandler() interface{}
	LocationHandler() interface{}
	LinkHandler() interface{}
	SubscribeHandler() interface{}
	LocationEventHandler() interface{}
	ClickHandler() interface{}
	ViewHandler() interface{}
}

func (c *Config) HandlerMessage(content string) error {
	msg := model.Message{}
	err := xml.Unmarshal([]byte(content), &msg)
	if err != nil {
		return err
	}
	c.SetCurrentMsg(msg.MsgType)
	if msg.MsgType == constant.EventMsg.Type {
		c.SetCurrentMsg(msg.Event)
		//事件消息
		if msg.Event == constant.ClickEvent.Type {
			//点击事件
			c.SetClickEventMessage(model.ClickEventMessage{
				ToUserName:   msg.ToUserName,
				FromUserName: msg.FromUserName,
				CreateTime:   msg.CreateTime,
				MsgType:      msg.MsgType,
				Event:        msg.Event,
				EventKey:     msg.EventKey,
			})
			return nil
		} else if msg.Event == constant.Subscribe.Type {
			//关注/取消关注事件
			c.SetSubscribeEventMessage(model.SubscribeEventMessage{
				ToUserName:   msg.ToUserName,
				FromUserName: msg.FromUserName,
				CreateTime:   msg.CreateTime,
				MsgType:      msg.MsgType,
				Event:        msg.Event,
				EventKey:     msg.EventKey,
			})
			return nil
		} else if msg.Event == constant.LocationEvent.Type {
			//上报位置事件
			c.SetLocationEventMessage(model.LocationEventMessage{
				ToUserName:   msg.ToUserName,
				FromUserName: msg.FromUserName,
				CreateTime:   msg.CreateTime,
				MsgType:      msg.MsgType,
				Event:        msg.Event,
				Latitude:     msg.LocationX,
				Longitude:    msg.LocationY,
			})
			return nil
		} else if msg.Event == constant.ViewEvent.Type {
			//视图事件
			c.SetViewEventMessage(model.ViewEventMessage{
				ToUserName:   msg.ToUserName,
				FromUserName: msg.FromUserName,
				CreateTime:   msg.CreateTime,
				MsgType:      msg.MsgType,
				Event:        msg.Event,
				EventKey:     msg.EventKey,
			})
			return nil
		}
	} else if msg.MsgType == constant.TextMsg.Type {
		//文本消息
		c.SetTxtMessage(model.TxtMessage{
			ToUserName:   msg.ToUserName,
			FromUserName: msg.FromUserName,
			CreateTime:   msg.CreateTime,
			MsgType:      msg.MsgType,
			Content:      msg.Content,
			MsgId:        msg.MsgId,
		})
		return nil
	} else if msg.MsgType == constant.ImageMsg.Type {
		//图片
		c.SetImageMessage(model.ImageMessage{
			ToUserName:   msg.ToUserName,
			FromUserName: msg.FromUserName,
			CreateTime:   msg.CreateTime,
			MsgType:      msg.MsgType,
			PicUrl:       msg.PicUrl,
			MediaId:      msg.MediaId,
			MsgId:        msg.MsgId,
		})
		return nil
	} else if msg.MsgType == constant.VoiceMsg.Type {
		//语音
		c.SetVoiceMessage(model.VoiceMessage{
			ToUserName:   msg.ToUserName,
			FromUserName: msg.ToUserName,
			CreateTime:   msg.CreateTime,
			MsgType:      msg.MsgType,
			MediaId:      msg.MediaId,
			MsgId:        msg.MsgId,
			Format:       msg.Format,
		})
		return nil
	} else if msg.MsgType == constant.VideoMsg.Type {
		//视频
		c.SetVideoMessage(model.VideoMessage{
			ToUserName:   msg.ToUserName,
			FromUserName: msg.FromUserName,
			CreateTime:   msg.CreateTime,
			MsgType:      msg.MsgType,
			MediaId:      msg.MediaId,
			ThumbMediaId: msg.ThumbMediaId,
			MsgId:        msg.MsgId,
		})
		return nil
	} else if msg.MsgType == constant.ShortVideoMsg.Type {
		//小视频
	} else if msg.MsgType == constant.LocationMsg.Type {
		//位置消息
		c.SetLocationMessage(model.LocationMessage{
			ToUserName:   msg.ToUserName,
			FromUserName: msg.FromUserName,
			CreateTime:   msg.CreateTime,
			MsgType:      msg.MsgType,
			LocationX:    msg.LocationX,
			LocationY:    msg.LocationY,
			Scale:        msg.Scale,
			Label:        msg.Label,
			MsgId:        msg.MsgId,
		})
		return nil
	} else if msg.MsgType == constant.LinkMsg.Type {
		//链接消息
		c.SetLinkMessage(model.LinkMessage{
			ToUserName:   msg.ToUserName,
			FromUserName: msg.FromUserName,
			CreateTime:   msg.CreateTime,
			MsgType:      msg.MsgType,
			Title:        msg.Title,
			Description:  msg.Description,
			Url:          msg.Url,
			MsgId:        msg.MsgId,
		})
		return nil
	}
	return errors.New("unknown message type")
}

func (c *Config) MessageHandler(messageHandler IMessageHandler) interface{} {
	if c.GetCurrentMsg() == constant.ClickEvent.Type {
		return messageHandler.ClickHandler()
	} else if c.GetCurrentMsg() == constant.ViewEvent.Type {
		return messageHandler.ViewHandler()
	} else if c.GetCurrentMsg() == constant.Subscribe.Type {
		return messageHandler.SubscribeHandler()
	} else if c.GetCurrentMsg() == constant.LocationEvent.Type {
		return messageHandler.LocationEventHandler()
	} else if c.GetCurrentMsg() == constant.TextMsg.Type {
		return messageHandler.TxtHandler()
	} else if c.GetCurrentMsg() == constant.ImageMsg.Type {
		return messageHandler.ImageHandler()
	} else if c.GetCurrentMsg() == constant.VoiceMsg.Type {
		return messageHandler.VoiceHandler()
	} else if c.GetCurrentMsg() == constant.VideoMsg.Type {
		return messageHandler.VideoHandler()
	} else if c.GetCurrentMsg() == constant.LocationMsg.Type {
		return messageHandler.LocationHandler()
	} else if c.GetCurrentMsg() == constant.LinkMsg.Type {
		return messageHandler.LinkHandler()
	}
	return nil
}
