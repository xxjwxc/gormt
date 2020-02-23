package dlg

import "os/exec"

func openURL(url string) {
	exec.Command(`xdg-open`, url).Start()
}
