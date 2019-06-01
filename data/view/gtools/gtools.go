package gtools

import (
	"fmt"
	"os/exec"

	"github.com/xie1xiao1jun/gormt/data/config"
	"github.com/xie1xiao1jun/public/mysqldb"
	"github.com/xie1xiao1jun/public/tools"
)

//开始执行
func Execute() {

	orm := mysqldb.OnInitDBOrm(config.GetMysqlConStr())
	defer orm.OnDestoryDB()

	// var tt oauth_db.UserInfoTbl
	// tt.Nickname = "ticket_001"
	// orm.Where("nickname = ?", "ticket_001").Find(&tt)
	// fmt.Println(tt)

	pkg := OnGetPackageInfo(orm, OnGetTables(orm))
	pkg.SetPackage(config.GetMysqlDbInfo().Database)
	str := pkg.Generate()

	path := config.GetOutDir() + "/" + config.GetMysqlDbInfo().Database + "/" + config.GetMysqlDbInfo().Database + ".go"
	tools.WriteFile(path,
		[]string{str}, true)

	cmd, _ := exec.Command("gofmt", "-l", "-w", path).Output()
	fmt.Println(string(cmd))
}
