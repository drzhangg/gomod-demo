package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var s = `{"resultcode":"200","reason":"查询成功","result":{"Country":"","Province":"","City":"内网IP","Isp":"内网IP"},"error_code":0}`
	m := unmarshalMap([]byte(s))

	mr := make(map[string]interface{})
	for k, _ := range m {
		if k == "result" {
			mr[k] = m[k]
		}
		//fmt.Println(k, v)
	}
	fmt.Println(mr)
}

func unmarshalMap(data []byte) map[string]interface{} {
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return nil
	}
	//fmt.Println(m)
	return m
}
