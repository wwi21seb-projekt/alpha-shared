package db

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	log "github.com/sirupsen/logrus"
)

type DB struct {
	Pool *pgxpool.Pool
}

type TxFunc func(tx pgx.Tx) error

func NewDB(connString string) (*DB, error) {
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, err
	}

	return &DB{Pool: pool}, nil
}

func (db *DB) Close() {
	db.Pool.Close()
}

// ----------------- Transaction Functions -----------------

// Begin starts a new transaction and returns a transaction object
func (db *DB) Begin(ctx context.Context) (pgx.Tx, error) {
	log.Info("Beginning new transaction...")
	return db.Pool.Begin(ctx)
}

// Commit commits the transaction
func (db *DB) Commit(ctx context.Context, tx pgx.Tx) error {
	log.Info("Committing transaction...")
	return tx.Commit(ctx)
}

// Rollback rolls back the transaction
func (db *DB) Rollback(ctx context.Context, tx pgx.Tx) error {
	err := tx.Rollback(ctx)

	// Ignore the error if the transaction was already committed
	if err != nil && errors.Is(err, pgx.ErrTxClosed) {
		return nil
	}

	// Log late, so the consumer doesn't get a log for noops (when the transaction was already committed)
	log.Info("Rolling back transaction...")
	return err
}
