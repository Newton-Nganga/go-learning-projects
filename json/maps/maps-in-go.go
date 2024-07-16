package main


import (
	"fmt"
	"encoding/json"
)

func main(){
	data := map[string]interface{}{
		"name":"zhangsan",
		"age":18,
		"sex":"male",
		"boolValue": true,
		"floatValue": 1.2,
		"nullValue": nil,
		"arrayValue": []string{"a","b","c"},
		"mapValue": map[string]interface{}{
			"key1":"value1",
			"key2":"value2",
			},
		"objectValue": map[string]interface{}{
			"arrayValue": []int{1,2,3,4,5},
		},
	}
	jsonData ,err := json.Marshal(data)

	if err != nil{
		fmt.Printf("could not marshal json: %s\n",err)
	}

	fmt.Printf("json data: %s\n",jsonData)
}

/*


>>. go run maps-in-go.go 

json data: {"age":18,
"arrayValue":["a","b","c"],
"boolValue":true,
"floatValue":1.2,
"mapValue":{"key1":"value1","key2":"value2"},
"name":"zhangsan",
"nullValue":null,
"objectValue":{"arrayValue":[1,2,3,4,5]},
"sex":"male"}

*/
