package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Printf("server listening on %v\n", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("error starting server:", err)
	}
}
