package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/robertlestak/procx-go/pkg/procx"
)

type ExampleData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Work string `json:"work"`
}

func main() {
	args := []string{
		"-driver",
		"redis-list",
		"-redis-host",
		"localhost",
		"-redis-port",
		"6379",
		"-redis-key",
		"test-key",
	}
	data, err := procx.Procx(args)
	if err != nil && err != procx.ErrNoData {
		panic(err)
	} else if err == procx.ErrNoData {
		println("no data")
		os.Exit(0)
	}
	bd, err := ioutil.ReadAll(data)
	if err != nil {
		panic(err)
	}
	var exampleData ExampleData
	err = json.Unmarshal(bd, &exampleData)
	if err != nil {
		panic(err)
	}
	println(exampleData.Name)
}
