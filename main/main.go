package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
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

func GetPersonEndPoint(w http.ResponseWriter, req * http.Request){

}


func GetPeopleEndPoint(w http.ResponseWriter, req * http.Request){

}


func CreatePersonEndPoint(w http.ResponseWriter, req * http.Request){

}


func DeletePersonEndPoint(w http.ResponseWriter, req * http.Request){

}

func main(){
	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
	router.HandleFunc("/people" , GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndPoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndPoint).Methods("DELETE")
	
}