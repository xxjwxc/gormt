package dlg

import (
	"fmt"
	"os/exec"

	"github.com/jroimartin/gocui"
	"github.com/xxjwxc/gormt/data/view/model"
	"github.com/xxjwxc/gormt/data/view/model/genmysql"
	"github.com/xxjwxc/gormt/data/view/model/gensqlite"

	"github.com/xxjwxc/gormt/data/config"

	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"
)

func division(a int, b float32) int {
	r := float32(a) / b
	return (int)(r)
}

func (dlg *menuDetails) nextButton(g *gocui.Gui, v *gocui.View) error {
	dlg.btnList[dlg.active].UnFocus()
	dlg.active = (dlg.active + 1) % len(dlg.btnList)
	menuFocusButton(g)
	return nil
}
func menuFocusButton(g *gocui.Gui) {
	setlog(g, SLocalize(btnLogArr[menuDlg.active]))
	menuDlg.btnList[menuDlg.active].Focus()
}

func (dlg *menuDetails) prevButton(g *gocui.Gui, v *gocui.View) error {
	dlg.btnList[dlg.active].UnFocus()
	if dlg.active == 0 {
		dlg.active = len(dlg.btnList)
	}
	dlg.active--
	menuFocusButton(g)
	return nil
}

func (dlg *menuDetails) Draw() {
	for _, b := range dlg.btnList {
		b.Draw()
	}
}

// OnDestroy destroy windows
func OnDestroy(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func requireValidator(value string) bool {
	if value == "" {
		return false
	}
	return true
}

func getBool(bstr string) bool {
	if bstr == "true" || bstr == " 是" {
		return true
	}

	return false
}

func generate(g *gocui.Gui, v *gocui.View) {
	var modeldb model.IModel
	switch config.GetDbInfo().Type {
	case 0:
		modeldb = genmysql.GetModel()
	case 1:
		modeldb = gensqlite.GetModel()
	}
	if modeldb == nil {
		mylog.Error(fmt.Errorf("modeldb not fund : please check db_info.type (0:mysql , 1:sqlite , 2:mssql) "))
		return
	}

	pkg := modeldb.GenModel()
	// just for test
	// out, _ := json.Marshal(pkg)
	// tools.WriteFile("test.txt", []string{string(out)}, true)

	list, mo := model.Generate(pkg)

	addlog(g, "\n \033[32;7m 开 始 : begin \033[0m\n")

	for _, v := range list {
		path := config.GetOutDir() + "/" + v.FileName
		tools.WriteFile(path, []string{v.FileCtx}, true)

		addlog(g, " formatting differs from goimport's:")
		cmd, _ := exec.Command("goimports", "-l", "-w", path).Output()
		addlog(g, " "+string(cmd))
	}

	addlog(g, "\033[32;7m 所 有 已 完 成 :  ALL completed!! \033[0m\n")
	// build item
	gPkg = mo.GetPackage()
	buildList(g, v)
}
