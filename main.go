package main

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request){

}

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request){

}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request){

}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request){

}

func main(){
  #routers
  router := mux.NewRouter()



  #Endpoints
  router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
  router.HandleFunc("/people/{id}", GetPersonEnpoint).Methods("GET")
  router.HandleFunc("/people/{id}", CreatePeopleEndpoint).Methods("POST")
  router.HandleFunc("/people/{id}", DeletePeopleEndpoint).Methods("DELETE")


  #ListenHTTP
  log.Fatal(http.ListenAndServe(":3000", router))
}
