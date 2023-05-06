package unicfg

import "testing"

func TestNew_Empty(t *testing.T) {
	_, err := New("")
	if err == nil {
		t.Error("expected error, but got nil")
	}
}

func TestNew_INI(t *testing.T) {
	CreateIniFile()
	defer RemoveIniFile()
	_, err := New("test.ini")
	if err != nil {
		t.Errorf("expected nil, but got error, err: %s", err)
	}
}

func TestNew_JSON(t *testing.T) {
	CreateJsonFile()
	defer RemoveJsonFile()
	_, err := New("test.json")
	if err != nil {
		t.Errorf("expected nil, but got error, err: %s", err)
	}
}
