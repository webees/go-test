package main

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
)

var arrStr []string

func handler(w http.ResponseWriter, r *http.Request) {
	// deJson := json.NewDecoder(req.Body)
	// var params map[string]string
	// deJson.Decode(&params)
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] %q\n", k, v)
	}
	if err := r.ParseForm(); err!= nil {
			log.Fatal(err.Error())
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] %q\n", k, v)
	}
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServeTLS("localhost:1024", "./server.crt", "./server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err.Error())
	}
}
