package main

import (
	"context"
	"github.com/Deansquirrel/goClientManager/common"
	"github.com/Deansquirrel/goClientManager/global"
	"github.com/Deansquirrel/goClientManager/webServer"
	log "github.com/Deansquirrel/goToolLog"
)

func main() {
	//==================================================================================================================
	config, err := common.GetServerConfig("server.toml")
	if err != nil {
		log.Error("加载配置文件时遇到错误：" + err.Error())
		return
	}
	config.FormatConfig()
	global.ServerConfig = config
	err = common.RefreshServerConfig(*global.ServerConfig)
	if err != nil {
		log.Error("刷新配置时遇到错误：" + err.Error())
		return
	}
	global.Ctx, global.Cancel = context.WithCancel(context.Background())
	//==================================================================================================================
	log.Warn("程序启动")
	defer log.Warn("程序退出")
	//==================================================================================================================
	ws := webServer.NewWebServer(global.ServerConfig.Iris.Port)
	ws.StartWebService()
	//==================================================================================================================
	select {
	case <-global.Ctx.Done():
	}
}
