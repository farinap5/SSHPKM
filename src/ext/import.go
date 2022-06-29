package ext

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func Module(name string) {
	cont, e := readfile(name)
	if e {
		return
	}

	data := make(map[string]interface{})
	err := yaml.Unmarshal(cont, &data)
	if err != nil {
		fmt.Println("[\u001B[1;31m!\u001B[0;0m] " + err.Error())
		return
	}

	//H := data["Host"]
	// U := data["User"]
	for key, val := range data {
		s := fmt.Sprintf("%s=%s", key, val)
		fmt.Println(s)
	}
}

func readfile(name string) ([]byte, bool) {
	body, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("[\u001B[1;31m!\u001B[0;0m] " + err.Error())
		return []byte(""), true
	}
	return body, false
}
