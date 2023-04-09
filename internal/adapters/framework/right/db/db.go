package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	// connect to database
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("db connection fialure: %v", err)
	}

	// testing the connection
	if err := db.Ping(); err != nil {
		log.Fatalf("db test or ping failure: %v", err)
	}
	return &Adapter{db: db}, nil
}

func (da Adapter) CloseDbConnection() {
	if err := da.db.Close(); err != nil {
		log.Fatalf("error in closing database: %v", err)
	}
}

func (da Adapter) AddToHistroy(answer int, operation string) error {
	query := `
		INSERT INTO arith_histroy (
			date,
			answer,
			operation
	  ) VALUES (
			$1,
			$2,
			$3
		);
	`
	rows, err := da.db.Exec(
		query,
		time.Now(),
		answer,
		operation,
	)
	if err != nil {
		return err
	}
	rows_affected, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if rows_affected == 0 {
		return fmt.Errorf("no change made")
	}
	return nil
}
