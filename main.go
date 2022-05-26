package main

import (
	"fmt"
	"github.com/russellcxl/go-translator/server"
	"log"
	"net/http"
)

func generalHTTP() {
	http.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "bye")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "hello")
	})

	log.Println("server starting on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {

	//generalHTTP()

	err := server.NewTranslator("images", "output", "en").Execute()
	if err != nil {
		log.Fatalln(err)
	}

}