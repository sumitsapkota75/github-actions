package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health-check", helloHandler)

	log.Println("Listning for requests at http://localhost:8000/health-check")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func helloHandler(w http.ResponseWriter, req *http.Request) {

	io.WriteString(w, "Tada, 👻 Server is up and running!\n")
}
