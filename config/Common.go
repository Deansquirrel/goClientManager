package config

import (
	"encoding/json"
	"strings"
)

func getConfigStr(v interface{}) (string, error) {
	sConfig, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(sConfig), nil
}

//检查并删除字符串的最后一个字符
func checkAndDeleteLastChar(s string, deleteChar string) string {
	if strings.HasSuffix(s, deleteChar) {
		s = s[:len(s)-len(deleteChar)]
		return checkAndDeleteLastChar(s, deleteChar)
	} else {
		return s
	}
}
