package model

const (
	_tagGorm = "gorm"
	_tagJSON = "json"
)

// ColumusKey Columus type elem. 类型枚举
type ColumusKey int

const (
	// ColumusKeyDefault default
	ColumusKeyDefault = iota
	// ColumusKeyPrimary primary key.主键
	ColumusKeyPrimary
	// ColumusKeyUnique unique key.唯一索引
	ColumusKeyUnique
	// ColumusKeyIndex index key.复合索引
	ColumusKeyIndex
	// ColumusKeyUniqueIndex unique index key.唯一复合索引
	ColumusKeyUniqueIndex
)

// DBInfo database default info
type DBInfo struct {
	DbName      string    // database name
	PackageName string    // package name
	TabList     []TabInfo // table list .表列表
}

// TabInfo database table default attribute
type TabInfo struct {
	BaseInfo
	SQLBuildStr string        // Create SQL statements.创建sql语句
	Em          []ColumusInfo // Columus list .表列表组合
}

// ColumusInfo Columus list .表列信息
type ColumusInfo struct {
	BaseInfo
	Type           string       // Type.类型标记
	Index          []KList      // index list.index列表
	IsNull         bool         // null if db is set null
	ForeignKeyList []ForeignKey // Foreign key list . 表的外键信息
}

// ForeignKey Foreign key of db info . 表的外键信息
type ForeignKey struct {
	TableName  string // Affected tables . 该索引受影响的表
	ColumnName string // Which column of the affected table.该索引受影响的表的哪一列
}

// KList database index /unique_index list.数据库index /unique_index 列表
type KList struct {
	Key     ColumusKey // non_unique of (show keys from [table])
	KeyName string     // key_name of (show keys from [table])
}

// BaseInfo base common attribute. 基础属性
type BaseInfo struct {
	Name  string // table name.表名
	Notes string // table comment . 表注释
}

// GenOutInfo generate file list. 生成的文件列表
type GenOutInfo struct {
	FileName string // output file name .输出文件名
	FileCtx  string // output file context. 输出文件内容
}

// def func sturct

// PreloadInfo 预加载列表
type PreloadInfo struct {
	IsMulti              bool
	Notes                string // 注释
	ForeignkeyStructName string // 外键类目
	ForeignkeyTableName  string // 外键表名
	ForeignkeyCol        string // 外键列表
	ColName              string // 表名
	ColStructName        string // 表结构体
}

// EmInfo func 表结构定义
type EmInfo struct {
	IsMulti       bool
	Notes         string // 注释
	Type          string // 类型
	ColName       string // 表名
	ColStructName string // 表结构体
}

type funDef struct {
	StructName  string
	TableName   string
	PreloadList []PreloadInfo
	Em          []EmInfo
}

//
