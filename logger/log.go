package logger

import (
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"path"
	"time"
)

type logFields map[string]interface{}

type Log struct {
	Fields logFields
}

func Module(moduleName string) *Log {
	fields := logFields{
		"module_name": moduleName, //"gin_weChat",
	}
	return &Log{Fields: fields}
}

func CtxModule(ctx *gin.Context, moduleName string) *Log {
	fields := logFields{
		"app_id": ctx.Param("app_id"),
		//"method": 	   ctx.Request.Method,
		//"http_status": ctx.Writer.Status(),
		"module_name": moduleName, //"gin_weChat",
	}
	return &Log{Fields: fields}
}

func dailyLogWriter(logPath, level string, save uint) io.Writer {
	logFullPath := path.Join(logPath, level)
	hook, err := rotatelogs.New(
		logFullPath+".%Y%m%d.log", // 没有使用go风格反人类的format格式
		//rotatelogs.WithLinkName(logFullPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithRotationCount(save),        // 文件最大保存份数
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		panic(err)
	}
	return hook
}
