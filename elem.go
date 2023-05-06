package unicfg

import (
	"strconv"
	"strings"
)

type Elem struct {
	value string
	child map[string]*Elem
	next  *Elem
}

func (e *Elem) Key(key string) *Elem {
	keys := strings.Split(key, ".")
	if len(keys) > 1 {
		return e.Key(keys[0]).Key(strings.Join(keys[1:], "."))
	} else {
		v, ok := e.child[key]
		if !ok {
			return newElem("")
		}

		return v
	}
}

func (e *Elem) Next() *Elem {
	return e.next
}

func (e *Elem) Children() map[string]*Elem {
	return e.child
}

func (e *Elem) String() string {
	return e.value
}

func (e *Elem) MustString(defaultValue string) string {
	if e.value == "" {
		return defaultValue
	}
	return e.value
}

func (e *Elem) Int64() int64 {
	return e.MustInt64(0)
}

func (e *Elem) MustInt64(defaultValue int64) int64 {
	v, err := strconv.ParseInt(e.value, 0, 64)
	if err != nil {
		return defaultValue
	}
	return v
}

func (e *Elem) Int() int {
	return int(e.Int64())
}

func (e *Elem) MustInt(defaultValue int) int {
	return int(e.MustInt64(int64(defaultValue)))
}

func (e *Elem) Bool() bool {
	return e.MustBool(false)
}

func (e *Elem) MustBool(defaultValue bool) bool {
	v, err := strconv.ParseBool(e.value)
	if err != nil {
		return defaultValue
	}
	return v
}

func newElem(v string) *Elem {
	return &Elem{value: v, child: map[string]*Elem{}}
}
