package main

import (
	"fmt" // printing formatting
	"log" // logging errors
	"net/http" // web server package
)

func main() { 
	fmt.Println("admin node starting on port:- 8080")
	// basic routing
	http.HandleFunc("/submit", func(w http.ResponseWriter,r *http.Request){ // if you reach /submit, execute the function below
		fmt.Fprintf(w, "Request received at Admin")
	})
	if err := http.ListenAndServe(":8080", nil); err != nil { // if error is not nil, log it and exit !!
		log.Fatal(err)
	}
}