package config

import (
	"os"
	"strings"
)

const (
	testFile = `
	serial_number : "1.0" #版本号
service_name :     #服务名
service_displayname :  #服务显示名
sercice_desc :    #服务描述
is_dev : false  # 是否开发者模式
out_dir : ./db  # 输出目录
simple : true #简单输出
isJsonTag : true #是否打json标记
mysql_info:
    host : 127.0.0.1
    port : 3306
    username : root
    password : qwer
    database : oauth_db
`
)

// IsRunTesting Determine whether to use it in a test environment.判断是否在测试环境下使用
func IsRunTesting() bool {
	if len(os.Args) > 1 {
		return strings.HasPrefix(os.Args[1], "-test")
	}
	return false
}
