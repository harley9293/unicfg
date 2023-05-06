package unicfg

import (
	"io/ioutil"
	"os"
)

func CreateIniFile() {
	iniString := ""
	iniString += "str_field = hello world!\n"
	iniString += "int_field = 1234\n"
	iniString += "int64_field = 22222222222\n"
	iniString += "bool_field = true\n"
	iniString += "[sub_field]\n"
	iniString += "sub_str_field = hello world!\n"
	iniString += "sub_int_field = 1234\n"
	iniString += "sub_int64_field = 22222222222\n"
	iniString += "sub_bool_field = true\n"

	ioutil.WriteFile("test.ini", []byte(iniString), 0644)
}

func RemoveIniFile() {
	os.Remove("test.ini")
}
