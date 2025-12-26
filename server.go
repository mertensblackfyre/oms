package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/oms/db"
	"github.com/oms/handlers"
	"github.com/oms/utils"
)

func main() {

	db := db.DBs()
	addr := ":8000"
	fmt.Printf("Starting server on %v\n", addr)
	http.ListenAndServe(addr, router(db))

	defer db.Close()
}

func router(db *sql.DB) http.Handler {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(utils.TokenAuth))
		r.Use(jwtauth.Authenticator(utils.TokenAuth))

		r.Get("/order", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Authorization:", r.Header.Get("Authorization"))
			_, claims, _ := jwtauth.FromContext(r.Context())
			w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["user_id"])))
		})
	})

	// Public routes
	r.Group(func(r chi.Router) {
		r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
			utils.GenerateJWT(w)
		})
		r.Post("/register", handlers.RegisterUser(db))
	})

	return r
}
