package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var s = `{"resultcode":"200","reason":"查询成功","result":{"Country":"","Province":"","City":"内网IP","Isp":"内网IP"},"error_code":0}`
	m := unmarshalMap([]byte(s))
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func unmarshalMap(data []byte) map[string]interface{} {
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return nil
	}
	fmt.Println(m)
	return m
}
