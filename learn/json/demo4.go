package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s := `{"name":"jerry","score":59.0}`

	user := struct {
		Name  string  `json:"name"`
		Score float64 `json:"score"`
	}{}

	json.Unmarshal([]byte(s), &user)
	fmt.Print(user)
}
