package dlg

import "os/exec"

func openURL(url string) {
	exec.Command(`cmd`, `/c`, `start`, url).Start() // 有GUI调用
}
