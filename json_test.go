package unicfg

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

type PersonTestUnicfg struct {
	Name    string             `unicfg:"name"`
	Age     int                `unicfg:"age"`
	Address map[string]string  `unicfg:"address"`
	Family  []PersonTestUnicfg `unicfg:"family"`
}

type PersonTestJson struct {
	Name    string            `json:"name"`
	Age     int               `json:"age"`
	Address map[string]string `json:"address"`
	Family  []PersonTestJson  `json:"family"`
}

func CreateJsonFile() {
	person1 := PersonTestJson{
		Name:    "person1",
		Age:     10,
		Address: nil,
		Family:  nil,
	}

	person2 := PersonTestJson{
		Name:    "person2",
		Age:     20,
		Address: nil,
		Family:  nil,
	}

	person3 := PersonTestJson{
		Name: "person3",
		Age:  30,
		Address: map[string]string{
			"country": "china",
		},
		Family: []PersonTestJson{person1, person2},
	}
	jsonData, _ := json.Marshal(person3)

	ioutil.WriteFile("test.json", jsonData, 0644)
}

func RemoveJsonFile() {
	os.Remove("test.json")
}

func Test_jsonLoader_load_fail(t *testing.T) {
	_, err := jsonLoader{}.load("fail.json")
	if err == nil {
		t.Error("expected error, but got nil")
	}
}

func Test_jsonLoader_load_fail2(t *testing.T) {
	CreateJsonFile()
	defer RemoveJsonFile()

	// append data to test.json file
	f, err := os.OpenFile("test.json", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		t.Errorf("expected nil, but got error, err: %s", err)
	}
	defer f.Close()
	if _, err := f.WriteString("}"); err != nil {
		t.Errorf("expected nil, but got error, err: %s", err)
	}

	_, err = jsonLoader{}.load("test.json")
	if err == nil {
		t.Error("expected error, but got nil")
	}
}
