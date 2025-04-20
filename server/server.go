package server

import (
	"babybetgo/handlers"
	"babybetgo/middleware"
	"log"
	"net/http"

	// Make sure handlers.DB is initialized before calling ServerStart()

	"github.com/go-chi/chi"
)

func ServerStart() {

	r := chi.NewRouter()

	//Middleware
	r.Use(middleware.LoggingMiddleware)

	//Routes
	r.Get("/", handlers.IndexHandler)
	r.Post("/register", handlers.RegisterHandler)
	r.Post("/login", handlers.LoginHandler)
	r.Get("/users/{id}", handlers.UserProfileHandler) //dynamic Routes

	fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))

	r.Handle("/static/", fileServer)

	log.Println("Starting Server on :8040")
	if err := http.ListenAndServe(":8040", r); err != nil {
		log.Println("server shutdown")
	}
}
