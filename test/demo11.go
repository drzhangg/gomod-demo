package main

import (
	"encoding/json"
	"fmt"
)

type Test2 struct {
	Timestamp string `json:"timestamp"`
	Value     float64
}

func main() {
	a := `[{"timestamp":"1583132460000","Value":1.06462E7}]`
	test := []Test2{}
	json.Unmarshal([]byte(a),&test)
	fmt.Println(test)

	var m map[string]interface{}
	for _,v := range test{
		j,_:=json.Marshal(v)
		json.Unmarshal(j,&m)
	}

	fmt.Println(m)
}
