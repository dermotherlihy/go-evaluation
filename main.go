package main

import (
	"log"
	"net/http"

	"github.com/dermotherlihy/go-evaluation/events-service"
)

func main() {
	events-service.AddServices()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
