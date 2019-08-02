package config

import "fmt"

//Config .
type Config struct {
	CfgBase
	MySQLInfo     MysqlDbInfo `toml:"mysql_info"`
	OutDir        string      `toml:"out_dir"`
	Simple        bool        `toml:"simple"`
	IsJSONTag     bool        `toml:"isJsonTag"`
	SingularTable bool        `toml:"singular_table"`
}

//MysqlDbInfo mysql 数据库信息
type MysqlDbInfo struct {
	Host     string `validate:"required"` //地址
	Port     int    `validate:"required"` //端口号
	Username string `validate:"required"` //用户名
	Password string `validate:"required"` //密码
	Database string `validate:"required"` //数据库名
}

//SetMysqlDbInfo 更新mysql配置信息
func SetMysqlDbInfo(info *MysqlDbInfo) {
	_map.MySQLInfo = *info
}

//GetMysqlDbInfo 获取mysql配置信息
func GetMysqlDbInfo() MysqlDbInfo {
	return _map.MySQLInfo
}

//GetMysqlConStr 获取mysql 连接字符串
func GetMysqlConStr() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		_map.MySQLInfo.Username,
		_map.MySQLInfo.Password,
		_map.MySQLInfo.Host,
		_map.MySQLInfo.Port,
		_map.MySQLInfo.Database,
	)
}

//SetOutDir 设置输出目录
func SetOutDir(outDir string) {
	_map.OutDir = outDir
}

//GetOutDir 获取输出目录
func GetOutDir() string {
	return _map.OutDir
}

//SetSingularTable 设置禁用表名复数
func SetSingularTable(b bool) {
	_map.SingularTable = b
}

//GetSingularTable 获取禁用表名复数
func GetSingularTable() bool {
	return _map.SingularTable
}

//GetSimple 简单输出
func GetSimple() bool {
	return _map.Simple
}

//GetIsJSONTag json标记
func GetIsJSONTag() bool {
	return _map.IsJSONTag
}
