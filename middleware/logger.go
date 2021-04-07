package middleware

import (
	"github.com/kataras/iris/v12/middleware/accesslog"
)

func Logger() *accesslog.AccessLog {
	filePath := "access.log"
	
	ac := accesslog.File(filePath)
	
	ac.Flush()
	// 输出到控制台
	//ac.SetOutput(ac.Writer)
	
	// The default configuration:
	ac.Delim = '|'
	ac.TimeFormat = "2006-01-02 15:04:05"
	ac.Async = true
	ac.IP = true
	ac.BytesReceivedBody = true
	ac.BytesSentBody = true
	ac.BytesReceived = false
	ac.BytesSent = false
	ac.BodyMinify = true
	ac.RequestBody = true
	ac.ResponseBody = true
	ac.KeepMultiLineError = true
	ac.PanicLog = accesslog.LogHandler
	
	// 日志格式化成SVC
	//ac.SetFormatter(&accesslog.CSV{})
	
	// 日志格式化成json
	//ac.SetFormatter(&accesslog.JSON{
	//	Indent:    "  ",
	//	HumanTime: true,
	//})
	
	return ac
}
