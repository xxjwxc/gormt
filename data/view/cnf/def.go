package cnf

// EImportsHead imports head options. import包含选项
var EImportsHead = map[string]string{
	"stirng":     `"string"`,
	"time.Time":  `"time"`,
	"gorm.Model": `"github.com/jinzhu/gorm"`,
	"fmt":        `"fmt"`,
}

// TypeMysqlDicMp Accurate matching type.精确匹配类型
var TypeMysqlDicMp = map[string]string{
	"smallint":            "int16",
	"smallint unsigned":   "uint16",
	"int":                 "int",
	"int unsigned":        "uint",
	"bigint":              "int64",
	"bigint unsigned":     "uint64",
	"varchar":             "string",
	"char":                "string",
	"date":                "time.Time",
	"datetime":            "time.Time",
	"bit(1)":              "int8",
	"tinyint":             "int8",
	"tinyint unsigned":    "uint8",
	"tinyint(1)":          "int8",
	"tinyint(1) unsigned": "int8",
	"json":                "string",
	"text":                "string",
	"timestamp":           "time.Time",
	"double":              "float64",
	"mediumtext":          "string",
	"longtext":            "string",
	"float":               "float32",
	"tinytext":            "string",
	"enum":                "string",
}

// TypeMysqlMatchMp Fuzzy Matching Types.模糊匹配类型
var TypeMysqlMatchMp = map[string]string{
	`^(tinyint)[(]\d+[)]`:     "int8",
	`^(smallint)[(]\d+[)]`:    "int16",
	`^(int)[(]\d+[)]`:         "int",
	`^(bigint)[(]\d+[)]`:      "int64",
	`^(char)[(]\d+[)]`:        "string",
	`^(enum)[(](.)+[)]`:       "string",
	`^(varchar)[(]\d+[)]`:     "string",
	`^(varbinary)[(]\d+[)]`:   "[]byte",
	`^(decimal)[(]\d+,\d+[)]`: "float64",
	`^(mediumint)[(]\d+[)]`:   "string",
	`^(double)[(]\d+,\d+[)]`:  "float64",
	`^(float)[(]\d+,\d+[)]`:   "float64",
	`^(datetime)[(]\d+[)]`:    "time.Time",
}
