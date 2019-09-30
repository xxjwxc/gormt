package config

import "fmt"

// Config custom config struct
type Config struct {
	CfgBase
	MySQLInfo     MysqlDbInfo `toml:"mysql_info"`
	OutDir        string      `toml:"out_dir"`
	Simple        bool        `toml:"simple"`
	IsJSONTag     bool        `toml:"isJsonTag"`
	SingularTable bool        `toml:"singular_table"`
}

// MysqlDbInfo mysql database information. mysql 数据库信息
type MysqlDbInfo struct {
	Host     string `validate:"required"` // Host. 地址
	Port     int    `validate:"required"` // Port 端口号
	Username string `validate:"required"` // Username 用户名
	Password string `validate:"required"` // Password 密码
	Database string `validate:"required"` // Database 数据库名
}

// SetMysqlDbInfo Update MySQL configuration information
func SetMysqlDbInfo(info *MysqlDbInfo) {
	_map.MySQLInfo = *info
}

// GetMysqlDbInfo Get MySQL configuration information .获取mysql配置信息
func GetMysqlDbInfo() MysqlDbInfo {
	return _map.MySQLInfo
}

// GetMysqlConStr Get MySQL connection string.获取mysql 连接字符串
func GetMysqlConStr() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		_map.MySQLInfo.Username,
		_map.MySQLInfo.Password,
		_map.MySQLInfo.Host,
		_map.MySQLInfo.Port,
		_map.MySQLInfo.Database,
	)
}

// SetOutDir Setting Output Directory.设置输出目录
func SetOutDir(outDir string) {
	_map.OutDir = outDir
}

// GetOutDir Get Output Directory.获取输出目录
func GetOutDir() string {
	return _map.OutDir
}

// SetSingularTable Set Disabled Table Name Plurals.设置禁用表名复数
func SetSingularTable(b bool) {
	_map.SingularTable = b
}

// GetSingularTable Get Disabled Table Name Plurals.获取禁用表名复数
func GetSingularTable() bool {
	return _map.SingularTable
}

// GetSimple simple output.简单输出
func GetSimple() bool {
	return _map.Simple
}

// GetIsJSONTag json tag.json标记
func GetIsJSONTag() bool {
	return _map.IsJSONTag
}
