package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

const (
	maxOpenDBConn = 10
	maxIdleDBConn = 5
	maxLifeDBTime = 5 * time.Minute
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dsn string) (*Adapter, error) {
	dbPool, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, err
	}

	dbPool.SetMaxOpenConns(maxOpenDBConn)
	dbPool.SetMaxIdleConns(maxIdleDBConn)
	dbPool.SetConnMaxLifetime(maxLifeDBTime)

	if err := dbPool.Ping(); err != nil {
		return nil, err
	}
	return &Adapter{
		db: dbPool,
	}, nil
}

func (da *Adapter) CloseDbConnection() {
	if err := da.db.Close(); err != nil {
		log.Fatalf("db close faliure: %v", err)
	}
}

func (da *Adapter) AddToHistroy(answer int32, operation string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `INSERT INTO arith-histroy (answer, operation) values $1, $2, $3`
	_, err := da.db.ExecContext(ctx, stmt, answer, operation)
	if err != nil {
		return fmt.Errorf("failed to insert the arith-histroy: %s", err)
	}
	return nil
}
