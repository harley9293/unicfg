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

func TestParse_Fail(t *testing.T) {
	err := Parse("", nil)
	if err == nil {
		t.Error("expected error, but got nil")
	}
}

func TestParse_Success(t *testing.T) {
	CreateJsonFile()
	defer RemoveJsonFile()

	person := PersonTestUnicfg{}
	err := Parse("test.json", &person)
	if err != nil {
		t.Errorf("expected nil, but got error, err: %s", err)
	}

	if person.Name != "person3" {
		t.Errorf("expected person3, but got %s", person.Name)
	}

	if person.Age != 30 {
		t.Errorf("expected 30, but got %d", person.Age)
	}

	if len(person.Address) != 1 {
		t.Errorf("expected address size 1, but got %d", len(person.Address))
	}

	if person.Address["country"] != "china" {
		t.Errorf("expected address country china, but got %s", person.Address["country"])
	}

	if len(person.Family) != 2 {
		t.Errorf("expected family size 2, but got %d", len(person.Family))
	}

	if person.Family[0].Name != "person1" {
		t.Errorf("expected family[0].name person1, but got %s", person.Family[0].Name)
	}

	if person.Family[0].Age != 10 {
		t.Errorf("expected family[0].age 10, but got %d", person.Family[0].Age)
	}

	if person.Family[1].Name != "person2" {
		t.Errorf("expected family[1].name person2, but got %s", person.Family[1].Name)
	}

	if person.Family[1].Age != 20 {
		t.Errorf("expected family[1].age 20, but got %d", person.Family[1].Age)
	}
}
