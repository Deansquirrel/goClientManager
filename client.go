package main

import (
	"context"
	"github.com/Deansquirrel/goClientManager/clientJob"
	"github.com/Deansquirrel/goClientManager/common"
	"github.com/Deansquirrel/goClientManager/global"
	log "github.com/Deansquirrel/goToolLog"
)

func main() {
	//==================================================================================================================
	config, err := common.GetClientConfig("client.toml")
	if err != nil {
		log.Error("加载配置文件时遇到错误：" + err.Error())
		return
	}
	config.FormatConfig()
	global.ClientConfig = config
	err = common.RefreshClientConfig(*config)
	if err != nil {
		log.Error("刷新配置时遇到错误：" + err.Error())
		return
	}
	global.Ctx, global.Cancel = context.WithCancel(context.Background())
	//==================================================================================================================
	log.Warn("程序启动")
	defer log.Warn("程序退出")
	//==================================================================================================================
	if config.Total.ServerUrl == "" {
		log.Warn("服务端地址不能为空")
		return
	}
	cj := clientJob.NewClientJob(config.Total.Interval)
	cj.StartClientJob()
	//==================================================================================================================
	select {
	case <-global.Ctx.Done():
	}
}
