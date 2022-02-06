package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server starting..")

	http.HandleFunc("/", HelloWorld)
	http.HandleFunc("/people", getAllPeople)
	http.ListenAndServe(":8080", nil)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func getAllPeople(w http.ResponseWriter, r *http.Request) {
	people := GetPeopleData()
	json.NewEncoder(w).Encode(people)
}

func GetPeopleData() []Person {
	fakeData := []Person{
		Person{FirstName: "Jacob"},
		Person{FirstName: "Andrew"},
	}
	return fakeData
}

type Person struct {
	FirstName string
}
