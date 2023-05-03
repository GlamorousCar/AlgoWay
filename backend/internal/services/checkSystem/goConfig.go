package checkSystem

import "os/exec"

type goConfig struct{}

func (c *goConfig) getExtension() string {
	return "go"
}

func (c *goConfig) getCMD() *exec.Cmd {
	return exec.Command("go", "run", filePath)
}
