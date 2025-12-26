package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/oms/utils"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var DB *sql.DB

func DBs() *sql.DB {

	url := utils.TURSO_DATABASE_URL + "?authToken=" + utils.TURSO_AUTH_TOKEN

	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", utils.TURSO_DATABASE_URL, err)
		os.Exit(1)
	}

	ctx := context.Background()

	schema := `

	PRAGMA foreign_keys = ON;

	CREATE TABLE IF NOT EXISTS users (
    	id TEXT PRIMARY KEY,              -- UUID
    	email TEXT UNIQUE NOT NULL,
    	name TEXT,
    	password TEXT,
    	created_at TEXT DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS products (
    	id TEXT PRIMARY KEY,              -- UUID
    	name TEXT NOT NULL,
    	price REAL NOT NULL,
    	out_of_stock INTEGER DEFAULT 0,    -- 0 = false, 1 = true
    	created_at TEXT DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS orders (
    	id TEXT PRIMARY KEY,              -- UUID
    	user_id TEXT NOT NULL,
    	product_id TEXT NOT NULL,
    	quantity INTEGER NOT NULL,
    	total_price REAL NOT NULL,
    	created_at TEXT DEFAULT CURRENT_TIMESTAMP,

    	FOREIGN KEY (user_id) REFERENCES users(id),
    	FOREIGN KEY (product_id) REFERENCES products(id)
	);
`
	_, err = db.ExecContext(ctx, schema)
	if err != nil {
		log.Fatalln(err)
	}

	DB = db
	return db
}
