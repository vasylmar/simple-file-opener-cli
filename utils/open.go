package utils

import "os/exec"

func OpenFile(path string) error {
	cmd := exec.Command("open", path)
	return cmd.Run()
}
