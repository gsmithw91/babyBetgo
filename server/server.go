package server

import (
	"babybetgo/handlers"
	"log"
	"net/http"
)

func ServerStart() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.IndexHandler)

	server := http.Server{
		Addr:    ":8040",
		Handler: mux,
	}

	log.Println("Starting server on :8040")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
