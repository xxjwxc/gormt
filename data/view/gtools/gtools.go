package gtools

import (
	"fmt"

	"github.com/xie1xiao1jun/gorm-tools/data/config"
	"github.com/xie1xiao1jun/public/mysqldb"
)

//开始执行
func Execute() {
	orm := mysqldb.OnInitDBOrm(config.GetMysqlConStr())
	defer orm.OnDestoryDB()

	//获取列名
	var tables []string
	rows, err := orm.Raw("show tables").Rows()
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var table string
		rows.Scan(&table)
		tables = append(tables, table)
	}
	fmt.Println(tables)

}
