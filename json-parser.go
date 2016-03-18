package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

func main() {
	config, err := LoadConfig("config.json")

	if err != nil {
		panic(err)
	}

	fmt.Println(config)


}

func LoadConfig(path string) (map[string]interface{}, error) {

	var m map[string]interface{}
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return m, err
	}

	err = json.Unmarshal(data, &m)

	return m, err

}
