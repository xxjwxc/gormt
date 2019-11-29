package config

import "fmt"

// Config custom config struct
type Config struct {
	CfgBase       `yaml:"base"`
	MySQLInfo     MysqlDbInfo `yaml:"mysql_info"`
	OutDir        string      `yaml:"out_dir"`
	Simple        bool        `yaml:"simple"`
	IsJSONTag     bool        `yaml:"is_json_tag"`
	SingularTable bool        `yaml:"singular_table"`
	IsForeignKey  bool        `yaml:"is_foreign_key"`
	IsOutSQL      bool        `yaml:"is_out_sql"`
}

// MysqlDbInfo mysql database information. mysql 数据库信息
type MysqlDbInfo struct {
	Host     string `validate:"required"` // Host. 地址
	Port     int    `validate:"required"` // Port 端口号
	Username string `validate:"required"` // Username 用户名
	Password string // Password 密码
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

// GetIsForeignKey if is foreign key
func GetIsForeignKey() bool {
	return _map.IsForeignKey
}

// SetForeignKey Set if is foreign key.设置是否外键关联
func SetForeignKey(b bool) {
	_map.IsForeignKey = b
}

// GetIsOutSQL if is output sql .
func GetIsOutSQL() bool {
	return _map.IsOutSQL
}
