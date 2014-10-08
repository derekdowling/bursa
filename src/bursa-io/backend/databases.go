package backend

import (
	_ "github.com/lib/pq"
	"log"
)

func pg_connect() *driver.Conn {
	db, err := sql.Open("postgres", "postgres://bursa:securemebaby@localhost/bursa")
	if err != nil {
		log.Fatal(err)
		return null
	}
	return db
}
