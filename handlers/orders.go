package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	d "github.com/oms/db"
)

func PostOrders(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var order d.Order

		w.Header().Set("Content-Type", "application/json")
		// Decode JSON from Request
		if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}
		fmt.Println(order.UserID)
		fmt.Println(order.ProductID)

	}
}
