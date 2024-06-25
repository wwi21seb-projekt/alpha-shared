package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/exaring/otelpgx"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wwi21seb-projekt/alpha-shared/config"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
)

type DB struct {
	pool   *pgxpool.Pool
	logger *zap.SugaredLogger
}

func NewDB(ctx context.Context, dbCfg config.DatabaseConfig, logger *zap.SugaredLogger) (*DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable search_path=%s",
		dbCfg.PostgresHost, dbCfg.PostgresPort, dbCfg.PostgresDB, dbCfg.PostgresUser, dbCfg.PostgresPassword, dbCfg.SchemaName)

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database config: %w", err)
	}

	options := []otelpgx.Option{otelpgx.WithIncludeQueryParameters(), otelpgx.WithTrimSQLInSpanName()}
	if os.Getenv("ENVIRONMENT") == "development" {
		options = append(options, otelpgx.WithAttributes()) // Include query attributes in spans
	}
	cfg.ConnConfig.Tracer = otelpgx.NewTracer(options...)

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	return &DB{pool: pool, logger: logger}, nil
}

func (db *DB) Close() {
	db.pool.Close()
}

func (db *DB) Ping(ctx context.Context) error {
	return db.pool.Ping(ctx)
}

// ----------------- Connection and Transaction Management -----------------

func (db *DB) Acquire(ctx context.Context) (*pgxpool.Conn, error) {
	db.logger.Debug("Acquiring connection from pool...")
	conn, err := db.pool.Acquire(ctx)
	if err != nil {
		db.logger.Errorw("Failed to acquire connection from pool", "error", err)
		return nil, status.Error(codes.Internal, "failed to acquire connection")
	}
	return conn, nil
}

func (db *DB) BeginTx(ctx context.Context, conn *pgxpool.Conn) (pgx.Tx, error) {
	db.logger.Debug("Starting transaction...")
	tx, err := conn.Begin(ctx)
	if err != nil {
		db.logger.Errorw("Failed to start transaction", "error", err)
		conn.Release()
		return nil, status.Error(codes.Internal, "failed to acquire connection")
	}
	return tx, nil
}

func (db *DB) CommitTx(ctx context.Context, tx pgx.Tx) error {
	db.logger.Debug("Committing transaction...")
	err := tx.Commit(ctx)
	if err != nil {
		return HandlePGError(err, db.logger, "CommitTx")
	}
	db.logger.Debug("Committed transaction")
	return nil
}

func (db *DB) RollbackTx(ctx context.Context, tx pgx.Tx) error {
	err := tx.Rollback(ctx)
	if err != nil && !errors.Is(err, pgx.ErrTxClosed) {
		return HandlePGError(err, db.logger, "RollbackTx")
	}
	db.logger.Debug("Rolled back transaction")
	return nil
}

func (db *DB) Exec(ctx context.Context, conn *pgxpool.Conn, query string, args ...interface{}) (pgconn.CommandTag, error) {
	tag, err := conn.Exec(ctx, query, args...)
	if err != nil {
		return tag, HandlePGError(err, db.logger, "Exec")
	}
	return tag, nil
}
