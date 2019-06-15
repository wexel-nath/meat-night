package database

import (
	"database/sql"
	"log"

	"github.com/wexel-nath/meat-night/pkg/config"

	_ "github.com/lib/pq"
)

var (
	connection *sql.DB
)

func getConnection() *sql.DB {
	if connection == nil || connection.Ping() != nil {
		var err error
		connection, err = sql.Open("postgres", config.GetDatabaseURL())
		if err != nil {
			log.Fatalf("Error opening database: %v", err)
		}
	}
	return connection
}

// SetMockConnection is used for testing only
func SetMockConnection(db *sql.DB) {
	connection = db
}
