package router

import (
	"github.com/gin-gonic/gin"
	"github/wbellmelodyw/gin-wechat/handler"
)

func Create() *gin.Engine {
	g := gin.New()
	g.NoRoute(handler.NotFound())
	translates := g.Group("/translate")
	{
		translates.GET("chinese", handler.Chinese)
		translates.GET("english", handler.English)
		translates.GET("audio", handler.Audio)
	}
	wechat := g.Group("/wechat")
	{
		wechat.GET("create_menu", handler.CreateMenu)
		wechat.GET("auth", handler.WeChatAuth)
		wechat.POST("auth", handler.WeChatAuth)
	}
	return g
}
