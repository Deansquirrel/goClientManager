package webServer

import (
	"context"
	"github.com/Deansquirrel/goClientManager/global"
	log "github.com/Deansquirrel/goToolLog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"strconv"
	"time"
)

import wsRouter "github.com/Deansquirrel/goClientManager/webServer/router"

type webServer struct {
	app  *iris.Application
	port int
}

func NewWebServer(port int) *webServer {
	return &webServer{
		app:  iris.New(),
		port: port,
	}
}

//启动Web服务
func (ws *webServer) StartWebService() {
	log.Info("StartWebService")
	app := iris.New()
	ws.irisInit(app)

	iris.RegisterOnInterrupt(func() {
		log.Info("StopWebService")
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		_ = app.Shutdown(ctx)
	})

	app.Use(recover.New())
	app.Use(logger.New())

	wsRouter.AddWebPartyClientInfo(app)

	go func() {
		_ = app.Run(iris.Addr(":"+strconv.Itoa(ws.port)), iris.WithoutInterruptHandler)
	}()
}

//iris初始化
func (ws *webServer) irisInit(app *iris.Application) {
	ws.setIrisLogLevel(app)
	ws.setIrisLogTimeFormat(app)
	ws.setIrisLogFile(app)
}

//校验SysConfig中iris日志级别设置
func (ws *webServer) checkIrisLogLevel(level string) string {
	switch level {
	case "debug", "info", "warn", "error":
		return level
	default:
		return "warn"
	}
}

//设置iris日志级别
func (ws *webServer) setIrisLogLevel(app *iris.Application) {
	app.Logger().SetLevel(ws.checkIrisLogLevel(global.SysConfig.Iris.LogLevel))
}

//设置iris日志时间格式
func (ws *webServer) setIrisLogTimeFormat(app *iris.Application) {
	app.Logger().SetTimeFormat("2006-01-02 15:04:05")
}

//设置iris日志输出
func (ws *webServer) setIrisLogFile(app *iris.Application) {
	//TODO 设置日志输出文件对象,需可按日期自动拆分
}
