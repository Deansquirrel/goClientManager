package common

import (
	"github.com/BurntSushi/toml"
	"github.com/Deansquirrel/goClientManager/object"
	"github.com/Deansquirrel/goToolCommon"
	log "github.com/Deansquirrel/goToolLog"
	"strings"
)

//获取配置
func GetSysConfig(fileName string) (*object.SysConfig, error) {
	path, err := goToolCommon.GetCurrPath()
	if err != nil {
		return nil, err
	}
	var config object.SysConfig
	fileFullPath := path + "\\" + fileName
	b, err := goToolCommon.PathExists(fileFullPath)
	if err != nil {
		log.Warn("检查路径是否存在时遇到错误:" + err.Error() + ",使用默认配置;filePath:" + fileFullPath)
		config = object.SysConfig{}
	} else if !b {
		log.Info("未发现配置文件,使用默认配置" + ";filePath:" + fileFullPath)
		config = object.SysConfig{}
	} else {
		_, err = toml.DecodeFile(fileFullPath, &config)
		if err != nil {
			return nil, err
		}
	}
	return &config, nil
}

//刷新配置
func RefreshConfig(config object.SysConfig) error {
	setLogLevel(strings.ToLower(config.Total.LogLevel))
	setStdOut(config.Total.StdOut)
	return nil
}

//设置标准输出
func setStdOut(isStdOut bool) {
	log.StdOut = isStdOut
}

//设置日志级别
func setLogLevel(logLevel string) {
	switch logLevel {
	case "debug":
		log.Level = log.LevelDebug
		return
	case "info":
		log.Level = log.LevelInfo
		return
	case "warn":
		log.Level = log.LevelWarn
		return
	case "error":
		log.Level = log.LevelError
		return
	default:
		log.Level = log.LevelWarn
	}
}
