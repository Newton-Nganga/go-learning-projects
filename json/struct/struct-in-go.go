package main


import (
	"fmt"
	"encoding/json"
	"time"
)

type myJSON struct {
	StringValue string `json:"stringValue"`
	IntValue int `json:"intValue"`
	BoolValue bool `json:"boolValue`
	DateValue time.Time `json:"dateValue`
	ObjectValue *myObject `json:"objectValue"`
	NullIntValue *int `json:"nullIntValue"`
	NullStringValue *string `json:nullStringValue`
}
type myObject struct{
	ArrrayValue []int `json:"arrayValue"`
}

func main(){
	otherInt := 4321
	data := &myJSON{
		IntValue:1234,
		BoolValue: true,
		DateValue: time.Now(),
		ObjectValue: &myObject{
			ArrrayValue:[]int{1,2,3,4,5},
			},
			NullIntValue: &otherInt,
			NullStringValue: nil,
			StringValue : "hello!",
	}
	jsonData ,err := json.Marshal(data)

	if err != nil{
		fmt.Printf("could not marshal json: %s\n",err)
	}

	fmt.Printf("json data: %s\n",jsonData)
}


/*
json data: {
"stringValue":"hello!"
,"intValue":1234,
"BoolValue":true,
"DateValue":"2024-07-16T11:05:03.46437663+03:00",
"objectValue":{"arrayValue":[1,2,3,4,5]},
"nullIntValue":4321,
"NullStringValue":null}



*/
