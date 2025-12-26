package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	d "github.com/oms/db"
)

func GetProductsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		products, err := d.FetchAllProducts(db)
		if err != nil {
			log.Println(err)
		}

		products_json, err := json.Marshal(products)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(string(products_json))
		w.Header().Add("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(products_json)

		if err != nil {
			log.Println(err)
		}
	}
}

/*
func SeedProducts(db *sql.DB) error {
	// 1. Create a context with a timeout (e.g., 5 seconds)
	// Even though you have Background(), a timeout ensures
	// the process doesn't hang forever if the DB is busy.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	adjectives := []string{"Pro", "Ultra", "Basic", "Gaming", "Cloud"}
	nouns := []string{"Laptop", "Mouse", "Keyboard", "Monitor", "Desk"}

	fmt.Println("Seeding 20 products...")

	for i := 0; i < 20; i++ {
		p := dd.Product{
			ID:         uuid.NewString(),
			Name:       fmt.Sprintf("%s %s %d", adjectives[rand.Intn(len(adjectives))], nouns[rand.Intn(len(nouns))], i),
			Price:      float64(rand.Intn(1000)) + rand.Float64(),
			OutOfStock: rand.Intn(2) == 0, // Random true/false
			CreatedAt:  time.Now(),
		}

		// 2. Use ExecContext.
		// Note: Change '?' to '$1, $2...' if you are using PostgreSQL.
		query := `INSERT INTO products (id, name, price, out_of_stock, created_at)
				  VALUES (?, ?, ?, ?, ?)`

		_, err := db.ExecContext(ctx, query, p.ID, p.Name, p.Price, p.OutOfStock, p.CreatedAt)
		if err != nil {
			return fmt.Errorf("failed to insert product %d: %w", i, err)
		}
	}

	fmt.Println("Successfully added 20 products!")
	return nil
}*/
