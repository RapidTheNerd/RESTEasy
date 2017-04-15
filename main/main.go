package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

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
	router.HandleFunc("/people/{id}", CreatePersonEndPoint).Methods("GET")
}