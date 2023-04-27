package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string
	Password string
	Tags     []string
}

func main() {
	fmt.Println("Json file")
	// EncodeJson()
	DecodeJson()

}

// func EncodeJson() {
// 	lco := []course{
// 		{"Reactjs", 299, "LCO", "abc123", []string{"web", "fe", "js"}},
// 		{"Python", 399, "LCO", "abc123", []string{"web", "be"}},
// 		{"Angular", 399, "LCO", "htc123", nil},
// 	}

// 	finalJson, err := json.MarshalIndent(lco, "", "\t")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%s\n", finalJson)

// }

func DecodeJson() {
	jsonData := []byte(`
				{
					"coursename": "Python",
					"Price": 399,
					"Platform": "LCO",
					"Password": "abc123",
					"Tags": ["web","be"]
				}
	`)
	var lcoCourse course
	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("JS code")
		json.Unmarshal(jsonData, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)

	} else {
		fmt.Println("Json was not valid")
	}

	//some case where we want to add data to key value pair

	var myData map[string]interface{}
	json.Unmarshal(jsonData, &myData)
	fmt.Printf("%#v\n", myData)

}
