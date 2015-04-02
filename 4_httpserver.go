package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r* http.Request) {
	fmt.Fprint(w, "Hello! You will see me on all urls.")
}

func main() {
	var rc error

	http.HandleFunc("/", handler);

	fmt.Print("Initiating http server\n")
	rc = http.ListenAndServe(":8080", nil);

	if rc != nil {
		fmt.Print("Error: ", rc.Error(), "\n")
	}
}
