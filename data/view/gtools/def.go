package gtools

const (
	_tag_gorm = "gorm"
	_tag_json = "json"
)

//精确匹配类型
var TypeDicMp = map[string]string{
	"int":                 "int",
	"bigint":              "int64",
	"varchar":             "string",
	"datetime":            "time.Time",
	"tinyint(1)":          "bool",
	"tinyint(1) unsigned": "bool",
	"json":                "string",
}

//模糊匹配类型
var TypeMatchMp = map[string]string{
	`int[(]\d+[)]`:     "int",
	`bigint[(]\d+[)]`:  "int64",
	`varchar[(]\d+[)]`: "string",
}
