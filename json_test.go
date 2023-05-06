package unicfg

import (
	"io/ioutil"
	"os"
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
