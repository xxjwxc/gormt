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

	modeldb := GetMysqlModel()
	pkg := modeldb.GenModel()
	list := model.Generate(pkg)

	for _, v := range list {
		path := config.GetOutDir() + "/" + v.FileName
		tools.WriteFile(path, []string{v.FileCtx}, true)

		fmt.Println("formatting differs from goimport's:")
		cmd, _ := exec.Command("goimports", "-l", "-w", path).Output()
		fmt.Println(string(cmd))

		fmt.Println("formatting differs from gofmt's:")
		cmd, _ = exec.Command("gofmt", "-l", "-w", path).Output()
		fmt.Println(string(cmd))
	}
}
