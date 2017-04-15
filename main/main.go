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

}


func GetPeopleEndPoint(w http.ResponseWriter, req * http.Request){
	json.NewEncoder(w).Encode(people)
}


func CreatePersonEndPoint(w http.ResponseWriter, req * http.Request){

}


func DeletePersonEndPoint(w http.ResponseWriter, req * http.Request){

}

func main(){
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", FirstName: "Somename", LastName:"RandomName", Address:
	&Address{City: "Some city", State: "Somestate"}})
	log.Fatal(http.ListenAndServe(":8080", router))
	router.HandleFunc("/people" , GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndPoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndPoint).Methods("DELETE")
	
}