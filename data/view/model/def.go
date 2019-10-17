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
	PackageName string
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
	Type string // Type.类型标记
	//Index   string     // index or unique_index or ''
	Index  []KList // index list.index列表
	IsNull bool    // null if db is set null

	//
	// Tags map[string][]string // tages.标记
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
