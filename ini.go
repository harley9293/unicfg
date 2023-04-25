package unicfg

import (
	"gopkg.in/ini.v1"
	"strings"
)

type iniLoader struct {
}

func (i iniLoader) load(configPath string) (*Elem, error) {
	f, err := ini.Load(configPath)
	if err != nil {
		return nil, err
	}

	r := newElem("")
	i.parseSection(r, f.Sections())

	return r, nil
}

func (i iniLoader) parseSection(e *Elem, sections []*ini.Section) {
	for _, section := range sections {
		if section.Name() == "DEFAULT" {
			for _, key := range section.Keys() {
				e.child[key.Name()] = i.parseValue(key.Value())
			}
		} else {
			e.child[section.Name()] = newElem("")
			i.parseSection(e.child[section.Name()], section.ChildSections())

			for _, key := range section.Keys() {
				e.child[section.Name()].child[key.Name()] = i.parseValue(key.Value())
			}
		}
	}
}

func (i iniLoader) parseValue(value string) *Elem {
	startElem := newElem(value)
	splitStr := strings.Split(value, ",")
	if len(splitStr) > 1 {

		e := startElem
		for _, v := range splitStr {
			e.next = newElem(v)
			e = e.next
		}
		return startElem
	}
	return startElem
}
