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
	text := ctx.Query("text")
	logger.Module("translate").Sugar().Error("english translate fail", text)
	//初始化
	googleTranslator := translate.GetGoogle(language.English, language.Chinese)
	value, err := googleTranslator.Text(text)
	if err != nil {
		logger.Module("translate").Sugar().Error("english translate fail", err)
	}
	logger.Module("translate").Sugar().Error("english translate fail", value)
	ctx.JSON(http.StatusOK, &model.ApiResult{
		Code: 200,
		Msg:  "success",
		Data: value,
	})
}

func English(ctx *gin.Context) {
	text := ctx.Param("text")
	googleTranslator := translate.GetGoogle(language.Chinese, language.English)
	value, err := googleTranslator.Text(text)
	if err != nil {
		logger.Module("translate").Sugar().Error("english translate fail", err)
	}
	ctx.JSON(http.StatusOK, &model.ApiResult{
		Code: 200,
		Msg:  "success",
		Data: value,
	})
}

func Audio(ctx *gin.Context) {
	text := ctx.Query("text")
	text1 := ctx.Param("text")
	logger.Module("translate").Sugar().Error("english translate fail", text+text1)
	googleTranslator := translate.GetGoogle(language.English, language.English)
	value, err := googleTranslator.Audio(text)
	if err != nil {
		logger.Module("translate").Sugar().Error("english translate fail", err)
	}
	ctx.Header("Content-Type", "audio/mpeg")
	ctx.Header("transfer-encoding", "identity")
	ctx.Writer.Write(value)
}
