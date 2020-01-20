package dlg

import (
	"github.com/xxjwxc/public/mycui"
)

const (
	_menuDefine = "menu"
	_listDefine = "list"
	_viewDefine = "view"
	_run        = "run"
	_set        = "set"
)

var (
	uiPart      = []float32{4, 3}                                 // x,y 对应列表
	mainViewArr = []string{_menuDefine, _listDefine, _viewDefine} // 主菜单列表
	mainIndex   = 0

	btnLogArr = []string{"log_run", "log_set"} // 主菜单列表
	formPart  = []int{14, 28, 10}
)

// menu 内容
type menuDetails struct {
	active  int
	btnList []*mycui.Button
}

var menuDlg *menuDetails
var form *mycui.Form
