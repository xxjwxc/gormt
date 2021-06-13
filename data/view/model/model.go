package model

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/xxjwxc/gormt/data/view/cnf"

	"github.com/xxjwxc/public/mybigcamel"

	"github.com/xxjwxc/gormt/data/config"
	"github.com/xxjwxc/gormt/data/view/genfunc"
	"github.com/xxjwxc/gormt/data/view/genstruct"
)

type _Model struct {
	info DBInfo
	pkg  *genstruct.GenPackage
}

// Generate build code string.生成代码
func Generate(info DBInfo) (out []GenOutInfo, m _Model) {
	m = _Model{
		info: info,
	}

	// struct
	if config.GetIsOutFileByTableName() {
		outByTable := m.GenerateByTableName()
		out = append(out, outByTable...)
	} else {
		var stt GenOutInfo
		stt.FileCtx = m.generate()
		stt.FileName = info.DbName + ".go"

		if name := config.GetOutFileName(); len(name) > 0 {
			stt.FileName = name + ".go"
		}
		out = append(out, stt)
	}

	// ------end

	// gen function
	if config.GetIsOutFunc() {
		out = append(out, m.generateFunc()...)
	}
	// -------------- end
	return
}

// GetPackage gen struct on table
func (m *_Model) GetPackage() genstruct.GenPackage {
	if m.pkg == nil {
		var pkg genstruct.GenPackage
		pkg.SetPackage(m.info.PackageName) //package name

		tablePrefix := config.GetTablePrefix()

		for _, tab := range m.info.TabList {
			var sct genstruct.GenStruct

			sct.SetTableName(tablePrefix + tab.Name)

			//如果设置了表前缀
			// if tablePrefix != "" {
			// 	tab.Name = strings.TrimLeft(tab.Name, tablePrefix)
			// }

			sct.SetStructName(getCamelName(tab.Name)) // Big hump.大驼峰
			sct.SetNotes(tab.Notes)
			sct.AddElement(m.genTableElement(tab.Em)...) // build element.构造元素
			sct.SetCreatTableStr(tab.SQLBuildStr)
			pkg.AddStruct(sct)
		}
		m.pkg = &pkg
	}

	return *m.pkg
}

// GetPackageByTableName Generate multiple model files based on the table name. 根据表名生成多个model文件
func (m *_Model) GenerateByTableName() (out []GenOutInfo) {
	if m.pkg == nil {
		for _, tab := range m.info.TabList {
			var pkg genstruct.GenPackage
			pkg.SetPackage(m.info.PackageName) //package name
			var sct genstruct.GenStruct
			sct.SetStructName(getCamelName(tab.Name)) // Big hump.大驼峰
			sct.SetNotes(tab.Notes)
			sct.AddElement(m.genTableElement(tab.Em)...) // build element.构造元素
			sct.SetCreatTableStr(tab.SQLBuildStr)
			sct.SetTableName(tab.Name)
			pkg.AddStruct(sct)
			var stt GenOutInfo
			stt.FileCtx = pkg.Generate()
			stt.FileName = tab.Name + ".go"
			out = append(out, stt)
		}
	}
	return
}

func (m *_Model) generate() string {
	m.pkg = nil
	m.GetPackage()
	return m.pkg.Generate()
}

// genTableElement Get table columns and comments.获取表列及注释
func (m *_Model) genTableElement(cols []ColumnsInfo) (el []genstruct.GenElement) {
	_tagGorm := config.GetDBTag()
	_tagJSON := config.GetURLTag()

	for _, v := range cols {
		var tmp genstruct.GenElement
		var isPK bool
		if strings.EqualFold(v.Type, "gorm.Model") { // gorm model
			tmp.SetType(v.Type) //
		} else {
			tmp.SetName(getCamelName(v.Name))
			tmp.SetNotes(v.Notes)
			tmp.SetType(getTypeName(v.Type, v.IsNull))
			// not simple output. 默认不输出gorm标签
			if !config.GetSimple() {
				for _, v1 := range v.Index {
					switch v1.Key {
					// case ColumnsKeyDefault:
					case ColumnsKeyPrimary: // primary key.主键
						tmp.AddTag(_tagGorm, "primaryKey")
						isPK = true
					case ColumnsKeyUnique: // unique key.唯一索引
						tmp.AddTag(_tagGorm, "unique")
					case ColumnsKeyIndex: // index key.复合索引
						uninStr := getUninStr("index", ":", v1.KeyName)
						// 兼容 gorm 本身 sort 标签
						if v1.KeyName == "sort" {
							uninStr = "index"
						}
						if v1.KeyType == "FULLTEXT" {
							uninStr += ",class:FULLTEXT"
						}
						tmp.AddTag(_tagGorm, uninStr)
					case ColumnsKeyUniqueIndex: // unique index key.唯一复合索引
						tmp.AddTag(_tagGorm, getUninStr("uniqueIndex", ":", v1.KeyName))
					}
				}
			}
		}

		if len(v.Name) > 0 {
			// not simple output
			if !config.GetSimple() {
				tmp.AddTag(_tagGorm, "column:"+v.Name)
				tmp.AddTag(_tagGorm, "type:"+v.Type)
				if !v.IsNull {
					tmp.AddTag(_tagGorm, "not null")
				}
				// default tag
				if len(v.Gormt) > 0 {
					tmp.AddTag(_tagGorm, v.Gormt)
				}
			}

			// json tag
			if config.GetIsWEBTag() {
				if isPK && config.GetIsWebTagPkHidden() {
					tmp.AddTag(_tagJSON, "-")
				} else {
					if config.GetWebTagType() == 0 {
						tmp.AddTag(_tagJSON, mybigcamel.UnSmallMarshal(mybigcamel.Marshal(v.Name)))
					} else {
						tmp.AddTag(_tagJSON, mybigcamel.UnMarshal(v.Name))
					}
				}
			}

		}

		tmp.ColumnName = v.Name // 列名
		el = append(el, tmp)

		// ForeignKey
		if config.GetIsForeignKey() && len(v.ForeignKeyList) > 0 {
			fklist := m.genForeignKey(v)
			el = append(el, fklist...)
		}
		// -----------end
	}

	return
}

// genForeignKey Get information about foreign key of table column.获取表列外键相关信息
func (m *_Model) genForeignKey(col ColumnsInfo) (fklist []genstruct.GenElement) {
	_tagGorm := config.GetDBTag()
	_tagJSON := config.GetURLTag()

	for _, v := range col.ForeignKeyList {
		isMulti, isFind, notes := m.getColumnsKeyMulti(v.TableName, v.ColumnName)
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

			tmp.AddTag(_tagGorm, "joinForeignKey:"+col.Name) // association_foreignkey
			tmp.AddTag(_tagGorm, "foreignKey:"+v.ColumnName)

			// json tag
			if config.GetIsWEBTag() {
				if config.GetWebTagType() == 0 {
					tmp.AddTag(_tagJSON, mybigcamel.UnSmallMarshal(mybigcamel.Marshal(v.TableName))+"List")
				} else {
					tmp.AddTag(_tagJSON, mybigcamel.UnMarshal(v.TableName)+"_list")
				}
			}

			fklist = append(fklist, tmp)
		}
	}

	return
}

func (m *_Model) getColumnsKeyMulti(tableName, col string) (isMulti bool, isFind bool, notes string) {
	var haveGomod bool
	for _, v := range m.info.TabList {
		if strings.EqualFold(v.Name, tableName) {
			for _, v1 := range v.Em {
				if strings.EqualFold(v1.Name, col) {
					for _, v2 := range v1.Index {
						switch v2.Key {
						case ColumnsKeyPrimary, ColumnsKeyUnique, ColumnsKeyUniqueIndex: // primary key unique key . 主键，唯一索引
							{
								if !v2.Multi { // 唯一索引
									return false, true, v.Notes
								}
							}
							// case ColumnsKeyIndex: // index key. 复合索引
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

// ///////////////////////// func
func (m *_Model) generateFunc() (genOut []GenOutInfo) {
	// getn base
	tmpl, err := template.New("gen_base").Funcs(template.FuncMap{"GetVV": func() string { return "`%v`" }}).Parse(genfunc.GetGenBaseTemp())
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	tmpl.Execute(&buf, m.info)
	genOut = append(genOut, GenOutInfo{
		FileName: "gen.base.go",
		FileCtx:  buf.String(),
	})
	//tools.WriteFile(outDir+"gen_router.go", []string{buf.String()}, true)
	// -------end------

	for _, tab := range m.info.TabList {
		var pkg genstruct.GenPackage
		pkg.SetPackage(m.info.PackageName) //package name
		pkg.AddImport(`"fmt"`)
		pkg.AddImport(`"context"`) // 添加import信息
		pkg.AddImport(cnf.EImportsHead["gorm.Model"])

		// wxw 2021.2.26 17:17
		var data funDef
		data.TableName = tab.Name
		// tablePrefix := config.GetTablePrefix()
		// //如果设置了表前缀
		// if tablePrefix != "" {
		// 	tab.Name = strings.TrimLeft(tab.Name, tablePrefix)
		// }
		data.StructName = getCamelName(tab.Name)

		var primary, unique, uniqueIndex, index []FList
		for _, el := range tab.Em {
			if strings.EqualFold(el.Type, "gorm.Model") {
				data.Em = append(data.Em, getGormModelElement()...)
				pkg.AddImport(`"time"`)
				buildFList(&primary, ColumnsKeyPrimary, "", "int64", "id")
			} else {
				typeName := getTypeName(el.Type, el.IsNull)
				isMulti := (len(el.Index) == 0)
				isUniquePrimary := false
				for _, v1 := range el.Index {
					if v1.Multi {
						isMulti = v1.Multi
					}

					switch v1.Key {
					// case ColumnsKeyDefault:
					case ColumnsKeyPrimary: // primary key.主键
						isUniquePrimary = !v1.Multi
						buildFList(&primary, ColumnsKeyPrimary, v1.KeyName, typeName, el.Name)
					case ColumnsKeyUnique: // unique key.唯一索引
						buildFList(&unique, ColumnsKeyUnique, v1.KeyName, typeName, el.Name)
					case ColumnsKeyIndex: // index key.复合索引
						buildFList(&index, ColumnsKeyIndex, v1.KeyName, typeName, el.Name)
					case ColumnsKeyUniqueIndex: // unique index key.唯一复合索引
						buildFList(&uniqueIndex, ColumnsKeyUniqueIndex, v1.KeyName, typeName, el.Name)
					}
				}

				if isMulti && isUniquePrimary { // 主键唯一
					isMulti = false
				}

				data.Em = append(data.Em, EmInfo{
					IsMulti:       isMulti,
					Notes:         fixNotes(el.Notes),
					Type:          typeName, // Type.类型标记
					ColName:       el.Name,
					ColNameEx:     fmt.Sprintf("`%v`", el.Name),
					ColStructName: getCamelName(el.Name),
				})
				if v2, ok := cnf.EImportsHead[typeName]; ok {
					if len(v2) > 0 {
						pkg.AddImport(v2)
					}
				}
			}

			// 外键列表
			for _, v := range el.ForeignKeyList {
				isMulti, isFind, notes := m.getColumnsKeyMulti(v.TableName, v.ColumnName)
				if isFind {
					var info PreloadInfo
					info.IsMulti = isMulti
					info.Notes = fixNotes(notes)
					info.ForeignkeyTableName = v.TableName
					info.ForeignkeyCol = v.ColumnName
					info.ForeignkeyStructName = getCamelName(v.TableName)
					info.ColName = el.Name
					info.ColStructName = getCamelName(el.Name)
					data.PreloadList = append(data.PreloadList, info)
				}
			}
			// ---------end--
		}

		data.Primary = append(data.Primary, primary...)
		data.Primary = append(data.Primary, unique...)
		data.Primary = append(data.Primary, uniqueIndex...)
		data.Index = append(data.Index, index...)
		tmpl, err := template.New("gen_logic").
			Funcs(template.FuncMap{"GenPreloadList": GenPreloadList, "GenFListIndex": GenFListIndex, "CapLowercase": CapLowercase, "GetTablePrefixName": GetTablePrefixName}).
			Parse(genfunc.GetGenLogicTemp())
		if err != nil {
			panic(err)
		}
		var buf bytes.Buffer
		tmpl.Execute(&buf, data)

		pkg.AddFuncStr(buf.String())
		genOut = append(genOut, GenOutInfo{
			FileName: fmt.Sprintf(m.info.DbName+".gen.%v.go", tab.Name),
			FileCtx:  pkg.Generate(),
		})
	}

	return
}
