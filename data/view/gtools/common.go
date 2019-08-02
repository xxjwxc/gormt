package gtools

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/xxjwxc/gormt/data/config"
	"github.com/xxjwxc/gormt/data/view/gtools/generate"

	"github.com/xxjwxc/public/mybigcamel"
	"github.com/xxjwxc/public/mysqldb"
)

//OnGetTables 获取表列及注释
func OnGetTables(orm *mysqldb.MySqlDB) map[string]string {
	tbDesc := make(map[string]string)

	//获取列名
	var tables []string
	rows, err := orm.Raw("show tables").Rows()
	if err != nil {
		fmt.Println(err)
		return tbDesc
	}
	defer rows.Close()

	for rows.Next() {
		var table string
		rows.Scan(&table)
		tables = append(tables, table)
		tbDesc[table] = ""
	}

	//获取表注释
	rows, err = orm.Raw("SELECT TABLE_NAME,TABLE_COMMENT FROM information_schema.TABLES WHERE table_schema=?;",
		config.GetMysqlDbInfo().Database).Rows()
	if err != nil {
		fmt.Println(err)
		return tbDesc
	}

	for rows.Next() {
		var table, desc string
		rows.Scan(&table, &desc)
		tbDesc[table] = desc
	}

	return tbDesc
}

//OnGetPackageInfo 获取包信息
func OnGetPackageInfo(orm *mysqldb.MySqlDB, tabls map[string]string) generate.GenPackage {
	var pkg generate.GenPackage
	for tab, desc := range tabls {
		var sct generate.GenStruct
		sct.SetStructName(OnGetCamelName(tab)) //大驼峰
		sct.SetNotes(desc)
		//构造元素
		ems := OnGetTableElement(orm, tab)
		//--------end
		sct.AddElement(ems...)
		//获取表注释
		rows, err := orm.Raw("show create table " + tab).Rows()
		defer rows.Close()
		if err == nil {
			if rows.Next() {
				var table, CreateTable string
				rows.Scan(&table, &CreateTable)
				sct.SetCreatTableStr(CreateTable)
			}
		}
		//----------end

		pkg.AddStruct(sct)
	}

	return pkg
}

//OnGetTableElement 获取表列及注释
func OnGetTableElement(orm *mysqldb.MySqlDB, tab string) []generate.GenElement {
	var el []generate.GenElement

	keyNums := make(map[string]int)
	//获取keys
	var Keys []struct {
		NonUnique  int    `gorm:"column:Non_unique"`
		KeyName    string `gorm:"column:Key_name"`
		ColumnName string `gorm:"column:Column_name"`
	}
	orm.Raw("show keys from " + tab).Find(&Keys)
	for _, v := range Keys {
		keyNums[v.KeyName]++
	}
	//----------end

	var list []struct {
		Field string `gorm:"column:Field"`
		Type  string `gorm:"column:Type"`
		Key   string `gorm:"column:Key"`
		Desc  string `gorm:"column:Comment"`
		Null  string `gorm:"column:Null"`
	}

	//获取表注释
	orm.Raw("show FULL COLUMNS from " + tab).Find(&list)
	//过滤 gorm.Model
	if OnHaveModel(&list) {
		var tmp generate.GenElement
		tmp.SetType("gorm.Model")
		el = append(el, tmp)
	}
	//-----------------end

	for _, v := range list {
		var tmp generate.GenElement
		tmp.SetName(OnGetCamelName(v.Field))
		tmp.SetNotes(v.Desc)
		tmp.SetType(OnGetTypeName(v.Type))

		if strings.EqualFold(v.Key, "PRI") { //设置主键
			tmp.AddTag(_tagGorm, "primary_key")
		} else if strings.EqualFold(v.Key, "UNI") { //unique
			tmp.AddTag(_tagGorm, "unique")
		} else {
			//index
			for _, v1 := range Keys {
				if strings.EqualFold(v1.ColumnName, v.Field) {
					_val := ""
					if v1.NonUnique == 1 { //index
						_val += "index"
					} else {
						_val += "unique_index"
					}
					if keyNums[v1.KeyName] > 1 { //复合索引？
						_val += ":" + v1.KeyName
					}

					tmp.AddTag(_tagGorm, _val)
				}
			}
		}

		//simple output
		if !config.GetSimple() {
			tmp.AddTag(_tagGorm, "column:"+v.Field)
			tmp.AddTag(_tagGorm, "type:"+v.Type)
			if strings.EqualFold(v.Null, "NO") {
				tmp.AddTag(_tagGorm, "not null")
			}
		}

		//json tag
		if config.GetIsJSONTag() {
			if strings.EqualFold(v.Field, "id") {
				tmp.AddTag(_tagJSON, "-")
			} else {
				tmp.AddTag(_tagJSON, v.Field)
			}
		}

		el = append(el, tmp)
	}

	return el
}

//OnHaveModel 过滤 gorm.Model
func OnHaveModel(list *[]struct {
	Field string `gorm:"column:Field"`
	Type  string `gorm:"column:Type"`
	Key   string `gorm:"column:Key"`
	Desc  string `gorm:"column:Comment"`
	Null  string `gorm:"column:Null"`
}) bool {
	var _temp []struct {
		Field string `gorm:"column:Field"`
		Type  string `gorm:"column:Type"`
		Key   string `gorm:"column:Key"`
		Desc  string `gorm:"column:Comment"`
		Null  string `gorm:"column:Null"`
	}

	num := 0
	for _, v := range *list {
		if strings.EqualFold(v.Field, "id") ||
			strings.EqualFold(v.Field, "created_at") ||
			strings.EqualFold(v.Field, "updated_at") ||
			strings.EqualFold(v.Field, "deleted_at") {
			num++
		} else {
			_temp = append(_temp, v)
		}
	}

	if num >= 4 {
		*list = _temp
		return true
	}

	return false
}

//OnGetTypeName 类型获取过滤
func OnGetTypeName(name string) string {
	//先精确匹配
	if v, ok := TypeDicMp[name]; ok {
		return v
	}

	//模糊正则匹配
	for k, v := range TypeMatchMp {
		if ok, _ := regexp.MatchString(k, name); ok {
			return v
		}
	}

	panic(fmt.Sprintf("type (%v) not match in any way.", name))
}

//OnGetCamelName 大驼峰或者首字母大写
func OnGetCamelName(name string) string {
	if config.GetSingularTable() { //如果全局禁用表名复数
		return TitleCase(name)
	}

	return mybigcamel.Marshal(name)
}

//TitleCase 首字母大写
func TitleCase(name string) string {
	vv := []rune(name)
	if len(vv) > 0 {
		if bool(vv[0] >= 'a' && vv[0] <= 'z') { //首字母大写
			vv[0] -= 32
		}
	}

	return string(vv)
}
