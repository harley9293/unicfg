package unicfg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type jsonLoader struct {
}

func (j jsonLoader) load(configPath string) (*Elem, error) {
	byteValue, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var jsonData any
	err = json.Unmarshal(byteValue, &jsonData)
	if err != nil {
		return nil, err
	}

	return j.parseValue(jsonData), nil
}

func (j jsonLoader) parseValue(v any) *Elem {
	startElem := newElem("")
	switch v.(type) {
	case map[string]any:
		for key, value := range v.(map[string]any) {
			startElem.child[key] = j.parseValue(value)
		}
	case []any:
		e := startElem
		for _, value := range v.([]any) {
			e.next = j.parseValue(value)
			e = e.next
		}
	default:
		startElem.value = fmt.Sprintf("%v", v)
	}
	return startElem
}
