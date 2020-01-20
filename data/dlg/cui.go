package dlg

import (
	"fmt"
	"log"

	"github.com/xxjwxc/public/tools"

	"github.com/xxjwxc/gormt/data/config"

	"github.com/jroimartin/gocui"
	"github.com/xxjwxc/public/mycui"
)

func nextView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (mainIndex + 1) % len(mainViewArr)
	name := mainViewArr[nextIndex]

	err := setlog(g, "Going from view "+v.Name()+" to "+name)
	if err != nil {
		return err
	}
	if _, err := g.SetCurrentView(name); err != nil { // 设置选中
		return err
	}
	g.SelFgColor = gocui.ColorGreen // 设置边框颜色
	g.FgColor = gocui.ColorWhite

	switch name {
	case _menuDefine:
		g.Cursor = false // 光标
		// g.FgColor = gocui.ColorGreen
		menuDlg.btnList[menuDlg.active].Focus()
	case _listDefine:
		g.Cursor = false
		menuDlg.btnList[menuDlg.active].UnFocus()
	case _viewDefine:
		g.Cursor = true
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
		v.Wrap = true
		v.Autoscroll = true
		if _, err := g.SetCurrentView(_menuDefine); err != nil {
			return err
		}
	}

	if v, err := g.SetView(_viewDefine, division(maxX, uiPart[0]), 1, maxX-1, maxY-1); err != nil {
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
	if err := g.SetKeybinding("main_title", gocui.MouseLeft, gocui.ModNone, about); err != nil {
		log.Panicln(err)
	}
	// if err := g.SetKeybinding(_run, gocui.MouseLeft, gocui.ModNone, about); err != nil {
	// 	log.Panicln(err)
	// }
	// if err := g.SetKeybinding(_set, gocui.MouseLeft, gocui.ModNone, about); err != nil {
	// 	log.Panicln(err)
	// }
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

func enterRun(g *gocui.Gui, v *gocui.View) error {
	setlog(g, "run .... ing")
	return nil
}

func enterSet(g *gocui.Gui, v *gocui.View) error {
	maxX, _ := g.Size()
	setlog(g, "")
	// new form
	form = mycui.NewForm(g, "set_ui", "Sign Up", division(maxX, uiPart[0])+3, 3, 0, 0)

	// add input field
	form.AddInputField("out_dir", SLocalize("out_dir"), formPart[0], formPart[1]).SetText(config.GetOutDir()).
		AddValidate("required input", requireValidator)
	form.AddInputField("db_host", SLocalize("db_host"), formPart[0], formPart[1]).SetText(config.GetMysqlDbInfo().Host).
		AddValidate("required input", requireValidator)
	form.AddInputField("db_port", SLocalize("db_port"), formPart[0], formPart[1]).SetText(tools.AsString(config.GetMysqlDbInfo().Port)).
		AddValidate("required input", requireValidator)
	form.AddInputField("db_usename", SLocalize("db_usename"), formPart[0], formPart[1]).SetText(config.GetMysqlDbInfo().Username).
		AddValidate("required input", requireValidator)
	form.AddInputField("db_pwd", SLocalize("db_pwd"), formPart[0], formPart[1]).SetText(config.GetMysqlDbInfo().Password).
		SetMask().SetMaskKeybinding(gocui.KeyCtrlA).
		AddValidate("required input", requireValidator)
	form.AddInputField("db_name", SLocalize("db_name"), formPart[0], formPart[1]).SetText(config.GetMysqlDbInfo().Database).
		AddValidate("required input", requireValidator)

	// add select
	form.AddSelect("is_dev", SLocalize("is_dev"), formPart[0], formPart[2]).
		AddOptions(SLocalize("true"), SLocalize("false")).SetSelected(SLocalize(tools.AsString(config.GetIsDev())))
	form.AddSelect("is_simple", SLocalize("is_simple"), formPart[0], formPart[2]).
		AddOptions(SLocalize("true"), SLocalize("false")).SetSelected(SLocalize(tools.AsString(config.GetSimple())))
	form.AddSelect("is_singular", SLocalize("is_singular"), formPart[0], formPart[2]).
		AddOptions(SLocalize("true"), SLocalize("false")).SetSelected(SLocalize(tools.AsString(config.GetSingularTable())))
	form.AddSelect("is_out_sql", SLocalize("is_out_sql"), formPart[0], formPart[2]).
		AddOptions(SLocalize("true"), SLocalize("false")).SetSelected(SLocalize(tools.AsString(config.GetIsOutSQL())))
	form.AddSelect("is_out_func", SLocalize("is_out_func"), formPart[0], formPart[2]).
		AddOptions(SLocalize("true"), SLocalize("false")).SetSelected(SLocalize(tools.AsString(config.GetIsOutFunc())))
	form.AddSelect("is_foreign_key", SLocalize("is_foreign_key"), formPart[0], formPart[2]).
		AddOptions(SLocalize("true"), SLocalize("false")).SetSelected(SLocalize(tools.AsString(config.GetIsForeignKey())))
	form.AddSelect("url_tag", SLocalize("url_tag"), formPart[0], formPart[2]).
		AddOptions("json", "url").SetSelected(tools.AsString(config.GetURLTag()))
	form.AddSelect("db_tag", SLocalize("db_tag"), formPart[0], formPart[2]).
		AddOptions("gorm", "db").SetSelected(config.GetDBTag())
	form.AddSelect("language", SLocalize("language"), formPart[0], formPart[2]).
		AddOptions("English", "中 文").SetSelected(config.GetLG())

	// add button
	form.AddButton("save", SLocalize("save"), buttonSave).AddHandler(gocui.MouseLeft, buttonSave)
	form.AddButton("cancel", SLocalize("cancel"), buttonCancel).AddHandler(gocui.MouseLeft, buttonCancel)

	form.Draw()

	return nil
}

func buttonCancel(g *gocui.Gui, v *gocui.View) error {
	menuFocusButton(g)
	if form != nil {
		return form.Close(g, nil)
	}
	return nil
}

func buttonSave(g *gocui.Gui, v *gocui.View) error {
	return nil
}

///////////////////////////////////////////

// OnInitDialog init main loop
func OnInitDialog() {
	g, err := gocui.NewGui(gocui.OutputNormal)

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
