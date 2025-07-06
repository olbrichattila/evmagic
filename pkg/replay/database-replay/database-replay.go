package databaseReplay

import (
	"database/sql"
	"fmt"
	"sync"

	frameworkAction "github.com/olbrichattila/evmagic/pkg/actions/framework-action"
	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
)

func New(publisher contracts.Publisher, db *sql.DB) contracts.Replay {
	return &replay{
		publisher: publisher,
		db:        db,
	}
}

type replay struct {
	mu        sync.Mutex
	db        *sql.DB
	publisher contracts.Publisher
}

// Register implements contracts.Replay.
func (r *replay) Register(parentActionIdentifier string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	// TODO may need to check if exist before
	_, err = r.db.Exec("INSERT INTO event_history (event_id) VALUES (?)", parentActionIdentifier)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return err
}

// Replay event if it is a duplicate
func (r *replay) Replay(actionId string) (bool, error) {
	rows, err := r.db.Query("SELECT event_id FROM event_history where event_id = ? limit 1", actionId)
	if err != nil {
		return false, err
	}

	// Nothing to replay
	if !rows.Next() {
		return false, nil
	}

	rows, err = r.db.Query("SELECT event_payload FROM event_replay WHERE event_id = ?", actionId)
	if err != nil {
		return true, err
	}

	var data []byte
	for rows.Next() {
		if err := rows.Scan(&data); err != nil {
			return true, err
		}
		aInfo, err := frameworkAction.ActionInfoFromPayload(data)
		if err != nil {
			return true, err
		}
		fmt.Println("Duplicate event found", aInfo.MessageIdentifier)
		err = r.publisher.Publish(aInfo.Topic, data)
		if err != nil {
			return true, err
		}
	}

	return true, nil
}

// Store event after success if need replaying later
func (r *replay) Store(parentActionIdentifier string, replayData []byte) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	_, err = r.db.Exec("INSERT INTO event_replay (event_id, event_payload) VALUES (?,?)", parentActionIdentifier, replayData)
	if err != nil {
		return err
	}

	tx.Commit()
	return err
}
