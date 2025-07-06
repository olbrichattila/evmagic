package replay

import (
	"database/sql"
	"os"

	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
	databaseReplay "github.com/olbrichattila/evmagic/pkg/replay/database-replay"
	memoryReplay "github.com/olbrichattila/evmagic/pkg/replay/memory-replay"
)

// TODO only replay the event if it is newer then the last action, otherwise acknolege and skip

func New(publisher contracts.Publisher, db *sql.DB) contracts.Replay {
	// TODO create factory to create replay by config or env (switch)
	switch os.Getenv("REPLAY") {
	case "memory":
		return memoryReplay.New(publisher)
	case "database":
		return databaseReplay.New(publisher, db)
	default:
		return databaseReplay.New(publisher, db)
	}
}
