package main

import (
	"encoding/json"
	"fmt"
)

type shape struct {
	Height int `json:"height"`
	Width  int `Json:"width"`
}

func main() {
	s := []byte(`[{"height": 24, "width": 12},{"height": 10, "width": 20}]`)
	var shapes []shape
	err := json.Unmarshal(s, &shapes)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Printf("%+v\n", shapes)
}
