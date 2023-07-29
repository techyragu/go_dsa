package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// b, _ := json.Marshal(true)
	// fmt.Println(string(b))

	// slcD := []string{"apple", "peach", "pear"}
	// slcB, _ := json.Marshal(slcD)
	// fmt.Println(slcB)

	// mapD := map[string]int{"apple": 5, "lettuce": 7}
	// mapB, _ := json.Marshal(mapD)
	// fmt.Println(string(mapB))

	type response1 struct {
		Page   int
		Fruits []string
	}

	res1D := response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	type response2 struct {
		Page   int      `json:"page"`
		Fruits []string `json:"fruits"`
	}

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))
}
