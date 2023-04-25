package unicfg

import (
	"errors"
	"strings"
)

type Loader interface {
	load(string) (*Elem, error)
}

type emptyLoader struct {
}

func (e emptyLoader) load(configPath string) (*Elem, error) {
	return nil, errors.New("unsupported config file type")
}

func New(configPath string) (*Elem, error) {
	var loader Loader
	if strings.HasSuffix(configPath, ".ini") {
		loader = iniLoader{}
	} else if strings.HasSuffix(configPath, ".json") {
		loader = jsonLoader{}
	} else {
		loader = emptyLoader{}
	}

	return loader.load(configPath)
}
