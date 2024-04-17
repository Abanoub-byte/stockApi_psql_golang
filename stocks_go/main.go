package main

import (
	"fmt"
	"log"
	"net/http"
	
	"github.com/bob/stocks_go/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on the port 8081")

	log.Fatal(http.ListenAndServe(":8081", r))
}
