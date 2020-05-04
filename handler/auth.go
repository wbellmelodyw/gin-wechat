package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
	myconfig "github/wbellmelodyw/gin-wechat/config"
	"github/wbellmelodyw/gin-wechat/logger"
	"github/wbellmelodyw/gin-wechat/translate"
	"github/wbellmelodyw/gin-wechat/utils"
)

func WeChatAuth(ctx *gin.Context) {
	//logger.Module("wechat").Sugar().Error("serve error", "come")
	//配置微信参数
	config := &wechat.Config{
		AppID:          myconfig.GetString("APP_ID"),
		AppSecret:      myconfig.GetString("APP_SECRET"),
		Token:          myconfig.GetString("TOKEN"),
		EncodingAESKey: myconfig.GetString("ENCODING_AES_KEY"),
	}
	logger.Module("wechat").Sugar().Info("serve error", config)
	wc := wechat.NewWechat(config)

	// 传入request和responseWriter
	server := wc.GetServer(ctx.Request, ctx.Writer)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {

		//回复消息：演示回复用户发送的消息
		form, to := utils.GetLanguageTag(msg.Content)
		translator := translate.GetGoogle(form, to)
		t, err := translator.Text(msg.Content)
		if t == nil || err != nil {
			logger.Module("wechat").Sugar().Error("serve error", err)
		}
		text := message.NewText(t.Mean)
		//发送其他的给他
		openId := server.GetOpenID()
		c := message.NewMessageManager(wc.Context)
		for a, attr := range t.Attr {
			for _, aa := range attr {
				err := c.Send(message.NewCustomerTextMessage(openId, a+":"+aa))
				logger.Module("wechat").Sugar().Error("message error", err)
			}
		}
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		logger.Module("wechat").Sugar().Error("serve error", err)
		return
	}
	//发送回复的消息
	server.Send()
}
