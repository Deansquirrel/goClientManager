package main

import (
	"github.com/Deansquirrel/goClientManager/common"
	"github.com/Deansquirrel/goClientManager/global"
)

func main() {
	//==================================================================================================================
	config, err := common.GetSysConfig("client.toml")
	if err != nil {
		common.PrintAndLog("加载配置文件时遇到错误：" + err.Error())
		return
	}
	global.SysConfig = config
	err = common.RefreshConfig(*global.SysConfig)
	if err != nil {
		common.PrintAndLog("刷新配置时遇到错误：" + err.Error())
		return
	}
	//==================================================================================================================
	common.PrintOrLog("程序启动")
	defer common.PrintOrLog("程序退出")
	//==================================================================================================================
}
