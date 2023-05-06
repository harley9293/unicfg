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
