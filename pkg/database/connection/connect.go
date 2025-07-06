package connection

import (
	"database/sql"
	"fmt"

	"github.com/olbrichattila/evmagic/pkg/database/config"
)

func Open() (*sql.DB, error) {
	fmt.Println("connecting to db", config.GetConnectionString())
	db, err := sql.Open(config.GetConnectionName(), config.GetConnectionString())
	if err != nil {
		return nil, err
	}

	return db, nil
}
