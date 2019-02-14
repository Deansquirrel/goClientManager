package router

import (
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
		//_,_ = ctx.WriteString(err.Error())
	} else {
		_, _ = ctx.WriteString(GetErrReturn("OK"))
	}
	//_,_ = ctx.Writef("Received: %#+v\n",info)

	//body := ctx.Request().Body
	//defer func() {
	//	_ = body.Close()
	//}()
	//b, err := ioutil.ReadAll(body)
	//if err != nil {
	//	msg := "读取请求内容时遇到错误"
	//	log.Error(msg + err.Error())
	//	_, err = ctx.WriteString(GetErrReturn(msg))
	//	if err != nil {
	//		log.Error(err.Error())
	//	}
	//	return
	//}
	//log.Debug(string(b))
	//_, err = ctx.WriteString(GetMsgReturn("Complete"))
	//if err != nil {
	//	log.Error(err.Error())
	//}
	//return
}
