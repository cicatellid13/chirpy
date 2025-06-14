package main

import (
	"fmt"
	"net/http"
	// "os"
)

func main() {
	mux := http.NewServeMux()

	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	data, err := os.ReadFile("index.html")
	// 	if err != nil {
	// 		http.Error(w, "Could not read index.html", http.StatusInternalServerError)
	// 		return
	// 	}
	// 	w.Header().Set("Content-Type", "text/html")
	// 	w.Write(data)
	// })
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
