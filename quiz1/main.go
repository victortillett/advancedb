package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello this is quiz 1"))

	weekday := time.Now().Weekday()

	// Prints the current weekday
	fmt.Println("Weekday is: ", weekday)
}

func displayDow(w http.ResponseWriter, r *http.Request) {
	day := time.Now().Weekday()

	ts, _ := template.ParseFiles("display.day.page.tpml")
	ts.Execute(w, day)
}

func main() {
	// server mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
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
