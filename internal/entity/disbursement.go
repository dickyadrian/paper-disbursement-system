package entity

import "time"

const (
	DisbursementStatusPending   = "pending"
	DisbursementStatusCompleted = "completed"
	DisbursementStatusFailed    = "failed"
)

type Disbursement struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Amount    int64     `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
