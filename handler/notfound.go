package handler

import (
	"github.com/gin-gonic/gin"
	"github/wbellmelodyw/gin-wechat/model"
	"net/http"
)

func NotFound() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusNotFound, &model.ApiResult{
			Code: http.StatusNotFound,
			Msg:  "page not found!test",
			Data: nil,
		})
	}
}
