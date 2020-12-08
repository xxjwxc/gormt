package config

import (
	"fmt"

	"github.com/xxjwxc/public/tools"
)

// Config custom config struct
type Config struct {
	CfgBase          `yaml:"base"`
	DBInfo           DBInfo `yaml:"db_info"`
	OutDir           string `yaml:"out_dir"`
	URLTag           string `yaml:"url_tag"`  // url tag
	Language         string `yaml:"language"` // language
	DbTag            string `yaml:"db_tag"`   // 数据库标签（gormt,db）
	Simple           bool   `yaml:"simple"`
	IsWEBTag         bool   `yaml:"is_web_tag"`
	IsWebTagPkHidden bool   `yaml:"is_web_tag_pk_hidden"` // web标记是否隐藏主键
	IsForeignKey     bool   `yaml:"is_foreign_key"`
	IsOutSQL         bool   `yaml:"is_out_sql"`
	IsOutFunc        bool   `yaml:"is_out_func"`
	IsGUI            bool   `yaml:"is_gui"` //
	IsTableName      bool   `yaml:"is_table_name"`
	IsNullToPoint    bool   `yaml:"is_null_to_point"` // null to porint
}

// DBInfo mysql database information. mysql 数据库信息
type DBInfo struct {
	Host     string `validate:"required"` // Host. 地址
	Port     int    // Port 端口号
	Username string // Username 用户名
	Password string // Password 密码
	Database string // Database 数据库名
	Type     int    // 数据库类型: 0:mysql , 1:sqlite , 2:mssql
}

// SetMysqlDbInfo Update MySQL configuration information
func SetMysqlDbInfo(info *DBInfo) {
	_map.DBInfo = *info
}

// GetDbInfo Get configuration information .获取数据配置信息
func GetDbInfo() DBInfo {
	return _map.DBInfo
}

// GetMysqlConStr Get MySQL connection string.获取mysql 连接字符串
func GetMysqlConStr() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&interpolateParams=True",
		_map.DBInfo.Username,
		_map.DBInfo.Password,
		_map.DBInfo.Host,
		_map.DBInfo.Port,
		_map.DBInfo.Database,
	)
}

// SetOutDir Setting Output Directory.设置输出目录
func SetOutDir(outDir string) {
	_map.OutDir = outDir
}

// GetOutDir Get Output Directory.获取输出目录
func GetOutDir() string {
	if len(_map.OutDir) == 0 {
		_map.OutDir = "./model"
	}

	return _map.OutDir
}

// // SetSingularTable Set Disabled Table Name Plurals.设置禁用表名复数
// func SetSingularTable(b bool) {
// 	_map.SingularTable = b
// }

// // GetSingularTable Get Disabled Table Name Plurals.获取禁用表名复数
// func GetSingularTable() bool {
// 	return _map.SingularTable
// }

// GetSimple simple output.简单输出
func GetSimple() bool {
	return _map.Simple
}

// SetSimple simple output.简单输出
func SetSimple(b bool) {
	_map.Simple = b
}

// GetIsWEBTag json tag.json标记
func GetIsWEBTag() bool {
	return _map.IsWEBTag
}

// GetIsWebTagPkHidden web tag是否隐藏主键
func GetIsWebTagPkHidden() bool {
	return _map.IsWebTagPkHidden
}

// GetIsForeignKey if is foreign key
func GetIsForeignKey() bool {
	return _map.IsForeignKey
}

// SetForeignKey Set if is foreign key.设置是否外键关联
func SetForeignKey(b bool) {
	_map.IsForeignKey = b
}

// SetIsOutSQL if is output sql .
func SetIsOutSQL(b bool) {
	_map.IsOutSQL = b
}

// GetIsOutSQL if is output sql .
func GetIsOutSQL() bool {
	return _map.IsOutSQL
}

// GetIsOutFunc if is output func .
func GetIsOutFunc() bool {
	return _map.IsOutFunc
}

// SetIsOutFunc if is output func .
func SetIsOutFunc(b bool) {
	_map.IsOutFunc = b
}

// GetIsGUI if is gui show .
func GetIsGUI() bool {
	return _map.IsGUI
}

// SetIsGUI if is gui show .
func SetIsGUI(b bool) {
	_map.IsGUI = b
}

// GetIsTableName if is table name .
func GetIsTableName() bool {
	return _map.IsTableName
}

// SetIsTableName if is table name .
func SetIsTableName(b bool) {
	_map.IsTableName = b
}

// GetURLTag get url tag.
func GetURLTag() string {
	if _map.URLTag != "json" && _map.URLTag != "url" {
		_map.URLTag = "json"
	}

	return _map.URLTag
}

// SetURLTag set url tag.
func SetURLTag(s string) {
	_map.URLTag = s
}

// GetLG get language tag.
func GetLG() string {
	if _map.Language != "English" && _map.Language != "中 文" {
		if tools.GetLocalSystemLang(true) == "en" {
			_map.Language = "English"
		} else {
			_map.Language = "中 文"
		}
	}

	return _map.Language
}

// SetLG set url tag.
func SetLG(s string) {
	_map.Language = s
}

// GetDBTag get database tag.
func GetDBTag() string {
	if _map.DbTag != "gorm" && _map.DbTag != "db" {
		_map.DbTag = "gorm"
	}

	return _map.DbTag
}

// SetDBTag get database tag.
func SetDBTag(s string) {
	_map.DbTag = s
}

// SetIsNullToPoint if with null to porint in struct
func SetIsNullToPoint(b bool) {
	_map.IsNullToPoint = b
}

// GetIsNullToPoint get if with null to porint in sturct
func GetIsNullToPoint() bool {
	return _map.IsNullToPoint
}
