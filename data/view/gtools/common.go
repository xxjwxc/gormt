package gtools

import (
	"fmt"

	"github.com/xie1xiao1jun/gorm-tools/data/config"
	"github.com/xie1xiao1jun/gorm-tools/data/view/gtools/generate"
	"github.com/xie1xiao1jun/public/mybigcamel"
	"github.com/xie1xiao1jun/public/mysqldb"
)

//获取表列及注释
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
		tables = append(tables, table)
		tbDesc[table] = desc
	}

	return tbDesc
}

//获取包信息
func OnGetPackageInfo(orm *mysqldb.MySqlDB, tabls map[string]string) generate.GenPackage {
	var pkg generate.GenPackage
	for tab, desc := range tabls {
		var sct generate.GenStruct
		sct.SetStructName(OnGetCamelName(tab)) //大驼峰
		sct.SetNotes(desc)
		//构造元素

		//--------end

		pkg.AddStruct(sct)
	}

	return pkg
}

// //获取表列及注释
func OnGetTables(orm *mysqldb.MySqlDB, tab string) []generate.GenElement {
	var el []generate.GenElement

}

//大驼峰或者首字母大写
func OnGetCamelName(name string) string {
	if config.GetSingularTable() { //如果全局禁用表名复数
		return TitleCase(name)
	}

	return mybigcamel.Marshal(name)
}

/*
	首字母大写
*/
func TitleCase(name string) string {
	vv := []rune(name)
	if len(vv) > 0 {
		if bool(vv[0] >= 'a' && vv[0] <= 'z') { //首字母大写
			vv[0] -= 32
		}
	}

	return string(vv)
}
