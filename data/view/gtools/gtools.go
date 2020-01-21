package gtools

import (
	"fmt"
	"os/exec"

	"github.com/xxjwxc/gormt/data/dlg"
	"github.com/xxjwxc/gormt/data/view/model"

	"github.com/xxjwxc/gormt/data/config"

	"github.com/xxjwxc/gormt/data/view/model/genmysql"
	"github.com/xxjwxc/public/tools"
)

// Execute exe the cmd
func Execute() {
	if config.GetIsGUI() {
		dlg.WinMain()
	} else {
		showCmd()
	}
}

func showCmd() {
	// var tt oauth_db.UserInfoTbl
	// tt.Nickname = "ticket_001"
	// orm.Where("nickname = ?", "ticket_001").Find(&tt)
	// fmt.Println(tt)
	modeldb := genmysql.GetMysqlModel()
	pkg := modeldb.GenModel()
	// just for test
	// out, _ := json.Marshal(pkg)
	// tools.WriteFile("test.txt", []string{string(out)}, true)

	list, _ := model.Generate(pkg)

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
