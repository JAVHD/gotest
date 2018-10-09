package main

import (
	"encoding/json"
	"fmt"
)

type aStruct struct {
	Name string
	Url  string
}

type bStruct struct {
	Bıutton []aStruct
}

func main() {
	//var aSlice

	//var aSlice aStruct
	var bSlice bStruct

	jsonStr := `{
	    "button": [
	      {
	            "name": "中星9号",
	            "url": "http://www.weixingcanshu.com/topic-92.2.html"
	          },
	       {
	            "name": "中星6B",
	            "url": "http://www.weixingcanshu.com/topic-115.5.html"
	          }
	    ]
	}`

	/*jsonStr := ` {
	  "name": "中星9号",
	  "url": "http://www.weixingcanshu.com/topic-92.2.html"
	}`*/

	json.Unmarshal([]byte(jsonStr), &bSlice)
	fmt.Println(bSlice)

}
