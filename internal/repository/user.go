package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/dickyadrian/paper-disbursement-system/internal/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, name string, balance int64) (*entity.User, error) {
	sql, args, err := psql.Insert("users").
		Columns("name", "balance").
		Values(name, balance).
		Suffix("RETURNING id, name, balance, created_at, updated_at").
		ToSql()
	if err != nil {
		return nil, err
	}

	var u entity.User
	if err := r.db.QueryRow(ctx, sql, args...).Scan(&u.ID, &u.Name, &u.Balance, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *UserRepository) Get(ctx context.Context, id int64) (*entity.User, error) {
	sql, args, err := psql.Select("id", "name", "balance", "created_at", "updated_at").
		From("users").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	var u entity.User
	if err := r.db.QueryRow(ctx, sql, args...).Scan(&u.ID, &u.Name, &u.Balance, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return nil, err
	}

	return &u, nil
}
