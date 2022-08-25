package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var f *os.File

func main() {
	s := mux.NewRouter()
	f, _ = os.OpenFile("./data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	s.HandleFunc("/", handler).Methods("Get")
	s.HandleFunc("/about", handlerAbout).Methods("Get")
	s.HandleFunc("/post", handlerPost).Methods("Post")
	http.ListenAndServe(":8080", s)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func handlerAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello There. This is just a simple code for makeing a Server :-)")
}

func handlerPost(w http.ResponseWriter, r *http.Request) {
	// get the body from the requeest.

	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	f.Write(b)
	f.Write([]byte("\n"))
}
