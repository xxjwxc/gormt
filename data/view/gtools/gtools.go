package gtools

import (
	"github.com/xie1xiao1jun/gorm-tools/data/config"
	"github.com/xie1xiao1jun/public/mysqldb"
)

//开始执行
func Execute() {
	orm := mysqldb.OnInitDBOrm(config.GetMysqlConStr())
	defer orm.OnDestoryDB()

	OnGetPackageInfo(orm, OnGetTables(orm))

}
