package util

import (
	"strings"
)

// Intercept 返回点后字符串
func Intercept(str string) string {
	comma := strings.LastIndex(str, ".")
	return str[comma:]
}

func Intercept2(str string) string {
	comma := strings.LastIndex(str, ".")
	return str[:comma]
}
