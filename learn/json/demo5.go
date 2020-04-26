package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	js := "{\"addresses\":{\"db83e561-1939-4f1e-9f9d-e98c79e1171f\":[{\"addr\":\"10.10.30.57\"}]}}"

	jsm := make(map[string]interface{})
	err := json.Unmarshal([]byte(js), &jsm)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(jsm["addresses"])
	d := jsm["addresses"].(map[string]interface{})

	var key string
	for k, _ := range d {
		key = k

	}
	addr := d[key].([]interface{})

	for _,v := range addr{
		add := v.(map[string]interface{})
		fmt.Println(add["addr"])
	}

}
