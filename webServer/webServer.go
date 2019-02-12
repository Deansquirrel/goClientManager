package webServer

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"strconv"
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

func (ws *webServer) StartWebService() {
	app := iris.New()
	//app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())

	wsRouter.AddWebPartyClientInfo(app)

	_ = app.Run(iris.Addr(":" + strconv.Itoa(ws.port)))
}
