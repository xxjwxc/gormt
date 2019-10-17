package model

import (
	"strings"

	"github.com/xxjwxc/gormt/data/config"
	"github.com/xxjwxc/gormt/data/view/genstruct"
)

// Generate build code string.生成代码
func Generate(info DBInfo) string {
	var pkg genstruct.GenPackage
	pkg.SetPackage(info.PackageName) //package name
	for _, tab := range info.TabList {
		var sct genstruct.GenStruct
		sct.SetStructName(getCamelName(tab.Name)) // Big hump.大驼峰
		sct.SetNotes(tab.Notes)
		sct.AddElement(getTableElement(tab.Em)...) // build element.构造元素
		sct.SetCreatTableStr(tab.SQLBuildStr)
		pkg.AddStruct(sct)
	}

	return pkg.Generate()
}

// getTableElement Get table columns and comments.获取表列及注释
func getTableElement(tabs []ColumusInfo) (el []genstruct.GenElement) {
	for _, v := range tabs {
		var tmp genstruct.GenElement
		if strings.EqualFold(v.Type, "gorm.Model") { // gorm model
			tmp.SetType(v.Type) //
		} else {
			tmp.SetName(getCamelName(v.Name))
			tmp.SetNotes(v.Notes)
			tmp.SetType(getTypeName(v.Type))
			for _, v1 := range v.Index {
				switch v1.Key {
				// case ColumusKeyDefault:
				case ColumusKeyPrimary: // primary key.主键
					tmp.AddTag(_tagGorm, "primary_key")
				case ColumusKeyUnique: // unique key.唯一索引
					tmp.AddTag(_tagGorm, "unique")
				case ColumusKeyIndex: // index key.复合索引
					if len(v1.KeyName) > 0 {
						tmp.AddTag(_tagGorm, "index:"+v1.KeyName)
					} else {
						tmp.AddTag(_tagGorm, "index")
					}
				case ColumusKeyUniqueIndex: // unique index key.唯一复合索引
					if len(v1.KeyName) > 0 {
						tmp.AddTag(_tagGorm, "unique_index:"+v1.KeyName)
					} else {
						tmp.AddTag(_tagGorm, "unique_index")
					}
				}
			}
		}

		// not simple output
		if !config.GetSimple() && len(v.Name) > 0 {
			tmp.AddTag(_tagGorm, "column:"+v.Name)
			tmp.AddTag(_tagGorm, "type:"+v.Type)
			if !v.IsNull {
				tmp.AddTag(_tagGorm, "not null")
			}
		}

		// json tag
		if config.GetIsJSONTag() {
			if strings.EqualFold(v.Name, "id") {
				tmp.AddTag(_tagJSON, "-")
			} else if len(v.Name) > 0 {
				tmp.AddTag(_tagJSON, v.Name)
			}
		}
		el = append(el, tmp)
	}

	return
}
