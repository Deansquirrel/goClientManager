package object

import (
	"encoding/json"
	"strings"
)

type ClientConfig struct {
	Total clientTotal `toml:"total"`
}

type clientTotal struct {
	StdOut    bool   `toml:"stdOut"`
	LogLevel  string `toml:"logLevel"`
	Interval  int    `toml:"interval"`
	ServerUrl string `toml:"serverUrl"`
}

func (sc *ClientConfig) GetConfigStr() (string, error) {
	sConfig, err := json.Marshal(sc)
	if err != nil {
		return "", err
	}
	return string(sConfig), nil
}

//配置检查并格式化
func (sc *ClientConfig) FormatConfig() {
	//===================================================================
	sc.Total.LogLevel = strings.ToLower(sc.Total.LogLevel)
	if sc.Total.Interval < 60 {
		sc.Total.Interval = 60
	}
	sc.Total.ServerUrl = sc.checkAndDeleteLastChar(sc.Total.ServerUrl, "/")
	sc.Total.ServerUrl = sc.checkAndDeleteLastChar(sc.Total.ServerUrl, "\\")
	sc.Total.ServerUrl = sc.checkAndDeleteLastChar(sc.Total.ServerUrl, ".")
	//===================================================================
}

//检查并删除字符串的最后一个字符
func (sc *ClientConfig) checkAndDeleteLastChar(s string, deleteChar string) string {
	if strings.HasSuffix(sc.Total.ServerUrl, deleteChar) {
		s = s[:len(s)-len(deleteChar)]
		return sc.checkAndDeleteLastChar(s, deleteChar)
	} else {
		return s
	}
}
