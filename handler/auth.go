package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/material"
	"github.com/silenceper/wechat/message"
	"github/wbellmelodyw/gin-wechat/cache"
	myconfig "github/wbellmelodyw/gin-wechat/config"
	"github/wbellmelodyw/gin-wechat/db"
	"github/wbellmelodyw/gin-wechat/logger"
	"github/wbellmelodyw/gin-wechat/model"
	"github/wbellmelodyw/gin-wechat/translate"
	"github/wbellmelodyw/gin-wechat/utils"
	"golang.org/x/text/language"
	"strings"
)

const ATTR = "1"
const EXPLAIN = "2"
const EXAMPLE = "3"
const AUDIO = "4"
const LAST_WORD_KEY = "last:word"

func WeChatAuth(ctx *gin.Context) {
	//logger.Module("wechat").Sugar().Error("serve error", "come")
	//配置微信参数
	weCache := cache.NewCache()
	config := &wechat.Config{
		AppID:          myconfig.GetString("APP_ID"),
		AppSecret:      myconfig.GetString("APP_SECRET"),
		Token:          myconfig.GetString("TOKEN"),
		EncodingAESKey: myconfig.GetString("ENCODING_AES_KEY"),
		Cache:          weCache,
	}
	logger.Module("wechat").Sugar().Info("config info", config)
	wc := wechat.NewWechat(config)

	// 传入request和responseWriter
	server := wc.GetServer(ctx.Request, ctx.Writer)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		//没办法自定义菜单,用序号定义
		switch msg.Content {
		case ATTR:
			return getAttr(weCache.Get(LAST_WORD_KEY))
		case EXPLAIN:
			return getExplain(weCache.Get(LAST_WORD_KEY))
		case EXAMPLE:
			return getExample(weCache.Get(LAST_WORD_KEY))
		case AUDIO:
			return getAudio(weCache.Get(LAST_WORD_KEY), wc)
		}
		//回复消息
		//先从数据库查,找不到再去调google
		w := model.Word{
			SrcContent: msg.Content,
		}

		ok, err := db.WeChat.Get(&w)
		if err != nil {
			logger.Module("db").Sugar().Panic("insert error", err)
		}
		if ok {
			text := message.NewText(w.DstContent)
			err = weCache.Set(LAST_WORD_KEY, w.SrcContent, -1)
			logger.Module("wechat").Sugar().Info("redis set error", err)
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
		}

		form, to := utils.GetLanguageTag(msg.Content)
		translator := translate.GetGoogle(form, to)
		t, err := translator.Text(msg.Content)

		err = weCache.Set(LAST_WORD_KEY, msg.Content, -1)
		logger.Module("wechat").Sugar().Info("redis set error", err)

		if t == nil || err != nil {
			logger.Module("wechat").Sugar().Error("serve error", err)
		}
		//异步存入sql
		tChan := make(chan *model.Text)
		go insert(tChan)
		tChan <- t

		//异步获取音频文件,中文大家都会，只获取英语读音
		audioText := make(chan string)
		go fetchAudio(audioText)
		if form == language.English {
			audioText <- msg.Content
		} else {
			audioText <- t.Mean
		}
		//发送其他的给他
		//openId := server.GetOpenID()
		//c := message.NewMessageManager(wc.Context)
		//for a, attr := range t.Attr {
		//	for _, aa := range attr {
		//		err := c.Send(message.NewCustomerTextMessage(openId, a+":"+aa))
		//		logger.Module("wechat").Sugar().Error("message error", err)
		//	}
		//}
		text := message.NewText(t.Mean)
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

func getAttr(content interface{}) *message.Reply {
	c := content.(string)
	w := model.Word{
		SrcContent: c,
	}

	ok, err := db.WeChat.Get(&w)
	if err != nil {
		logger.Module("db").Sugar().Panic("insert error", err)
	}
	if ok {
		text := message.NewText(w.DstAttr)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	}
	return &message.Reply{MsgType: message.MsgTypeText, MsgData: message.NewText("找不到")}
}

func getExplain(content interface{}) *message.Reply {
	c := content.(string)
	w := model.Word{
		SrcContent: c,
	}

	ok, err := db.WeChat.Get(&w)
	logger.Module("db").Sugar().Info("insert error", ok)
	logger.Module("db").Sugar().Info("insert error", w)
	if err != nil {
		logger.Module("db").Sugar().Panic("insert error", err)
	}
	if ok {
		text := message.NewText(w.DstExplain)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	}
	return &message.Reply{MsgType: message.MsgTypeText, MsgData: message.NewText("找不到")}
}

func getExample(content interface{}) *message.Reply {
	c := content.(string)
	w := model.Word{
		SrcContent: c,
	}

	ok, err := db.WeChat.Get(&w)
	if err != nil {
		logger.Module("db").Sugar().Panic("db error", err)
	}
	if ok {
		//text := message.NewText(w.DstExample)
		text := message.NewText(w.DstExample[0:520])
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	}
	return &message.Reply{MsgType: message.MsgTypeText, MsgData: message.NewText("找不到")}
}

func getAudio(content interface{}, wc *wechat.Wechat) *message.Reply {
	//检查是不是超过了最大限制

	c := content.(string)
	w := model.Word{
		SrcContent: c,
	}
	ok, err := db.WeChat.Get(&w)
	if err != nil {
		logger.Module("db").Sugar().Panic("db error", err)
	}
	if ok {
		uploadText := c
		if utils.IsHan(c) {
			uploadText = w.DstContent
		}
		if w.MediaId != "" {
			text := message.NewVoice(w.MediaId) //"5zHKjknSbOcaVTCYQKNJXEto6jr36ceXPKzTboLl0F8"
			return &message.Reply{MsgType: message.MsgTypeVoice, MsgData: text}
		}
		//写完开始上传
		m := material.NewMaterial(wc.Context)
		mid, urll, err := m.AddMaterial(material.MediaTypeVoice, "media/"+uploadText+".mp3")
		//m.DeleteMaterial("5zHKjknSbOcaVTCYQKNJXDgAUMO2SlL945idx28hYso")
		logger.Module("audio").Sugar().Info("material upload", mid)
		logger.Module("audio").Sugar().Info("material upload", urll)
		logger.Module("audio").Sugar().Info("material err", err)
		//上传完更新mediaId
		var newW model.Word
		newW.MediaId = mid
		row, err := db.WeChat.Id(w.Id).Update(&newW)
		logger.Module("audio").Sugar().Info("material update row", row)
		if err != nil {
			logger.Module("audio").Sugar().Panic("material update err", err)
		}
		//text := message.NewText(w.DstExample)
		//text := message.NewVoice(w.MediaId)
		text := message.NewVoice(mid) //"5zHKjknSbOcaVTCYQKNJXEto6jr36ceXPKzTboLl0F8"
		return &message.Reply{MsgType: message.MsgTypeVoice, MsgData: text}
	}
	return &message.Reply{MsgType: message.MsgTypeText, MsgData: message.NewText("找不到")}
}

func insert(textChan <-chan *model.Text) {
	text := <-textChan
	word := new(model.Word)
	word.SrcContent = text.Content
	word.DstContent = text.Mean
	attrs := make([]string, 2)
	for name, attr := range text.Attr {
		attrs = append(attrs, name+":"+strings.Join(attr, ";"))
	}
	word.DstAttr = strings.Join(attrs, "\n")
	explain := make([]string, 2)
	for name, e := range text.Explain {
		explain = append(explain, name+strings.Join(e, ";"))
	}
	word.DstExplain = strings.Join(explain, "\n")
	word.DstExample = strings.Join(text.Example, "\n")
	row, err := db.WeChat.Insert(word)
	if err != nil {
		logger.Module("db").Sugar().Panic("insert error", err)
	}
	logger.Module("db").Sugar().Info("insert row", row)
}

//异步提取音频
func fetchAudio(text chan string) {
	t := <-text
	//isDone := make(chan int)
	googleTranslator := translate.GetGoogle(language.English, language.English)
	googleTranslator.AudioSaveFile(t)

}
