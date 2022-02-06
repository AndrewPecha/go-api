package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server starting..")

	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/people", getAllPeople)
	http.HandleFunc("/test", test)
	http.ListenAndServe(":8080", nil)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func getAllPeople(w http.ResponseWriter, r *http.Request) {
	people := getPeopleData()
	json.NewEncoder(w).Encode(people)
}

func getPeopleData() []Person {
	fakeData := []Person{
		Person{FirstName: "Jacob"},
		Person{FirstName: "Andrew"},
	}
	return fakeData
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test succeeded at /test")
}

type Person struct {
	FirstName string
}
