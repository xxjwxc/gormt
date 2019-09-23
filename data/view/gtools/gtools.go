package gtools

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/xxjwxc/gormt/data/config"

	"github.com/xxjwxc/public/mysqldb"
	"github.com/xxjwxc/public/tools"
)

//Execute 开始执行
func Execute() {

	orm := mysqldb.OnInitDBOrm(config.GetMysqlConStr())
	defer orm.OnDestoryDB()

	// var tt oauth_db.UserInfoTbl
	// tt.Nickname = "ticket_001"
	// orm.Where("nickname = ?", "ticket_001").Find(&tt)
	// fmt.Println(tt)

	pkg := OnGetPackageInfo(orm, OnGetTables(orm))
	pkg.SetPackage(getPkgName())
	str := pkg.Generate()

	path := config.GetOutDir() + "/" + config.GetMysqlDbInfo().Database + ".go"
	tools.WriteFile(path,
		[]string{str}, true)

	fmt.Println("formatting differs from goimport's:")
	cmd, _ := exec.Command("goimports", "-l", "-w", path).Output()
	fmt.Println(string(cmd))

	fmt.Println("formatting differs from gofmt's:")
	cmd, _ = exec.Command("gofmt", "-l", "-w", path).Output()
	fmt.Println(string(cmd))
}

// 通过config outdir 配置获取报名
func getPkgName() string {
	dir := config.GetOutDir()
	dir = strings.Replace(dir, "\\", "/", -1)
	if len(dir) > 0 {
		if dir[len(dir)-1] == '/' {
			dir = dir[:(len(dir) - 1)]
		}
	}
	var pkgName string
	list := strings.Split(dir, "/")
	if len(list) > 0 {
		pkgName = list[len(list)-1]
	}

	if len(pkgName) == 0 || pkgName == "." {
		list = strings.Split(tools.GetModelPath(), "/")
		if len(list) > 0 {
			pkgName = list[len(list)-1]
		}
	}

	return pkgName
}
