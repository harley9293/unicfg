package unicfg

import (
	"reflect"
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

func (e *Elem) Parse(st any) {
	stValue := reflect.ValueOf(st)
	e.parse(stValue)
}

func (e *Elem) Next() *Elem {
	return e.next
}

func (e *Elem) Len() int {
	l := 0
	p := e
	for p.Next() != nil {
		l++
		p = p.Next()
	}
	return l
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

func (e *Elem) Float() float64 {
	return e.MustFloat(0)
}

func (e *Elem) MustFloat(defaultValue float64) float64 {
	v, err := strconv.ParseFloat(e.value, 64)
	if err != nil {
		return defaultValue
	}
	return v
}

func newElem(v string) *Elem {
	return &Elem{value: v, child: map[string]*Elem{}}
}

func (e *Elem) parse(v reflect.Value) {
	if v.Type().Kind() == reflect.Pointer {
		v = v.Elem()
	}

	switch v.Type().Kind() {
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			tag, ok := v.Type().Field(i).Tag.Lookup("unicfg")
			if !ok {
				continue
			}
			e.Key(tag).parse(v.Field(i))
		}
	case reflect.Map:
		if v.IsNil() {
			v.Set(reflect.MakeMap(v.Type()))
		}

		for name, child := range e.Children() {
			newValue := reflect.New(v.Type().Elem())
			child.parse(newValue)
			v.SetMapIndex(reflect.ValueOf(name), newValue.Elem())
		}
	case reflect.Array, reflect.Slice:
		if v.IsNil() {
			v.Set(reflect.New(v.Type()).Elem())
		}
		tmp := reflect.New(v.Type()).Elem()
		for n := e.Next(); n != nil; n = n.Next() {
			newValue := reflect.New(v.Type().Elem())
			n.parse(newValue)
			tmp = reflect.Append(tmp, newValue.Elem())
		}
		v.Set(tmp)
	case reflect.String:
		v.SetString(e.String())
	case reflect.Int, reflect.Int64:
		v.SetInt(e.Int64())
	case reflect.Bool:
		v.SetBool(e.Bool())
	}
}
