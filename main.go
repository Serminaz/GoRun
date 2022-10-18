package main

import (
	"log"
	"net/http"
	//	"os"
)

//	func indexHandler(w http.ResponseWriter, r *http.Request) {
//		w.Write([]byte("<h1>Hello World!</h1>"))
//	}
//
//	func main() {
//		port := os.Getenv("PORT")
//		if port == "" {
//			port = "3001"
//		}
//
//		mux := http.NewServeMux()
//
//		mux.HandleFunc("/", indexHandler)
//		http.ListenAndServe(":"+port, mux)
//	}
func main() {
	http.HandleFunc("/hello", viewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello, web")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}
