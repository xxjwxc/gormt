package dlg

import "github.com/jroimartin/gocui"

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
