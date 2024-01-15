package main

import (
	"net/http"

	"log"

	"github.com/samueltuyizere/query-builder/app/views"
)

func main(){
	http.HandleFunc("/", views.ShowMainPage)
	log.Fatal(http.ListenAndServe(":8001", nil))

}
