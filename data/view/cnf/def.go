package cnf

// EImportsHead imports head options. import包含选项
var EImportsHead = map[string]string{
	"stirng":     `"string"`,
	"time.Time":  `"time"`,
	"gorm.Model": `"github.com/jinzhu/gorm"`,
}

// TypeMysqlDicMp Accurate matching type.精确匹配类型
var TypeMysqlDicMp = map[string]string{
	"int":                 "int",
	"bigint":              "int64",
	"varchar":             "string",
	"char":                "string",
	"date":                "time.Time",
	"datetime":            "time.Time",
	"bit(1)":              "bool",
	"tinyint(1)":          "bool",
	"tinyint(1) unsigned": "bool",
	"json":                "string",
	"text":                "string",
	"timestamp":           "time.Time",
	"double":              "float64",
	"mediumtext":          "string",
}

// TypeMysqlMatchMp Fuzzy Matching Types.模糊匹配类型
var TypeMysqlMatchMp = map[string]string{
	`^(tinyint)[(]\d+[)]`:     "int8",
	`^(smallint)[(]\d+[)]`:    "int8",
	`^(int)[(]\d+[)]`:         "int",
	`^(bigint)[(]\d+[)]`:      "int64",
	`^(char)[(]\d+[)]`:        "string",
	`^(varchar)[(]\d+[)]`:     "string",
	`^(decimal)[(]\d+,\d+[)]`: "float64",
	`^(mediumint)[(]\d+[)]`:   "string",
	`^(double)[(]\d+,\d+[)]`:  "float64",
	`^(float)[(]\d+,\d+[)]`:   "float64",
}
