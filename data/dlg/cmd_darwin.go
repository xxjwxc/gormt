package dlg

import "os/exec"

func openURL(url string) {
	exec.Command(`open`, url).Start()
}
