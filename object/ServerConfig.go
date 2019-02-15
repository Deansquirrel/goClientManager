package object

import (
	"encoding/json"
	"strings"
)

type ServerConfig struct {
	Total serverTotal `toml:"total"`
	//debug info warn error
	Iris iris `toml:"iris"`
}

type serverTotal struct {
	Port     int    `toml:"port"`
	StdOut   bool   `toml:"stdOut"`
	LogLevel string `toml:"logLevel"`
}

type iris struct {
	LogLevel string `toml:"logLevel"`
}

func (sc *ServerConfig) GetConfigStr() (string, error) {
	sConfig, err := json.Marshal(sc)
	if err != nil {
		return "", err
	}
	return string(sConfig), nil
}

//配置检查并格式化
func (sc *ServerConfig) FormatConfig() {
	//===================================================================
	if sc.Total.Port == 0 {
		sc.Total.Port = 8000
	}
	sc.Total.LogLevel = strings.ToLower(sc.Total.LogLevel)
	//===================================================================
	sc.Iris.LogLevel = strings.ToLower(sc.Iris.LogLevel)
	//===================================================================
}
