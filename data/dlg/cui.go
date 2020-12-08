package dlg

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/xxjwxc/public/tools"

	"github.com/xxjwxc/gormt/data/config"

	"github.com/jroimartin/gocui"
	"github.com/xxjwxc/public/myclipboard"
	"github.com/xxjwxc/public/mycui"
)

func nextView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (mainIndex + 1) % len(mainViewArr)
	name := mainViewArr[nextIndex]

	if _, err := g.SetCurrentView(name); err != nil { // 设置选中
		return err
	}
	g.SelFgColor = gocui.ColorGreen // 设置边框颜色
	g.FgColor = gocui.ColorWhite

	switch name {
	case _menuDefine:
		//g.Cursor = false // 光标
		// g.FgColor = gocui.ColorGreen
		menuDlg.btnList[menuDlg.active].Focus()
	case _listDefine:
		//g.Cursor = false
		menuDlg.btnList[menuDlg.active].UnFocus()
	case _viewDefine:
		//g.Cursor = true
		menuDlg.btnList[menuDlg.active].UnFocus()
	}

	mainIndex = nextIndex
	return nil
}

func mainLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("main_title", maxX/2-16, -1, maxX/2+16, 1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SelFgColor = gocui.ColorGreen | gocui.AttrUnderline
		fmt.Fprintln(v, "https://github.com/xxjwxc/gormt")
	}

	if v, err := g.SetView(_menuDefine, 0, 1, division(maxX, uiPart[0])-1, division(maxY, uiPart[1])-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = SLocalize(_menuDefine)
		// v.Editable = true // 是否可以编辑
		v.Wrap = true
		v.Autoscroll = true
		// g.FgColor = gocui.ColorGreen
	}

	if v, err := g.SetView(_listDefine, 0, division(maxY, uiPart[1]), division(maxX, uiPart[0])-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = SLocalize(_listDefine)
		// v.Wrap = true
		// v.Autoscroll = true
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		// if _, err := g.SetCurrentView(_menuDefine); err != nil {
		// 	return err
		// }
	}

	if v, err := g.SetView(_viewDefine, division(maxX, uiPart[0]), 1, maxX-1, maxY-3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = SLocalize(_viewDefine)
		v.Wrap = true
		v.Autoscroll = true
		v.Editable = true
	}

	nemuLayOut(g) // menuLayOut
	return nil
}

func nemuLayOut(g *gocui.Gui) {
	menuDlg = &menuDetails{}
	menuDlg.btnList = append(menuDlg.btnList,
		mycui.NewButton(g, _run, SLocalize(_run), 0, 2, 3).SetTextColor(gocui.ColorRed|gocui.AttrReverse, gocui.ColorWhite).
			AddHandler(gocui.KeyArrowUp, menuDlg.prevButton).AddHandler(gocui.KeyArrowDown, menuDlg.nextButton).
			AddHandler(gocui.KeyEnter, enterRun).AddHandler(gocui.MouseLeft, enterRun))

	menuDlg.btnList = append(menuDlg.btnList,
		mycui.NewButton(g, _set, SLocalize(_set), 0, 4, 3).
			AddHandler(gocui.KeyArrowUp, menuDlg.prevButton).AddHandler(gocui.KeyArrowDown, menuDlg.nextButton).
			AddHandler(gocui.KeyEnter, enterSet).AddHandler(gocui.MouseLeft, enterSet))

	maxX, maxY := g.Size() // division(maxY, uiPart[1])
	clipboardBtn = mycui.NewButton(g, _clipboardBtn, SLocalize(_clipboardBtn), division(maxX, uiPart[0])+2, maxY-3, 5).
		AddHandler(gocui.KeyEnter, enterClipboard).AddHandler(gocui.MouseLeft, enterClipboard)
	clipboardBtn.Draw()

	menuDlg.Draw()
	menuFocusButton(g)
}

func keybindings(g *gocui.Gui) {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, OnDestroy); err != nil { // 退出事件
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, nextView); err != nil { // tab next事件
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyEsc, gocui.ModNone, buttonCancel); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlQ, gocui.ModNone, buttonCancel); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("main_title", gocui.MouseLeft, gocui.ModNone, about); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding(_listDefine, gocui.KeyArrowDown, gocui.ModNone, listDown); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding(_listDefine, gocui.KeyArrowUp, gocui.ModNone, listUp); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding(_listDefine, gocui.KeyEnter, gocui.ModNone, showStruct); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding(_listDefine, gocui.MouseLeft, gocui.ModNone, showStruct); err != nil {
		log.Panicln(err)
	}

}

///////////////////signal slot ///////////
func about(g *gocui.Gui, v *gocui.View) error {
	openURL("https://github.com/xxjwxc/gormt")
	return nil
}

func setlog(g *gocui.Gui, str string) error {
	logView, err := g.View(_viewDefine)
	if err == nil {
		logView.Clear()
		fmt.Fprintln(logView, str)
	}

	return err
}

func addlog(g *gocui.Gui, str string) error {
	logView, err := g.View(_viewDefine)
	if err == nil {
		fmt.Fprintln(logView, str)
	}

	return err
}

func enterClipboard(g *gocui.Gui, v *gocui.View) error {
	myclipboard.Set(copyInfo)

	maxX, _ := g.Size()
	modal := mycui.NewModal(g, division(maxX, uiPart[0])+5, 10, division(maxX, uiPart[0])+35).
		SetTextColor(gocui.ColorRed).SetText("copy success \n 已 复 制 到 剪 切 板 ")
	modal.Mouse = true
	//	modal.SetBgColor(gocui.ColorRed)
	_handle := func(g *gocui.Gui, v *gocui.View) error {
		modal.Close()
		return nil
	}
	modal.AddButton("ok", "OK", gocui.KeyEnter, _handle).AddHandler(gocui.MouseLeft, _handle)
	modal.Draw()

	return nil
}

func enterRun(g *gocui.Gui, v *gocui.View) error {
	setlog(g, "run .... ing")
	generate(g, v)
	return nil
}

func enterSet(g *gocui.Gui, v *gocui.View) error {
	maxX, _ := g.Size()
	setlog(g, "")
	// new form
	form = mycui.NewForm(g, "set_ui", "Sign Up", division(maxX, uiPart[0])+2, 2, 0, 0)

	// add input field
	form.AddInputField("out_dir", SLocalize("out_dir"), formPart[0], formPart[1]).SetText(config.GetOutDir()).
		AddValidate("required input", requireValidator)
	form.AddInputField("db_host", SLocalize("db_host"), formPart[0], formPart[1]).SetText(config.GetDbInfo().Host).
		AddValidate("required input", requireValidator)
	form.AddInputField("db_port", SLocalize("db_port"), formPart[0], formPart[1]).SetText(tools.AsString(config.GetDbInfo().Port)).
		AddValidate("required input", requireValidator)
	form.AddInputField("db_usename", SLocalize("db_usename"), formPart[0], formPart[1]).SetText(config.GetDbInfo().Username).
		AddValidate("required input", requireValidator)
	form.AddInputField("db_pwd", SLocalize("db_pwd"), formPart[0], formPart[1]).SetText(config.GetDbInfo().Password).
		SetMask().SetMaskKeybinding(gocui.KeyCtrlA).
		AddValidate("required input", requireValidator)
	form.AddInputField("db_name", SLocalize("db_name"), formPart[0], formPart[1]).SetText(config.GetDbInfo().Database).
		AddValidate("required input", requireValidator)
	form.AddSelect("db_type", SLocalize("db_type"), formPart[0], formPart[2]).AddOptions(getDBTypeList()...).
		SetSelected(GetDBTypeStr(config.GetDbInfo().Type))
	// add select
	form.AddSelect("is_dev", SLocalize("is_dev"), formPart[0], formPart[2]).
		AddOptions(SLocalize("true"), SLocalize("false")).SetSelected(SLocalize(tools.AsString(config.GetIsDev())))
	form.AddSelect("is_simple", SLocalize("is_simple"), formPart[0], formPart[2]).
		AddOptions(SLocalize("true"), SLocalize("false")).SetSelected(SLocalize(tools.AsString(config.GetSimple())))
	form.AddSelect("is_out_sql", SLocalize("is_out_sql"), formPart[0], formPart[2]).
		AddOptions(SLocalize("true"), SLocalize("false")).SetSelected(SLocalize(tools.AsString(config.GetIsOutSQL())))
	form.AddSelect("is_out_func", SLocalize("is_out_func"), formPart[0], formPart[2]).
		AddOptions(SLocalize("true"), SLocalize("false")).SetSelected(SLocalize(tools.AsString(config.GetIsOutFunc())))
	form.AddSelect("is_foreign_key", SLocalize("is_foreign_key"), formPart[0], formPart[2]).
		AddOptions(SLocalize("true"), SLocalize("false")).SetSelected(SLocalize(tools.AsString(config.GetIsForeignKey())))
	form.AddSelect("is_gui", SLocalize("is_gui"), formPart[0], formPart[2]).
		AddOptions(SLocalize("true"), SLocalize("false")).SetSelected(SLocalize(tools.AsString(config.GetIsGUI())))
	form.AddSelect("is_table_name", SLocalize("is_table_name"), formPart[0], formPart[2]).
		AddOptions(SLocalize("true"), SLocalize("false")).SetSelected(SLocalize(tools.AsString(config.GetIsTableName())))
	form.AddSelect("url_tag", SLocalize("url_tag"), formPart[0], formPart[2]).
		AddOptions("json", "url").SetSelected(tools.AsString(config.GetURLTag()))
	form.AddSelect("db_tag", SLocalize("db_tag"), formPart[0], formPart[2]).
		AddOptions("gorm", "db").SetSelected(config.GetDBTag())
	form.AddSelect("language", SLocalize("language"), formPart[0], formPart[2]).
		AddOptions("English", "中 文").SetSelected(config.GetLG())

	// add button
	form.AddButton("save", SLocalize("save"), buttonSave).AddHandler(gocui.MouseLeft, buttonSave)
	form.AddButton("cancel", SLocalize("cancel"), buttonCancel).AddHandler(gocui.MouseLeft, buttonCancel)
	form.AddButton("about", SLocalize("about"), about).AddHandler(gocui.MouseLeft, about)

	form.Draw()

	return nil
}

func buttonCancel(g *gocui.Gui, v *gocui.View) error {
	menuFocusButton(g)
	if form != nil {
		err := form.Close(g, nil)
		form = nil
		return err
	}
	return nil
}

func buttonSave(g *gocui.Gui, v *gocui.View) error {
	maxX, _ := g.Size()
	mp := form.GetFieldTexts()
	config.SetOutDir(mp["out_dir"])

	var dbInfo config.DBInfo
	dbInfo.Host = mp["db_host"]
	port, err := strconv.Atoi(mp["db_port"])
	if err != nil {
		modal := mycui.NewModal(g, division(maxX, uiPart[0])+5, 10, division(maxX, uiPart[0])+35).SetTextColor(gocui.ColorRed).SetText("port error")

		_handle := func(g *gocui.Gui, v *gocui.View) error {
			modal.Close()
			form.SetCurrentItem(form.GetCurrentItem())
			return nil
		}
		//	modal.SetBgColor(gocui.ColorRed)
		modal.AddButton("ok", "OK", gocui.KeyEnter, _handle).AddHandler(gocui.MouseLeft, _handle)

		modal.Draw()
		return nil
	}

	dbInfo.Port = port
	dbInfo.Username = mp["db_usename"]
	dbInfo.Password = mp["db_pwd"]
	dbInfo.Database = mp["db_name"]
	mp = form.GetSelectedOpts()

	dbInfo.Type = GetDBTypeID(mp["db_type"])
	config.SetMysqlDbInfo(&dbInfo)

	config.SetIsDev(getBool(mp["is_dev"]))
	config.SetSimple(getBool(mp["is_simple"]))
	config.SetIsOutSQL(getBool(mp["is_out_sql"]))
	config.SetIsOutFunc(getBool(mp["is_out_func"]))
	config.SetForeignKey(getBool(mp["is_foreign_key"]))
	config.SetIsGUI(getBool(mp["is_gui"]))
	config.SetIsTableName(getBool(mp["is_table_name"]))
	config.SetURLTag(mp["url_tag"])
	config.SetDBTag(mp["db_tag"])
	config.SetLG(mp["language"])

	config.SaveToFile()
	modal := mycui.NewModal(g, division(maxX, uiPart[0])+5, 10, division(maxX, uiPart[0])+35).SetText("save success")
	_handle := func(g *gocui.Gui, v *gocui.View) error {
		modal.Close()
		buttonCancel(g, v)
		return nil
	}
	modal.AddButton("ok", "OK", gocui.KeyEnter, _handle).AddHandler(gocui.MouseLeft, _handle)
	modal.Draw()

	return nil
}

func buildList(g *gocui.Gui, v *gocui.View) error {
	listView, err := g.View(_listDefine)
	if err != nil {
		panic(err)
	}

	listView.Clear()
	for _, info := range gPkg.Structs {
		fmt.Fprintln(listView, info.Name)
	}

	return nil
}

func showStruct(g *gocui.Gui, v *gocui.View) error {
	var l string
	var err error

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}

	var out, out1 []string
	for _, v := range gPkg.Structs {
		if v.Name == l {
			out = v.GeneratesColor()
			out1 = v.Generates()
			break
		}
	}

	setlog(g, "\n\n\n")
	for _, v := range out {
		addlog(g, v)
	}

	copyInfo = strings.Join(out1, "\n")

	return nil
}

func listDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if cy < len(gPkg.Structs)-1 {
			if err := v.SetCursor(cx, cy+1); err != nil {
				ox, oy := v.Origin()
				if err := v.SetOrigin(ox, oy+1); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func listUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

///////////////////////////////////////////

// OnInitDialog init main loop
func OnInitDialog() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	g.ASCII = true

	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = false // 光标
	g.Mouse = true
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen // 设置边框颜色

	mainLayout(g)
	//g.SetManagerFunc(mainLayout) // 主布局
	keybindings(g)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit { // 主循环
		log.Panicln(err)
	}
}

// GetDBTypeStr 0:mysql , 1:sqlite , 2:mssql
func GetDBTypeStr(tp int) string {
	switch tp {
	case 0:
		return "mysql"
	case 1:
		return "sqlite"
	case 2:
		return "mssql"
	}
	// default
	return "mysql"
}

// GetDBTypeID 0:mysql , 1:sqlite , 2:mssql
func GetDBTypeID(name string) int {
	switch name {
	case "mysql":
		return 0
	case "sqlite":
		return 1
	case "mssql":
		return 2
	}
	// default
	return 0
}

func getDBTypeList() []string {
	return []string{"mysql", "sqlite", "mssql"}
}
