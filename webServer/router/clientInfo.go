package router

import (
	log "github.com/Deansquirrel/goToolLog"
	"github.com/kataras/iris"
	"io/ioutil"
)

func AddWebPartyClientInfo(app *iris.Application) {
	clientInfo := app.Party("/ClientInfo", clientInfoHandler)
	clientInfo.Post("/Info", clientInfoInfoHandler)
}

func clientInfoHandler(ctx iris.Context) {
	ctx.Next()
}

func clientInfoInfoHandler(ctx iris.Context) {
	body := ctx.Request().Body
	defer func() {
		_ = body.Close()
	}()
	b, err := ioutil.ReadAll(body)
	if err != nil {
		msg := "读取请求内容时遇到错误"
		log.Error(msg + err.Error())
		_, err = ctx.WriteString(GetErrReturn(msg))
		if err != nil {
			log.Error(err.Error())
		}
		return
	}
	log.Debug(string(b))
	_, err = ctx.WriteString(GetMsgReturn("Complete"))
	if err != nil {
		log.Error(err.Error())
	}
	return
}
