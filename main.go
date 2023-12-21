package main

import (
	"log"
	"net/http"
	"os"
)

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	country := r.Header.Get("CF-IPCountry")
	if country == "IN" {
		http.Redirect(w, r, "https://satyarthi.org.in", 302)
	} else {
		http.Redirect(w, r, "https://satyarthi-us.org", 302)
	}
}

func main() {
	http.HandleFunc("/", handleRedirect)	

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"

	}
	log.Fatal(http.ListenAndServe("0.0.0.0:" + port, nil))
}
