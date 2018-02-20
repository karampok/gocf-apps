package db

import (
	"database/sql"
	"fmt"
)

func DeployV1Schema(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS schema_info (
               version INTEGER
             )`)
	if err != nil {
		return err
	}

	version, err := getSchemaVersion(db)
	if err != nil {
		return err
	}

	if version == 1 {
		return nil
	}

	_, err = db.Exec(`INSERT INTO schema_info VALUES (1)`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS products (
               name      TEXT,
               value   INTEGER
             )`)
	if err != nil {
		return err
	}
	return nil
}

func getSchemaVersion(db *sql.DB) (int, error) {
	r, err := db.Query(`SELECT version FROM schema_info LIMIT 1`)
	if err != nil {
		if err.Error() == "no such table: schema_info" {
			return 0, nil
		}
		return 0, err
	}
	defer r.Close()

	// no records = no schema
	if !r.Next() {
		return 0, nil
	}

	var v int
	err = r.Scan(&v)
	// failed unmarshall is an actual error
	if err != nil {
		return 0, err
	}

	// invalid (negative) schema version is an actual error
	if v < 0 {
		return 0, fmt.Errorf("Invalid schema version %d found", v)
	}

	return int(v), nil
}
