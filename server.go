package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/oms/utils"
)

func main() {
	addr := ":8000"
	fmt.Printf("Starting server on %v\n", addr)
	http.ListenAndServe(addr, router())

}

func router() http.Handler {
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
			utils.Generate(w)
		})
	})

	return r
}
