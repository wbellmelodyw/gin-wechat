package handler

import (
	"github.com/gin-gonic/gin"
	"github/wbellmelodyw/gin-wechat/logger"
	"github/wbellmelodyw/gin-wechat/model"
	"github/wbellmelodyw/gin-wechat/translate"
	"golang.org/x/text/language"
	"net/http"
)

func Chinese(ctx *gin.Context) {
	text := ctx.Param("text")
	value, err := translate.Text(text, language.English, language.Chinese)
	if err != nil {
		logger.Module("translate").Sugar().Error("english translate fail", err)
	}
	ctx.JSON(http.StatusOK, &model.ApiResult{
		Code: 200,
		Msg:  "success",
		Data: value,
	})
}

func English(ctx *gin.Context) {
	text := ctx.Param("text")
	value, err := translate.Text(text, language.Chinese, language.English)
	if err != nil {
		logger.Module("translate").Sugar().Error("english translate fail", err)
	}
	ctx.JSON(http.StatusOK, &model.ApiResult{
		Code: 200,
		Msg:  "success",
		Data: value,
	})
}
