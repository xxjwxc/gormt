package config

import (
	"fmt"
	"strings"

	"github.com/xxjwxc/public/tools"
)

// Config custom config struct
type Config struct {
	CfgBase              `yaml:"base"`
	DBInfo               DBInfo            `yaml:"db_info"`
	OutDir               string            `yaml:"out_dir"`
	URLTag               string            `yaml:"url_tag"`  // url tag
	Language             string            `yaml:"language"` // language
	DbTag                string            `yaml:"db_tag"`   // 数据库标签（gormt,db）
	Simple               bool              `yaml:"simple"`
	IsWEBTag             bool              `yaml:"is_web_tag"`
	IsWebTagPkHidden     bool              `yaml:"is_web_tag_pk_hidden"` // web标记是否隐藏主键
	IsForeignKey         bool              `yaml:"is_foreign_key"`
	IsOutSQL             bool              `yaml:"is_out_sql"`
	IsOutFunc            bool              `yaml:"is_out_func"`
	IsGUI                bool              `yaml:"is_gui"` //
	IsTableName          bool              `yaml:"is_table_name"`
	IsNullToPoint        bool              `yaml:"is_null_to_point"` // null to porint
	TablePrefix          string            `yaml:"table_prefix"`     // 表前缀
	SelfTypeDef          map[string]string `yaml:"self_type_define"`
	OutFileName          string            `yaml:"out_file_name"`
	WebTagType           int               `yaml:"web_tag_type"`              // 默认小驼峰
	TableNames           string            `yaml:"table_names"`               // 表名（多个表名用","隔开）
	IsColumnName         bool              `yaml:"is_column_name"`            //是否输出列名
	IsOutFileByTableName bool              `yaml:"is_out_file_by_table_name"` //是否根据表名生成文件(多个表名生成多个文件)
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

// SetIsWEBTag json tag.json标记
func SetIsWEBTag(b bool) {
	_map.IsWEBTag = b
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

// SetTablePrefix set table prefix
func SetTablePrefix(t string) {
	_map.TablePrefix = t
}

// GetTablePrefix get table prefix
func GetTablePrefix() string {
	return _map.TablePrefix
}

// SetSelfTypeDefine 设置自定义字段映射
func SetSelfTypeDefine(data map[string]string) {
	_map.SelfTypeDef = data
}

// GetSelfTypeDefine 获取自定义字段映射
func GetSelfTypeDefine() map[string]string {
	return _map.SelfTypeDef
}

// SetOutFileName 设置输出文件名
func SetOutFileName(s string) {
	_map.OutFileName = s
}

// GetOutFileName 获取输出文件名
func GetOutFileName() string {
	return _map.OutFileName
}

// SetWebTagType 设置json tag类型
func SetWebTagType(i int) {
	_map.WebTagType = i
}

// GetWebTagType 获取json tag类型
func GetWebTagType() int {
	return _map.WebTagType
}

//GetTableNames get format tableNames by config. 获取格式化后设置的表名
func GetTableNames() string {
	var sb strings.Builder
	if _map.TableNames != "" {
		tableNames := _map.TableNames
		tableNames = strings.TrimLeft(tableNames, ",")
		tableNames = strings.TrimRight(tableNames, ",")
		if tableNames == "" {
			return ""
		}

		sarr := strings.Split(_map.TableNames, ",")
		if len(sarr) == 0 {
			fmt.Printf("tableNames is vailed, genmodel will by default global")
			return ""
		}

		for i, val := range sarr {
			sb.WriteString(fmt.Sprintf("'%s'", val))
			if i != len(sarr)-1 {
				sb.WriteString(",")
			}
		}
	}
	return sb.String()
}

//GetOriginTableNames get origin tableNames. 获取原始的设置的表名
func GetOriginTableNames() string {
	return _map.TableNames
}

//SetTableNames set tableNames. 设置生成的表名
func SetTableNames(tableNames string) {
	_map.TableNames = tableNames
}

//GetIsColumnName get  gen columnName config . 获取生成列名的config
func GetIsColumnName() bool {
	return _map.IsColumnName
}

//SetIsColumnName set gen ColumnName config. 设置生成列名的config
func SetIsColumnName(isColumnName bool) {
	_map.IsColumnName = isColumnName
}

//GetIsOutFileByTableName get  gen columnName config . 设置是否根据表名生成文件
func GetIsOutFileByTableName() bool {
	return _map.IsOutFileByTableName
}

//SetIsOutFileByTableName set gen ColumnName config. 设置是否根据表名生成文件
func SetIsOutFileByTableName(isOutFileByTableName bool) {
	_map.IsColumnName = isOutFileByTableName
}
