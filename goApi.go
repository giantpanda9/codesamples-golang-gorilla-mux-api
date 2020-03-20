package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"log"
	"io/ioutil"
	"github.com/gorilla/mux"
)

type jsonStruct struct {
	Url string
	Name string
	Date string `json:",omitempty"`
	Description string `json:",omitempty"`
}

func getData(w http.ResponseWriter, r *http.Request) {
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
	}
	
	outputJSON, errMarshal := json.Marshal(jsonMapping)
	if errMarshal != nil {
		fmt.Println("Error Marshalling JSON:\n")
		fmt.Println(errMarshal)
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(string(outputJSON))
}

func main() {
	muxRouter := mux.NewRouter()	
	muxRouter.HandleFunc("/goapi/getdata", getData).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", muxRouter))
}
