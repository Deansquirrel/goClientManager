package config

import "strings"

type ClientConfig struct {
	Total  Total        `toml:"total"`
	Config clientConfig `toml:"client"`
}

type clientConfig struct {
	LogLevel  string `toml:"logLevel"`
	Interval  int    `toml:"interval"`
	ServerUrl string `toml:"serverUrl"`
}

//返回配置字符串
func (cc *ClientConfig) GetConfigStr() (string, error) {
	return getConfigStr(cc)
}

//配置检查并格式化
func (cc *ClientConfig) FormatConfig() {
	cc.Total.FormatConfig()
	cc.Config.FormatConfig()
}

//格式化
func (c *clientConfig) FormatConfig() {
	//间隔时间为空或小于60时，默认设置为60
	if c.Interval < 60 {
		c.Interval = 60
	}
	//除去尾部特殊字符
	c.ServerUrl = checkAndDeleteLastChar(c.ServerUrl, "/")
	c.ServerUrl = checkAndDeleteLastChar(c.ServerUrl, "\\")
	c.ServerUrl = checkAndDeleteLastChar(c.ServerUrl, ".")
	//去除首尾空格
	c.LogLevel = strings.Trim(c.LogLevel, " ")
	//设置Iris默认日志级别
	if c.LogLevel == "" {
		c.LogLevel = "warn"
	}
	//设置字符串转换为小写
	c.LogLevel = strings.ToLower(c.LogLevel)
}
