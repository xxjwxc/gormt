package config

import (
	"os"
	"strings"
)

const (
	test_file = `
`
)

//判断是否在测试环境下使用
func IsRunTesting() bool {
	if len(os.Args) > 1 {
		return strings.HasPrefix(os.Args[1], "-test")
	}
	return false
}
