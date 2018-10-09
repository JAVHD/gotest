package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	apiData := make(map[string]interface{})
	apiData["username"] = "wy"
	apiData["password"] = "zhanglanlu"

	jsonStr, _ := json.Marshal(apiData)
	fmt.Println(string(jsonStr))
}
