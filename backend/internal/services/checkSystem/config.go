package checkSystem

import (
	"errors"
	"fmt"
	"os/exec"
)

const FileName = "./tests"

var configs = map[string]config{
	"go":     &goConfig{},
	"python": &pythonConfig{},
}

type config interface {
	getExtension() string
	getCMD() *exec.Cmd
}

func newConfig(lang string) (config, error) {
	config, found := configs[lang]
	if !found {
		errMessage := fmt.Sprintf("Config by language %s not found", lang)
		return nil, errors.New(errMessage)
	}

	return config, nil
}
