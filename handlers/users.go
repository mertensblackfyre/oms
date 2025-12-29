package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		
		// Decode JSON from Request
		var body map[string]string
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		// Extract Values
		email := body["email"]
		password := body["password"]
		name := body["name"]

		if email == "" || password == "" {
			http.Error(w, "email and password required", http.StatusBadRequest)
			return
		}

		// Hash password
		hashed, err := bcrypt.GenerateFromPassword(
			[]byte(password),
			bcrypt.DefaultCost,
		)
		if err != nil {
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		id := uuid.NewString()

		// Insert to Database
		_, err = db.Exec(`
			INSERT INTO users (id, email, name, password)
			VALUES (?, ?, ?, ?)
		`, id, email, name, string(hashed))

		if err != nil {
			if strings.Contains(err.Error(), "UNIQUE") {
				http.Error(w, "email already exists", http.StatusConflict)
				return
			}
			http.Error(w, "database error", http.StatusInternalServerError)
			return
		}

		// Write JSON
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"id":     id,
			"status": "Success",
		})
	}
}
