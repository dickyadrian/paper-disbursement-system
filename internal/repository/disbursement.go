package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/dickyadrian/paper-disbursement-system/internal/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DisbursementRepository struct {
	db *pgxpool.Pool
}

func NewDisbursementRepository(db *pgxpool.Pool) *DisbursementRepository {
	return &DisbursementRepository{db: db}
}

func (r *DisbursementRepository) Create(ctx context.Context, userID int64, amount int64) (*entity.Disbursement, error) {
	sql, args, err := psql.Insert("disbursements").
		Columns("user_id", "amount", "status").
		Values(userID, amount, entity.DisbursementStatusPending).
		Suffix("RETURNING id, user_id, amount, status, created_at, updated_at").
		ToSql()
	if err != nil {
		return nil, err
	}

	var d entity.Disbursement
	if err := r.db.QueryRow(ctx, sql, args...).Scan(&d.ID, &d.UserID, &d.Amount, &d.Status, &d.CreatedAt, &d.UpdatedAt); err != nil {
		return nil, err
	}

	return &d, nil
}

func (r *DisbursementRepository) GetByID(ctx context.Context, id int64) (*entity.Disbursement, error) {
	sql, args, err := psql.Select("id", "user_id", "amount", "status", "created_at", "updated_at").
		From("disbursements").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	var d entity.Disbursement
	if err := r.db.QueryRow(ctx, sql, args...).Scan(&d.ID, &d.UserID, &d.Amount, &d.Status, &d.CreatedAt, &d.UpdatedAt); err != nil {
		return nil, err
	}

	return &d, nil
}

func (r *DisbursementRepository) Process(ctx context.Context, disbursementID int64) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var d entity.Disbursement
	if err := tx.QueryRow(ctx,
		`SELECT id, user_id, amount, status FROM disbursements WHERE id = $1`,
		disbursementID,
	).Scan(&d.ID, &d.UserID, &d.Amount, &d.Status); err != nil {
		return err
	}

	if d.Status != entity.DisbursementStatusPending {
		return nil
	}

	var balance int64
	if err := tx.QueryRow(ctx,
		`SELECT balance FROM users WHERE id = $1 FOR UPDATE`,
		d.UserID,
	).Scan(&balance); err != nil {
		return err
	}

	if balance < d.Amount {
		if _, err := tx.Exec(ctx,
			`UPDATE disbursements SET status = $1, updated_at = NOW() WHERE id = $2`,
			entity.DisbursementStatusFailed, d.ID,
		); err != nil {
			return err
		}
		return tx.Commit(ctx)
	}

	newBalance := balance - d.Amount

	if _, err := tx.Exec(ctx,
		`UPDATE users SET balance = $1, updated_at = NOW() WHERE id = $2`,
		newBalance, d.UserID,
	); err != nil {
		return err
	}

	if _, err := tx.Exec(ctx,
		`INSERT INTO balance_transactions (user_id, disbursement_id, amount, balance_after) VALUES ($1, $2, $3, $4)`,
		d.UserID, d.ID, -d.Amount, newBalance,
	); err != nil {
		return err
	}

	if _, err := tx.Exec(ctx,
		`UPDATE disbursements SET status = $1, updated_at = NOW() WHERE id = $2`,
		entity.DisbursementStatusCompleted, d.ID,
	); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
