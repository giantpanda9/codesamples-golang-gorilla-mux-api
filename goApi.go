package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type jsonStruct struct {
	Url string
	Name string
	Date string `json:",omitempty"`
	Description string `json:",omitempty"`
}

func getJSON() string {
	inputJSON, err := ioutil.ReadFile("data/output.json")
	if err != nil {
		fmt.Println("Error opening file:\n")
		fmt.Println(err)
	}
	jsonMapping := make([]jsonStruct,0)
	errUnmarshal := json.Unmarshal([]byte(inputJSON), &jsonMapping)
	if errUnmarshal != nil {
		fmt.Println("Error Unmarshalling JSON:\n")
		fmt.Println(errUnmarshal)
		return ""
	}
	
	outputJSON, errMarshal := json.Marshal(jsonMapping)
	if errMarshal != nil {
		fmt.Println("Error Marshalling JSON:\n")
		fmt.Println(errMarshal)
		return ""
	}
	//fmt.Println(string(outputJSON))
	return string(outputJSON)
}

func main() {
	fmt.Println(getJSON())
}
