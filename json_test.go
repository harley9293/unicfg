package unicfg

import (
	"io/ioutil"
	"os"
	"testing"
)

func CreateJsonFile() {
	jsonString := ""
	jsonString += "{\n"
	jsonString += "  \"str_field\": \"hello world!\",\n"
	jsonString += "  \"int_field\": 1234,\n"
	jsonString += "  \"int64_field\": \"22222222222\",\n"
	jsonString += "  \"bool_field\": true,\n"
	jsonString += "  \"array_field\": [\n"
	jsonString += "    \"hello\",\n"
	jsonString += "    \"world\",\n"
	jsonString += "    \"!\"\n"
	jsonString += "  ]\n"
	jsonString += "}\n"

	ioutil.WriteFile("test.json", []byte(jsonString), 0644)
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
