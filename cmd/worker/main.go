package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Worker node starting on port:- 8081")
	http.HandleFunc("/infer", func(w http.ResponseWriter, r *http.Request) { // why * ? * means any method (GET, POST, etc.)
		fmt.Fprintf(w, "Inference request received at Worker")
	})

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}

}