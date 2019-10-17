package genmysql

type keys struct {
	NonUnique  int    `gorm:"column:Non_unique"`
	KeyName    string `gorm:"column:Key_name"`
	ColumnName string `gorm:"column:Column_name"`
}

// genColumns show full columns
type genColumns struct {
	Field string `gorm:"column:Field"`
	Type  string `gorm:"column:Type"`
	Key   string `gorm:"column:Key"`
	Desc  string `gorm:"column:Comment"`
	Null  string `gorm:"column:Null"`
}
