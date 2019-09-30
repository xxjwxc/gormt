package genmysql

const (
	_tagGorm = "gorm"
	_tagJSON = "json"
)

// GenColumns show full columns
type GenColumns struct {
	Field string `gorm:"column:Field"`
	Type  string `gorm:"column:Type"`
	Key   string `gorm:"column:Key"`
	Desc  string `gorm:"column:Comment"`
	Null  string `gorm:"column:Null"`
}
