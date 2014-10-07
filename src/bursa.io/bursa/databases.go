package bursa

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func pg_connect() {
	db, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
	if err != nil {
		log.Fatal(err)
	}
}
