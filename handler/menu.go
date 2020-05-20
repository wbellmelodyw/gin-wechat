package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/menu"
	"github/wbellmelodyw/gin-wechat/cache"
	myconfig "github/wbellmelodyw/gin-wechat/config"
	"github/wbellmelodyw/gin-wechat/logger"
	"github/wbellmelodyw/gin-wechat/model"
	"net/http"
)

func CreateMenu(ctx *gin.Context) {
	//配置微信参数
	config := &wechat.Config{
		AppID:          myconfig.GetString("APP_ID"),
		AppSecret:      myconfig.GetString("APP_SECRET"),
		Token:          myconfig.GetString("TOKEN"),
		EncodingAESKey: myconfig.GetString("ENCODING_AES_KEY"),
		Cache:          cache.NewCache(),
	}
	logger.Module("wechat").Sugar().Info("config info", config)
	wc := wechat.NewWechat(config)

	// 传入request和responseWriter
	server := wc.GetServer(ctx.Request, ctx.Writer)
	buttons := make([]*menu.Button, 3)
	buttons[0].SetClickButton("词性", "attr")
	buttons[1].SetClickButton("解释", "explain")
	buttons[2].SetClickButton("造句", "example")
	menus := menu.NewMenu(server.Context)
	err := menus.SetMenu(buttons)
	logger.Module("wechat").Sugar().Info("set menu err", err)
	ctx.JSON(http.StatusOK, &model.ApiResult{
		Code: 200,
		Msg:  "success",
		Data: "",
	})
}
