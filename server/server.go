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

	// Middleware for all
	r.Use(middleware.LoggingMiddleware)

	// Unprotected Routes
	r.Get("/", handlers.IndexHandler)
	r.Post("/register", handlers.RegisterHandler)
	r.Post("/login", handlers.LoginHandler)
	r.Get("/users/{id}", handlers.UserProfileHandler) //dynamic Routes

	// Protected Routes
	r.Group(func(protected chi.Router) {
		protected.Use(middleware.RequireAuth)
		protected.Get("/me", handlers.MeHandler)
		protected.Get("/user_info_partial", handlers.UserInfoPartialHandler)
		protected.Post("/create_pregnancy", handlers.CreatePregnancyHandler)
	})

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Starting Server on :8040")
	if err := http.ListenAndServe(":8040", r); err != nil {
		log.Println("server shutdown")
	}
}
