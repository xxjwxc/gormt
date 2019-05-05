package config

import "fmt"

//
type Config struct {
	CfgBase
	MySQLInfo MysqlDbInfo `toml:"mysql_info"`
}

//mysql 数据库信息
type MysqlDbInfo struct {
	Host     string `validate:"required"` //地址
	Port     int    `validate:"required"` //端口号
	Username string `validate:"required"` //用户名
	Password string `validate:"required"` //密码
	Database string `validate:"required"` //数据库名
}

//更新mysql配置信息
func SetMysqlDbInfo(info *MysqlDbInfo) {
	_map.MySQLInfo = *info
}

//获取mysql配置信息
func GetMysqlDbInfo() MysqlDbInfo {
	return _map.MySQLInfo
}

//获取mysql 连接字符串
func GetMysqlConStr() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		_map.MySQLInfo.Username,
		_map.MySQLInfo.Password,
		_map.MySQLInfo.Host,
		_map.MySQLInfo.Port,
		_map.MySQLInfo.Database,
	)
}
