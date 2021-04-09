package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"project/handlers"
)


func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/send_msg", handlers.BodyRequest)
	fmt.Println("Server is running at port 5050")
	log.Fatal(http.ListenAndServe(":5050", router))


}
