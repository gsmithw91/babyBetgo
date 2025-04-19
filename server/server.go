package server

import (
	"babybetgo/handlers"
	// Make sure handlers.DB is initialized before calling ServerStart()

	"log"
	"net/http"
)

func ServerStart() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/register", handlers.RegisterHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/get_user_balance", handlers.GetUserBalanceHandler)
	// Serve static files
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	server := http.Server{
		Addr:    ":8040",
		Handler: mux,
	}

	log.Println("Starting server on :8040")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
