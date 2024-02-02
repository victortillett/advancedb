package main

import (
	"log"
	"net/http"
)

func displayDow(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Stranger ! Today is monday"))
}

func main() {
	// server mux
	mux := http.NewServeMux()
	mux.HandleFunc("/dow", displayDow)

	// mux.HandleFunc("/time", displayTime)
	// mux.HandleFunc("/area-calculator", getValues)
	// mux.HandleFunc("/area-calculator-2", calculateArea)
	// fileserver creation
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/quiz1/", http.StripPrefix("/quiz1", fileServer))
	log.Println("Starting web server on port :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
