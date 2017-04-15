package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
)

type Person struct{
	ID string `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	
	Address *Address `json:"address"`
	
}

type Address struct {
	City string `json:"city"`
	State string `json:"state"`
	
}

var people []Person


func GetPersonEndPoint(w http.ResponseWriter, req * http.Request){
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["ID"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}


func GetPeopleEndPoint(w http.ResponseWriter, req * http.Request){
	json.NewEncoder(w).Encode(people)
}


func CreatePersonEndPoint(w http.ResponseWriter, req * http.Request){
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}


func DeletePersonEndPoint(w http.ResponseWriter, req * http.Request){
	params := mux.Vars(req)
	for index, item := range people{
		if item.ID == params["id"]{
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main(){
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", FirstName: "Somename", LastName:"RandomName", Address:
	&Address{City: "Some city", State: "Somestate"}})
	log.Fatal(http.ListenAndServe(":8080", router))
	router.HandleFunc("/people/" , GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndPoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndPoint).Methods("DELETE")
	
}