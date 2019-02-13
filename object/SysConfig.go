package object

import "encoding/json"

type SysConfig struct {
	Total total `toml:"total"`
	//debug info warn error
	Iris iris `toml:"iris"`
}

type total struct {
	StdOut   bool   `toml:"stdOut"`
	LogLevel string `toml:"logLevel"`
}

type iris struct {
	LogLevel string `toml:"logLevel"`
}

func (sc *SysConfig) GetConfigStr() (string, error) {
	sConfig, err := json.Marshal(sc)
	if err != nil {
		return "", err
	}
	return string(sConfig), nil
}
