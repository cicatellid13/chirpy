package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Optional: Add a handler
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//     fmt.Fprintln(w, "Hello from ServeMux!")
	// })

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Printf("server listening on %v\n", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("error starting server:", err)
	}
}
