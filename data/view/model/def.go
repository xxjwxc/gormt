package model

const (
// _tagGorm = "gorm"
// _tagJSON = "json"
)

// ColumnsKey Columns type elem. 类型枚举
type ColumnsKey int

const (
	// ColumnsKeyDefault default
	ColumnsKeyDefault = iota
	// ColumnsKeyPrimary primary key.主键
	ColumnsKeyPrimary // 主键
	// ColumnsKeyUnique unique key.唯一索引
	ColumnsKeyUnique // unix 唯一索引
	// ColumnsKeyIndex index key.复合索引
	ColumnsKeyIndex // 可重复 index 索引
	// ColumnsKeyUniqueIndex unique index key.唯一复合索引
	ColumnsKeyUniqueIndex // 唯一复合索引
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
	Em          []ColumnsInfo // Columns list .表列表组合
}

// ColumnsInfo Columns list .表列信息
type ColumnsInfo struct {
	BaseInfo
	IsNull         bool         // null if db is set null
	Type           string       // Type.类型标记
	Gormt          string       // 默认值
	Index          []KList      // index list.index列表
	ForeignKeyList []ForeignKey // Foreign key list . 表的外键信息
}

// ForeignKey Foreign key of db info . 表的外键信息
type ForeignKey struct {
	TableName  string // Affected tables . 该索引受影响的表
	ColumnName string // Which column of the affected table.该索引受影响的表的哪一列
}

// KList database index /unique_index list.数据库index /unique_index 列表
type KList struct {
	Key     ColumnsKey // non_unique of (show keys from [table])
	Multi   bool       // Multiple .是否多个(复合组建)
	KeyName string     // key_name of (show keys from [table])
	KeyType string     // Key_type of (show keys from [Index_type])
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

// FEm ...
type FEm struct {
	Type          string // 类型
	ColName       string // 列名
	ColStructName string // 列结构体
}

// FList index of list
type FList struct {
	Key     ColumnsKey // non_unique of (show keys from [table])
	KeyName string     // key_name of (show keys from [table])
	Kem     []FEm
}

// EmInfo element of func info
type EmInfo struct {
	IsMulti       bool
	Notes         string // 注释
	Type          string // 类型
	ColName       string // 列名
	ColStructName string // 列结构体
}

type funDef struct {
	StructName  string
	TableName   string
	PreloadList []PreloadInfo // 外键列表，(生成关联数据)
	Em          []EmInfo      // index 列表
	Primay      []FList       // primay unique
	Index       []FList       // index
}

//
