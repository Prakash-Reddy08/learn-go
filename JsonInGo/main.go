package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //giving alias
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // if we provide "-" , whenever someone consumes this json object password field will not appear
	Tags     []string `json:"tags,omitempty"` //if value is nil or null it is this filed is omitted
}

func main() {
	EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJs BootCamp", 299, "myCode.online", "abc123", []string{"javascript", "webDevelopment"}},
		{"MERN BootCamp", 399, "myCode.online", "abc123", []string{"javascript", "webDevelopment", "node.js", "mongoDB", "React.js"}},
		{"Go BootCamp", 0, "myCode.online", "abc123", nil},
	}

	//package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonData := []byte(`{
		"coursename": "ReactJs BootCamp",
		"Price": 299,
		"website": "myCode.online",
		"tags": ["javascript","webDevelopment"]
	}`)

	var lcoCourse course
	// check if the json is valid or not
	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("json was valid")
		json.Unmarshal(jsonData, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Json was not valid")
	}

	//some cases where you just want to add data to key value pairs

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonData, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v, value is %v and Types is:%T\n", k, v, v)
	}
}
