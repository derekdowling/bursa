package backend

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

// returns a driver.Conn
// http://godoc.org/database/sql/driver#Conn
func pg_connect() *sql.DB {
	db, err := sql.Open("postgres", "postgres://bursa:securemebaby@localhost/bursa")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}
