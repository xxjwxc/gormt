package model

import (
	"strings"

	"github.com/xxjwxc/public/mybigcamel"

	"github.com/xxjwxc/gormt/data/config"
	"github.com/xxjwxc/gormt/data/view/genstruct"
)

type _Model struct {
	info DBInfo
}

// Generate build code string.生成代码
func Generate(info DBInfo) (out []GenOutInfo) {
	m := _Model{
		info: info,
	}

	// struct
	var stt GenOutInfo
	stt.FileCtx = m.generate()
	stt.FileName = info.DbName + ".go"
	out = append(out, stt)
	// ------end

	// gen function
	// -------------- end
	return
}

func (m *_Model) generate() string {
	var pkg genstruct.GenPackage
	pkg.SetPackage(m.info.PackageName) //package name
	for _, tab := range m.info.TabList {
		var sct genstruct.GenStruct
		sct.SetStructName(getCamelName(tab.Name)) // Big hump.大驼峰
		sct.SetNotes(tab.Notes)
		sct.AddElement(m.genTableElement(tab.Em)...) // build element.构造元素
		sct.SetCreatTableStr(tab.SQLBuildStr)
		pkg.AddStruct(sct)
	}

	return pkg.Generate()
}

// genTableElement Get table columns and comments.获取表列及注释
func (m *_Model) genTableElement(cols []ColumusInfo) (el []genstruct.GenElement) {
	for _, v := range cols {
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
				tmp.AddTag(_tagJSON, mybigcamel.UnMarshal(v.Name))
			}
		}
		el = append(el, tmp)

		// ForeignKey
		if config.GetIsForeignKey() && len(v.ForeignKeyList) > 0 {
			fklist := m.genForeignKey(v)
			if len(fklist) > 0 {
				el = append(el, fklist...)
			}
		}
		// -----------end
	}

	return
}

// genForeignKey Get information about foreign key of table column.获取表列外键相关信息
func (m *_Model) genForeignKey(col ColumusInfo) (fklist []genstruct.GenElement) {
	for _, v := range col.ForeignKeyList {
		isMulti, isFind, notes := m.getColumusKeyMulti(v.TableName, v.ColumnName)
		if isFind {
			var tmp genstruct.GenElement
			tmp.SetNotes(notes)
			if isMulti {
				tmp.SetName(getCamelName(v.TableName) + "List")
				tmp.SetType("[]" + getCamelName(v.TableName))
			} else {
				tmp.SetName(getCamelName(v.TableName))
				tmp.SetType(getCamelName(v.TableName))
			}

			tmp.AddTag(_tagGorm, "association_foreignkey:"+col.Name)
			tmp.AddTag(_tagGorm, "foreignkey:"+v.ColumnName)

			// json tag
			if config.GetIsJSONTag() {
				tmp.AddTag(_tagJSON, mybigcamel.UnMarshal(v.TableName)+"_list")
			}

			fklist = append(fklist, tmp)
		}
	}

	return
}

func (m *_Model) getColumusKeyMulti(tableName, col string) (isMulti bool, isFind bool, notes string) {
	var haveGomod bool
	for _, v := range m.info.TabList {
		if strings.EqualFold(v.Name, tableName) {
			for _, v1 := range v.Em {
				if strings.EqualFold(v1.Name, col) {
					for _, v2 := range v1.Index {
						switch v2.Key {
						case ColumusKeyPrimary, ColumusKeyUnique, ColumusKeyUniqueIndex: // primary key unique key . 主键，唯一索引
							{
								return false, true, v.Notes
							}
							// case ColumusKeyIndex: // index key. 复合索引
							// 	{
							// 		isMulti = true
							// 	}
						}
					}
					return true, true, v.Notes
				} else if strings.EqualFold(v1.Type, "gorm.Model") {
					haveGomod = true
					notes = v.Notes
				}
			}
			break
		}
	}

	// default gorm.Model
	if haveGomod {
		if strings.EqualFold(col, "id") {
			return false, true, notes
		}

		if strings.EqualFold(col, "created_at") ||
			strings.EqualFold(col, "updated_at") ||
			strings.EqualFold(col, "deleted_at") {
			return true, true, notes
		}
	}

	return false, false, ""
	// -----------------end
}
