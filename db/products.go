package db

import (
	"database/sql"
	"fmt"
	"log"
)

func FetchAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var products = []Product{}

	for rows.Next() {
		var p Product
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Price,
			&p.OutOfStock,
			&p.CreatedAt,
		)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}

		products = append(products, p)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
	}
	return products, nil
}
