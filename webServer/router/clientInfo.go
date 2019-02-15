package router

import (
	"encoding/json"
	"github.com/Deansquirrel/goClientManager/object"
	log "github.com/Deansquirrel/goToolLog"
	"github.com/kataras/iris"
)

func AddWebPartyClientInfo(app *iris.Application) {
	clientInfo := app.Party("/ClientInfo", clientInfoHandler)
	clientInfo.Post("/Info", clientInfoInfoHandler)
}

func clientInfoHandler(ctx iris.Context) {
	ctx.Next()
}

func clientInfoInfoHandler(ctx iris.Context) {
	var info object.ClientInfo
	err := ctx.ReadJSON(&info)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.WriteString(GetErrReturn(err.Error()))
		log.Warn(err.Error())
		return
	}

	b, err := json.Marshal(info)
	if err != nil {
		log.Info(err.Error())
	} else {
		log.Info(string(b))
	}
	_, _ = ctx.WriteString(GetMsgReturn("OK"))
}
