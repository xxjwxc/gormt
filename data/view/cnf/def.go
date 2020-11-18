package cnf

// EImportsHead imports head options. import包含选项
var EImportsHead = map[string]string{
	"stirng":     `"string"`,
	"time.Time":  `"time"`,
	"gorm.Model": `"gorm.io/gorm"`,
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
	"bit(1)":              "[]uint8",
	"tinyint":             "int8",
	"tinyint unsigned":    "uint8",
	"tinyint(1)":          "bool", // tinyint(1) 默认设置成bool
	"tinyint(1) unsigned": "bool", // tinyint(1) 默认设置成bool
	"json":                "string",
	"text":                "string",
	"timestamp":           "time.Time",
	"double":              "float64",
	"double unsigned":     "float64",
	"mediumtext":          "string",
	"longtext":            "string",
	"float":               "float32",
	"float unsigned":      "float32",
	"tinytext":            "string",
	"enum":                "string",
	"time":                "time.Time",
	"tinyblob":            "[]byte",
	"blob":                "[]byte",
	"mediumblob":          "[]byte",
	"longblob":            "[]byte",
	"integer":             "int64",
}

// TypeMysqlMatchMp Fuzzy Matching Types.模糊匹配类型
var TypeMysqlMatchMp = map[string]string{
	`^(tinyint)[(]\d+[)] unsigned`:    "uint8",
	`^(smallint)[(]\d+[)] unsigned`:   "uint16",
	`^(int)[(]\d+[)] unsigned`:        "uint32",
	`^(bigint)[(]\d+[)] unsigned`:     "uint64",
	`^(tinyint)[(]\d+[)]`:             "int8",
	`^(smallint)[(]\d+[)]`:            "int16",
	`^(int)[(]\d+[)]`:                 "int",
	`^(bigint)[(]\d+[)]`:              "int64",
	`^(char)[(]\d+[)]`:                "string",
	`^(enum)[(](.)+[)]`:               "string",
	`^(varchar)[(]\d+[)]`:             "string",
	`^(varbinary)[(]\d+[)]`:           "[]byte",
	`^(blob)[(]\d+[)]`:                "[]byte",
	`^(binary)[(]\d+[)]`:              "[]byte",
	`^(decimal)[(]\d+,\d+[)]`:         "float64",
	`^(mediumint)[(]\d+[)]`:           "string",
	`^(double)[(]\d+,\d+[)]`:          "float64",
	`^(double)[(]\d+,\d+[)] unsigned`: "float64",
	`^(float)[(]\d+,\d+[)]`:           "float64",
	`^(float)[(]\d+,\d+[)] unsigned`:  "float64",
	`^(datetime)[(]\d+[)]`:            "time.Time",
	`^(bit)[(]\d+[)]`:                 "[]uint8",
	`^(text)[(]\d+[)]`:                "string",
	`^(integer)[(]\d+[)]`:             "int",
}
