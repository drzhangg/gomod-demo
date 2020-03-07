package main

import (
	"encoding/json"
	"fmt"
)

type Demo struct {
	User string `json:"user"`
}

func main() {
	var u Demo
	bytes, _:=json.Marshal(u)
	fmt.Println(string(bytes))
}
