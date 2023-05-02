package checkSystem

import "os/exec"

type pythonConfig struct{}

func (c *pythonConfig) getExtension() string {
	return "py"
}

func (c *pythonConfig) getCMD() *exec.Cmd {
	return exec.Command("python", "-v")
}
