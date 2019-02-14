package router

import (
	log "github.com/Deansquirrel/goToolLog"
	"github.com/kataras/iris"
)

type clientInfo struct {
	OsInfo  osInfo  `json:"os"`
	NetInfo netInfo `json:"net"`
	DbInfo  dbInfo  `json:"db"`
	ErpInfo erpInfo `json:"erp"`
}

type osInfo struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Ver      string `json:"ver"`
	HostName string `json:"hostname"`
}

type netInfo struct {
	InternetIp string `json:"internet"`
	IntranetIp string `json:"intranet"`
}

type dbInfo struct {
	Ver string `json:"ver"`
}

type erpInfo struct {
	Ver string `json:"ver"`
}

func AddWebPartyClientInfo(app *iris.Application) {
	clientInfo := app.Party("/ClientInfo", clientInfoHandler)
	clientInfo.Post("/Info", clientInfoInfoHandler)
}

func clientInfoHandler(ctx iris.Context) {
	ctx.Next()
}

func clientInfoInfoHandler(ctx iris.Context) {
	var info clientInfo

	err := ctx.ReadJSON(&info)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.WriteString(GetErrReturn(err.Error()))
	} else {
		log.Debug(info.OsInfo.Type)
		_, _ = ctx.WriteString(GetMsgReturn("OK"))
	}
}
