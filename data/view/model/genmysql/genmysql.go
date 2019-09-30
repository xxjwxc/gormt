package genmysql

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/xxjwxc/gormt/data/config"
	"github.com/xxjwxc/gormt/data/view/cnf"
	"github.com/xxjwxc/gormt/data/view/genstruct"
	"github.com/xxjwxc/public/mybigcamel"
	"github.com/xxjwxc/public/mysqldb"
)

// GenMysql 开始mysql解析
func GenMysql() genstruct.GenPackage {
	orm := mysqldb.OnInitDBOrm(config.GetMysqlConStr())
	defer orm.OnDestoryDB()

	return OnGetPackageInfo(orm, OnGetTables(orm))
}

// OnGetTables Get columns and comments.获取表列及注释
func OnGetTables(orm *mysqldb.MySqlDB) map[string]string {
	tbDesc := make(map[string]string)

	// Get column names.获取列名
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

	// Get table annotations.获取表注释
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

// OnGetPackageInfo getpackage info.获取包信息
func OnGetPackageInfo(orm *mysqldb.MySqlDB, tabls map[string]string) genstruct.GenPackage {
	var pkg genstruct.GenPackage
	for tab, desc := range tabls {
		var sct genstruct.GenStruct
		sct.SetStructName(OnGetCamelName(tab)) // Big hump.大驼峰
		sct.SetNotes(desc)
		// build element.构造元素
		ems := OnGetTableElement(orm, tab)
		// --------end
		sct.AddElement(ems...)
		// Get table annotations.获取表注释
		rows, err := orm.Raw("show create table " + tab).Rows()
		defer rows.Close()
		if err == nil {
			if rows.Next() {
				var table, CreateTable string
				rows.Scan(&table, &CreateTable)
				sct.SetCreatTableStr(CreateTable)
			}
		}
		// ----------end

		pkg.AddStruct(sct)
	}

	return pkg
}

// OnGetTableElement Get table columns and comments.获取表列及注释
func OnGetTableElement(orm *mysqldb.MySqlDB, tab string) []genstruct.GenElement {
	var el []genstruct.GenElement

	keyNums := make(map[string]int)
	// get keys
	var Keys []struct {
		NonUnique  int    `gorm:"column:Non_unique"`
		KeyName    string `gorm:"column:Key_name"`
		ColumnName string `gorm:"column:Column_name"`
	}
	orm.Raw("show keys from " + tab).Find(&Keys)
	for _, v := range Keys {
		keyNums[v.KeyName]++
	}
	// ----------end

	var list []GenColumns
	// Get table annotations.获取表注释
	orm.Raw("show FULL COLUMNS from " + tab).Find(&list)
	// filter gorm.Model.过滤 gorm.Model
	if OnHaveModel(&list) {
		var tmp genstruct.GenElement
		tmp.SetType("gorm.Model")
		el = append(el, tmp)
	}
	// -----------------end

	for _, v := range list {
		var tmp genstruct.GenElement
		tmp.SetName(OnGetCamelName(v.Field))
		tmp.SetNotes(v.Desc)
		tmp.SetType(OnGetTypeName(v.Type))

		if strings.EqualFold(v.Key, "PRI") { // Set primary key.设置主键
			tmp.AddTag(_tagGorm, "primary_key")
		} else if strings.EqualFold(v.Key, "UNI") { // unique
			tmp.AddTag(_tagGorm, "unique")
		} else {
			// index
			for _, v1 := range Keys {
				if strings.EqualFold(v1.ColumnName, v.Field) {
					_val := ""
					if v1.NonUnique == 1 { // index
						_val += "index"
					} else {
						_val += "unique_index"
					}
					if keyNums[v1.KeyName] > 1 { // Composite index.复合索引
						_val += ":" + v1.KeyName
					}

					tmp.AddTag(_tagGorm, _val)
				}
			}
		}

		// simple output
		if !config.GetSimple() {
			tmp.AddTag(_tagGorm, "column:"+v.Field)
			tmp.AddTag(_tagGorm, "type:"+v.Type)
			if strings.EqualFold(v.Null, "NO") {
				tmp.AddTag(_tagGorm, "not null")
			}
		}

		// json tag
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

// OnHaveModel filter.过滤 gorm.Model
func OnHaveModel(list *[]GenColumns) bool {
	var _temp []GenColumns

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

// OnGetTypeName Type acquisition filtering.类型获取过滤
func OnGetTypeName(name string) string {
	// Precise matching first.先精确匹配
	if v, ok := cnf.TypeMysqlDicMp[name]; ok {
		return v
	}

	// Fuzzy Regular Matching.模糊正则匹配
	for k, v := range cnf.TypeMysqlMatchMp {
		if ok, _ := regexp.MatchString(k, name); ok {
			return v
		}
	}

	panic(fmt.Sprintf("type (%v) not match in any way.maybe need to add on ()", name))
}

// OnGetCamelName Big Hump or Capital Letter.大驼峰或者首字母大写
func OnGetCamelName(name string) string {
	if config.GetSingularTable() { // If the table name plural is globally disabled.如果全局禁用表名复数
		return TitleCase(name)
	}

	return mybigcamel.Marshal(name)
}

// TitleCase title case.首字母大写
func TitleCase(name string) string {
	vv := []rune(name)
	if len(vv) > 0 {
		if bool(vv[0] >= 'a' && vv[0] <= 'z') { // title case.首字母大写
			vv[0] -= 32
		}
	}

	return string(vv)
}
