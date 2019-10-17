package gtools

import (
	"fmt"
	"os/exec"

	"github.com/xxjwxc/gormt/data/view/model"

	"github.com/xxjwxc/gormt/data/config"

	"github.com/xxjwxc/public/tools"
)

// Execute
func Execute() {

	// var tt oauth_db.UserInfoTbl
	// tt.Nickname = "ticket_001"
	// orm.Where("nickname = ?", "ticket_001").Find(&tt)
	// fmt.Println(tt)

	modeldb := GetModel()
	pkg := modeldb.GenModel()
	pkg.PackageName = modeldb.GetPkgName()
	str := model.Generate(pkg)

	path := config.GetOutDir() + "/" + modeldb.GetDbName() + ".go"
	tools.WriteFile(path,
		[]string{str}, true)

	fmt.Println("formatting differs from goimport's:")
	cmd, _ := exec.Command("goimports", "-l", "-w", path).Output()
	fmt.Println(string(cmd))

	fmt.Println("formatting differs from gofmt's:")
	cmd, _ = exec.Command("gofmt", "-l", "-w", path).Output()
	fmt.Println(string(cmd))
}
