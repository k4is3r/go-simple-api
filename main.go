package main

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

func GetPeopleEndpoint(){

}


func main(){
  #routers
  router := mux.NewRouter()



  #Endpoints
  router.HandleFunc("/people", GetPeopleEndpoint)



}
