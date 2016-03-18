package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Config map[string]interface{}

func (c Config) Name() string {
	return c["name"].(string)
	
}

func main() {
	config, err := LoadConfig("config.json")

	if err != nil {
		panic(err)
	}

	fmt.Println(config)
	fmt.Println(config.Name())

}

func LoadConfig(path string) (Config, error) {

	var m map[string]interface{}
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return m, err
	}

	err = json.Unmarshal(data, &m)

	return m, err

}
