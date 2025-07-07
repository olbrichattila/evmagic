package dbHelper

// Todo add sync, to avoid database locks as multiple go routines can use it the same time
import (
	"database/sql"
	"fmt"
	"strings"
)

type DBHelper interface {
	LowerCaseResult()
	OriginalCaseResult()
	QueryAll(string, ...any) <-chan map[string]interface{}
	QueryOne(string, ...any) (map[string]interface{}, error)
	Execute(string, ...any) (int64, error)
	GetLastError() error
}

type dBase struct {
	// l         logger.Logger
	lowerCaseResult bool
	tx              *sql.Tx
	lastError       error
}

func New(tx *sql.Tx) DBHelper {
	db := &dBase{
		tx:              tx,
		lowerCaseResult: true,
	}

	return db
}

func (d *dBase) LowerCaseResult() {
	d.lowerCaseResult = true
}

func (d *dBase) OriginalCaseResult() {
	d.lowerCaseResult = false
}

func (d *dBase) QueryAll(sql string, pars ...any) <-chan map[string]interface{} {
	ch := make(chan map[string]interface{}, 1)
	d.lastError = nil
	if d.tx == nil {
		d.lastError = fmt.Errorf("db not open")
		close(ch)
		return ch
	}

	stmt, err := d.tx.Prepare(sql)
	if err != nil {
		d.lastError = err
		close(ch)
		return ch
	}

	rows, err := stmt.Query(pars...)
	if err != nil {
		d.lastError = err
		stmt.Close()
		close(ch)
		return ch
	}

	cols, err := rows.Columns()
	if err != nil {
		d.lastError = err
		rows.Close()
		stmt.Close()
		close(ch)
		return ch
	}

	colCount := len(cols)

	row := make([]interface{}, colCount)
	for i := range row {
		row[i] = new(interface{})
	}

	go func() {
		for rows.Next() {
			err := rows.Scan(row...)
			if err != nil {
				d.lastError = err
				break
			}
			result := make(map[string]interface{}, colCount)
			for i, colName := range cols {
				if d.lowerCaseResult {
					colName = strings.ToLower(colName)
				}
				value := *(row[i].(*interface{}))

				switch v := value.(type) {
				case string:
					result[colName] = v
				case []byte:
					result[colName] = string(v)
				default:
					result[colName] = v
				}
			}
			ch <- result
			result = nil
		}
		rows.Close()
		stmt.Close()
		close(ch)
	}()

	return ch
}

func (d *dBase) QueryOne(sql string, pars ...any) (map[string]interface{}, error) {
	if d.tx == nil {
		return nil, fmt.Errorf("db not open")
	}

	stmt, err := d.tx.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(pars...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	colCount := len(cols)
	row := make([]any, colCount)
	for i := range row {
		row[i] = new(any)
	}

	if !rows.Next() {
		return nil, fmt.Errorf("row cannot be found")
	}

	err = rows.Scan(row...)
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{}, colCount)
	for i, colName := range cols {
		if d.lowerCaseResult {
			colName = strings.ToLower(colName)
		}
		value := *(row[i].(*interface{}))
		switch v := value.(type) {
		case string:
			result[colName] = v
		case []byte:
			result[colName] = string(v)
		default:
			result[colName] = v
		}
	}

	return result, nil
}

func (d *dBase) Execute(sql string, pars ...any) (int64, error) {
	if d.tx == nil {
		return 0, fmt.Errorf("db not open")
	}

	stmt, err := d.tx.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(pars...)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (d *dBase) GetLastError() error {
	return d.lastError
}

func (d *dBase) logError(message string) {
	// TODO implement
	// if d.l != nil {
	// 	d.l.Error(message)
	// }
}
