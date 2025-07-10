package connection

import (
	"database/sql"

	"github.com/olbrichattila/evmagic/pkg/database/config"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open(config.GetConnectionName(), config.GetConnectionString())
	if err != nil {
		return nil, err
	}

	// db.SetMaxOpenConns(100)                 // hard cap, e.g., 100 total DB connections
	// db.SetMaxIdleConns(20)                  // idle connections to keep alive
	// db.SetConnMaxLifetime(time.Minute * 10) // avoid using stale connections

	return db, nil
}
