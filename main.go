package main

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "fmt"
)

type Person struct {
  ID string `json:"id,omitempty"`
  FirstName string `json:"firstname,omitempty"`
  LastName string `json:"lastname,omitempty"`
  Address *Address `json:"address,omitempty"`
}

type Address struct {
  City string `json:"city,omitempty"`
  State string `json:"state,omitempty"`
}

var people []Person


func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request){
 json.NewEncoder(w).Encode(people)
}

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request){
  params := mux.Vars(req)
  for _, item := range people {
    if item.ID == params["id"]{
     json.NewEncoder(w).Encode(item)
     return
   }
  }
  json.NewEncoder(w).Encode(&Person{})

}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request){
  params := mux.Vars(req)

  var person Person

  _ = json.NewDecoder(req.Body).Decode(&person)

  person.ID = params["id"]
  people = append(people, person)
  json.NewEncoder(w).Encode(people)
}


func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request){
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
  //routers
  router := mux.NewRouter()

  people = append(people, Person{ID: "1", FirstName: "Edward", LastName: "Ime", Address: &Address{ City: "Dublin", State: "Casto"}})

people = append(people, Person{ID: "2", FirstName: "Caro", LastName: "Hernandez", Address: &Address{ City: "SanFrancisco", State: "California"}})

people = append(people, Person{ID: "3", FirstName: "Samantha", LastName: "Rueda", Address: &Address{ City: "Spain", State: "Madrid"}})

people = append(people, Person{ID: "4", FirstName: "MariaDelCarmen", LastName: "De Abreu", Address: &Address{ City: "Santiago", State: "Chile"}})

people = append(people, Person{ID: "5", FirstName: "Eduardo", LastName: "Wintetdaal", Address: &Address{ City: "Holanda", State: "Amsterdan"}})

  //endpoints
  router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
  router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
  router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
  router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")


  //ListenHTTP
  fmt.Println("Welcome to the API")
  fmt.Println("API Listen in localhost:3000")
  log.Fatal(http.ListenAndServe(":3000", router))
}
