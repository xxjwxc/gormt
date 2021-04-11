package genmssql

import "regexp"

// genColumns show full columns
type genColumns struct {
	Field   string  `gorm:"column:Field"`
	Type    string  `gorm:"column:Type"`
	Key     string  `gorm:"column:Key"`
	Desc    string  `gorm:"column:Comment"`
	Null    string  `gorm:"column:Null"`
	Default *string `gorm:"column:Default"`
}

//select table_schema,table_name,column_name,referenced_table_schema,referenced_table_name,referenced_column_name from INFORMATION_SCHEMA.KEY_COLUMN_USAGE
// where table_schema ='matrix' AND REFERENCED_TABLE_NAME IS NOT NULL AND TABLE_NAME = 'credit_card' ;
// genForeignKey Foreign key of db info . 表的外键信息
type genForeignKey struct {
	TableSchema           string `gorm:"column:table_schema"`            // Database of columns.列所在的数据库
	TableName             string `gorm:"column:table_name"`              // Data table of column.列所在的数据表
	ColumnName            string `gorm:"column:column_name"`             // Column names.列名
	ReferencedTableSchema string `gorm:"column:referenced_table_schema"` // The database where the index is located.该索引所在的数据库
	ReferencedTableName   string `gorm:"column:referenced_table_name"`   // Affected tables . 该索引受影响的表
	ReferencedColumnName  string `gorm:"column:referenced_column_name"`  // Which column of the affected table.该索引受影响的表的哪一列
}

/////////////////////////////////////////////////////////////////////////

// TableDescription 表及表注释
type TableDescription struct {
	Name  string `gorm:"column:name"`  // 表名
	Value string `gorm:"column:value"` // 表注释
}

type ColumnKeys struct {
	ID     int    `gorm:"column:id"`
	Name   string `gorm:"column:name"`   // 列名
	Pk     int    `gorm:"column:pk"`     // 是否主键
	Type   string `gorm:"column:tp"`     // 类型
	Length int    `gorm:"column:len"`    // 长度
	Isnull int    `gorm:"column:isnull"` // 是否为空
	Desc   string `gorm:"column:des"`    // 列注释
}

var noteRegex = regexp.MustCompile(`^\[@gorm\s(\S+)+\]`)
var foreignKeyRegex = regexp.MustCompile(`^\[@fk\s(\S+)+\]`)
