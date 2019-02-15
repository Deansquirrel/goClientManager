package common

import (
	"github.com/BurntSushi/toml"
	"github.com/Deansquirrel/goClientManager/object"
	"github.com/Deansquirrel/goToolCommon"
	log "github.com/Deansquirrel/goToolLog"
)

//获取配置
func GetServerConfig(fileName string) (*object.ServerConfig, error) {
	path, err := goToolCommon.GetCurrPath()
	if err != nil {
		return nil, err
	}
	var config object.ServerConfig
	fileFullPath := path + "\\" + fileName
	b, err := goToolCommon.PathExists(fileFullPath)
	if err != nil {
		log.Warn("检查路径是否存在时遇到错误:" + err.Error() + ",使用默认配置;filePath:" + fileFullPath)
		config = object.ServerConfig{}
	} else if !b {
		log.Info("未发现配置文件,使用默认配置" + ";filePath:" + fileFullPath)
		config = object.ServerConfig{}
	} else {
		_, err = toml.DecodeFile(fileFullPath, &config)
		if err != nil {
			return nil, err
		}
	}
	return &config, nil
}

//获取配置
func GetClientConfig(fileName string) (*object.ClientConfig, error) {
	path, err := goToolCommon.GetCurrPath()
	if err != nil {
		return nil, err
	}
	var config object.ClientConfig
	fileFullPath := path + "\\" + fileName
	b, err := goToolCommon.PathExists(fileFullPath)
	if err != nil {
		log.Warn("检查路径是否存在时遇到错误:" + err.Error() + ",使用默认配置;filePath:" + fileFullPath)
		config = object.ClientConfig{}
	} else if !b {
		log.Info("未发现配置文件,使用默认配置" + ";filePath:" + fileFullPath)
		config = object.ClientConfig{}
	} else {
		_, err = toml.DecodeFile(fileFullPath, &config)
		if err != nil {
			return nil, err
		}
	}
	return &config, nil
}

//刷新服务端配置
func RefreshServerConfig(config object.ServerConfig) error {
	setLogLevel(config.Total.LogLevel)
	setStdOut(config.Total.StdOut)
	return nil
}

//刷新客户端配置
func RefreshClientConfig(config object.ClientConfig) error {
	setLogLevel(config.Total.LogLevel)
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
