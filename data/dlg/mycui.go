package dlg

import (
	"runtime"

	"github.com/xxjwxc/public/myi18n"
)

// WinMain windows main loop
func WinMain() {
	if runtime.GOOS == "windows" {
		myi18n.SetLocalLG("en")
	}
	OnInitDialog()
}
