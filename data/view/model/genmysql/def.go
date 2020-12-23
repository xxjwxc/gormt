package genmysql

import "regexp"

type keys struct {
	NonUnique  int    `gorm:"column:Non_unique"`
	KeyName    string `gorm:"column:Key_name"`
	ColumnName string `gorm:"column:Column_name"`
	IndexType  string `gorm:"column:Index_type"`
}

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

var noteRegex = regexp.MustCompile(`^\[@gormt\s(\S+)+\]`)
