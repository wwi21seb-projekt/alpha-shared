package db

import (
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandlePGError processes PostgreSQL errors and logs them appropriately
func HandlePGError(err error, logger *zap.SugaredLogger, operation string) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case pgerrcode.ForeignKeyViolation:
			switch pgErr.ConstraintName {
			case "user1_fk", "user2_fk", "users_fk", "sender_fk", "recipient_fk", "username_fk", "subscribee_fk", "subscriber_fk":
				return status.Errorf(codes.NotFound, "user does not exist")
			case "posts_fk", "post_id_posts_fk":
				return status.Errorf(codes.NotFound, "post does not exist")
			case "hashtags_fk":
				return status.Errorf(codes.NotFound, "hashtag does not exist")
			case "chats_fk":
				return status.Errorf(codes.NotFound, "chat does not exist")
			default:
				return status.Errorf(codes.NotFound, "foreign key violation")
			}
		case pgerrcode.UniqueViolation:
			switch pgErr.ConstraintName {
			case "users_uq":
				return status.Errorf(codes.AlreadyExists, "chat already exists")
			case "hashtags_uq":
				return status.Errorf(codes.AlreadyExists, "hashtag already exists")
			case "email_uq":
				return status.Errorf(codes.AlreadyExists, "email already exists")
			case "subscriptions_uq":
				return status.Errorf(codes.AlreadyExists, "subscription already exists")
			case "user_token_type_uq":
				return status.Errorf(codes.AlreadyExists, "token already exists for user")
			default:
				return status.Errorf(codes.AlreadyExists, "unique constraint violation")
			}
		case pgerrcode.CheckViolation:
			return status.Errorf(codes.InvalidArgument, "check constraint violation")
		default:
			logger.Warnw("Error executing SQL query", "error", err)
			return status.Errorf(codes.InvalidArgument, "sql query error")
		}
	}

	logger.Errorw("Error executing SQL query", "error", err)
	return status.Errorf(codes.Internal, "internal error executing sql query")
}
