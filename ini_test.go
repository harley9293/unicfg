package unicfg

import (
	"io/ioutil"
	"os"
	"testing"
)

func CreateIniFile() {
	iniString := ""
	iniString += "str_field = hello world!\n"
	iniString += "int_field = 1234\n"
	iniString += "int64_field = 22222222222\n"
	iniString += "bool_field = true\n"
	iniString += "[sub_field]\n"
	iniString += "sub_array_field = hello,world,!\n"

	ioutil.WriteFile("test.ini", []byte(iniString), 0644)
}

func RemoveIniFile() {
	os.Remove("test.ini")
}

func Test_iniLoader_load_fail(t *testing.T) {
	_, err := iniLoader{}.load("fail.ini")
	if err == nil {
		t.Error("expected error, but got nil")
	}
}
