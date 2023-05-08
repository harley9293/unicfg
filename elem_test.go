package unicfg

import "testing"

func TestElem_Int(t *testing.T) {
	e := newElem("123")
	if e.Int() != 123 {
		t.Errorf("expected 123, but got %d", e.Int())
	}
}

func TestElem_MustInt(t *testing.T) {
	e := newElem("123")
	if e.MustInt(0) != 123 {
		t.Errorf("expected 123, but got %d", e.MustInt(0))
	}
}

func TestElem_Int64(t *testing.T) {
	e := newElem("123")
	if e.Int64() != 123 {
		t.Errorf("expected 123, but got %d", e.Int64())
	}
}

func TestElem_MustInt64_Nil(t *testing.T) {
	e := newElem("")
	if e.MustInt64(123) != 123 {
		t.Errorf("expected 123, but got %d", e.MustInt64(123))
	}
}

func TestElem_MustInt64(t *testing.T) {
	e := newElem("123")
	if e.MustInt64(0) != 123 {
		t.Errorf("expected 123, but got %d", e.MustInt64(0))
	}
}

func TestElem_String(t *testing.T) {
	e := newElem("123")
	if e.String() != "123" {
		t.Errorf("expected 123, but got %s", e.String())
	}
}

func TestElem_MustString(t *testing.T) {
	e := newElem("123")
	if e.MustString("0") != "123" {
		t.Errorf("expected 123, but got %s", e.MustString("0"))
	}
}

func TestElem_MustString_Nil(t *testing.T) {
	e := newElem("")
	if e.MustString("0") != "0" {
		t.Errorf("expected 0, but got %s", e.MustString("0"))
	}
}

func TestElem_Bool(t *testing.T) {
	e := newElem("true")
	if !e.Bool() {
		t.Errorf("expected true, but got false")
	}
}

func TestElem_MustBool(t *testing.T) {
	e := newElem("true")
	if !e.MustBool(false) {
		t.Errorf("expected true, but got false")
	}
}

func TestElem_MustBool_Nil(t *testing.T) {
	e := newElem("")
	if e.MustBool(true) != true {
		t.Errorf("expected true, but got %t", e.MustBool(true))
	}
}

func TestElem_Float(t *testing.T) {
	e := newElem("123.456")
	if e.Float() != 123.456 {
		t.Errorf("expected 123.456, but got %f", e.Float())
	}
}

func TestElem_MustFloat(t *testing.T) {
	e := newElem("123.456")
	if e.MustFloat(0) != 123.456 {
		t.Errorf("expected 123.456, but got %f", e.MustFloat(0))
	}
}

func TestElem_MustFloat_Nil(t *testing.T) {
	e := newElem("")
	if e.MustFloat(123.456) != 123.456 {
		t.Errorf("expected 123.456, but got %f", e.MustFloat(123.456))
	}
}

func TestElem_Next(t *testing.T) {
	e := newElem("123")
	e.next = newElem("456")
	if e.Next().String() != "456" {
		t.Errorf("expected 456, but got %s", e.Next().String())
	}
}

func TestElem_Children(t *testing.T) {
	e := newElem("123")
	e.child["456"] = newElem("456")
	if e.Children()["456"].String() != "456" {
		t.Errorf("expected 456, but got %s", e.Children()["456"].String())
	}
}

func TestElem_Key(t *testing.T) {
	e := newElem("123")
	e.child["456"] = newElem("456")
	e.child["456"].child["789"] = newElem("789")
	if e.Key("456").String() != "456" {
		t.Errorf("expected 456, but got %s", e.Key("456").String())
	}

	if e.Key("456.789").String() != "789" {
		t.Errorf("expected 789, but got %s", e.Key("456.789").String())
	}

	if e.Key("456.789.10").String() != "" {
		t.Errorf("expected empty, but got %s", e.Key("456.789.10").String())
	}
}

func TestElem_Len(t *testing.T) {
	e := newElem("123")
	e.next = newElem("456")
	e.next.next = newElem("789")
	if e.Len() != 2 {
		t.Errorf("expected 2, but got %d", e.Len())
	}
}
